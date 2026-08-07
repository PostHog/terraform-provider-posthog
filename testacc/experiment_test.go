package tests

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/posthog/terraform-provider/internal/httpclient"
)

const testExperimentResourceName = "posthog_experiment.test"

// TestExperiment_Lifecycle drives a single experiment through the full lifecycle
// (draft → running → paused → running → stopped) via the status block, asserting the
// server-derived state after each transition. Metric-less on purpose: it exercises the
// state machine without tripping the unseen-event validation on a fresh stack.
func TestExperiment_Lifecycle(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{ // create draft
				Config: testAccExperimentConfig(name, `status { state = "draft" }`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testExperimentResourceName, "name", name),
					resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "draft"),
					resource.TestCheckResourceAttr(testExperimentResourceName, "variant.#", "2"),
					resource.TestCheckResourceAttrSet(testExperimentResourceName, "id"),
				),
			},
			{ // launch
				Config: testAccExperimentConfig(name, `status { state = "running" }`),
				Check:  resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "running"),
			},
			{ // pause
				Config: testAccExperimentConfig(name, `status { state = "paused" }`),
				Check:  resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "paused"),
			},
			{ // resume
				Config: testAccExperimentConfig(name, `status { state = "running" }`),
				Check:  resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "running"),
			},
			{ // end
				Config: testAccExperimentConfig(name, `status {
    state = "stopped"
    stopped { conclusion = "won" }
  }`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "stopped"),
					resource.TestCheckResourceAttr(testExperimentResourceName, "status.stopped.conclusion", "won"),
				),
			},
		},
	})
}

func testAccExperimentConfig(name, statusBlock string) string {
	return testAccExperimentConfigWith(name, "", statusBlock)
}

// testAccExperimentConfigWith is the shared scaffold (provider + two variants); extraAttrs are
// injected before the status block so tests can add attributes like exposure_criteria or metrics.
func testAccExperimentConfigWith(name, extraAttrs, statusBlock string) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_experiment" "test" {
  name             = %q
  feature_flag_key = %q

  variant {
    key                = "control"
    name               = "Control"
    rollout_percentage = 50
  }
  variant {
    key                = "test"
    rollout_percentage = 50
  }
%s
  %s
}
`, name, name, extraAttrs, statusBlock)
}

// testAccCheckExperimentDestroy confirms every experiment in state is soft-deleted upstream.
func testAccCheckExperimentDestroy(s *terraform.State) error {
	client := httpclient.NewDefaultClient(
		os.Getenv("POSTHOG_HOST"),
		os.Getenv("POSTHOG_API_KEY"),
		"acceptance-test",
	)
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "posthog_experiment" {
			continue
		}
		exp, status, err := client.GetExperiment(context.Background(), projectID, rs.Primary.ID)
		if err != nil {
			// A 404 (or any not-found) means it's gone — acceptable.
			if status == 404 {
				continue
			}
			return fmt.Errorf("unexpected error checking experiment %s: %w", rs.Primary.ID, err)
		}
		if exp.Deleted == nil || !*exp.Deleted {
			return fmt.Errorf("experiment %s still exists (not soft-deleted)", rs.Primary.ID)
		}
	}
	return nil
}

// TestExperiment_ShipVariant creates a running experiment then stops it by shipping the
// winning variant, asserting the terminal state.
func TestExperiment_ShipVariant(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-ship")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExperimentConfig(name, `status { state = "running" }`),
				Check:  resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "running"),
			},
			{
				Config: testAccExperimentConfig(name, `status {
    state = "stopped"
    stopped {
      ship_variant = "test"
      conclusion   = "won"
    }
  }`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "stopped"),
					resource.TestCheckResourceAttr(testExperimentResourceName, "status.stopped.ship_variant", "test"),
				),
			},
		},
	})
}

// TestExperiment_Import verifies a project-scoped import round-trip.
func TestExperiment_Import(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-imp")
	projectID := os.Getenv("POSTHOG_PROJECT_ID")

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExperimentConfig(name, `status { state = "running" }`),
			},
			{
				ResourceName:            testExperimentResourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"allow_unknown_events", "update_feature_flag_params", "status.stopped", "variant"},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					rs, ok := s.RootModule().Resources[testExperimentResourceName]
					if !ok {
						return "", fmt.Errorf("resource not found: %s", testExperimentResourceName)
					}
					return fmt.Sprintf("%s/%s", projectID, rs.Primary.ID), nil
				},
			},
		},
	})
}

// TestExperiment_ExposureCriteriaNoDrift guards against perpetual diff: a JSON blob set in config must round-trip
// with no perpetual diff even though the API echoes computed fields back.
func TestExperiment_ExposureCriteriaNoDrift(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-exp")
	cfg := testAccExperimentConfigWith(name,
		`  exposure_criteria = jsonencode({ filterTestAccounts = true })`,
		`status { state = "draft" }`)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check:  resource.TestCheckResourceAttrSet(testExperimentResourceName, "exposure_criteria"),
			},
			{ // no changes → plan must be empty (no perpetual diff)
				Config:   cfg,
				PlanOnly: true,
			},
		},
	})
}

// TestExperiment_LinkExistingFlag omits the variant blocks and points feature_flag_key at a
// feature flag managed by a separate posthog_feature_flag resource — the experiment links that
// existing multivariate flag instead of creating one, with no perpetual diff (variants live on
// the flag, not the experiment).
func TestExperiment_LinkExistingFlag(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-link")
	cfg := fmt.Sprintf(`
provider "posthog" {}

resource "posthog_feature_flag" "backing" {
  key = %q
  filters = jsonencode({
    multivariate = { variants = [
      { key = "control", rollout_percentage = 50 },
      { key = "test", rollout_percentage = 50 },
    ] }
    groups = [{ properties = [], rollout_percentage = 100 }]
  })
}

resource "posthog_experiment" "test" {
  name             = %q
  feature_flag_key = posthog_feature_flag.backing.key
  # no variant blocks -> link the existing flag
  status { state = "draft" }
}
`, name, name)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(testExperimentResourceName, "id"),
					resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "draft"),
					resource.TestCheckResourceAttr(testExperimentResourceName, "variant.#", "0"),
				),
			},
			{Config: cfg, PlanOnly: true}, // no perpetual diff — variants belong to the linked flag
		},
	})
}

// TestExperiment_MetricNoDrift guards the no-perpetual-diff behaviour for metrics specifically: the API injects a `uuid`
// and a `fingerprint` into every metric that are absent from config. The whitelist normalizer
// must drop them so the metric round-trips with no perpetual diff. Uses allow_unknown_events so
// the metric's event needn't already be ingested on a fresh stack.
func TestExperiment_MetricNoDrift(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-metric")
	cfg := testAccExperimentConfigWith(name, `  allow_unknown_events = true

  metrics = jsonencode([{
    kind        = "ExperimentMetric"
    metric_type = "mean"
    name        = "Pageviews"
    source = {
      kind  = "EventsNode"
      event = "$pageview"
      math  = "total"
    }
  }])`, `status { state = "draft" }`)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: cfg,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(testExperimentResourceName, "metrics"),
					resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "draft"),
				),
			},
			{ // no changes → plan must be empty despite server-injected uuid + fingerprint
				Config:   cfg,
				PlanOnly: true,
			},
		},
	})
}

// TestExperiment_NoStatus asserts omitting the status block fails at plan time with a clear
// message rather than the cryptic inconsistent-result error the materialized status would cause.
func TestExperiment_NoStatus(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-nostatus")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config:      testAccExperimentConfigWith(name, "", ""),
				ExpectError: regexp.MustCompile(`(?i)status.*block.*required|required.*status`),
			},
		},
	})
}

// meanMetricHCL / funnelMetricHCL are valid ExperimentMetric shapes verified against a live stack
// (both accrete server uuid/fingerprint, exercising the normalizer).
const meanMetricHCL = `jsonencode([{
    kind        = "ExperimentMetric"
    metric_type = "mean"
    name        = "%s"
    source      = { kind = "EventsNode", event = "$pageview", math = "total" }
  }])`

const funnelMetricHCL = `jsonencode([{
    kind        = "ExperimentMetric"
    metric_type = "funnel"
    name        = "Signup funnel"
    series      = [{ kind = "EventsNode", event = "$pageview" }, { kind = "EventsNode", event = "signed_up" }]
  }])`

// TestExperiment_DefinitionUpdate updates non-status definition fields (description, metrics, and
// adds metrics_secondary) on an existing draft — the BuildUpdateRequest field-mapping path that
// the status-only lifecycle test never exercises.
func TestExperiment_DefinitionUpdate(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-defupd")
	step1 := testAccExperimentConfigWith(name, `  description          = "first"
  allow_unknown_events = true
  metrics              = `+fmt.Sprintf(meanMetricHCL, "Primary")+`
`, `status { state = "draft" }`)
	step2 := testAccExperimentConfigWith(name, `  description          = "second"
  allow_unknown_events = true
  metrics              = `+fmt.Sprintf(meanMetricHCL, "Primary renamed")+`
  metrics_secondary    = `+funnelMetricHCL+`
`, `status { state = "draft" }`)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: step1,
				Check:  resource.TestCheckResourceAttr(testExperimentResourceName, "description", "first"),
			},
			{
				Config: step2,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testExperimentResourceName, "description", "second"),
					resource.TestCheckResourceAttrSet(testExperimentResourceName, "metrics_secondary"),
				),
			},
			{Config: step2, PlanOnly: true}, // no perpetual diff after the update
		},
	})
}

// testAccExperimentVariantConfig builds a config with a specific control/test split (draft).
func testAccExperimentVariantConfig(name string, controlPct, testPct int) string {
	return fmt.Sprintf(`
provider "posthog" {}

resource "posthog_experiment" "test" {
  name             = %q
  feature_flag_key = %q

  variant {
    key                = "control"
    rollout_percentage = %d
  }
  variant {
    key                = "test"
    rollout_percentage = %d
  }

  status { state = "draft" }
}
`, name, name, controlPct, testPct)
}

// TestExperiment_VariantUpdate changes the variant split on a draft — the only exerciser of the
// variantsChanged -> buildFeatureFlagConfig-on-update path.
func TestExperiment_VariantUpdate(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-varupd")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccExperimentVariantConfig(name, 50, 50),
				Check:  resource.TestCheckResourceAttr(testExperimentResourceName, "variant.0.rollout_percentage", "50"),
			},
			{
				Config: testAccExperimentVariantConfig(name, 70, 30),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(testExperimentResourceName, "variant.0.rollout_percentage", "70"),
					resource.TestCheckResourceAttr(testExperimentResourceName, "variant.1.rollout_percentage", "30"),
				),
			},
		},
	})
}

// TestExperiment_ShipLater stops an experiment with a plain end, then ships a winner on the
// already-stopped experiment in a second apply — exercising computeTransition's same-state ship
// branch and the ship-idempotency guard (a re-plan after shipping must be empty).
func TestExperiment_ShipLater(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-shiplater")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{Config: testAccExperimentConfig(name, `status { state = "running" }`)},
			{ // stop, no ship
				Config: testAccExperimentConfig(name, `status {
    state = "stopped"
    stopped { conclusion = "inconclusive" }
  }`),
				Check: resource.TestCheckResourceAttr(testExperimentResourceName, "status.state", "stopped"),
			},
			{ // ship a winner on the already-stopped experiment
				Config: testAccExperimentConfig(name, `status {
    state = "stopped"
    stopped {
      ship_variant = "test"
      conclusion   = "won"
    }
  }`),
				Check: resource.TestCheckResourceAttr(testExperimentResourceName, "status.stopped.ship_variant", "test"),
			},
		},
	})
}

// TestExperiment_IllegalBackwardTransition asserts a forward-only violation (running -> draft)
// fails with the provider's structural error rather than an invalid API call.
func TestExperiment_IllegalBackwardTransition(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-backward")
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{Config: testAccExperimentConfig(name, `status { state = "running" }`)},
			{
				Config:      testAccExperimentConfig(name, `status { state = "draft" }`),
				ExpectError: regexp.MustCompile(`(?i)no transition available`),
			},
		},
	})
}

// TestExperiment_FunnelMetric exercises a second metric_type (funnel) and asserts no perpetual
// diff, guarding that normalization is metric-shape-agnostic.
func TestExperiment_FunnelMetric(t *testing.T) {
	skipIfNotAcceptance(t)

	name := acctest.RandomWithPrefix("tf-acc-exp-funnel")
	cfg := testAccExperimentConfigWith(name, `  allow_unknown_events = true
  metrics              = `+funnelMetricHCL+`
`, `status { state = "draft" }`)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckExperimentDestroy,
		Steps: []resource.TestStep{
			{Config: cfg, Check: resource.TestCheckResourceAttrSet(testExperimentResourceName, "metrics")},
			{Config: cfg, PlanOnly: true},
		},
	})
}

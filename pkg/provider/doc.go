// Package provider implements the PostHog Terraform provider. New returns the
// constructor that providerserver.Serve runs in this module's main package, and
// that in-process bridges such as Crossplane's upjet or Pulumi's Terraform
// bridge can call directly.
//
// Only the Terraform interface is versioned: resource and data source names,
// their schemas, and state. This package and the other packages under pkg/ are
// not a stable Go API and can change in any release, including patch releases.
package provider

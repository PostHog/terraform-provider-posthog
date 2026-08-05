# Import using: project_id/cohort_id
terraform import posthog_cohort.example 12345/678

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_cohort.example 678

# Import using: project_id/experiment_id
terraform import posthog_experiment.pricing 12345/678

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_experiment.pricing 678

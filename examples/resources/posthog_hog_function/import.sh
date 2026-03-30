# Import using: project_id/hog_function_uuid
terraform import posthog_hog_function.example 12345/your-hog-function-uuid

# If project_id is configured at the provider level, you can omit it:
terraform import posthog_hog_function.example your-hog-function-uuid

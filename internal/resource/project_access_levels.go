package resource

// These apply to project-level access for roles and organization members.
const (
	ProjectAccessLevelNone   = "none"
	ProjectAccessLevelMember = "member"
	ProjectAccessLevelAdmin  = "admin"
)

var ProjectAccessLevels = []string{ProjectAccessLevelNone, ProjectAccessLevelMember, ProjectAccessLevelAdmin}

package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type OrganizationMemberUser struct {
	UUID            string  `json:"uuid"`
	Email           *string `json:"email,omitempty"`
	FirstName       *string `json:"first_name,omitempty"`
	LastName        *string `json:"last_name,omitempty"`
	IsEmailVerified *bool   `json:"is_email_verified,omitempty"`
}

type OrganizationMember struct {
	ID           string                  `json:"id"`
	User         *OrganizationMemberUser `json:"user,omitempty"`
	Level        *int                    `json:"level,omitempty"`
	JoinedAt     *string                 `json:"joined_at,omitempty"`
	Is2FAEnabled *bool                   `json:"is_2fa_enabled,omitempty"`
}

type OrganizationMembersList struct {
	Results []OrganizationMember `json:"results"`
	Next    *string              `json:"next,omitempty"`
}

type OrganizationMemberRequest struct {
	Level *int `json:"level,omitempty"`
}

func (c *PosthogClient) ListOrganizationMembers(ctx context.Context, organizationID string) ([]OrganizationMember, error) {
	var allMembers []OrganizationMember
	path := fmt.Sprintf("/api/organizations/%s/members/", organizationID)

	for path != "" {
		result, _, err := doGet[OrganizationMembersList](c, ctx, path)
		if err != nil {
			return nil, err
		}

		allMembers = append(allMembers, result.Results...)

		if result.Next == nil || *result.Next == "" {
			break
		}

		// The next URL may be absolute or relative - extract just the path
		nextURL := *result.Next
		if strings.HasPrefix(nextURL, "http") {
			// Extract path from full URL (everything after the host)
			if idx := strings.Index(nextURL, "/api/"); idx != -1 {
				path = nextURL[idx:]
			} else {
				tflog.Warn(ctx, "unexpected pagination URL format, stopping pagination", map[string]any{
					"next_url": nextURL,
				})
				break
			}
		} else {
			path = nextURL
		}
	}

	return allMembers, nil
}

// GetOrganizationMember retrieves a specific member by user UUID.
// Since there's no direct GET endpoint, we list all members and filter.
func (c *PosthogClient) GetOrganizationMember(ctx context.Context, organizationID, userUUID string) (OrganizationMember, HTTPStatusCode, error) {
	members, err := c.ListOrganizationMembers(ctx, organizationID)
	if err != nil {
		return OrganizationMember{}, 0, err
	}

	for _, member := range members {
		if member.User != nil && member.User.UUID == userUUID {
			return member, http.StatusOK, nil
		}
	}

	return OrganizationMember{}, http.StatusNotFound, fmt.Errorf("member with user UUID %s not found in organization %s", userUUID, organizationID)
}

// UpdateOrganizationMember updates a member's level.
func (c *PosthogClient) UpdateOrganizationMember(ctx context.Context, organizationID, userUUID string, input OrganizationMemberRequest) (OrganizationMember, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/members/%s/", organizationID, userUUID)
	return doPatch[OrganizationMember](c, ctx, path, input)
}

// DeleteOrganizationMember removes a member from the organization.
func (c *PosthogClient) DeleteOrganizationMember(ctx context.Context, organizationID, userUUID string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/members/%s/", organizationID, userUUID)
	return doDelete(c, ctx, path)
}

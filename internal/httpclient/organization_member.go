package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

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

		// Parse the next URL - handles both absolute and relative URLs
		parsed, err := url.Parse(*result.Next)
		if err != nil {
			tflog.Warn(ctx, "failed to parse pagination URL", map[string]any{
				"next_url": *result.Next,
				"error":    err.Error(),
			})
			break
		}
		path = parsed.Path
		if parsed.RawQuery != "" {
			path += "?" + parsed.RawQuery
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

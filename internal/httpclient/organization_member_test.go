package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListOrganizationMembers_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/api/organizations/org-123/members/", r.URL.Path, "request path should match expected endpoint")

		resp := OrganizationMembersList{
			Results: []OrganizationMember{
				{ID: "member-1", User: &OrganizationMemberUser{UUID: "user-1"}},
				{ID: "member-2", User: &OrganizationMemberUser{UUID: "user-2"}},
			},
			Next: nil,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	members, err := client.ListOrganizationMembers(context.Background(), "org-123")

	require.NoError(t, err, "listing organization members should not return an error")
	assert.Len(t, members, 2, "should return all members from single page")
}

func TestListOrganizationMembers_MultiplePages(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")

		requestURI := r.URL.RequestURI()

		switch requestURI {
		case "/api/organizations/org-123/members/":
			nextURL := "/api/organizations/org-123/members/?cursor=page2"
			resp := OrganizationMembersList{
				Results: []OrganizationMember{
					{ID: "member-1", User: &OrganizationMemberUser{UUID: "user-1"}},
					{ID: "member-2", User: &OrganizationMemberUser{UUID: "user-2"}},
				},
				Next: &nextURL,
			}
			json.NewEncoder(w).Encode(resp)

		case "/api/organizations/org-123/members/?cursor=page2":
			nextURL := "/api/organizations/org-123/members/?cursor=page3"
			resp := OrganizationMembersList{
				Results: []OrganizationMember{
					{ID: "member-3", User: &OrganizationMemberUser{UUID: "user-3"}},
				},
				Next: &nextURL,
			}
			json.NewEncoder(w).Encode(resp)

		case "/api/organizations/org-123/members/?cursor=page3":
			resp := OrganizationMembersList{
				Results: []OrganizationMember{
					{ID: "member-4", User: &OrganizationMemberUser{UUID: "user-4"}},
				},
				Next: nil,
			}
			json.NewEncoder(w).Encode(resp)

		default:
			t.Errorf("unexpected request URI: %s", requestURI)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	members, err := client.ListOrganizationMembers(context.Background(), "org-123")

	require.NoError(t, err, "listing organization members should not return an error")
	assert.Len(t, members, 4, "should return all members from all pages")
	assert.Equal(t, 3, requestCount, "should make exactly 3 requests for 3 pages")

	// Verify all members are present
	uuids := make(map[string]bool)
	for _, m := range members {
		if m.User == nil {
			continue
		}
		uuids[m.User.UUID] = true
	}
	for i := 1; i <= 4; i++ {
		expected := fmt.Sprintf("user-%d", i)
		assert.True(t, uuids[expected], "should contain %s in results", expected)
	}
}

func TestListOrganizationMembers_AbsoluteNextURL(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.Header().Set("Content-Type", "application/json")

		if requestCount == 1 {
			// First page with absolute URL - tests that url.Parse extracts the path correctly
			nextURL := "https://us.posthog.com/api/organizations/org-123/members/?cursor=page2"
			resp := OrganizationMembersList{
				Results: []OrganizationMember{
					{ID: "member-1", User: &OrganizationMemberUser{UUID: "user-1"}},
				},
				Next: &nextURL,
			}
			json.NewEncoder(w).Encode(resp)
		} else {
			resp := OrganizationMembersList{
				Results: []OrganizationMember{
					{ID: "member-2", User: &OrganizationMemberUser{UUID: "user-2"}},
				},
				Next: nil,
			}
			json.NewEncoder(w).Encode(resp)
		}
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	members, err := client.ListOrganizationMembers(context.Background(), "org-123")

	require.NoError(t, err, "listing organization members should not return an error")
	assert.Len(t, members, 2, "should return members from both pages")
	assert.Equal(t, 2, requestCount, "should make exactly 2 requests when absolute URL is used for pagination")
}

func TestGetOrganizationMember_Found(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := "alice@example.com"
		resp := OrganizationMembersList{
			Results: []OrganizationMember{
				{ID: "member-1", User: &OrganizationMemberUser{UUID: "user-1", Email: &email}},
				{ID: "member-2", User: &OrganizationMemberUser{UUID: "user-2"}},
			},
			Next: nil,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	member, status, err := client.GetOrganizationMember(context.Background(), "org-123", "user-1")

	require.NoError(t, err, "getting existing organization member should not return an error")
	assert.Equal(t, http.StatusOK, int(status), "should return 200 OK status")
	assert.Equal(t, "member-1", member.ID, "should return the correct member ID")
	require.NotNil(t, member.User, "member should have user data")
	assert.Equal(t, "user-1", member.User.UUID, "should return the correct user UUID")
}

func TestGetOrganizationMember_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := OrganizationMembersList{
			Results: []OrganizationMember{
				{ID: "member-1", User: &OrganizationMemberUser{UUID: "user-1"}},
			},
			Next: nil,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	client := NewClient(server.Client(), server.URL, "test-key", "test")
	_, status, err := client.GetOrganizationMember(context.Background(), "org-123", "user-999")

	require.Error(t, err, "getting non-existent organization member should return an error")
	assert.Equal(t, http.StatusNotFound, int(status), "should return 404 Not Found status")
}

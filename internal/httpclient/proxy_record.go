package httpclient

import (
	"context"
	"fmt"
)

type ProxyRecord struct {
	ID          string  `json:"id"`
	Domain      string  `json:"domain"`
	TargetCNAME string  `json:"target_cname"`
	Status      string  `json:"status"`
	Message     *string `json:"message,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	CreatedBy   *int64  `json:"created_by,omitempty"`
}

type ProxyRecordRequest struct {
	Domain string `json:"domain"`
}

type ProxyRecordListResponse struct {
	MaxProxyRecords *int64        `json:"max_proxy_records,omitempty"`
	Results         []ProxyRecord `json:"results"`
}

func (c *PosthogClient) ListProxyRecords(ctx context.Context, organizationID string) ([]ProxyRecord, error) {
	path := fmt.Sprintf("/api/organizations/%s/proxy_records/", organizationID)
	resp, _, err := doGet[ProxyRecordListResponse](c, ctx, path)
	if err != nil {
		return nil, err
	}
	return resp.Results, nil
}

func (c *PosthogClient) CreateProxyRecord(ctx context.Context, organizationID string, input ProxyRecordRequest) (ProxyRecord, error) {
	path := fmt.Sprintf("/api/organizations/%s/proxy_records/", organizationID)
	result, _, err := doPost[ProxyRecord](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetProxyRecord(ctx context.Context, organizationID, proxyRecordID string) (ProxyRecord, HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/proxy_records/%s/", organizationID, proxyRecordID)
	return doGet[ProxyRecord](c, ctx, path)
}

func (c *PosthogClient) DeleteProxyRecord(ctx context.Context, organizationID, proxyRecordID string) (HTTPStatusCode, error) {
	path := fmt.Sprintf("/api/organizations/%s/proxy_records/%s/", organizationID, proxyRecordID)
	return doDelete(c, ctx, path)
}

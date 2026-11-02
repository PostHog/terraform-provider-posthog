package httpclient

import (
	"context"
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

func (c *PosthogClient) ListProxyRecords(ctx context.Context, orgIDOrSlug string) ([]ProxyRecord, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/proxy_records/")
	if err != nil {
		return nil, err
	}
	resp, _, err := doGet[ProxyRecordListResponse](c, ctx, path)
	if err != nil {
		return nil, err
	}
	return resp.Results, nil
}

func (c *PosthogClient) CreateProxyRecord(ctx context.Context, orgIDOrSlug string, input ProxyRecordRequest) (ProxyRecord, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/proxy_records/")
	if err != nil {
		return ProxyRecord{}, err
	}
	result, _, err := doPost[ProxyRecord](c, ctx, path, input)
	return result, err
}

func (c *PosthogClient) GetProxyRecord(ctx context.Context, orgIDOrSlug, proxyRecordID string) (ProxyRecord, HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/proxy_records/%s/", proxyRecordID)
	if err != nil {
		return ProxyRecord{}, 0, err
	}
	return doGet[ProxyRecord](c, ctx, path)
}

func (c *PosthogClient) DeleteProxyRecord(ctx context.Context, orgIDOrSlug, proxyRecordID string) (HTTPStatusCode, error) {
	path, err := c.orgPath(ctx, orgIDOrSlug, "/api/organizations/%s/proxy_records/%s/", proxyRecordID)
	if err != nil {
		return 0, err
	}
	return doDelete(c, ctx, path)
}

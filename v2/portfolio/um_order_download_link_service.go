package portfolio

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetUMOrderDownloadLinkService get UM futures order download link by Id
type GetUMOrderDownloadLinkService struct {
	c          *Client
	downloadID string
	recvWindow *int64
}

// DownloadID set downloadId
func (s *GetUMOrderDownloadLinkService) DownloadID(downloadID string) *GetUMOrderDownloadLinkService {
	s.downloadID = downloadID
	return s
}

// RecvWindow set recvWindow
func (s *GetUMOrderDownloadLinkService) RecvWindow(recvWindow int64) *GetUMOrderDownloadLinkService {
	s.recvWindow = &recvWindow
	return s
}

// Do send request
func (s *GetUMOrderDownloadLinkService) Do(ctx context.Context, opts ...RequestOption) (*UMOrderDownloadLink, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/papi/v1/um/order/asyn/id",
		secType:  secTypeSigned,
	}
	r.setParam("downloadId", s.downloadID)
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	data, _, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res := new(UMOrderDownloadLink)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// UMOrderDownloadLink define download link response
type UMOrderDownloadLink struct {
	DownloadID          string `json:"downloadId"`
	Status              string `json:"status"` // Enum：completed，processing
	URL                 string `json:"url"`    // The link is mapped to download id
	S3Link              string `json:"s3Link"`
	Notified            bool   `json:"notified"`            // ignore
	ExpirationTimestamp int64  `json:"expirationTimestamp"` // The link would expire after this timestamp
	IsExpired           *bool  `json:"isExpired"`
}

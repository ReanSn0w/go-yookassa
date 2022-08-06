package models

import (
	"fmt"
	"net/url"
	"time"
)

type ListRequest struct {
	CreatedAtGTE   time.Time `json:"created_at.gte,omitempty"`
	CreatedAtGT    time.Time `json:"created_at.gt,omitempty"`
	CreatedAtLTE   time.Time `json:"created_at.lte,omitempty"`
	CreatedAtLT    time.Time `json:"created_at.lt,omitempty"`
	CapturedAtGTE  time.Time `json:"captured_at.gte,omitempty"`
	CapturedAtGT   time.Time `json:"captured_at.gt,omitempty"`
	CapturedAtLTE  time.Time `json:"captured_at.lte,omitempty"`
	CapturedAtLT   time.Time `json:"captured_at.lt,omitempty"`
	PaymentMethod  string    `json:"payment_method,omitempty"`
	PaymentID      string    `json:"payment_id,omitempty"`
	RefoundID      string    `json:"refound_id,omitempty"`
	Status         string    `json:"status,omitempty"`
	Limit          int       `json:"limit,omitempty"`
	FullTextSearch string    `json:"full_text_search"`
	Cursor         string    `json:"cursor,omitempty"`
}

func (p *ListRequest) Query() string {
	u := url.Values{}
	if p.CreatedAtGTE.Unix() != 0 {
		u.Add("created_at.gte", p.CreatedAtGTE.String())
	}

	if p.CreatedAtGT.Unix() != 0 {
		u.Add("created_at.gt", p.CreatedAtGT.String())
	}

	if p.CreatedAtLTE.Unix() != 0 {
		u.Add("created_at.lte", p.CreatedAtLTE.String())
	}

	if p.CreatedAtLT.Unix() != 0 {
		u.Add("created_at.lt", p.CreatedAtLT.String())
	}

	if p.CapturedAtGTE.Unix() != 0 {
		u.Add("captured_at.gte", p.CapturedAtGTE.String())
	}

	if p.CapturedAtGT.Unix() != 0 {
		u.Add("captured_at.gt", p.CapturedAtGT.String())
	}

	if p.CapturedAtLTE.Unix() != 0 {
		u.Add("captured_at.lte", p.CapturedAtLTE.String())
	}

	if p.CapturedAtLT.Unix() != 0 {
		u.Add("captured_at.lt", p.CapturedAtLT.String())
	}

	if p.PaymentMethod != "" {
		u.Add("payment_method", p.PaymentMethod)
	}

	if p.Status != "" {
		u.Add("status", p.Status)
	}

	if p.Limit != 0 {
		u.Add("limit", fmt.Sprint(p.Limit))
	}

	if p.PaymentID != "" {
		u.Add("payment_id", p.PaymentID)
	}

	if p.RefoundID != "" {
		u.Add("refound_id", p.RefoundID)
	}

	if p.Cursor != "" {
		u.Add("cursor", p.Cursor)
	}

	if p.FullTextSearch != "" {
		u.Add("full_text_search", p.FullTextSearch)
	}

	return u.Encode()
}

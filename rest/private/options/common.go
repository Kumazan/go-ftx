package options

import "time"

type Quote struct {
	ID          int64   `json:"id"`
	RequestID   int64   `json:"requestId"`
	Status      string  `json:"status"`
	QuoterSide  string  `json:"quoterSide"`
	RequestSide string  `json:"requestSide"`
	Price       float64 `json:"price"`
	Size        float64 `json:"size"`
	Collateral  float64 `json:"collateral"`

	Option Option `json:"option"`

	QuoteExpiry time.Time `json:"quoteExpiry"`
	Time        time.Time `json:"time"`
}

type Option struct {
	Underlying string    `json:"underlying"`
	Type       string    `json:"type"`
	Strike     float64   `json:"strike"`
	Expiry     time.Time `json:"expiry"`
}

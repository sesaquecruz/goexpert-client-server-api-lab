package model

type Quote struct {
	Bid       float64 `json:"bid,string"`
	Timestamp int     `json:"timestamp,string"`
}

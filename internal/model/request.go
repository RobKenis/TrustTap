package model

import "time"

type State string

const (
	Pending  State = "pending"
	Approved State = "approved"
	Denied   State = "denied"
)

type AccessRequest struct {
	IpAddress string
	Timestamp time.Time
	State     State
}

func NewAccessRequest(ip string) AccessRequest {
	return AccessRequest{
		IpAddress: ip,
		Timestamp: time.Now(),
		State:     Pending,
	}
}

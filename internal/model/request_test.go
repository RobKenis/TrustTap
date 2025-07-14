package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewAccessRequest(t *testing.T) {
	ip := "192.168.1.1"
	req := NewAccessRequest(ip)

	assert.Equal(t, ip, req.IpAddress, "IpAddress should match input")
	assert.Equal(t, Pending, req.State, "State should be Pending")
	assert.LessOrEqual(t, time.Since(req.Timestamp), time.Second, "Timestamp should be recent")
}

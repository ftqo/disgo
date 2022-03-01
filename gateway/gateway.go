package gateway

import (
	"context"
	"io"
	"time"

	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/log"
)

// Status is the state that the client is currently in
type Status int

// IsConnected returns whether the Gateway is connected
func (s Status) IsConnected() bool {
	switch s {
	case StatusWaitingForHello, StatusIdentifying, StatusWaitingForReady, StatusReady:
		return true
	default:
		return false
	}
}

// Indicates how far along the client is to connecting
const (
	StatusUnconnected Status = iota
	StatusConnecting
	StatusWaitingForHello
	StatusIdentifying
	StatusResuming
	StatusWaitingForReady
	StatusReady
	StatusDisconnected
)

type EventHandlerFunc func(gatewayEventType discord.GatewayEventType, sequenceNumber discord.GatewaySequence, shardID int, payload io.Reader)

// Gateway is what is used to connect to discord
type Gateway interface {
	Logger() log.Logger
	Config() Config
	ShardID() int
	ShardCount() int
	Open(ctx context.Context) error
	ReOpen(ctx context.Context, delay time.Duration) error
	Close(ctx context.Context)
	CloseWithCode(ctx context.Context, code int, message string)
	Status() Status
	Send(ctx context.Context, command discord.GatewayCommand) error
	Latency() time.Duration
}

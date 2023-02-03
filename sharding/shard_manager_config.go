package sharding

import (
	"github.com/disgoorg/log"

	"github.com/disgoorg/disgo/gateway"
)

// DefaultConfig returns a Config with sensible defaults.
func DefaultConfig() *Config {
	return &Config{
		Logger:            log.Default(),
		GatewayCreateFunc: gateway.New,
		ShardSplitCount:   2,
	}
}

// Config lets you configure your ShardManager instance.
type Config struct {
	Logger                    log.Logger
	ShardIDs                  map[int]struct{}
	ShardCount                int
	ShardSplitCount           int
	AutoScaling               bool
	GatewayCreateFunc         gateway.CreateFunc
	GatewayConfigOpts         []gateway.ConfigOpt
	RateLimiter               RateLimiter
	RateRateLimiterConfigOpts []RateLimiterConfigOpt
}

// ConfigOpt is a type alias for a function that takes a Config and is used to configure your Server.
type ConfigOpt func(config *Config)

// Apply applies the given ConfigOpt(s) to the Config
func (c *Config) Apply(opts []ConfigOpt) {
	for _, opt := range opts {
		opt(c)
	}
	if c.RateLimiter == nil {
		c.RateLimiter = NewRateLimiter(c.RateRateLimiterConfigOpts...)
	}
}

// WithLogger sets the logger of the ShardManager.
func WithLogger(logger log.Logger) ConfigOpt {
	return func(config *Config) {
		config.Logger = logger
	}
}

// WithShardIDs sets the shardIDs the ShardManager should manage.
func WithShardIDs(shardIDs ...int) ConfigOpt {
	return func(config *Config) {
		if config.ShardIDs == nil {
			config.ShardIDs = map[int]struct{}{}
		}
		for _, shardID := range shardIDs {
			config.ShardIDs[shardID] = struct{}{}
		}
	}
}

// WithShardCount sets the shard count of the ShardManager.
func WithShardCount(shardCount int) ConfigOpt {
	return func(config *Config) {
		config.ShardCount = shardCount
	}
}

// WithShardSplitCount sets the count a shard should be split into if it is too large.
// This is only used if AutoScaling is enabled.
func WithShardSplitCount(shardSplitCount int) ConfigOpt {
	return func(config *Config) {
		config.ShardSplitCount = shardSplitCount
	}
}

// WithAutoScaling sets whether the ShardManager should automatically re-shard shards if they are too large.
func WithAutoScaling(autoScaling bool) ConfigOpt {
	return func(config *Config) {
		config.AutoScaling = autoScaling
	}
}

// WithGatewayCreateFunc sets the function which is used by the ShardManager to create a new gateway.Gateway.
func WithGatewayCreateFunc(gatewayCreateFunc gateway.CreateFunc) ConfigOpt {
	return func(config *Config) {
		config.GatewayCreateFunc = gatewayCreateFunc
	}
}

// WithGatewayConfigOpts lets you configure the gateway.Gateway created by the ShardManager.
func WithGatewayConfigOpts(opts ...gateway.ConfigOpt) ConfigOpt {
	return func(config *Config) {
		config.GatewayConfigOpts = append(config.GatewayConfigOpts, opts...)
	}
}

// WithRateLimiter lets you inject your own srate.RateLimiter into the ShardManager.
func WithRateLimiter(rateLimiter RateLimiter) ConfigOpt {
	return func(config *Config) {
		config.RateLimiter = rateLimiter
	}
}

// WithRateRateLimiterConfigOpt lets you configure the default srate.RateLimiter used by the ShardManager.
func WithRateRateLimiterConfigOpt(opts ...RateLimiterConfigOpt) ConfigOpt {
	return func(config *Config) {
		config.RateRateLimiterConfigOpts = append(config.RateRateLimiterConfigOpts, opts...)
	}
}

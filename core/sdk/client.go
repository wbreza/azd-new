// Package sdk provides the public Azure Dev SDK for extensions.
package sdk

import (
	"fmt"

	utils "github.com/wbreza/azd-new/core/internal"
)

// Client represents the main SDK client for Azure Dev operations.
type Client struct {
	logger *utils.Logger
	config map[string]string
}

// NewClient creates a new SDK client instance.
func NewClient(config map[string]string) *Client {
	return &Client{
		logger: utils.NewLogger("SDK"),
		config: config,
	}
}

// Deploy performs a deployment operation.
func (c *Client) Deploy(appName string) error {
	c.logger.Info(fmt.Sprintf("Starting deployment for app: %s", appName))

	if !utils.ValidateConfig(c.config) {
		return fmt.Errorf("invalid configuration")
	}

	// Simulate deployment logic
	c.logger.Info(fmt.Sprintf("Deployment completed for app: %s", appName))
	return nil
}

// GetConfig returns the current configuration.
func (c *Client) GetConfig() map[string]string {
	return c.config
}

// SetConfig updates the configuration.
func (c *Client) SetConfig(key, value string) {
	if c.config == nil {
		c.config = make(map[string]string)
	}
	c.config[key] = value
	c.logger.Info(fmt.Sprintf("Configuration updated: %s = %s", key, value))
}

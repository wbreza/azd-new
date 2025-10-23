// Package utils provides shared internal utilities for Azure Dev CLI and SDK.
// This package is not intended for external consumption.
package utils

import (
	"fmt"
	"time"
)

// Logger provides basic logging functionality for internal use.
type Logger struct {
	prefix string
}

// NewLogger creates a new logger instance with the given prefix.
func NewLogger(prefix string) *Logger {
	return &Logger{prefix: prefix}
}

// Info logs an informational message.
func (l *Logger) Info(msg string) {
	fmt.Printf("[%s] INFO: %s - %s\n", time.Now().Format("15:04:05"), l.prefix, msg)
}

// Error logs an error message.
func (l *Logger) Error(msg string) {
	fmt.Printf("[%s] ERROR: %s - %s\n", time.Now().Format("15:04:05"), l.prefix, msg)
}

// Helper functions for internal use
func ValidateConfig(config map[string]string) bool {
	return len(config) > 0
}

func FormatDuration(d time.Duration) string {
	return fmt.Sprintf("%.2fs", d.Seconds())
}
//go:build integration

package main

import (
	"testing"

	"github.com/slidebolt/sdk-integration-testing"
)

const pluginID = "plugin-test-flaky"

// TestIntegration_SelfHealing verifies that the launcher's supervision loop
// restarts the plugin after each panic, and that the plugin persists its attempt
// counter via the registry so it eventually starts cleanly on attempt 3.
func TestIntegration_SelfHealing(t *testing.T) {
	s := integrationtesting.New(t, "github.com/slidebolt/plugin-test-flaky", ".")

	// RequirePlugin waits up to 30s — enough time for 2 crash/restart cycles
	// (each with exponential backoff) before the third attempt succeeds.
	s.RequirePlugin(pluginID)

	plugins, err := s.Plugins()
	if err != nil {
		t.Fatalf("GET /api/plugins: %v", err)
	}
	reg, ok := plugins[pluginID]
	if !ok {
		t.Fatalf("plugin %q not in registry after self-healing", pluginID)
	}
	if reg.Manifest.Name != "Self-Healing Plugin" {
		t.Errorf("unexpected name %q", reg.Manifest.Name)
	}
	t.Logf("plugin registered after crash recovery: %s v%s", reg.Manifest.Name, reg.Manifest.Version)
}

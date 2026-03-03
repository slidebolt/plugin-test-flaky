### `plugin-test-flaky` repository

#### Project Overview

This repository contains the `plugin-test-flaky`, a simple plugin designed for testing the resilience and supervision capabilities of the Slidebolt launcher.

#### Architecture

This is a minimal Go plugin that implements the `runner.Plugin` interface. Its purpose is to simulate a "flaky" or unstable plugin that crashes on startup but is eventually able to "self-heal".

-   **Crash Simulation**: The plugin uses its persistent `Storage` to keep track of how many times it has been started. It is hard-coded to `panic` during the `OnReady` lifecycle hook for the first two attempts.
-   **Self-Healing**: On the third attempt, it no longer panics and will start successfully.

This behavior allows developers to test that the Slidebolt launcher's supervision system correctly identifies a crashed plugin, restarts it, and that the plugin can eventually reach a stable state.

The plugin does not create any devices or entities, nor does it handle any commands or events.

#### Key Files

| File | Description |
| :--- | :--- |
| `go.mod` | Defines the Go module and its dependencies on the `sdk-runner` and `sdk-types`. |
| `main.go` | Contains the complete implementation of the flaky test plugin. |

#### Available Commands

This plugin does not handle any commands. It is intended for internal testing of the Slidebolt system.

#### Standalone Discovery Mode

This plugin supports a standalone discovery mode for rapid testing and diagnostics without requiring the full Slidebolt stack (NATS, Gateway, etc.).

To run discovery and output the results to JSON:
```bash
./plugin-test-flaky -discover
```

**Note**: Ensure any required environment variables (e.g., API keys, URLs) are set before running.

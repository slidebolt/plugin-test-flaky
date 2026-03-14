module github.com/slidebolt/plugin-test-flaky

go 1.25.7

require (
	github.com/slidebolt/registry v0.0.2
	github.com/slidebolt/sdk-integration-testing v0.0.3
	github.com/slidebolt/sdk-runner v1.20.3
	github.com/slidebolt/sdk-types v1.20.6
)

require (
	github.com/klauspost/compress v1.18.4 // indirect
	github.com/nats-io/nats.go v1.49.0 // indirect
	github.com/nats-io/nkeys v0.4.15 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/slidebolt/sdk-entities v1.20.2 // indirect
	golang.org/x/crypto v0.48.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
)

replace github.com/slidebolt/sdk-types => ../sdk-types

replace github.com/slidebolt/registry => ../registry

replace github.com/slidebolt/sdk-integration-testing => ../sdk-integration-testing

replace github.com/slidebolt/sdk-runner => ../sdk-runner

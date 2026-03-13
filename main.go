package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	regsvc "github.com/slidebolt/registry"
	runner "github.com/slidebolt/sdk-runner"
	"github.com/slidebolt/sdk-types"
)

type FlakyPlugin struct {
	attempts int
	reg      *regsvc.Registry
}

func (p *FlakyPlugin) Initialize(ctx runner.PluginContext) (types.Manifest, error) {
	p.reg = ctx.Registry

	state, _ := ctx.Registry.LoadState()
	attempts, _ := strconv.Atoi(state.Meta)
	attempts++
	p.attempts = attempts

	if err := ctx.Registry.SaveState(types.Storage{Meta: strconv.Itoa(attempts)}); err != nil {
		return types.Manifest{}, fmt.Errorf("save attempt count: %w", err)
	}

	return types.Manifest{
		ID:      "plugin-test-flaky",
		Name:    "Self-Healing Plugin",
		Version: "1.0.0",
		Schemas: types.CoreDomains(),
	}, nil
}

func (p *FlakyPlugin) Start(ctx context.Context) error {
	if p.attempts < 3 {
		panic("Deterministic Crash")
	}
	return nil
}

func (p *FlakyPlugin) Stop() error { return nil }

func (p *FlakyPlugin) OnReset() error {
	return p.reg.DeleteState()
}

func (p *FlakyPlugin) OnCommand(req types.Command, entity types.Entity) error { return nil }

func main() {
	if err := runner.RunCLI(func() runner.Plugin { return &FlakyPlugin{} }); err != nil {
		log.Fatal(err)
	}
}

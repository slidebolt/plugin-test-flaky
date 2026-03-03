package main

import (
	"context"
	"log"
	"strconv"

	runner "github.com/slidebolt/sdk-runner"
	"github.com/slidebolt/sdk-types"
)

type FlakyPlugin struct{ attempts int }

func (p *FlakyPlugin) OnInitialize(config runner.Config, state types.Storage) (types.Manifest, types.Storage) {
	attempts, _ := strconv.Atoi(state.Meta)
	attempts++
	p.attempts = attempts
	state.Meta = strconv.Itoa(attempts)
	return types.Manifest{ID: "plugin-test-flaky", Name: "Self-Healing Plugin", Version: "1.0.0"}, state
}

func (p *FlakyPlugin) OnReady() {
	if p.attempts < 3 {
		panic("Deterministic Crash")
	}
}

func (p *FlakyPlugin) WaitReady(ctx context.Context) error {
	return nil
}

func (p *FlakyPlugin) OnShutdown() {}

func (p *FlakyPlugin) OnHealthCheck() (string, error) { return "perfect", nil }
func (p *FlakyPlugin) OnStorageUpdate(current types.Storage) (types.Storage, error) {
	return current, nil
}

func (p *FlakyPlugin) OnDeviceCreate(dev types.Device) (types.Device, error) {
	return dev, nil
}
func (p *FlakyPlugin) OnDeviceUpdate(dev types.Device) (types.Device, error) { return dev, nil }
func (p *FlakyPlugin) OnDeviceDelete(id string) error                        { return nil }
func (p *FlakyPlugin) OnDevicesList(current []types.Device) ([]types.Device, error) {
	return current, nil
}
func (p *FlakyPlugin) OnDeviceSearch(q types.SearchQuery, res []types.Device) ([]types.Device, error) {
	return res, nil
}

func (p *FlakyPlugin) OnEntityCreate(e types.Entity) (types.Entity, error) { return e, nil }
func (p *FlakyPlugin) OnEntityUpdate(e types.Entity) (types.Entity, error) { return e, nil }
func (p *FlakyPlugin) OnEntityDelete(d, e string) error                    { return nil }
func (p *FlakyPlugin) OnEntitiesList(d string, c []types.Entity) ([]types.Entity, error) {
	return c, nil
}

func (p *FlakyPlugin) OnCommandTyped(req types.CommandRequest[types.GenericPayload], entity types.Entity) (types.Entity, error) {
	return entity, nil
}

func (p *FlakyPlugin) OnEventTyped(evt types.EventTyped[types.GenericPayload], entity types.Entity) (types.Entity, error) {
	return entity, nil
}

func main() {
	r, err := runner.NewRunner(&FlakyPlugin{})
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

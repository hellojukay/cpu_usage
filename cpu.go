package main

import (
	"context"
	"encoding/json"

	"github.com/shirou/gopsutil/cpu"
)

type CPU struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewCPU() *CPU {
	return &CPU{}
}

// startup is called at application startup
func (b *CPU) startup(ctx context.Context) {
	// Perform your setup here
	b.ctx = ctx
}

// domReady is called after the front-end dom has been loaded
func (b *CPU) domReady(ctx context.Context) {
	// Add your action here
}

// shutdown is called at application termination
func (b *CPU) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Greet returns a greeting for the given name
func (b *CPU) Usage() string {
	precents, err := cpu.Percent(0, true)
	if err != nil {
		return "[]"
	}
	v, err := json.Marshal(precents)
	if err != nil {
		return "[]"
	}
	return string(v)
}

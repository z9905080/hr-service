// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package api

import (
	"context"
)

// Injectors from wire.go:

func Initialize(ctx context.Context) (*App, error) {
	app := &App{}
	return app, nil
}

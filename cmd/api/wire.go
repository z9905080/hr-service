//go:build wireinject
// +build wireinject

package api

import (
	"context"
	"github.com/google/wire"
)

func Initialize(ctx context.Context) (*App, error) {
	wire.Build(ProviderSet)

	return &App{}, nil
}

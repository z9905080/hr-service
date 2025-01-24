//go:build wireinject
// +build wireinject

package api

import (
	"context"
	"github.com/google/wire"
)

func Initialize(ctx context.Context) (*App, func(), error) {
	wire.Build(wire.Struct(new(App), "*"),
		ProviderSet)
	return &App{}, nil, nil
}

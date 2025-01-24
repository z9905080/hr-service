//go:build wireinject
// +build wireinject

package migrate

import (
	"context"
	"github.com/google/wire"
)

func Initialize(ctx context.Context) (*App, error) {
	wire.Build(ProviderSet)

	return &App{}, nil
}

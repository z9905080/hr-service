//go:build wireinject
// +build wireinject

package api

import (
	"context"

	"github.com/google/wire"
)

func Initialize(ctx context.Context) (*App, func(), error) {
	wire.Struct(new(App), "*")
}

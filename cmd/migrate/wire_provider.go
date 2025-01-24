package migrate

import (
	"github.com/google/wire"
	"github.com/z9905080/hr_service/environment"
	"github.com/z9905080/hr_service/internal/infra"
	"github.com/z9905080/hr_service/pkg/logger"
)

var ProviderSet = wire.NewSet(

	NewApp,
	environment.StaticPath,
	environment.NewConfig,
	logger.NewLogger,
	infra.NewGormDB,
)

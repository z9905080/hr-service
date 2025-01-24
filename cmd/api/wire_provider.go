package api

import (
	"github.com/google/wire"
	"github.com/z9905080/hr_service/environment"
	attendanceImpl "github.com/z9905080/hr_service/internal/implement/attendance_impl"
	departmentImpl "github.com/z9905080/hr_service/internal/implement/department_impl"
	employeeImpl "github.com/z9905080/hr_service/internal/implement/employee_impl"
	"github.com/z9905080/hr_service/internal/interface/adapter"
	"github.com/z9905080/hr_service/internal/interface/handler"
	"github.com/z9905080/hr_service/internal/usecase"
	"github.com/z9905080/hr_service/pkg/logger"
)

var ProviderSet = wire.NewSet(

	NewApp,
	environment.StaticPath,
	environment.NewConfig,
	logger.NewLogger,
	adapter.NewHTTPAdapter,
	handler.NewGinHandler,
	usecase.NewUsecase,
	employeeImpl.NewEmployeeImpl,
	departmentImpl.NewDepartmentImpl,
	attendanceImpl.NewAttendanceImpl,
)

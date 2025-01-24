package migrate

import (
	"context"
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	migrates "github.com/z9905080/hr_service/cmd/migrate/migrates"
	"gorm.io/gorm"
)

func NewMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			migrate()
		},
	}
	return cmd
}

func NewRollbackCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rollback",
		Short: "rollback",
		Run: func(cmd *cobra.Command, args []string) {
			rollback()
		},
	}
	return cmd
}

func migrate() {
	app, err := Initialize(context.Background())

	if err != nil {
		panic(err)
	}

	if runErr := app.Migrate(); runErr != nil {
		panic(runErr)
	}
}

func rollback() {
	app, err := Initialize(context.Background())

	if err != nil {
		panic(err)
	}

	if runErr := app.RollbackLast(); runErr != nil {
		panic(runErr)
	}
}

type App struct {
	migrator *gormigrate.Gormigrate
}

func NewApp(db *gorm.DB) *App {

	migrateList := []*gormigrate.Migration{
		migrates.Ver20250124_INIT,
	}
	return &App{
		migrator: gormigrate.New(db, gormigrate.DefaultOptions, migrateList),
	}
}

func (a *App) Migrate() error {
	err := a.migrator.Migrate()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) RollbackLast() error {
	err := a.migrator.RollbackLast()
	if err != nil {
		return err
	}

	return nil
}

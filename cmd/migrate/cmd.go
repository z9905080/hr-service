package api

import (
	gormigrate "github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
	"os"
)

func NewMigrateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			execute()
		},
	}
	return cmd
}

func execute() {
	cmd := NewMigrateCmd()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

type App struct {
	migrator *gormigrate.Gormigrate
}

func NewApp(db *gorm.DB) *App {
	return &App{
		migrator: gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{}),
	}
}

//
//func (a *App) Run() error {
//	//return a.srv.ListenAndServe()
//}

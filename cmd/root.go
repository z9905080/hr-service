package cmd

import (
	"github.com/z9905080/hr_service/cmd/api"
	"github.com/z9905080/hr_service/cmd/migrate"
	"github.com/z9905080/hr_service/environment"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "root",
		Short: "root",
	}
	return rootCmd
}

func Execute() {
	rootCmd := NewRootCmd()
	rootCmd.AddCommand(api.NewApiCmd())
	rootCmd.AddCommand(migrate.NewMigrateCmd())
	rootCmd.AddCommand(migrate.NewRollbackCmd())

	rootCmd.PersistentFlags().StringVar((*string)(&environment.CnfPath), "config", "./environment/local_dev/config.json", "config file")

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

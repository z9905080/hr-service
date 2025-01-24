package cmd

import (
	"github.com/z9905080/hr_service/cmd/api"
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
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

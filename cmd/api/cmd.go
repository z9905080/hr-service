package api

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func NewApiCmd() *cobra.Command {
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "api",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return apiCmd
}

func ExecuteApi() {
	apiCmd := NewApiCmd()
	if err := apiCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

type App struct {
	srv *http.Server
}

func NewApp(srv *http.Server) *App {
	return &App{
		srv: srv,
	}
}

func (a *App) Run() error {
	return a.srv.ListenAndServe()
}

package api

import (
	"context"
	"github.com/spf13/cobra"
	"github.com/z9905080/hr_service/environment"
	"net/http"
)

func NewApiCmd() *cobra.Command {
	apiCmd := &cobra.Command{
		Use:   "api",
		Short: "api",
		Run: func(cmd *cobra.Command, args []string) {
			exec()
		},
	}
	apiCmd.PersistentFlags().StringVar((*string)(&environment.CnfPath), "config", "./environment/local_dev/config.json", "config file")
	return apiCmd
}

func exec() {
	app, err := Initialize(context.Background())

	if err != nil {
		panic(err)
	}

	if runErr := app.Run(); runErr != nil {
		panic(runErr)
	}

}

type App struct {
	srv *http.Server
}

func NewApp(srv http.Handler, config environment.Config) *App {
	return &App{
		srv: func() *http.Server {
			return &http.Server{
				Addr: func() string {
					if config.Http.Port == "" {
						return ":8080"
					}

					return ":" + config.Http.Port
				}(),
				Handler: srv,
			}
		}(),
	}
}

func (a *App) Run() error {
	return a.srv.ListenAndServe()
}

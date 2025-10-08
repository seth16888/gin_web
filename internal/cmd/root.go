// Package cmd provides the root command for the coauth OAuth2 server.
// Usage: coauth [command] [flags] [args]
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"seth16888/api/gw/internal/bootstrap"
	"seth16888/api/gw/internal/di"
	"syscall"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	configFile string
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c",
		"configs/config.yaml", "-c config file (default is configs/config.yaml)")
}

var rootCmd = &cobra.Command{
	Use:   "coauth [command] [flags] [args]",
	Short: "coauth is a OAuth2 server",
	Long:  `coauth is a OAuth2 server. It provides api for client to authorize and get access token.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Error: %v\n", r)
				os.Exit(1)
			}
		}()

		conf, err := bootstrap.InitConfig(configFile)
		if err != nil {
			return err
		}
		log := bootstrap.InitLogger(conf.Log)

		di.Get().Conf = conf
		di.Get().Log = log

		return err
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		defer func() {
			if di.Get() != nil && di.Get().Log != nil {
				di.Get().Log.Sync() // flushes buffer, if any
			}
		}()

		srv := bootstrap.InitServer(*di.Get().Conf, di.Get().Log)
		di.Get().Server = srv

		// Run server
		errChan := make(chan error, 1)
		srv.Run(errChan)

		quitFunc := func() {
			srv.Shutdown()
			di.Get().Log.Info("Server shutdown")
		}
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			quitFunc()
			return nil
		case err := <-errChan:
			di.Get().Log.Error("Server error", zap.Error(err))
			quitFunc()
			return err
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

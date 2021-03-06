package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/server/http"
	sqlstorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar App",
	Long:  `Calendar App`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfigFromYaml(configFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		logg := logger.New(config.Logger.Level)

		// storage := memorystorage.New()
		storage := sqlstorage.New(config.Database)
		if err := storage.Connect(context.Background()); err != nil {
			logg.Error(err.Error())
			os.Exit(0)
		}
		logg.Info("Database connected")

		calendar := app.New(logg, storage)
		server := internalhttp.NewServer(logg, calendar)

		ctx, cancel := signal.NotifyContext(context.Background(),
			syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		defer cancel()

		go func() {
			<-ctx.Done()

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			if err := server.Stop(ctx); err != nil {
				logg.Error("failed to stop http server: " + err.Error())
			}
		}()

		logg.Info("calendar is running...")

		if err := server.Start(ctx, net.JoinHostPort(config.Server.Host, config.Server.Port)); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1)
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

var configFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
}

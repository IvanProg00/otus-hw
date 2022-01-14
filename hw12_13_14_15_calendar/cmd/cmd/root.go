package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/memory"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "Calendar App",
	Long:  `Calendar App`,
	Run: func(cmd *cobra.Command, args []string) {
		config := config.NewConfig()
		logg := logger.New(config.Logger.Level)

		storage := memorystorage.New()
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

		if err := server.Start(ctx); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1) //nolint:gocritic
		}
	},
}

func Execute() error {
	return rootCmd.Execute()
}

var configFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "/etc/calendar/config.toml", "Path to configuration file")
}

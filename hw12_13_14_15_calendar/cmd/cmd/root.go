package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/app"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/logger"
	sqlstorage "github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/google/uuid"
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
		userId, err := uuid.NewRandom()
		if err != nil {
			logg.Debug(err.Error())
			os.Exit(0)
		}
		id, err := uuid.Parse("66df9cd1-3fbc-42c1-a26f-c192a2a32a6b")
		if err != nil {
			logg.Debug(err.Error())
			os.Exit(0)
		}
		if err := calendar.UpdateEvent(context.TODO(), id, "Title 1", "Description 1", time.Now(), time.Now().Add(24*time.Hour), userId); err != nil {
			logg.Error(err.Error())
			os.Exit(1)
		}

		// server := internalhttp.NewServer(logg, calendar)

		// ctx, cancel := signal.NotifyContext(context.Background(),
		// 	syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		// defer cancel()

		// go func() {
		// 	<-ctx.Done()

		// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		// 	defer cancel()

		// 	if err := server.Stop(ctx); err != nil {
		// 		logg.Error("failed to stop http server: " + err.Error())
		// 	}
		// }()

		// logg.Info("calendar is running...")

		// if err := server.Start(ctx); err != nil {
		// 	logg.Error("failed to start http server: " + err.Error())
		// 	cancel()
		// 	os.Exit(1)
		// }
	},
}

func Execute() error {
	return rootCmd.Execute()
}

var configFile string

func init() {
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
}

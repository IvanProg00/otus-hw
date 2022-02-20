package cmd

import (
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get Version",
	Long:  "Get Version",
	Run: func(cmd *cobra.Command, args []string) {
		config.PrintVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

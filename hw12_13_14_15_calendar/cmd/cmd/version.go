package cmd

import (
	"github.com/IvanProg00/otus-hw/hw12_13_14_15_calendar/internal/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.PrintVersion()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
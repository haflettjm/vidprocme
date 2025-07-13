package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vidctl",
	Short: "vidctl – coming soon",
	Long:  "vidctl – coming soon",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vidctl – coming soon")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vidctl version 1.0.0")
	},
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print help information",
	Long:  "Print help information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vidctl help")
	},
}

var submitCmd = &cobra.Command{
	Use:   "submit",
	Short: "Submit a video",
	Long:  "Submit a video",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vidctl submit")
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Print the status of a video",
	Long:  "Print the status of a video",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vidctl status")
	},
}

func main() {
	rootCmd.Execute()
}

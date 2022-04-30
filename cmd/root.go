package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "助手工具",
}

func init() {
	rootCmd.AddCommand(cssCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

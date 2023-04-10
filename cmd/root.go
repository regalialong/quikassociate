package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "quikassociate",
	Short: "Rapidly restore file associations",
	Long: `quikassociate - rapidly restore file associations
When I install new applications, they putsch my preferred apps out. Quikassociate helps bulk setting mime defaults.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

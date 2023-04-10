package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Dumps all supported file types for app",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mime := processMime(mimeType(args[0]))
		for _, app := range mime {
			fmt.Println(app)
		}
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
}

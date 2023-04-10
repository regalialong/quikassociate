package cmd

import (
	"github.com/spf13/cobra"
)

var thisisunsafe bool
var xdgbinlocation string

var associateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Set default for all supported file types",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		mime := processMime(mimeType(args[0]))
		associate(mime, args[0], thisisunsafe, xdgbinlocation)
	},
}

func init() {
	associateCmd.Flags().BoolVarP(&thisisunsafe, "thisisunsafe", "k", false, "Disable association warning")
	associateCmd.Flags().StringVarP(&xdgbinlocation, "xdgbinlocation", "b", "/usr/bin/xdg-mime", "Location to xdg-mime binary")
	rootCmd.AddCommand(associateCmd)
}

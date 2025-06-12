package cmd

import (
	"blog/pkg/bootrap"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(ServeCmd)
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the server",
	Long:  "App will be served over HTTP",
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {

	bootrap.Serve()
}

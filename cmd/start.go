package cmd

import "github.com/spf13/cobra"
import "BookApi/api"

var Port int

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start cmd starts the server on a port",
	Long: `It starts the server on a given port number, 
				Port number will be given in the cmd`,

	Run: func(cmd *cobra.Command, args []string) {
		api.Start(Port)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().IntVarP(&Port, "Port", "p", 8080, "default port for http server")
}

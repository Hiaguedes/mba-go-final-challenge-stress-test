package cli

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "stress-test",
	Short: "Stress test CLI",
	Long: `Stress test CLI is a command line interface for stress testing applications. You pass the url with --url flag and the number of requests with --requests flag. 
	The CLI will then send the specified number of requests to the specified url with --requests
	and you will specify the number of simultaneous requests with --concurrency.`,
}

func RootCmd() *cobra.Command {
	return rootCmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func AddRun(runE func(cmd *cobra.Command, args []string) error) {
	rootCmd.RunE = runE
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL to stress test")
	rootCmd.Flags().IntP("requests", "r", 1, "Number of requests to send")
	rootCmd.Flags().IntP("concurrency", "c", 1, "Number of simultaneous requests to send")
	rootCmd.MarkFlagRequired("url")
}
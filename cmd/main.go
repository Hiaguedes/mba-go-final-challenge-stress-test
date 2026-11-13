package main

import (
	"github.com/hiaguedes/mba-go-final-challenge-stress-test/internal/infra/cli"
	"github.com/hiaguedes/mba-go-final-challenge-stress-test/internal/infra/handlers"
)

func main() {

	rootCmd := cli.RootCmd()
	handler := &handlers.RequestHandler{}
	cli.AddRun(handler.HandleRequests)
	rootCmd.Execute()
}
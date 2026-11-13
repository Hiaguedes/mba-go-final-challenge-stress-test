package handlers

import (
	"fmt"

	"github.com/hiaguedes/mba-go-final-challenge-stress-test/internal/usecase"
	"github.com/spf13/cobra"
)

type RequestHandler struct {
}

func (h *RequestHandler) HandleRequests(cmd *cobra.Command, args []string) error {
	url := cmd.Flag("url").Value.String()
	requests, err := cmd.Flags().GetInt("requests")
	concurrency, err := cmd.Flags().GetInt("concurrency")
	if err != nil {
		return fmt.Errorf("failed to get flag values: %v", err)
	}

	// fmt.Printf("Sending %d requests to url %s with concurrency %d\n", requests, url, concurrency)

	reqs := usecase.NewMakeRequestsUseCase()
	inputDTO := usecase.RequestsInputDTO{
		Url: url,
		Requests: requests,
		Concurrency: concurrency,
	}

	res, err := reqs.Execute(inputDTO)

	if err != nil {
		return fmt.Errorf("failed to execute requests: %v", err)
	}

	generateReportUseCase := usecase.NewGenerateReportUseCase()
	
	reportInputDTO := usecase.GenerateReportInputDTO{
		SetHTTPStatusCodes: res.SetHTTPStatusCodes,
		Duration: res.Duration,
	}
	report, err := generateReportUseCase.Execute(reportInputDTO)
	if err != nil {
		return fmt.Errorf("failed to generate report: %v", err)
	}

	fmt.Print(report)

	return nil
}
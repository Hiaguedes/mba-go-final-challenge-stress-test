package usecase

import (
	"fmt"
	"strconv"
)

type GenerateReportInputDTO struct {
	SetHTTPStatusCodes map[int]int // Map of HTTP status codes and their counts
	Duration           float64
}

type GenerateReportUseCase struct {
}

func NewGenerateReportUseCase() *GenerateReportUseCase {
	return &GenerateReportUseCase{}
}

func (uc *GenerateReportUseCase) Execute(input GenerateReportInputDTO) (string, error) {
	report := "\n-------------\n"
	report += "Report:\n"
	report += "Duration: " + fmt.Sprintf("%.2f", input.Duration) + " seconds\n\n"
	report += "HTTP Status Codes:\n"
	report += "Number of status request: frequency counted\n"
	numberOfRequests := 0

	for statusCode, count := range input.SetHTTPStatusCodes {
		report += fmt.Sprintf("%d 			: %s\n", statusCode, strconv.Itoa(count) + " requests")
		numberOfRequests += count
	}
	report += fmt.Sprintf("\nTotal Requests: %d\n", numberOfRequests)

	return report, nil
}
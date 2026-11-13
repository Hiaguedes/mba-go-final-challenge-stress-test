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
	report += "Relatório:\n"
	report += "Duração: " + fmt.Sprintf("%.2f", input.Duration) + " segundos\n\n"
	report += "HTTP Status Codes:\n"
	report += "Número do HTTP Status Code: Frequencia de requisições com esse numero\n"
	numberOfRequests := 0

	for statusCode, count := range input.SetHTTPStatusCodes {
		report += fmt.Sprintf("%d 			  : %s\n", statusCode, strconv.Itoa(count) + " requisições")
		numberOfRequests += count
	}
	report += fmt.Sprintf("\nNúmero total de requisições: %d\n", numberOfRequests)

	return report, nil
}
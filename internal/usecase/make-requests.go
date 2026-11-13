package usecase

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type RequestsInputDTO struct {
	Url         string 
	Requests    int
	Concurrency int
}

type SingleRequestOutputDTO struct {
	HTTPStatusCode int
}

type RequestsOutputDTO struct {
	SetHTTPStatusCodes map[int]int // Map of HTTP status codes and their counts
	Duration        float64
}

type MakeRequestsUseCase struct {
}

func NewMakeRequestsUseCase() *MakeRequestsUseCase {
	return &MakeRequestsUseCase{}
}

func  SendGetRequest(input RequestsInputDTO) (SingleRequestOutputDTO, error) {
	req, err := http.NewRequest("GET", input.Url, nil)
	if err != nil {
		fmt.Println("error 1", err)
		return SingleRequestOutputDTO{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error on do", err)
		return SingleRequestOutputDTO{}, err
	}
	defer resp.Body.Close()


	return SingleRequestOutputDTO{
		HTTPStatusCode: resp.StatusCode,
	}, nil
}

func ConcurrentRequests(input RequestsInputDTO) (RequestsOutputDTO, error) {
	if input.Concurrency <= 0 || input.Requests <= 0 {
		return RequestsOutputDTO{}, fmt.Errorf("concurrency and requests must be greater than zero")
	}

	// Calcula a carga base por goroutine e o resto (requests que sobraram)
	baseRequests := input.Requests / input.Concurrency
	remainder := input.Requests % input.Concurrency

	results := make(chan SingleRequestOutputDTO, input.Requests)
	errors := make(chan error, input.Requests)
	
	setStatusCodes := make(map[int]int)
	
	var mu sync.Mutex // Protege o acesso concorrente ao map
	var wg sync.WaitGroup
	setStatusCodes[http.StatusOK] = 0
	
	start := time.Now()

	// for pra criar o numero de goroutines
	for i := 0; i < input.Concurrency; i++ {
		// pra cada goroutine crio um wait group
		wg.Add(1)

		// Distribui o resto da divisão entre as primeiras goroutines
		currentRequests := baseRequests
		if i < remainder { 
			// se o resto e 4 pra 6 goutines entao goroutine 0, 1 ,2 e 3 ganham mais uma requisicao
			currentRequests++
		}

		// Se essa goroutine não for processar nenhuma requisição, pula
		if currentRequests == 0 {
			// se nao tem mais requests encerro essa wait group
			wg.Done()
			continue
		}

		fmt.Printf("Rodando goroutina %v com %v chamadas\n", i, currentRequests)

		go func(nReqs int) {
			defer wg.Done()
			for j := 0; j < nReqs; j++ {
				result, err := SendGetRequest(input)
				if err != nil {

					errors <- err
					continue
				}

				// Protege o mapa de concorrência (Race Condition)
				mu.Lock()
				setStatusCodes[result.HTTPStatusCode]++
				mu.Unlock()

				results <- result
			}
		}(currentRequests)
	}

	// Aguarda todas as goroutines terminarem em background
	wg.Wait()
	close(results)
	close(errors)

	duration := time.Since(start).Seconds()

	return RequestsOutputDTO{
		SetHTTPStatusCodes: setStatusCodes,
		Duration:           duration,
	}, nil
}

func (m*MakeRequestsUseCase) Execute(input RequestsInputDTO) (RequestsOutputDTO, error) {


	res, err := ConcurrentRequests(input)


	if err != nil {
		return RequestsOutputDTO{}, err
	}

	return res, nil
}
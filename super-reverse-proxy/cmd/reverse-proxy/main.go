package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"super-revrse-proxy/config"
	"sync"
	"time"
)

type Job struct {
	ID       int
	Request  *http.Request
	Response http.ResponseWriter
}

type JsonResponse struct {
	Message string
	JobID   int
}

const numWorkers = 50

var responseChan = make(chan JsonResponse, 1000)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory: ", err)

	} else {
		fmt.Println("Working directory: ", wd)
	}
	configPath := filepath.Join(wd, "config", "config.yaml")
	fmt.Printf("Config path: %v\n", configPath)

	cfg := config.LoadConfig(configPath)
	fmt.Printf("Port: %v\n", cfg.Port)

	JobQueue := make(chan Job, 1000)
	var wg sync.WaitGroup

	for i := 0; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, JobQueue, &wg)
	}

	http.HandleFunc("/help", func(w http.ResponseWriter, r *http.Request) {
		job := Job{ID: int(time.Now().UnixNano()), Request: r, Response: w}
		select {
		case JobQueue <- job:
			response := <-responseChan
			w.Header().Set("Content-Type", "application/json")
			fmt.Println("response %d", response)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
				return
			}
			fmt.Printf("Job %s added to queue\n", job.ID)
		default:
			http.Error(w, "Job queue is full", http.StatusServiceUnavailable)
		}
	})
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf(cfg.Port), nil))
	close(JobQueue)
	wg.Wait()

}

func worker(i int, JobQueue chan Job, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range JobQueue {
		fmt.Printf("Worker %d processing job %d\n", i, job.ID)
		response := Process(job)
		responseChan <- response
	}
}

func Process(job Job) JsonResponse {
	//time.Sleep(2 * time.Second)
	response := JsonResponse{
		Message: "processed jon",
		JobID:   job.ID,
	}
	//fmt.Printf("Processing job %d\n", job.ID)
	//job.Response.Header().Set("Content-Type", "application/json")
	//err := json.NewEncoder(job.Response).Encode(response)
	//if err != nil {
	//	fmt.Printf("Error encoding response: %v\n", err)
	//}
	fmt.Printf("Successfully processed job %d\n", job.ID)
	return response
}

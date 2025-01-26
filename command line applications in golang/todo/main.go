package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	wd, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer wd.Close()

	doesCsvExist := false

	files, err := wd.ReadDir(-1)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.Name() == "Tasks.csv" {
			doesCsvExist = true
		}
	}

	if !doesCsvExist {
		_, err := os.Create("Tasks.csv")
		if err != nil {
			panic(err)
		}
	}
	taskCsv, err := os.OpenFile("Tasks.csv", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("There was some error while opening Tasks.csv")
		panic(err)
	}
	defer taskCsv.Close()

	//reader := csv.NewReader(taskCsv)
	//tasks, err := reader.ReadAll()
	//if err != nil {
	//	fmt.Println("Error reading taks from taskCsv")
	//}
	//for _, task := range tasks {
	//	fmt.Printf("tasks in csv %v", task)
	//}
	var taskToAdd string
	flag.StringVar(&taskToAdd, "task", "", "Task to add to csv")
	flag.Parse()
	writer := csv.NewWriter(taskCsv)
	defer writer.Flush()
	if err := writer.Write([]string{taskToAdd}); err != nil {
		fmt.Println("There was some error while writing to Tasks.csv")
		panic(err)
	}
	fmt.Printf(taskToAdd)

	fmt.Println("Task added to Tasks.csv successfully")
}

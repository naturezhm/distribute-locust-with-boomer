package main

import (
	"flag"
	"log"
	"strings"

	"github.com/naturezhm/boomer"
	"github.com/naturezhm/distribute-locust-with-boomer/docker/boomer/task"
)

// 启动入口
func main() {
	var taskParam string

	flag.StringVar(&taskParam, "task", "", "Load Test Task Need param [task]")
	flag.Parse()
	log.Printf(`HTTP benchmark is running with these args:taskParam: %s`, taskParam)

	jobs := strings.Split(taskParam, ",")
	tasks := make([]*boomer.Task, len(jobs))
	for i, jobName := range jobs {
		jobName = strings.TrimSpace(jobName)
		locustTask := task.Tasks[jobName]
		locustTask.Build()
		tasks[i] = locustTask.Task
	}
	boomer.Run(tasks...)

	log.Println(`Run Finish : %s`, taskParam)
}

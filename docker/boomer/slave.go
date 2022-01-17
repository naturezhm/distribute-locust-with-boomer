package main

import (
	"flag"
	"log"
	"strings"

	"github.com/myzhan/boomer"
	"github.com/naturezhm/distribute-locust-with-boomer/docker/boomer/task"
)

// 应用入口
func main() {
	var job string

	flag.StringVar(&job, "task", "", "Load Test Task ID:-{tracker-click}")
	flag.Parse()
	log.Printf(`HTTP benchmark is running with these args:job: %s`, job)

	jobs := strings.Split(job, ",")
	tasks := make([]*boomer.Task, len(jobs))
	for i, jobName := range jobs {
		jobName = strings.TrimSpace(jobName)
		locustTask := task.Tasks[jobName]
		locustTask.Build()
		tasks[i] = locustTask.Task
	}
	boomer.Run(tasks...)

}

package task

import (
	"github.com/naturezhm/boomer"
	"github.com/naturezhm/distribute-locust-with-boomer/docker/boomer/task/impl/fasthttp"
	"github.com/naturezhm/distribute-locust-with-boomer/docker/boomer/task/impl/http"
)

/**
* Data: used by task to execute
* ctx:  context to be shared across various task functions
* Build: build your dependency here
* Task: Boomer Task. Each task contains a function which is executed once
        per request
*/

type LocustTask struct {
	Task  *boomer.Task
	Data  interface{}
	Build func()
	Ctx   map[string]interface{}
}

//var TrackerClickTask *LocustTask

var FastHttpTask *LocustTask
var HttpTask *LocustTask
var Tasks map[string]*LocustTask

func init() {
	//TrackerClickTask = &LocustTask{
	//	Task: &boomer.Task{
	//		Name:   "tracker-click",
	//		Weight: 1000,
	//		Fn:     makeClick,
	//	},
	//	Data:  []interface{}{},
	//	Build: buildTrackerClickTask,
	//	ctx:   map[string]interface{}{},
	//}

	FastHttpTask = &LocustTask{
		Task: &boomer.Task{
			Name:   fasthttp.TaskName,
			Weight: 1000,
			Fn:     fasthttp.StartFastHttpTask,
		},
		Data:  []interface{}{},
		Build: fasthttp.BuildFastHttpTask,
		Ctx:   map[string]interface{}{},
	}

	HttpTask = &LocustTask{
		Task: &boomer.Task{
			Name:   fasthttp.TaskName,
			Weight: 1000,
			Fn:     http.StartRequest,
		},
		Data:  []interface{}{},
		Build: http.BuildHttpTask,
		Ctx:   map[string]interface{}{},
	}

	Tasks = map[string]*LocustTask{
		//"tracker-click": TrackerClickTask,
		fasthttp.TaskName: FastHttpTask,
	}
}

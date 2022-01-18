package task

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/naturezhm/boomer"
	"github.com/valyala/fasthttp"
)

var client *fasthttp.Client

type Result struct {
	Ret     string
	ErrMsg  string
	ErrCode string
}

func worker1() {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	req.Header.SetMethod("POST")
	req.SetRequestURI("https://")
	startTime := time.Now()
	err := client.Do(req, resp)
	elapsed := time.Since(startTime)
	if err != nil {
		boomer.RecordFailure("http", "unknown", elapsed.Nanoseconds()/int64(time.Millisecond), err.Error())
	}
	result := Result{}
	error1 := json.Unmarshal(resp.Body(), &result)
	if error1 != nil {
		log.Println("error: ", error1)
		log.Println(string(resp.Body()))
	}

	if result.Ret == "1" {
		boomer.RecordSuccess("http", strconv.Itoa(resp.StatusCode()), elapsed.Nanoseconds()/int64(time.Millisecond), int64(len(resp.Body())))
	} else {
		boomer.RecordFailure("http", result.ErrCode, elapsed.Nanoseconds()/int64(time.Millisecond), result.ErrMsg)
	}

}

func main() {

	client = &fasthttp.Client{
		MaxConnsPerHost: 2000,
	}

	task := &boomer.Task{
		Name:   "worker",
		Weight: 10,
		Fn:     worker1,
	}

	boomer.Run(task)
}

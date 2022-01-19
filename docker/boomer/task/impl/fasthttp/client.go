package fasthttp

import (
	"crypto/tls"
	"log"
	"strconv"
	"time"

	"github.com/naturezhm/boomer"
	"github.com/naturezhm/distribute-locust-with-boomer/docker/boomer/util"
	"github.com/valyala/fasthttp"
)

var TaskName = "fasthttp-task"
var convRate = 1 //percentage

var timeout time.Duration
var client *fasthttp.Client

func BuildFastHttpTask() {
	//links, err := data.GetData()
	//if err != nil {
	//	panic(fmt.Sprintf("Error while getting data: %+v\n", err))
	//}
	//task.FastHttpTask.Data = links

	timeout = 10 * time.Second

	client = &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NoDefaultUserAgentHeader: false,
		MaxIdleConnDuration:      5 * time.Minute,
		ReadTimeout:              2 * time.Second,
		MaxConnsPerHost:          1000,
	}
}
func StartFastHttpTask() {
	tracker, err := util.GetEnv("TRACKER_URL")
	// realIp := "113.212.121.11"
	//links := task.FastHttpTask.Data.([]*data.Empty)
	//testURL := fmt.Sprintf(tracker, links[0])
	//client := task.FastHttpTask.Ctx["http"].(*fasthttp.Client)
	log.Printf("Request URL %s\n", tracker)

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(tracker)
	request.Header.SetMethod("GET")
	// request.Header.Set("HTTP_X_FORWARDED_FOR", realIp)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36 Edg/97.0.1072.55")

	response := fasthttp.AcquireResponse()
	startTime := time.Now()
	err = client.DoTimeout(request, response, timeout)
	elapsed := time.Since(startTime)

	if err != nil {
		log.Fatalf("%v\n", err)
		boomer.RecordFailure(TaskName, "error", elapsed.Nanoseconds()/int64(time.Millisecond), err.Error())
	} else {
		boomer.RecordSuccess(TaskName, strconv.Itoa(response.StatusCode()), elapsed.Nanoseconds()/int64(time.Millisecond), int64(response.Header.ContentLength()))
	}

	fasthttp.ReleaseRequest(request)
	fasthttp.ReleaseResponse(response)

	log.Printf("Request FINISH URL %s\n", tracker)
}

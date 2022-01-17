package task

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strconv"
	"time"

	"github.com/myzhan/boomer"
	"github.com/naturezhm/distribute-locust-with-boomer/docker/boomer/data"
	"github.com/naturezhm/distribute-locust-with-boomer/docker/boomer/util"
	"github.com/valyala/fasthttp"
)

var taskName = "tracker-click"
var convRate = 1 //percentage

func buildTrackerClickTask() {
	links, err := data.GetData()
	if err != nil {
		panic(fmt.Sprintf("Error while getting data: %+v\n", err))
	}
	TrackerClickTask.Data = links
	TrackerClickTask.ctx["http"] = &fasthttp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		NoDefaultUserAgentHeader: false,
		MaxIdleConnDuration:      5 * time.Minute,
		ReadTimeout:              2 * time.Second,
		MaxConnsPerHost:          1000,
	}
}
func makeClick() {
	tracker, err := util.GetEnv("TRACKER_URL")
	// realIp := "113.212.121.11"
	links := TrackerClickTask.Data.([]*data.Empty)
	testURL := fmt.Sprintf(tracker, links[0])
	client := TrackerClickTask.ctx["http"].(*fasthttp.Client)

	request := fasthttp.AcquireRequest()
	request.SetRequestURI(testURL)
	request.Header.SetMethod("GET")
	// request.Header.Set("HTTP_X_FORWARDED_FOR", realIp)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36 Edg/97.0.1072.55")

	response := fasthttp.AcquireResponse()
	startTime := boomer.Now()
	err = client.Do(request, response)
	elapsed := boomer.Now() - startTime

	if err != nil {
		log.Fatalf("%v\n", err)
		boomer.RecordFailure(taskName, "error", 0.0, err.Error())
	} else {
		boomer.RecordSuccess(taskName, strconv.Itoa(response.StatusCode()), elapsed, int64(response.Header.ContentLength()))
	}

	if rand.Intn(100) < (convRate + 1) {
		byteArray := response.Header.Peek("Location")
		location := string(byteArray[:])
		addToClickRing(location)
	}
	fasthttp.ReleaseRequest(request)
}
func addToClickRing(location string) {
	ur, err := url.ParseRequestURI(location)
	if err != nil {
		log.Printf(`Unable to get ClickId from url: %s`, location)
		return
	}
	q := ur.Query()
	clickID := q.Get("click_id")
	if clickID == "" {
		return
	}
	ClickIds.Next()
	ClickIds.Value = clickID

}

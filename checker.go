package main

import (
	"net/http"
)

type Checker func(*MonitorConf) *MonitorLog

func checkHTTPStatus(mc *MonitorConf) *MonitorLog {
	resp, err := http.Head("http://" + mc.Url)
	if err != nil {
		return NewMonitorLog(false, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return NewMonitorLog(false, "Http status is "+resp.Status)
	}

	return NewMonitorLog(true, "Http status code is 200")
}

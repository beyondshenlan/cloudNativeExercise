package utils

import "time"

var limitQueue []MockReq
var ok bool

// MockReq 模拟请求信息
type MockReq struct {
	//请求时间
	reqTime ReqTime
	//请求IP地址
	reqIP  string
	reqUrl string
}

type ReqTime int64

// LimitFreq 参数count 次数，timeWindow 滑动窗口时间段
func LimitFreq(count uint, timeWindow int64) bool {
	currTime := time.Now().Unix()
	if limitQueue == nil {
		limitQueue = make([]MockReq, 0)
	}
	if uint(len(limitQueue)) < count {
		limitQueue = append(limitQueue, MockReq{reqTime: ReqTime(currTime)})
		return true
	}
	earlyTime := limitQueue[0].reqTime
	if currTime-int64(earlyTime) <= timeWindow {
		return false
	} else {
		limitQueue = limitQueue[1:]
		limitQueue = append(limitQueue, MockReq{reqTime: ReqTime(currTime)})
	}
	return true
}

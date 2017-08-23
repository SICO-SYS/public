/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package public

import (
	"math"
	"strconv"
	"time"
)

func Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CurrentTimeStamp() string {
	ts := time.Now().Unix()
	s := strconv.FormatInt(ts, 10)
	return s
}

func CurrentUTCISO8601() string {
	return time.Now().UTC().Format(time.RFC3339)
}

func CurrentYYYMMDD() string {
	return time.Now().UTC().Format("20060102")
}
func CurrentYYYMMDDTHHMMSSZ() string {
	return time.Now().UTC().Format("20060102T150405Z")
}

func TimesPer30s() (string, string, string) {
	currentTimeStamp := time.Now().Unix()
	currentTimes := int(math.Floor(float64(currentTimeStamp / 30)))
	stringCurrentTimes := strconv.Itoa(currentTimes)
	stringPrevTimes := strconv.Itoa(currentTimes - 1)
	stringNextTimes := strconv.Itoa(currentTimes + 1)

	return stringPrevTimes, stringCurrentTimes, stringNextTimes
}

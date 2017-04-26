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

func Per30sTimes() (string, string, string) {
	ts := time.Now().Unix()
	now := float64(ts / 30)
	snow := strconv.Itoa(int(math.Floor(now)))
	safter := strconv.Itoa(int(math.Floor(now + 1)))
	sbefore := strconv.Itoa(int(math.Floor(now - 1)))

	return sbefore, snow, safter
}

func TS() string {
	ts := time.Now().Unix()
	s := strconv.FormatInt(ts, 10)
	return s
}

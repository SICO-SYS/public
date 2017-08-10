/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package public

import (
	"encoding/hex"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func GenerateHexString() string {
	data, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	defer data.Close()
	buf := make([]byte, 16)
	data.Read(buf)
	v := hex.EncodeToString(buf)
	return v
}

func GenerateNonce() string {
	rand.Seed(time.Now().UnixNano())
	time := rand.Intn(10000) + 10000
	return strconv.Itoa(time)
}

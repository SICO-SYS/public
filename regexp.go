/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package public

import (
	"regexp"
)

func RegexpBool(pattern string, str string) (bool, error) {
	return regexp.Match(pattern, []byte(str))
}

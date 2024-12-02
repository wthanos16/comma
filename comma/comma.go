package comma

import (
	"fmt"
	"reflect"
	"regexp"
)

type Generic interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~float32 | ~float64
}

func Comma[T Generic](num T) string {
	var str string
	typeOf := reflect.TypeOf(num).Kind()
	if typeOf == reflect.Int {
		str = fmt.Sprintf("%d", num)
	} else if typeOf == reflect.Float64 {
		str = fmt.Sprintf("%.2f", num)
	}
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}

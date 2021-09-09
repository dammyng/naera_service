package helpers

import (
	"strings"
	"time"
)

func NowUnix() int64 {
	return time.Now().Unix()
}
func StartsWith(text, prefix string) bool {
	return strings.HasPrefix(text, prefix)
}
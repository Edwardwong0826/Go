package models

import (
	"fmt"
	"time"
)

// example of model is to put struct, for database interact, like spring model use by repository
// or put func to let other contoller to use
func UnixToTime(timestamp int) string {
	fmt.Println(timestamp)
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2066-01-021 15:04:05")
}

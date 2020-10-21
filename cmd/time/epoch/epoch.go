package main

import (
	"fmt"
	"time"
)

// A common requirement in programs is getting the number of seconds, millis, or nanos since the Unix epoch.
// Here's how to do it in Go

func main() {
	// Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch in seconds or nanos, respectively.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Note that there is no UnixMIllis, so to get milliseconds since epoch, manually device from nanos.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// You can also convert integer seconds or nanos since the epoch into the corresponding time.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}

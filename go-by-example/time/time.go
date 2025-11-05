package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
	/*
	   $ go run time.go
	   2025-11-05 15:17:13.464116583 +0700 +07 m=+0.000013479
	   2009-11-17 20:34:58.651387237 +0000 UTC
	   2009
	   November
	   17
	   20
	   34
	   58
	   651387237
	   UTC
	   Tuesday
	   true
	   false
	   false
	   139955h42m14.812729346s
	   139955.70411464703
	   8.397342246878823e+06
	   5.0384053481272936e+08
	   503840534812729346
	   2025-11-05 08:17:13.464116583 +0000 UTC
	   1993-11-30 08:52:43.838657891 +0000 UTC
	*/
}

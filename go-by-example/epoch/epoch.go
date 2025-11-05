package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
	/*
	   $ go run epoch.go
	   2025-11-05 15:23:39.12356895 +0700 +07 m=+0.000018272
	   1762331019
	   1762331019123
	   1762331019123568950
	   2025-11-05 15:23:39 +0700 +07
	   2025-11-05 15:23:39.12356895 +0700 +07
	*/
}

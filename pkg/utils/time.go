package utils

import (
	"fmt"
	"time"
)

func PrintElapsedTime(prefix string) {
	start := time.Now()
	for {
		total := int(time.Since(start).Seconds())
		hours := int(total / (60 * 60) % 24)
		minutes := int(total/60) % 60
		seconds := int(total % 60)

		fmt.Printf("\r%s %02d:%02d:%02d", prefix, hours, minutes, seconds)
		time.Sleep(1 * time.Second)
	}
}

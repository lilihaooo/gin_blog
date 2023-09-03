package main

import (
	"fmt"
	"time"
)

func main0() {
	// 创建一个持续时间
	duration := time.Duration(2)*time.Hour + time.Duration(30)*time.Minute

	// 输出持续时间的总分钟数
	fmt.Println("Total Minutes:", duration.Minutes())

	// 输出持续时间的总小时数
	fmt.Println("Total Hours:", duration.Hours())

	// 将持续时间转换为分钟
	durationInMinutes := duration.Minutes()
	fmt.Println("Duration in Minutes:", durationInMinutes)

	// 创建一个持续时间的切片
	durations := []time.Duration{
		2 * time.Hour,
		30 * time.Minute,
		1 * time.Hour,
	}

	// 计算切片中持续时间的总和
	totalDuration := time.Duration(0)
	for _, d := range durations {
		totalDuration += d
	}
	fmt.Println("Total Combined Duration:", totalDuration)

	// 进行时间单位的转换
	seconds := duration.Seconds()
	fmt.Println("Duration in Seconds:", seconds)

	// 格式化时间间隔
	fmt.Println("Formatted Duration:", duration.String())
}

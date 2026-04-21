package main

import "fmt"

func avgScore(scores map[string]int) (float64, int) {
	sum := 0
	count := 0
	for _, v := range scores {
		if v >= 60 {
			count++
		}
		sum += v
	}
	return float64(sum) / float64(len(scores)), count
}

func main() {
	stu_scores := map[string]int{
		"张三": 98,
	}
	stu_scores["李四"] = 80
	stu_scores["王五"] = 58
	avgScore, count := avgScore(stu_scores)
	fmt.Printf("平均成绩:%.2f,及格人数:%d", avgScore, count)
}

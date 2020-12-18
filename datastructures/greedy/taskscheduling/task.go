package taskscheduling

import (
	"fmt"
	"sort"
)

// !+
// 算法导论16.5 用拟阵求解任务调度问题
// 使用矩阵求解任务调度
// 方法:贪心

// !-

// Task 任务 deadline:截止时间 weight:在截止时间为完成前的惩罚
type Task struct {
	name     string
	deadline int
	weight   int
}

func TaskSchedule(tasks []Task) []Task {
	// 传过来的任务需要按照权值递减先排序
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].weight == tasks[j].weight {
			return tasks[i].deadline < tasks[j].deadline
		}
		return tasks[i].weight > tasks[j].weight
	})
	fmt.Printf("原先的权值:\n")
	fmt.Println(tasks)

	// 任务调度器
	res := taskSchedule(tasks)
	return res
}

// taskSchedule 进行任务调度并且输出结果
func taskSchedule(tasks []Task) []Task {
	t := make([]int, 0)
	onTime := make([]Task, 0)   //能在截止时间内完成的Task
	overTime := make([]Task, 0) //不能在截止时间内完成的Task
	for _, task := range tasks {
		deadline := task.deadline
		if len(t) <= deadline {
			l := len(t)
			for j := 0; j <= deadline-l; j++ {
				t = append(t, 0)
			}
		}
		if check(t, deadline) {
			t[deadline]++
			tempTask := task
			onTime = append(onTime, tempTask)
		} else {
			tempTask := task
			overTime = append(overTime, tempTask)
		}
	}
	sort.Slice(onTime, func(i, j int) bool {
		return onTime[i].deadline < onTime[j].deadline
	})
	// 结果函数
	showResult(onTime, overTime)
	onTime = append(onTime, overTime...)
	return onTime
}

// t 是指N(t), time是新加入的截止时间
func check(t []int, time int) bool {
	pre := 0
	for i := 1; i < len(t); i++ {
		pre += t[i]
		if i == time {
			pre++
		}
		if pre > i {
			return false
		}
	}
	return true
}

func showResult(onTime, overTime []Task) {
	fmt.Println("最优调度为:")
	for i := 0; i < len(onTime); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(onTime[i].name)
	}
	fmt.Print("<<-没有惩罚----有惩罚->>")
	punish := 0
	for i := 0; i < len(overTime); i++ {
		punish += overTime[i].weight
		fmt.Print(" ")
		fmt.Print(overTime[i].name)
	}
	fmt.Println("总惩罚为", punish)
}

package taskscheduling

import (
	"testing"
)

func TestTaskSchedule(t *testing.T) {
	//!+ case1
	var tasks1 = []Task{
		{"a1", 4, 70},
		{"a2", 2, 60},
		{"a3", 4, 50},
		{"a4", 3, 40},
		{"a5", 1, 30},
		{"a6", 4, 20},
		{"a7", 6, 10},
	}
	TaskSchedule(tasks1)
	//+- case1

	//!+ case2
	var tasks2 = []Task{
		{"a1", 4, 0},
		{"a2", 2, 10},
		{"a3", 4, 20},
		{"a4", 3, 30},
		{"a5", 1, 40},
		{"a6", 4, 50},
		{"a7", 6, 60},
	}
	//+- case2
	TaskSchedule(tasks2)
}

func isEqual(got, want []Task) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i].weight != want[i].weight {
			return false
		} else if got[i].deadline != want[i].deadline {
			return false
		} else if got[i].name != want[i].name {
			return false
		}
	}
	return true
}

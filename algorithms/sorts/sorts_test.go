package sorts

import "testing"

func testFramework(t *testing.T, sortingFunction func([]int) []int) {
	for _, test := range sortTests {
		t.Run(test.name, func(t *testing.T) {
			actual := sortingFunction(test.input)
			pos, sorted := compareSlices(actual, test.expected)
			if !sorted {
				if pos == -1 {
					t.Errorf("test %s failed due to slice length changing", test.name)
				}
				t.Errorf("test %s failed at index %d", test.name, pos)
			}
		})
	}
}

//BEGIN TESTS

func TestInsertion(t *testing.T) {
	testFramework(t, InsertionSort)
}

// very slow
//func TestSelectionSort(t *testing.T) {
//	testFramework(t, SelectionSort)
//}

func TestMergeSort(t *testing.T) {
	testFramework(t, MergeSort)
}

func TestQuickSort(t *testing.T) {
	testFramework(t, QuickSort)
}

func TestCountingSort(t *testing.T) {
	testFramework(t, CountingSort)
}

// slow
//func TestBubbleSort(t *testing.T) {
//	testFramework(t, BubbleSort)
//}

func TestRadixSort(t *testing.T) {
	testFramework(t, RadixSort)
}

func TestBucketSort(t *testing.T) {
	testFramework(t, BucketSort)
}

func TestHeapSort(t *testing.T) {
	testFramework(t, HeapSort)
}

func compareSlices(a []int, b []int) (int, bool) {
	if len(a) != len(b) {
		return -1, false
	}
	for pos := range a {
		if a[pos] != b[pos] {
			return pos, false
		}
	}
	return -1, true
}

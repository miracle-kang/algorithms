package sort_test

import (
	"math/rand"
	"testing"
	"time"

	s "../sort"
)

// randArr an util that help create a rand array
// max: 	The max value in the rand array
// size: 	The array size
func randArr(max int, size int) []int {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = r.Intn(max)
	}
	return arr[:]
}

// randOrderedArr an util that help create a rand and ordered array
// max: 	The max value in the rand array
// size: 	The array size
func randOrderedArr(max, size int) []int {
	arr := randArr(max, size)
	s.QuickSort(arr)
	return arr
}

func TestBubbleSort(t *testing.T) {
	arr := randArr(100, 20)

	t.Log("Before sort: ", arr)
	s.BubbleSort(arr)
	t.Log("After sort: ", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("Unexpected order at %d, %d, values: \n", i-1, i)
			t.Error(arr)
		}
	}
}

func TestInsertSort(t *testing.T) {
	arr := randArr(100, 20)

	t.Log("Before sort: ", arr)
	s.InsertSort(arr)
	t.Log("After sort: ", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("Unexpected order at %d, %d, values: \n", i-1, i)
			t.Error(arr)
		}
	}
}

func TestSelectSort(t *testing.T) {
	arr := randArr(100, 20)

	t.Log("Before sort: ", arr)
	s.SelectSort(arr)
	t.Log("After sort: ", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("Unexpected order at %d, %d, values: \n", i-1, i)
			t.Error(arr)
		}
	}
}

func TestMergeSort(t *testing.T) {
	arr := randArr(1000, 50)

	t.Log("Before sort: ", arr)
	s.MergeSort(arr)
	t.Log("After sort: ", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("Unexpected order at %d, %d, values: \n", i-1, i)
			t.Error(arr)
		}
	}
}

func TestQuickSort(t *testing.T) {
	arr := randArr(1000, 50)

	t.Log("Before sort: ", arr)
	s.QuickSort(arr)
	t.Log("After sort: ", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("Unexpected order at %d, %d, values: \n", i-1, i)
			t.Error(arr)
		}
	}
}

func TestCountingSort(t *testing.T) {
	arr := randArr(1000, 50)

	t.Log("Before sort: ", arr)
	s.CountingSort(arr)
	t.Log("After sort: ", arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			t.Errorf("Unexpected order at %d, %d, values: \n", i-1, i)
			t.Error(arr)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	searchValue := 30
	arr := randOrderedArr(100, 20)

	index := s.BinarySearch(arr, searchValue, false)

	if index != -1 {
		t.Logf("Search index %d, value %d", index, arr[index])
	} else {
		t.Logf("Search value %d not found.", searchValue)
	}
	if index != -1 && arr[index] != searchValue {
		t.Errorf("Unexpected index %d value %d, expect %d", index, arr[index], searchValue)
	}
}

func TestBinarySearchGt(t *testing.T) {
	searchValue := 30
	arr := randOrderedArr(100, 20)

	index := s.BinarySearchGt(arr, searchValue)

	if index != -1 {
		t.Logf("Search index %d, value %d", index, arr[index])
	} else {
		t.Logf("Search value %d not found.", searchValue)
	}
	if index != -1 && arr[index] < searchValue {
		t.Errorf("Unexpected index %d value %d", index, arr[index])
	}
}

func TestBinarySearchLt(t *testing.T) {
	searchValue := 30
	arr := randOrderedArr(100, 20)

	index := s.BinarySearchLt(arr, searchValue)

	if index != -1 {
		t.Logf("Search index %d, value %d", index, arr[index])
	} else {
		t.Logf("Search value %d not found.", searchValue)
	}
	if index != -1 && arr[index] > searchValue {
		t.Errorf("Unexpected index %d value %d", index, arr[index])
	}
}

package main

import "fmt"

func partition(l []int, debut, fin int) ([]int, []int) {
	if debut >= fin {
		return l, []int{}
	}
	pivot := l[debut]
	i := debut + 1
	j := fin
	for i <= j {
		for i <= fin && l[i] <= pivot {
			i++
		}
		for j > debut && l[j] >= pivot {
			j--
		}
		if i < j {
			l[i], l[j] = l[j], l[i]
		}
	}
	l[debut], l[j] = l[j], l[debut]
	return l[:j], l[j+1:]
}

func QuickSort(l []int) []int {
	if len(l) <= 1 {
		return l
	}
	left, right := partition(l, 0, len(l)-1)
	return append(QuickSort(left), append([]int{l[len(left)]}, QuickSort(right)...)...)
}

func main() {
	test := []int{3, 6, 8, 10, 1, 2, 1}
	sorted := QuickSort(test)
	fmt.Println(sorted)
}

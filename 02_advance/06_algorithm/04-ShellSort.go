package main

import "fmt"

func ShellSort(arr []int) {
	for inc := len(arr) / 2; inc > 0; inc /= 2 {
		for i := inc; i < len(arr); i++ {

			for j := i; j-inc >= 0; j -= inc {
				if arr[j] < arr[j-inc] {
					arr[j], arr[j-inc] = arr[j-inc], arr[j]
				} else {
					break
				}

			}
		}
	}

}
func main() {
	arr := []int{1, 3, 5, 6, 3, 4, 59, 100, 0, 7, 10, 1, 6, 3}
	ShellSort(arr)
	fmt.Println(arr)
}

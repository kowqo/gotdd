package main

func main() {

}

func Sum(numbers []int) (sum int) {
	for i := range numbers {
		sum += numbers[i]
	}

	return
}

func SumAll(numbers ...[]int) []int {
	var result = make([]int, len(numbers))

	for i := 0; i < len(numbers); i++ {
		result[i] = Sum(numbers[i])
	}

	return result
}

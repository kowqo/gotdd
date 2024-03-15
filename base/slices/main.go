package main

func main() {

}

func Sum(numbers []int) (sum int) {
	for i := range numbers {
		sum += numbers[i]
	}

	return
}

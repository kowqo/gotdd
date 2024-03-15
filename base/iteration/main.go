package main

func main() {

}

func Repeat(str string, cal int) (res string) {
	for i := 0; i < cal; i++ {
		res += str
	}
	return
}

package main

import "fmt"

func main() {

	a := "s"
	fmt.Println(a)
}

const prefixEn = "Hello "
const prefixRu = "Privet "

func Hello(name, lang string) string {
	var pref string

	if name == "" {
		if lang == "RU" {
			name = "Mir"
		} else {
			name = "World!"
		}
	}

	if lang == "RU" {
		pref = prefixRu
	} else {
		pref = prefixEn
	}

	return pref + name
}

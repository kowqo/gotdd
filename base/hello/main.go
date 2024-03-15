package main

import "fmt"

func main() {

	fmt.Println(Hello("", "EN"))
}

const prefixEn = "Hello "
const prefixRu = "Privet "
const prefixFr = "Bonjour "

func Hello(name, lang string) string {

	if name == "" {
		name = "World!"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (pref string) {
	switch lang {
	case "RU":
		pref = prefixRu
	case "EN":
		pref = prefixEn
	case "FR":
		pref = prefixFr

	}

	return
}

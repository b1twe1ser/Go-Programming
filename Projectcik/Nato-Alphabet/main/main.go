package main

import (
	"log"
	"strings"
)

/*
	@param 👉🏻
		word type string
		- word or sentence to encrypt
	@return 👈🏼
		encrypted type string
		- encrypted word or sentence in nato alphabet
*/
func encode(word string) string {

	var toEncode = strings.ToLower(word)

	var natoMap = make(map[string]string)
	var natoAlphabet = []string{"Alpha", "Beta", "Charlie", "Delta", "Echo",
		"Foxtrot", "Golf", "Hotel", "India", "Juliette", "Kilo",
		"Lima", "Mike", "November", "Oscar", "Papa", "Quebec",
		"Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey",
		"Xray", "Yankee", "Zulu"}

	for i, j := 0, 97; j < 122; i, j = i+1, j+1 {
		natoMap[string(byte(j))] = natoAlphabet[i]
	}

	var encrypted string

	for _, v := range toEncode {
		if v < 97 || v > 122 {
			encrypted += string(v)
		}
		encrypted += natoMap[string(v)] + " "
	}

	return encrypted
}

/*
	Script to encode a given string into nato alphabet ⚙️
*/
func main() {
	var toEncrypt = "Nico Rueckner is cool"

	log.Println(encode(toEncrypt))
}

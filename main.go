package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nbArgs := 0
	for range os.Args {
		nbArgs++
	}
	if nbArgs == 1 {
		fmt.Println("Missing arguments")
		return
	}
	if nbArgs > 1 {
		var filename string
		if nbArgs == 2 {
			filename = "standard.txt"
		} else {
			switch os.Args[2] {
			case "standard":
				filename = "standard.txt"
			case "shadow":
				filename = "shadow.txt"
			case "thinkertoy":
				filename = "thinkertoy.txt"
			default:
				filename = "standard.txt"
			}
		}
		file, _ := os.Open(filename)
		txt, _ := ioutil.ReadFile(filename)
		argument := Split(os.Args[1], "\\n")
		for a := 0; a < len(argument); a++ {
			for y := 0; y < 8; y++ {
				for _, tab := range argument[a] {
					pos1 := int(tab)
					pos2 := int(tab)
					pos1 = ((pos1 - 32) * 9) + 1
					pos2 = pos1 + 8
					start := LenOfTxt(txt, pos1+y)
					end := LenOfTxt(txt, pos2)
					CarriageReturn := false
					for CarriageReturn == false {
						for i := start; i <= end; i++ {
							if txt[i] == byte(13) || txt[i] == byte(10) { //retour à la ligne --> 13 || 10 --> pour que ça marche sur ios
								CarriageReturn = true
								break
							} else {
								z01.PrintRune(rune(txt[i]))
							}
						}
					}
				}
				print("\n")
			}
			file.Close()
		}
	}
}

//LenOfTxt Fonction qui compte le nombre de caractères jusqu'à la ligne nb d'un fichiers txt
func LenOfTxt(file []byte, nb int) int {
	nbLine := 0
	length := 0
	for _, x := range file {
		length++
		if x == '\n' {
			nbLine++
		}
		if nbLine == nb {
			break
		}
	}
	return length
}

//Split : Fonction qui permet de séprarer plusieurs bouts de string et les met dans un tableau de string
func Split(str, charset string) []string {
	answer := []string{}
	word := ""
	for i := 0; i < len(str); i++ {
		if isCharset(str, charset, i) && i < len(str)-1 {
			if word != "" {
				answer = app(answer, word)
				word = ""
				i = i + len(charset) - 1
			}
		} else {
			word = word + string(str[i])
		}
	}
	if word != "" {
		answer = app(answer, word)
	}
	return answer
}
//isCharset : Fonction qui permet l'exécution Split
func isCharset(str, charset string, i int) bool {
	j := 0
	for j < len(charset) && i < len(str) {
		if str[i] != charset[j] {
			return false
		}
		i++
		j++
	}
	return true
}

//app : Fonction qui permet l'exécution de Split
func app(arr []string, str string) []string {
	arr2 := make([]string, len(arr)+1)
	for i := 0; i <= len(arr)-1; i++ {
		arr2[i] = arr[i]
	}
	arr2[len(arr2)-1] = str
	return arr2
}


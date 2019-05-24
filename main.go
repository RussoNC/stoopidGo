package main

import (
    "fmt"
	"flag"
	"strings"
	"time"
)

func main() {

	//Keeping execution time
	start := time.Now()

	//Set the flag
	phraseFlag := flag.String("phrase", "foo", "a string")
	printExecTime := flag.Bool("showtime", true, "a bool")

	flag.Parse()

	//Set the string builder
	var chainString strings.Builder

	slicedFlagContent :=  strings.Split(*phraseFlag,"")

	wasUpper := false
	wasLower := false

	for i, _ := range slicedFlagContent {

		slicedCharacter := slicedFlagContent[i]

		//Detect special characters
		if getRune(slicedCharacter) < 'A' || getRune(slicedCharacter) > 'z'  {
			chainString.WriteString(slicedCharacter)
		} else {
			if !wasLower {
				modifiedChar := strings.ToLower(slicedCharacter)
				chainString.WriteString(modifiedChar)
				wasLower = true
				wasUpper = false
			} else if !wasUpper {
				modifiedChar := strings.ToUpper(slicedCharacter)
				chainString.WriteString(modifiedChar)
				wasUpper = true
				wasLower = false
			}
		}
	}

	fmt.Println(chainString.String())

	if *printExecTime {
		//Get the execution time in ms
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Println(elapsed.Seconds()*1e6)
	}

}

func getRune(stringToTransform string) rune {
	r := []rune(stringToTransform)
	return r[0]
}
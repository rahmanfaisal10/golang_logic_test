package main

import (
	"GITS_TEST/test_soal"
	"fmt"
)

func main() {
    	for {
		var arrayLong, numberQuestion int
		referseWord := make([]byte,0)
		var charToString, balanceBracket string
		fmt.Println("\n\n===================================================")
		fmt.Println("Welcome my Answer from GITS Question")
		fmt.Println("===================================================")
		fmt.Println("Please Choose Number Question from my answer :")
		fmt.Println("===================================================")
		fmt.Print("\n\n")
		fmt.Print("Your Number Question :")
		fmt.Scanln(&numberQuestion)

		switch numberQuestion {
		case 1:
			fmt.Print("\n\n")
			fmt.Println("Welcome Number Question ", numberQuestion)
			fmt.Print("\n")
			fmt.Println("Please enter the length array of the data you want: ")
			fmt.Scanln(&arrayLong)
			if arrayLong >= 10.000 {
				fmt.Println("Sorry, array length not allowed more than 10.000")
			}
			data := test_soal.QuestionNumber1(arrayLong)
			fmt.Println("My Answer is: ", data)
			fmt.Println("\n\nTERIMAKASIH")
			return
		case 2:
			fmt.Print("\n\n")
			fmt.Println("Welcome Number Question ", numberQuestion)
			fmt.Print("\n")
			fmt.Println("Please enter the length array of the data you want: ")
			fmt.Scanln(&arrayLong)
			fmt.Println("Please enter the Character you want to referse: ")
			for i := 0; i < arrayLong; i++ {
				fmt.Scanln(&referseWord)
				charToString += string(referseWord)
			}

			data := test_soal.QuestionNumber2(charToString)

			runeGo	:= make([]string,0)
			for _, v := range data {
				runeGo = append(runeGo, string(v))
			}
			fmt.Println("\nMy Answer is: ", runeGo)
			fmt.Println("\n\nTERIMAKASIH")
			return
		case 3:
			fmt.Print("\n\n")
			fmt.Println("Welcome Number Question ", numberQuestion)
			fmt.Print("\n")
			fmt.Println("Please enter balance Bracket you want: ")
			fmt.Scanln(&balanceBracket)
			data := test_soal.QuestionNumber3(balanceBracket)
			if data {
				fmt.Println("My Answer is: 'YES'")
			} else {
				fmt.Println("My Answer is: 'NO'")
			}
			fmt.Println("\n\nTERIMAKASIH")
			return
		default:
			fmt.Println("TERIMAKASIH")
			return
	    }
    }
}
package main

import (
	"fmt"

	"github.com/faint/CSVManager"
)

func main() {
	csvManager := new(CSVManager.CSVManager)
	csvManager.Load("banwords.csv")
	csvManager.Load("first_name.csv")
	csvManager.Load("last_name.csv")

	banwordsCSV, _ := csvManager.List.GetCSV("banwords")
	firstNameCSV, _ := csvManager.List.GetCSV("first_name")
	lastNameCSV, _ := csvManager.List.GetCSV("last_name")

	fmt.Println("LastName检查开始！")
	for _, last := range lastNameCSV.Lines {
		lastName, _ := last.Get("char")
		for _, ban := range banwordsCSV.Lines {
			banName, _ := ban.Get("name")

			if lastName == banName {
				fmt.Println("last name conflict:", lastName)
			}
		}
	}

	fmt.Println("LastName检查完毕！")

	fmt.Println("FirstName检查开始！")
	for _, first := range firstNameCSV.Lines {
		firstName, _ := first.Get("char")
		for _, ban := range banwordsCSV.Lines {
			banName, _ := ban.Get("name")

			if firstName == banName {
				fmt.Println("first name conflict:", firstName)
			}
		}
	}

	fmt.Println("FirstName检查完毕！")
}

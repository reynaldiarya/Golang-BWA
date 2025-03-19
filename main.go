package main

// import "fmt"
import (
	"fmt"
	"golang/calculation"
)

func main() {
	// Variable Declaration
	var name string = "Reynaldi"
	var belajarGolang bool = true
	age := 22

	// Print
	fmt.Println("Hello World")
	// sentence	:= Entity()
	// fmt.Println(sentence)

	fmt.Println("Saya", name)
	fmt.Println("Umur saya", age)

	// Conditional
	if age > 20 {
		fmt.Println("Umur saya lebih dari 20 tahun")
	} else {
		fmt.Println("Umur saya kurang dari 20 tahun")
	}
	fmt.Println("Belajar Golang", belajarGolang)

	// Calling Function
	addResult := calculation.Add(10, 5)
	fmt.Println(addResult)

	multiplyResult := calculation.Multiply(10, 5)
	fmt.Println(multiplyResult)

	// Looping
	for i := 0; i < 15; i++ {
		fmt.Println("Saya Ganteng", i+1)
	}

	title := "Golang the best language"
	for index, letter := range title {
		fmt.Println("index", index, "letter", string(letter))
	}

	// Looping with if
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index", index, "letter", string(letter))
		}
	}

	// Looping with switch
	for index, letter := range title {
		letterString := string(letter)

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index", index, "letter", letterString)
		}

	}

	// Array
	programmingLanguage := [...]string{"PHP", "Python", "Golang", "Javascript"}
	fmt.Println(programmingLanguage)
	lengthProgrammingLanguage := len(programmingLanguage)
	fmt.Println(lengthProgrammingLanguage)

	// Looping Array
	for index, lang := range programmingLanguage {
		fmt.Println("index", index, "language", lang)
	}

	// Slice
	var gamingConsole []string
	gamingConsole = append(gamingConsole, "Playstation2")
	gamingConsole = append(gamingConsole, "Playstation3")
	gamingConsole = append(gamingConsole, "Playstation4")

	for _, console := range gamingConsole {
		fmt.Println(console)
	}

	// Map
	myMap := map[string]string{"golang": "is beautiful", "php": "is lambo"}
	fmt.Println(myMap)

	// Check Key in Map
	value, isAvailable := myMap["ruby"]
	fmt.Println(value, isAvailable)

	// Slice of Map
	students := []map[string]string{
		{"name": "Reynaldi", "score": "100"},
		{"name": "Alex", "score": "80"},
	}
	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}

	// Average Calculation
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var total int
	for _, score := range scores {
		total = total + score
	}
	length := len(scores)
	average := float64(total) / float64(length)
	fmt.Println("Average", average)

	// Good Scores
	scores2 := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	var goodScores []int
	for _, score := range scores2 {
		if score >= 90 {
			goodScores = append(goodScores, score)
		}
	}
	fmt.Println("Good Scores", goodScores)

	// Call Function
	printMyResult("Hello World")
}

func printMyResult(sentence string) {
	fmt.Println(sentence)
}

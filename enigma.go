package main

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const enigma_art = "▓█████ ███▄    █   ██   ▄████   ███▄ ▄███▓ ▄▄▄     \n▓█   ▀ ██ ▀█   █ ▒▓██▒ ██▒ ▀█▒ ▓██▒▀█▀ ██▒▒████▄   \n▒███  ▓██  ▀█ ██▒░▒██░▒██░▄▄▄░ ▓██    ▓██░▒██  ▀█▄ \n▒▓█  ▄▓██▒  ▐▌██▒ ░██░░▓█  ██▓ ▒██    ▒██ ░██▄▄▄▄██\n░▒████▒██░   ▓██░ ░██░▒▓███▀▒░▒▒██▒   ░██▒ ▓█   ▓██\n░░ ▒░ ░ ▒░   ▒ ▒  ░▓  ░▒   ▒  ░░ ▒░   ░  ░ ▒▒   ▓▒█\n ░ ░  ░ ░░   ░ ▒░  ▒   ░   ░  ░░  ░      ░  ░   ▒▒ \n   ░     ░   ░ ░   ▒ ░ ░   ░ ░ ░      ░     ░   ▒  \n   ░           ░   ░       ░  ░       ░         ░  "
const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// converts input to int
func inputCheckerInt() int {

	var input string
	var result int
	var problem error

	//in event input doesn't work at all for some reason, restart loop
	for {
		_, problem = fmt.Scanln(&input)
		if problem != nil {
			fmt.Println("Invalid input. Please write something actually applicable")
			continue
		}

		//actually convert string to int
		result, problem = strconv.Atoi(input)
		//if there's a problem converting, restart loop
		if problem != nil {
			fmt.Println("Invalid input. Please write an actual number")
			continue
		}

		//if there's no problem, break and return the result
		break

	}

	return result

}

// pseudorandom alphabet map maker
func pseudorandom() map[int]string {

	//make array of numbers, and the map holding the int to string key
	numbers := make([]string, 26)
	result := make(map[int]string)

	//assign letters to numbers
	for i := range numbers {
		numbers[i] = string(alphabet[i])
	}
	//credit to openAI for this cure to my headache, black magic frickery that does pseudorandom
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	//make result in format "1, a"
	for value := range numbers {
		result[value] = numbers[value]
	}

	return result

}

// writes entire result of either cipher into a file
func fileWriter(outright_output *string, filename string) {

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("File can't be created. Cause: ", err)
	}
	defer file.Close()

	_, err = io.WriteString(file, *outright_output)

	if err != nil {
		fmt.Println("Can't write to file: ", err)
		return
	}
	fmt.Println("File written successfully!")

}

// ceasar cipher output
func ceasarOutput(mode byte, ceasar_salad_query map[int]string) {

	var ceasar_output string

	//output generation
	for position := 1; position <= 26; position++ {
		ceasar_output += strconv.Itoa(position) + ": " + ceasar_salad_query[position] + "\n"
	}

	//print mode
	if mode == 'p' {
		fmt.Println(ceasar_output)
	} else if mode == 'w' { //write to file mode
		fileWriter(&ceasar_output, "Ceasar Output.txt")
	}
	//no need to keep result, lessen memory consumption
	ceasar_output = ""
}

// generate ceasar cipher
func ceasarGenerator(og_alphabet map[int]string, spaces int) {

	ceasar_salad := make(map[int]string)
	var yes_or_no string
	placement := 0
	for counter := 0; counter <= 25; counter++ {

		placement = counter + spaces
		if placement > 26 {
			placement = placement - 26
		}

		ceasar_salad[placement] = string(alphabet[counter])

	}

	ceasarOutput('p', ceasar_salad)

	fmt.Println("\nNow, might you want to export this ceaser script to a txt file? (y/n)")
	fmt.Scanln(&yes_or_no)

	if yes_or_no == "n" {
		fmt.Println("Thank you for your honesty!")
	} else if yes_or_no == "y" {
		ceasarOutput('w', ceasar_salad)
	}

}

// generates a visible result of 4x4
func fourByFourOutput(four_by_four_query *[12][12]string, mode rune) {

	var four_by_four_output string
	//make output in string rendition, listerine = list of one line
	for _, listerine := range four_by_four_query {
		//loops trough each character of te line
		for yes := range listerine {
			//if current character is blank, make it fancy
			if listerine[yes] == "" {
				listerine[yes] = "#"
			}
			//add current char with space @ the end
			four_by_four_output += listerine[yes] + " "
		}
		four_by_four_output += "\n"
	}

	//mode is print
	if mode == 'p' {
		fmt.Println(four_by_four_output)
	} else if mode == 'w' { //mode is write to file
		fileWriter(&four_by_four_output, "Four By Four Output.txt")
	}

	//'tis a lot of space, can free it
	four_by_four_output = ""

}

// makes the 4x4 cipher grid
func fourByFourGenerator(result *[12][12]string) {

	row := 0
	column := 0
	randomized := make(map[int]string)

	//makes each of 4 squares/groups
	for square := 1; square <= 4; square++ {

		if square == 2 { //if on the second group, do again but in other corner
			row = 7
			column = 7
		} else if square == 3 { //if on third group, do ops in top right
			row = 0
			column = 7
			randomized = pseudorandom()
		} else if square == 4 { //if on fourth group, do ops in bottom left
			row = 7
			column = 0
			time.Sleep(10 * time.Nanosecond)
			rand.Seed(time.Now().UnixNano())
			randomized = pseudorandom()
		}
		for letters := 1; letters < 27; letters++ {
			//if we reached the letter J, skip it
			if letters == 10 && square < 3 {
				letters++
			}
			//in terms of the right-hand blocks, at the end of the block go to the next line
			if column == 5 {
				column = 0
				row++
			} else if column == 12 {
				column = 7
				row++
			}

			//if non-randomized groups (1 and 2)
			if square < 3 {
				result[row][column] = strings.ToLower(string(alphabet[letters-1]))
			} else { //if randomized groups
				if string(randomized[letters]) == "J" {
					letters++
				}
				result[row][column] = string(randomized[letters-1])
			}

			column++
		}
	}

}

func main() {

	var cipher, confirmation string
	var amt_of_spaces int
	rand.Seed(time.Now().UnixNano())
	alphabetDict := make(map[int]string)
	var four_by_four_result [12][12]string

	fmt.Println(enigma_art, "\nby Cyber Crusader\n\nFelicitations, partaker of Cryptography! May I ask if you are in the mood for:\na)Ceasar Cipher\nb)4x4 Square")
	fmt.Println("(Please enter corresponding letter):")
	fmt.Scanln(&cipher)

	//if the chosen cipher is ceaser
	if cipher == "a" {

		fmt.Println("How many spaces do you want to make the alphabet skip?")
		amt_of_spaces = inputCheckerInt()
		ceasarGenerator(alphabetDict, amt_of_spaces)
		//below is if the cipher chosen is the 4x4
	} else if cipher == "b" {
		fourByFourGenerator(&four_by_four_result)
		fourByFourOutput(&four_by_four_result, 'p')
		fmt.Println("Would you want to save this output to a txt file? (y/n)")
		fmt.Scanln(&confirmation)
		if confirmation == "y" {
			fourByFourOutput(&four_by_four_result, 'w')
		} else {
			fmt.Println("Understandable, have a good day")
		}

	} else {
		fmt.Println("Understandable, have a good day")
	}

	//let message be visible for 10 seconds
	time.Sleep(10 * time.Second)

}

//plans: rot13, 4x4
//4x4 dimensions: 11 high 11 wide, ended up 12x12

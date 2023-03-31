package runprog

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func RunProgram(input string, fontFamily string) string {
	var args []string
	args = append(args, input)
	args = append(args, selectFile(fontFamily))

	str := input

	argsRune := []rune(str)
	strTab := strings.Split(string(argsRune), `\n`)
	asciiTable := getAsciiValues(fontFamily)
	for _, v := range strTab {
		word := wordFormation([]rune(v), asciiTable)
		PrintWord(word)
	}
	return ReturnFile("result.txt")

}
func ReturnFile(file string) string {
	data, _ := ioutil.ReadFile(file)

	return string(data)
}

func wordFormation(argsRune []rune, asciiTable [][]string) [][]string {
	word := [][]string{}
	for _, v := range argsRune {
		if v >= 32 && v <= 126 {
			word = append(word, asciiTable[v-32])
		}
	}
	return word
}
func getAsciiValues(filename string) [][]string {
	// read the file and split it by line
	filename = selectFile(filename)
	asciiData, _ := os.ReadFile(filename)
	asciiDataStr := string(asciiData)
	scanner := bufio.NewScanner(strings.NewReader(asciiDataStr))
	scannerTable := []string{}
	for scanner.Scan() {
		scannerTable = append(scannerTable, scanner.Text())
	}
	asciiTable := strings.Split(strings.Join(scannerTable, "\n"), "\n")
	// parse the data and get each 8 lines as a letter
	bigTable := [][]string{}
	for i := 0; i < len(asciiTable)-9; i += 9 {
		bigTable = append(bigTable, asciiTable[i+1:i+9])
	}
	return bigTable
}

// This function print the ouput of your treatment
func PrintWord(words [][]string) {
	if len(words) == 0 {
		fmt.Println("test")
		return
	}
	line1 := ""
	line2 := ""
	line3 := ""
	line4 := ""
	line5 := ""
	line6 := ""
	line7 := ""
	line8 := ""
	for i := 0; i < len(words); i++ {
		line1 += words[i][0]
		line2 += words[i][1]
		line3 += words[i][2]
		line4 += words[i][3]
		line5 += words[i][4]
		line6 += words[i][5]
		line7 += words[i][6]
		line8 += words[i][7]
	}
	var filnalString = ""
	filnalString = concatString(line1, filnalString)
	filnalString = concatString(line2, filnalString)
	filnalString = concatString(line3, filnalString)
	filnalString = concatString(line4, filnalString)
	filnalString = concatString(line5, filnalString)
	filnalString = concatString(line6, filnalString)
	filnalString = concatString(line7, filnalString)
	filnalString = concatString(line8, filnalString)

	ioutil.WriteFile("result.txt", []byte(filnalString), 0666)

}
func concatString(s, str string) string {

	str += s + "\n"
	return str

}
func selectFile(filename string) string {
	switch filename {
	case "thinkertoy":
		filename = "fonts/thinkertoy.txt"
	case "shadow":
		filename = "fonts/shadow.txt"
	default:
		filename = "fonts/standard.txt"
	}
	return filename
}

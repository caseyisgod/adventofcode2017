package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/caseyisgod/routine"
)

func main() {
	input := readWordsInLines("input.txt")

	unicCount, noAnagramCount := 0, 0
	for _, a := range input {
		if ifThereOnlyUnicWords(a) {
			unicCount++
		}
		if ifThereNoAnagrams(a) {
			noAnagramCount++
		}
	}
	fmt.Printf("words inputed: %d, unic among them: %d, no anagrams: %d", len(input), unicCount, noAnagramCount)
	/*	for _, l := range "word" {
		fmt.Println(string(l))
	}*/
}

func ifThereNoAnagrams(a []string) bool {
	for i := 0; i < len(a)-1; i++ { //cicling over all words in line but last one
		for _, w := range a[i+1:] { //cicling over remaining words in line after chosen and comparing letters
			for j, l := range a[i] { //cicling over all letters in the word
				if !strings.ContainsAny(w, string(l)) {
					break //no such a letter? next word!
				} else { //there is such a letter? Remove it so it will not trigger detector again!
					w = strings.Replace(w, string(l), "", 1)
				}

				if w == "" && j == len(a[i])-1 {
					//so if all letters in the control word are passed
					//and all letters in checked word are gone, we have a winner!
					return false
				}
			}
		}
	}
	return true
}
func ifThereOnlyUnicWords(a []string) bool {
	for i := 0; i < len(a)-1; i++ {
		for _, w := range a[i+1:] {
			if strings.Compare(a[i], w) == 0 {
				return false
			}
		}
	}
	return true
}
func readWordsInLines(fn string) [][]string {
	file, err := os.Open(fn)
	defer file.Close()
	routine.Check_err(err)

	scanner := bufio.NewScanner(file)

	var temp []string
	var output [][]string
	for scanner.Scan() {
		temp = nil
		//here we split lines we just got to words
		fakeScanner := bufio.NewScanner(strings.NewReader(scanner.Text())) //probably could just use strings.Split
		// Set the split function for the scanning operation.
		fakeScanner.Split(bufio.ScanWords)
		//reading words in a line
		for fakeScanner.Scan() {
			temp = append(temp, fakeScanner.Text())
		}
		if err := fakeScanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading input:", err)
		}
		//appendings array of words we just got to the output 2dim array
		output = append(output, temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return output
}

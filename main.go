package main

import (
	"flag"
	"strings"
)

var morseCodes = map[string]string{
	"A": ".-",
	"B": "-...",
	"C": "-.-.",
	"D": "-..",
	"E": ".",
	"F": "..-.",
	"G": "--.",
	"H": "....",
	"I": "..",
	"J": ".---",
	"K": "-.-",
	"L": ".-..",
	"M": "--",
	"N": "-.",
	"O": "---",
	"P": ".--.",
	"Q": "--.-",
	"R": ".-.",
	"S": "...",
	"T": "-",
	"U": "..-",
	"V": "...-",
	"W": ".--",
	"X": "-..-",
	"Y": "-.--",
	"Z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
}

func main() {
	t := &Tree{}
	for k, v := range morseCodes {
		t.Insert(v, rune(k[0]))
	}

	var input string
	flag.StringVar(&input, "m", "", "morse code to decode")
	flag.Parse()

	words := strings.Split(input, "  ")

	output := ""

	for _, w := range words {
		letters := strings.Split(w, " ")
		for _, l := range letters {
			r, _ := t.Get(l)
			output += string(r)
		}
		output += " "
	}

	println(output)
}

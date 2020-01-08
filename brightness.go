package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
    "flag"
)

var MIN_VALUE = 1
var MAX_VALUE = 255

func main() {
    dec := flag.Int("dec", 0, "Increases brightness by the specified amount.")
    inc := flag.Int("inc", 0, "Decreases brightness by the specified amount.")
    flag.Parse()
    // A value that is not preceeded by a flag is most likely a file
    args := flag.Args()
    // With no flags specified or without a file we've got nothing to do
    if (*dec == 0 && *inc == 0) || len(args) != 1 {
        PrintUsage()
        return
    }
    f := args[0]
    numericValue := ReadBrightnessValue(f)
    if *dec != 0 {
        numericValue -= *dec
    } else {
        numericValue += *inc
    }
	if numericValue < MIN_VALUE {
		numericValue = MIN_VALUE
	} else if numericValue > MAX_VALUE {
        numericValue = MAX_VALUE
    }
    WriteBrightnessValue(f, numericValue)
    return
}

func ReadBrightnessValue(file string) int {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	numericValue, err := strconv.Atoi(string(dat[:len(dat)-1]))
	if err != nil {
		log.Fatal(err)
	}
    return numericValue
}

func WriteBrightnessValue(file string, value int) {
	d1 := []byte(strconv.Itoa(value) + "\n")
    err := ioutil.WriteFile(file, d1, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintUsage() {
    fmt.Println("USAGE:")
    fmt.Printf("\tbrightness -inc|-dec VALUE FILE\n")
}

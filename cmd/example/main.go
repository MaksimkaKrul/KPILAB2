package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/MaksimkaKrul/KPILAB2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file with expression")
	outputFile      = flag.String("o", "", "Output file for result")
)

func main() {
	flag.Parse()

	if (*inputExpression != "" && *inputFile != "") || (*inputExpression == "" && *inputFile == "") {
		fmt.Fprintln(os.Stderr, "Error: Use either -e or -f, not both")
		os.Exit(1)
	}

	var input io.Reader
	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	var output io.Writer = os.Stdout
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	handler := &lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}

	if err := handler.Compute(); err != nil {
		fmt.Fprintf(os.Stderr, "Computation error: %v\n", err)
		os.Exit(1)
	}
}

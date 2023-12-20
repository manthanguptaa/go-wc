package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/cobra"
)

const newLineAdjustment int = 2

type Counts struct {
	lines int
	words int
	bytes int
	chars int
}

func countBytesLinesWordsChars(scanner *bufio.Scanner, countLines, countBytes,
	countWords, countChars bool) Counts {
	var counts Counts
	for scanner.Scan() {
		line := scanner.Text()
		arrayOfBytes := scanner.Bytes()
		if countLines {
			counts.lines++
		}
		if countBytes {
			// Potential optimization: count bytes directly from arrayOfBytes
			counts.bytes += len(arrayOfBytes) + newLineAdjustment
		}
		if countWords {
			words := strings.Fields(line)
			counts.words += len(words)
		}
		if countChars {
			counts.chars += utf8.RuneCountInString(line) + newLineAdjustment
		}
	}
	return counts
}

var rootCmd = &cobra.Command{
	Use:  "go-wc [flags] [filename]",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var filename string
		var input io.Reader

		if len(args) == 0 {
			input = os.Stdin
		} else {
			filename = args[0]
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not open the file: %v", err)
				os.Exit(-1)
			}
			defer file.Close()
			input = file
		}

		scanner := bufio.NewScanner(input)
		getBytesToggle, _ := cmd.Flags().GetBool("c")
		getLinesToggle, _ := cmd.Flags().GetBool("l")
		getWordsToggle, _ := cmd.Flags().GetBool("w")
		getCharsToggle, _ := cmd.Flags().GetBool("m")

		// default option when no flag is given in the command
		if !getBytesToggle && !getLinesToggle && !getWordsToggle && !getCharsToggle {
			getBytesToggle = true
			getLinesToggle = true
			getWordsToggle = true
		}

		var counts Counts = countBytesLinesWordsChars(scanner, getLinesToggle, getBytesToggle, getWordsToggle, getCharsToggle)

		// if there is an issue in the middle of reading the file, handling the error
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "something bad happened with the file %v", err)
			os.Exit(-1)
		}

		if getBytesToggle && getLinesToggle && getWordsToggle {
			fmt.Fprintf(os.Stdout, "%9d %9d %9d %s\n", counts.lines, counts.words, counts.bytes, filename)
		} else if getBytesToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", counts.bytes, filename)
		} else if getLinesToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", counts.lines, filename)
		} else if getWordsToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", counts.words, filename)
		} else if getCharsToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", counts.chars, filename)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("c", "c", false, "Outputs the numbers of bytes in the file")
	rootCmd.Flags().BoolP("l", "l", false, "Outputs the numbers of lines in the file")
	rootCmd.Flags().BoolP("w", "w", false, "Outputs the numbers of words in the file")
	rootCmd.Flags().BoolP("m", "m", false, "Outputs the numbers of characters in the file")
}

var RootCmd = rootCmd

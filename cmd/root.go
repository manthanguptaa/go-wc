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

func FormatOutput(getBytesToggle, getLinesToggle, getWordsToggle, getCharsToggle int) {

}

var rootCmd = &cobra.Command{
	Use:  "go-wc [flags] [filename]",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var numberOfBytes, numberOfLines, numberOfWords, numberOfChars int
		var filename string
		var input io.Reader

		// checking if there is an input file or pipe input
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
		if !(getBytesToggle && getLinesToggle && getWordsToggle && getCharsToggle) {
			getBytesToggle = true
			getLinesToggle = true
			getWordsToggle = true
		}

		for scanner.Scan() {
			line := scanner.Text()
			arrayOfBytes := scanner.Bytes()
			if getLinesToggle {
				numberOfLines++
			}
			if getBytesToggle {
				numberOfBytes += len(arrayOfBytes) + 2 // +2 is to account for newline bytes
			}
			if getWordsToggle {
				words := strings.Fields(line)
				numberOfWords += len(words)
			}
			if getCharsToggle {
				numberOfChars += utf8.RuneCountInString(line) + 2 // +2 is to account for newline chars
			}
		}
		if getBytesToggle && getLinesToggle && getWordsToggle {
			fmt.Fprintf(os.Stdout, "%9d %9d %9d %s\n", numberOfLines, numberOfWords, numberOfBytes, filename)
		} else if getBytesToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", numberOfBytes, filename)
		} else if getLinesToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", numberOfLines, filename)
		} else if getWordsToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", numberOfWords, filename)
		} else if getCharsToggle {
			fmt.Fprintf(os.Stdout, "%9d %s\n", numberOfChars, filename)
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

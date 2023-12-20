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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "go-wc [filename]",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var numberOfBytes, numberOfLines, numberOfWords, numberOfChars int
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
		getBytesToggle, _ := cmd.Flags().GetBool("getBytesToggle")
		getLinesToggle, _ := cmd.Flags().GetBool("getLinesToggle")
		getWordsToggle, _ := cmd.Flags().GetBool("getWordsToggle")
		getCharsToggle, _ := cmd.Flags().GetBool("getCharsToggle")

		if !getBytesToggle && !getLinesToggle && !getWordsToggle {
			getBytesToggle = true
			getLinesToggle = true
			getWordsToggle = true
		}

		for scanner.Scan() {
			if getLinesToggle {
				numberOfLines++
			}
			if getBytesToggle {
				arrayOfBytes := scanner.Bytes()
				numberOfBytes += len(arrayOfBytes) + 2 // +2 is to account for newline bytes
			}
			if getWordsToggle {
				line := scanner.Text()
				words := strings.Fields(line)
				numberOfWords += len(words)
			}
			if getCharsToggle {
				line := scanner.Text()
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
	rootCmd.Flags().BoolP("getBytesToggle", "c", false, "Outputs the numbers of bytes in the file")
	rootCmd.Flags().BoolP("getLinesToggle", "l", false, "Outputs the numbers of lines in the file")
	rootCmd.Flags().BoolP("getWordsToggle", "w", false, "Outputs the numbers of words in the file")
	rootCmd.Flags().BoolP("getCharsToggle", "m", false, "Outputs the numbers of characters in the file")
}

var RootCmd = rootCmd

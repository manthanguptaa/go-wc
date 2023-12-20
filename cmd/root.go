package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:  "go-wc [filename]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var numberOfBytes, numberOfLines, numberOfWords int
		filename := args[0]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not open the file: %v", err)
			os.Exit(-1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		getBytesToggle, _ := cmd.Flags().GetBool("getBytesToggle")
		getLinesToggle, _ := cmd.Flags().GetBool("getLinesToggle")
		getWordsToggle, _ := cmd.Flags().GetBool("getWordsToggle")

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
		}
		if getBytesToggle {
			fmt.Fprintf(os.Stdout, "%6d %s\n", numberOfBytes, filename)
		} else if getLinesToggle {
			fmt.Fprintf(os.Stdout, "%6d %s\n", numberOfLines, filename)
		} else if getWordsToggle {
			fmt.Fprintf(os.Stdout, "%6d %s\n", numberOfWords, filename)
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
}

var RootCmd = rootCmd

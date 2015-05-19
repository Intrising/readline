// test program for the readline package
package main

import (
	"fmt"
	"github.com/gobs/readline"
	"os"
	"os/signal"
	"strings"
)

func CtrlC_Handler(sc <-chan os.Signal) {
	sigtimes := 0
	for s := range sc {
		fmt.Println("Got signal:", s)
		sigtimes += 1
		if sigtimes >= 3 {
			os.Exit(0)
		}
	}
}

var (
	words   = []string{"alpha", "beta", "charlie", "delta", "another", "banana", "carrot", "delimiter"}
	matches = make([]string, 0, len(words))
)

//
// this will use CompletionEntry to match the "command" name (first word in the line
//
func AttemptedCompletion(text string, start, end int) []string {
	if start == 0 { // this is the command to match
		return readline.CompletionMatches(text, CompletionEntry)
	} else {
		return nil
	}
}

//
// this return a match in the "words" array
//
func CompletionEntry(prefix string, index int) string {
	if index == 0 {
		matches = matches[:0]

		for _, w := range words {
			if strings.HasPrefix(w, prefix) {
				matches = append(matches, w)
			}
		}
	}

	if index < len(matches) {
		return matches[index]
	} else {
		return ""
	}
}

func setupMetaKey() {
	readline.ClearSignals()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go CtrlC_Handler(c)
}

func main() {
	setupMetaKey()
	prompt := "by your command> "

	// loop until ReadLine returns nil (signalling EOF)
L:
	for {
		result := readline.ReadLine(&prompt)
		if result == nil { // exit loop
			break L
		}

		line := *result

		switch line {
		case "": // ignore blank lines
			continue

		case "exit", "quit":
			break L

		case "help":
			fmt.Println("Available commands:")
			fmt.Println("help exit attempted compentry nocompletion prompt")
			fmt.Println()
			fmt.Println("Try completion for these words:")
			fmt.Println(words)

		case "att", "attempted":
			readline.SetAttemptedCompletionFunction(AttemptedCompletion)
			readline.SetCompletionEntryFunction(nil)

		case "compentry":
			readline.SetCompletionEntryFunction(CompletionEntry)
			readline.SetAttemptedCompletionFunction(nil)

		case "nocomp", "nocompletion":
			readline.SetCompletionEntryFunction(nil)
			readline.SetAttemptedCompletionFunction(nil)

		default:
			if strings.HasPrefix(line, "prompt ") {
				prompt = strings.TrimPrefix(line, "prompt ")
			} else {
				fmt.Println(line)
			}

			readline.AddHistory(line) //allow user to recall this line
		}
	}
}

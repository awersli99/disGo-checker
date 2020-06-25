package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

var (
	inputfile     string
	outputfile    string
	threads       int
	workingTokens []string
	brokenTokens  []string
	brokenfile    string
	wg            sync.WaitGroup
	amountDone    int
)
var tokens = make([]string, 0)

func init() {

	flag.StringVar(&inputfile, "i", "", "[REQUIRED] Your input file containing unchecked tokens")
	flag.IntVar(&threads, "t", 1, "[REQUIRED] The amount of threads the app should run on (at least one!)")
	flag.StringVar(&outputfile, "o", "", "[REQUIRED] Your output file name that will contain checked tokens")
	flag.StringVar(&brokenfile, "b", "", "[OPTIONAL] The output file for broken tokens")
	flag.Parse()

	if outputfile == "" || inputfile == "" || threads < 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func checkToken(token string) {
	dg, _ := discordgo.New("Bot " + token)
	err := dg.Open()
	time.Sleep(time.Millisecond * 10)
	if err != nil {
		fmt.Println("[BROKEN]", token)
		brokenTokens = append(brokenTokens, token)
		return
	}
	dg.Close()
	fmt.Println("[WORKING]", token)
	workingTokens = append(workingTokens, token)
	amountDone++
}

func worker() {
	for i := len(tokens); i > 0; i-- {
		if len(tokens) != 0 {
			var token string
			token, tokens = tokens[0], tokens[1:]
			checkToken(token)
		} else {
			wg.Done()
			return
		}
	}
}

func readFile(inputfile string) {
	file, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	wg.Add(threads)
	readFile(inputfile)
	fmt.Println("---------------------- Discord token checker ------------------------")
	time.Sleep(time.Second * 2)
	start := time.Now()
	for i := threads; i > 0; i-- {
		go worker()
	}
	wg.Wait()
	err := writeLines(workingTokens, outputfile)
	err = writeLines(brokenTokens, brokenfile)
	if err != nil {
		log.Fatal(err)
	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("---Checked", amountDone, "tokens in", elapsed, "Tokens saved to:", outputfile, "---")

}

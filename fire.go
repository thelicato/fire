/*
	Code by @thelicato
  Happy hacking :D
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func log(verbose bool, s string, a ...any) {
	if verbose {
		fmt.Printf(s+"\n", a...)
	}
}

func printBanner(verbose bool) {
	if verbose {
		fmt.Println(`
    ██████   ███                    
    ███░░███ ░░░                     
   ░███ ░░░  ████  ████████   ██████ 
  ███████   ░░███ ░░███░░███ ███░░███
 ░░░███░     ░███  ░███ ░░░ ░███████ 
   ░███      ░███  ░███     ░███░░░  
   █████     █████ █████    ░░██████ 
  ░░░░░     ░░░░░ ░░░░░      ░░░░░░  
                                     
  `)
		fmt.Println("\033[32mfi\033[0mlter \033[32mre\033[0msolved - Made by @thelicato - https://github.com/thelicato/fire")
		fmt.Println()
	}
}

func worker(wg *sync.WaitGroup, jobs chan string) {
	defer wg.Done()
	for domain := range jobs {
		_, err := net.ResolveIPAddr("ip4", domain)
		if err != nil {
			continue
		}
		fmt.Println(domain)
	}
}

func main() {
	concurrency := 20
	verbose := false
	flag.IntVar(&concurrency, "c", 20, "Specify concurrency level")
	flag.BoolVar(&verbose, "v", false, "Set to VERBOSE")

	flag.Parse()

	// A banner should never be missed
	printBanner(verbose)

	// Initial check to be sure the stdin is not empty
	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if fi.Mode()&os.ModeNamedPipe == 0 {
		fmt.Println("[main] Nothing in stdin :(")
		os.Exit(1)
	}

	// Set 'jobs' as channel of strings
	jobs := make(chan string)

	var wg sync.WaitGroup

	// Run the workers
	log(verbose, "[main]: starting %d workers", concurrency)
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go worker(&wg, jobs)
	}

	log(verbose, "[main]: Workers started.")

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if len(line) == 0 {
			close(jobs)
			break
		}
		jobs <- line
	}

	close(jobs)

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	wg.Wait()
	log(verbose, "[main]: Execution completed")
}

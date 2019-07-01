package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var bold = color.New(color.FgGreen).Add(color.Bold)
var red = color.New(color.FgRed).Add(color.Bold)

func main() {

	welcome()
	repo := getRepoOrPanic()

	url := fmt.Sprint("https://api.github.com/users/", repo, "/repos")
	results := getRepoListOrPanic(url)
	for idx, r := range getRepoListOrPanic(url) {
		bold.Println(fmt.Sprintf("[%v] - %s", idx, r.Name))
	}
	bold.Println("[x]", " - Exit")

	text := blockForInput()
	if text == "x" {
		return
	}

	i, err := strconv.Atoi(text)
	if err != nil {
		red.Println(err)
		return
	}

	logrus.Println("Cloning: ", results[i].CloneURL)
	done := showLoading()
	cmd := exec.Command("git", "clone", results[i].CloneURL)
	if _, err = cmd.Output(); err != nil {
		red.Println(err)
		done <- true
		return
	}
	done <- true
	fmt.Println("✅  Cloned ✅  ")
	os.Exit(0)
}

func showLoading() chan bool {
	c := make(chan bool)
	go func() {
		fmt.Println("")
		for {
			select {
			case <-c:
				fmt.Println("")
				return
			case <-time.After(500 * time.Millisecond):
				fmt.Print("🦖")
			}
		}
	}()
	return c
}

func blockForInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")
	return text
}

func getRepoOrPanic() string {
	a := os.Args
	if len(a) != 2 {
		red.Println("Expected 2 arguments - the github user, got,", len(a))
		for _, ai := range os.Args {
			red.Println(ai)
		}
		panic("Expected args")
	}
	return a[1]
}

func welcome() {
	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	fmt.Print("🦖🦖🦖  ")
	c.Print("Welcome to Ghulp. My little Github clone helper")
	fmt.Println("  🦖🦖🦖")

}

type Repo struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
}

func getRepoListOrPanic(url string) []Repo {
	resp, err := http.Get(url)
	if err != nil {
		red.Println(err)
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		red.Println(err)
		panic(err)
	}
	results := make([]Repo, 0)
	if err = json.Unmarshal(body, &results); err != nil {
		red.Println(err)
		panic(err)
	}
	return results

}

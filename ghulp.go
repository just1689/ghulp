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
)

func main() {

	// Create a new color object
	c := color.New(color.FgCyan).Add(color.Underline)
	fmt.Print("🦖🦖🦖  ")
	c.Print("Welcome to Ghulp. My little Github clone helper")
	fmt.Println("  🦖🦖🦖")

	green := color.New(color.FgGreen)
	bold := green.Add(color.Bold)

	rc := color.New(color.FgRed)
	red := rc.Add(color.Bold)

	a := os.Args
	if len(a) != 2 {
		red.Println("Expected 2 arguments - the github user, got,", len(a))
		for _, ai := range os.Args {
			red.Println(ai)
		}
		return
	}

	url := fmt.Sprint("https://api.github.com/users/", a[1], "/repos")
	results, err := getRepoList(url)
	if err != nil {
		red.Println("Could not get users repos")
		return
	}

	for idx, r := range results {
		bold.Println(fmt.Sprintf("[%v] - %s", idx, r.Name))
	}
	bold.Println("[x]", " - Exit")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")
	if text == "x" {
		return
	}

	i, err := strconv.Atoi(text)
	if err != nil {
		red.Println(err)
		return
	}

	logrus.Println("Cloning: ", results[i].CloneURL)
	cmd := exec.Command("git", "clone", results[i].CloneURL)
	_, err = cmd.Output()
	if err != nil {
		red.Println(err)
		return
	}

	fmt.Println("✅  Cloned ✅  ")

}

type Repo struct {
	Name     string `json:"name"`
	CloneURL string `json:"clone_url"`
}

func getRepoList(url string) (results []Repo, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	results = make([]Repo, 0)
	err = json.Unmarshal(body, &results)
	if err != nil {
	}
	return

}

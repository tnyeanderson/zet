package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Generate a zettel id (zid) from the current time
//   Format: YYYYMMDDHHMMSS
func GenerateZid() string {
	return time.Now().Format("20060102150305")
}

func ParseZid(s string) string {
	re := regexp.MustCompile("[0-9]{14}")
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return ""
	}
	return matches[0]
}

func FormatSearchResult(zid string) string {
	return fmt.Sprint(zid, " : ", GetTitle(zid))
}

func GetTitle(zid string) string {
	readme := ReadmePath(zid)
	output, err := exec.Command("head", "-n", "1", readme).Output()
	if err != nil {
		fmt.Println(err)
	}
	o := string(output)
	if strings.HasPrefix(o, "# ") {
		o = strings.Replace(o, "# ", "", 1)
	}
	return strings.ReplaceAll(o, "\n", "")
}

func ReadmePath(zid string) string {
	return filepath.Join(zid, "README.md")
}

func FormatTitle(title string) string {
	return strings.ToUpper(string(title[0])) + title[1:]
}

func cmdCreate(args []string) {
	zid := GenerateZid()
	title := FormatTitle(strings.Join(args, " "))
	CreateZettel(zid, title)
}

func cmdSearch(args []string) {
	zids := SearchZettels(args)
	for _, zid := range zids {
		fmt.Println(FormatSearchResult(zid))
	}
}

func cmdPrint(args []string) {
	zid := SearchAndSelectZettel(args)
	fmt.Println("----")
	PrintZettel(zid)
}

func cmdEdit(args []string) {
	zid := SearchAndSelectZettel(args)
	EditZettel(zid)
}

func cmdView(args []string) {
	zid := SearchAndSelectZettel(args)
	ViewZettel(zid)
}

func SearchAndSelectZettel(terms []string) string {
	return SelectZettelPrompt(SearchZettels(terms))
}

func SelectZettelPrompt(zids []string) string {
	var selectionIndex string
	for i, zid := range zids {
		fmt.Printf("%d) %s\n", i+1, FormatSearchResult(zid))
	}
	fmt.Print("Select a zet: ")
	fmt.Scan(&selectionIndex)
	i, err := strconv.Atoi(selectionIndex)
	if err != nil {
		fmt.Println(err)
	}
	return zids[i-1]
}

func CreateZettel(zid string, title string) {
	readme := ReadmePath(zid)
	if err := os.Mkdir(zid, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	file, err := os.Create(readme)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString("# " + title)
	file.Close()
	EditZettel(zid)
}

func runInteractive(name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func EditZettel(zid string) {
	readme := ReadmePath(zid)
	runInteractive("vim", readme)
	OnSave(zid)
}

func ViewZettel(zid string) {
	readme := ReadmePath(zid)
	runInteractive("view", readme)
	OnSave(zid)
}

func PrintZettel(zid string) {
	readme := ReadmePath(zid)
	runInteractive("cat", readme)
	fmt.Println()
	fmt.Println("Zettel:", zid)
}

func SearchZettels(terms []string) (zids []string) {
	query := strings.Join(terms, "|")
	r, err := exec.Command("git", "grep", "-i", "--name-only", "-E", query).Output()
	if err != nil {
		fmt.Println("No zettels found")
	}
	results := strings.Split(string(r[:]), "\n")
	for _, f := range results {
		zid := ParseZid(f)
		if zid != "" {
			zids = append(zids, zid)
		}
	}
	return
}

func OnSave(zid string) {
	readme := ReadmePath(zid)
	contents, err := ioutil.ReadFile(readme)
	if err != nil {
		fmt.Println(err)
	}
	if len(contents) == 0 {
		// File is empty
		fmt.Println("empty file!")
		err := os.Remove(zid)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func RunZetCommand(command string, args []string) {
	switch command {
	case "c", "create":
		cmdCreate(args)
	case "e", "edit":
		cmdEdit(args)
	//case "h","help":
	case "p", "print":
		cmdPrint(args)
	case "s", "search":
		cmdSearch(args)
	case "v", "view":
		cmdView(args)
	default:
		fmt.Println("Command not found:", command)
	}
}

func cli() {
	if len(os.Args) < 2 {
		fmt.Println("Missing command argument")
		os.Exit(1)
	}
	command := os.Args[1]
	RunZetCommand(command, os.Args[2:])
}

func main() {
	cli()
}

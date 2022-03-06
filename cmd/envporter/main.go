package main

import (
	"bufio"
	"flag"
	"github.com/darkjinnee/envporter/internal/app/envporter"
	"github.com/darkjinnee/go-err"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Args struct {
	File *os.File
	Path string
	Reg  string
	Min  int
	Max  int
}

var args Args

func init() {
	flag.StringVar(
		&args.Path,
		"file",
		".env",
		"path to .env file",
	)
	flag.IntVar(
		&args.Min,
		"min",
		3000,
		"minimum port number range",
	)
	flag.IntVar(
		&args.Max,
		"max",
		9999,
		"maximum port number range",
	)
	flag.StringVar(
		&args.Reg,
		"reg",
		`\BPORT\b`,
		"regular expression for substring search",
	)
	flag.Parse()

	if args.Min > args.Max {
		log.Fatalf("[Error] envporter.init: Range incorrect")
	}

	dirPath, err := os.Getwd()
	goerr.Fatal(
		err,
		"[Error] envporter.init: Failed to return the root path of directory",
	)

	if strings.EqualFold(args.Path, ".env") {
		args.Path = strings.Join([]string{
			dirPath,
			args.Path,
		}, "/")
	}

	args.File, err = os.OpenFile(
		args.Path,
		os.O_RDWR,
		os.ModePerm,
	)
	goerr.Fatal(
		err,
		"[Error] envporter.init: Failed to open file",
	)
}

func main() {
	var content string
	fileScanner := bufio.NewScanner(args.File)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if ok, _ := regexp.MatchString(args.Reg, line); !ok {
			content += strings.Join([]string{
				line,
				"\n",
			}, "")
			continue
		}

		port := envporter.FreePort(args.Min, args.Max)
		lineArr := strings.Split(line, "=")
		content += strings.Join([]string{
			lineArr[0],
			"=",
			strconv.Itoa(port),
			"\n",
		}, "")
	}

	envporter.Overwrite(
		args.File,
		content,
	)
}

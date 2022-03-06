package main

import (
	"flag"
	"fmt"
	"github.com/darkjinnee/envporter/internal/app/envporter"
	"github.com/darkjinnee/go-err"
	"os"
	"strings"
)

type Args struct {
	File *os.File
	Path string
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
	flag.Parse()

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
	for min, max := range map[int]int{6000: 7000, 8000: 9999} {
		fmt.Print(envporter.FreePort(min, max))
		fmt.Print("\n")
	}
}

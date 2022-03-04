package main

import (
	"flag"
	"fmt"
	"github.com/darkjinnee/envporter/pkg/porter"
	"github.com/darkjinnee/go-err"
	"os"
	"strings"
)

var file *os.File

func init() {
	path := flag.String(
		"f",
		".env",
		"path to .env file",
	)
	flag.Parse()
	filePath := fmt.Sprintf("%s", *path)

	dirPath, err := os.Getwd()
	goerr.Fatal(
		err,
		"[Error] envporter.init: Failed to return the root path of directory",
	)

	if strings.EqualFold(filePath, ".env") {
		filePath = strings.Join([]string{
			dirPath,
			filePath,
		}, "/")
	}

	file, err = os.OpenFile(
		filePath,
		os.O_RDWR,
		os.ModePerm,
	)
	goerr.Fatal(
		err,
		"[Error] envporter.init: Failed to open file",
	)
}

func main() {
	fmt.Print(file.Name() + "\n")
	fmt.Print(porter.CheckIpv4Tcp(80))
}

package main

import (
	"flag"
	"fmt"
	"github.com/darkjinnee/envporter/pkg/porter"
	"strconv"
	"strings"
)

type Args struct {
	Port     int
	Protocol string
	Ipv      int
}

var args Args

func init() {
	flag.IntVar(
		&args.Port,
		"port",
		80,
		"port number",
	)
	flag.StringVar(
		&args.Protocol,
		"protocol",
		"tcp",
		"data transfer protocol",
	)
	flag.IntVar(
		&args.Ipv,
		"ipv",
		4,
		"ip address version",
	)
	flag.Parse()
}

func main() {
	result := porter.CheckPort(
		args.Port,
		args.Protocol,
		args.Ipv,
	)

	port := strconv.Itoa(args.Port)
	if result {
		fmt.Print(strings.Join([]string{
			"port",
			port,
			"used",
			"\n",
		}, " "))
	} else {
		fmt.Print(strings.Join([]string{
			"port",
			port,
			"free",
			"\n",
		}, " "))
	}
}

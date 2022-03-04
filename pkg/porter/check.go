package porter

import (
	"bytes"
	"github.com/b4b4r07/go-pipe"
	"os/exec"
	"strconv"
	"strings"
)

func CheckPort(port int, protocol string, ipv int) bool {
	buffer := new(bytes.Buffer)
	_ = pipe.Command(
		buffer,
		exec.Command(
			"ss",
			"-ln",
			"--"+protocol,
			"--ipv"+strconv.Itoa(ipv),
		),
		exec.Command(
			"grep",
			strconv.Itoa(port),
		),
	)

	output := strings.TrimSpace(buffer.String())
	return !strings.EqualFold(output, "")
}

func CheckIpv4Tcp(port int) bool {
	return CheckPort(
		port,
		"tcp",
		4,
	)
}

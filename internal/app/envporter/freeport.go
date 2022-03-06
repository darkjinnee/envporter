package envporter

import (
	"github.com/darkjinnee/envporter/pkg/porter"
	"math/rand"
	"time"
)

var usedPorts []int

func NumberRange(min, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func FreePort(min, max int) int {
	var port int
	for i := 1; i < 10; i++ {
		port = NumberRange(min, max)
		for _, value := range usedPorts {
			if port == value || !porter.CheckIpv4Tcp(port) {
				break
			}
		}
	}

	usedPorts = append(
		usedPorts,
		port,
	)
	return port
}

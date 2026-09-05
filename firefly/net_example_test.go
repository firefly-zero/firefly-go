package firefly_test

import "github.com/firefly-zero/firefly-go/firefly"

func ExamplePeers_Scanner() {
	scanner := firefly.GetPeers().Scanner()
	for scanner.Scan() {
		peer := scanner.Peer()
		_ = peer
	}
}

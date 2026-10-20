package netnslab

import (
	"fmt"
	"net"
	"runtime"

	"github.com/vishvananda/netns"
)

func createNamespace(name string) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	originalNS, err := netns.Get()
	if err != nil {
		return fmt.Errorf("failed to get current network namespace: %v", err)
	}
	defer originalNS.Close()

	newNS, err := netns.NewNamed(name)
	if err != nil {
		return fmt.Errorf("failed to create new network namespace %s: %v", name, err)
	}
	defer newNS.Close()

	
	return nil
}

package helpers

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

/**
 *  Clearing environment after Application shutdown
 */
func HandleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		for _ = range c {
			fmt.Printf("\nIt was amazing time! Adjö\n\n")
			os.Exit(0)
		}
	}()
}

// vim: noai:ts=4:sw=4

package main

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func TestName(t *testing.T) {
	cmd := exec.Command("osascript", "-e", "display notification \"10.0.0.1\" with title \"iHBUT_LOGIN\" subtitle \"Success\"")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

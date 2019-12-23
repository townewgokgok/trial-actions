// +build mage

package main

import (
	"os"
	"runtime"

	"github.com/magefile/mage/sh"
)

func init() {
	os.Setenv("GO111MODULE", "on")
}

// Build program
func Build() error {
	dname := "release/trial-actions-" + runtime.GOOS + "-" + runtime.GOARCH
	if err := os.MkdirAll(dname, 0755); err != nil {
		return err
	}
	fname := dname + "/trial-actions"
	if runtime.GOOS == "windows" {
		fname += ".exe"
	}
	if err := sh.RunV("go", "build", "-o", fname, "."); err != nil {
		return err
	}
	if err := sh.RunV("tar", "cvzf", dname+".tar.gz", dname); err != nil {
		return err
	}
	return nil
}

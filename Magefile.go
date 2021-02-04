//+build mage

package main

import (
	// mage:import
	build "github.com/grafana/grafana-plugin-sdk-go/build"
	"github.com/magefile/mage/mg"
)

// Default configures the default target.
func Build64() {
	var arch64 = build.Build{}
	mg.Deps(arch64.Linux, arch64.Windows, arch64.Darwin, arch64.LinuxARM64)
}

var Default = Build64

package main

import (
	nbdkit "github.com/mrnold/nbdkit-testing"
	"k8s.io/klog"
)

func main() {
	klog.Info("testing 123")
	nbdkit.PluginInitialize("test123", nil)
}

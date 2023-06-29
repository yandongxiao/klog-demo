package main

import (
	"flag"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	flag.Set("v", "2")
	flag.Parse()

	klog.Info("klog.Info")
	// in order to log to stderr, you need to call flag.Parse() and klog.InitFlags(nil) before
	// go run main.go  -v 4 to log to stderr
	klog.V(4).Info("klog.V(4)")
}

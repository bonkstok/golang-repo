// just a sample script to deploy on openshift with s2i
// gets the container hostname and keep priting 
//deploy it like oc new-app golang~https://github.com/bonkstok/golang-repo --context-dir=openshift/hello-hostname
// oc new-app https://github.com/bonkstok/golang-repo --context-dir=openshift/hello-hostname
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var containerHostname, err = os.Hostname()
	if err != nil {
		fmt.Printf("Failed to get hostname err: %s", err)
		os.Exit(1)
	}
	for true {
		fmt.Println("Hello from container", containerHostname)
		time.Sleep(2 * time.Second)
	}
}

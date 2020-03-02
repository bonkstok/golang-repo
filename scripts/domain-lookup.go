package main

import (
	"fmt"
	"net"
	"strings"
	"os"
	"time"
)

const LOOKUP_ROUNDS = 10
const FAILURE_ALLOWED = 5

func main () {
	//var s []int
	domains := []string{"google.nl", "google.com", "doesnotexit404.failed","johnnyvanveen.com","kpn.com","atos.net"}
    failedMap := make(map[string]int) //make that has a domain as k and num of failed lookups as value

    domains = append(domains, "append.io")
    if len(os.Args) > 1 {
    	for _ , k := range os.Args[1:]{
	    	domains = append(domains, k)
    	}
    }

    //printSliceStr(str, failedMap) //call the func
    for  i := 0; i < LOOKUP_ROUNDS ; i++ {
    	printSliceStr(domains, failedMap)
    	time.Sleep(2 * time.Second)
    }
    for domain, failCount := range failedMap {
    	fmt.Printf("Failure count for domain %s is %d\n",domain, failCount)
    }
    //failedMap = printSliceStr(str, failedMap)
}


func printSliceStr(mySlice []string, failed map[string]int ){
	//fmt.Println(mySlice)
	for _ , k := range mySlice {
		ips, err := net.LookupIP(k) //returns slice
		if err != nil {
			if failed[k] >= FAILURE_ALLOWED {
				fmt.Println(k,"Failed too many times!")
				// need to find a good way to not lookup anymore..
			} else {
				failed[k]++
				fmt.Println("Failed lookup for domain",k)				
			}		
		}
		//fmt.Println(ips)
		for _ , ip := range ips {
			if failed[k] < FAILURE_ALLOWED {
				fmt.Printf("ip for %s:%s -- address is of type: %s\n", k,ip, ip4Or6(ip))
			}
		}
	}
}


func ip4Or6(ip net.IP) string { //very simplicit check to see if address is ipv4 or 6
	splited := strings.Split(ip.String(),".")
	if len(splited) == 4 {
		//fmt.Println("ipv4")
		return "ipv4"
	} else {
		if len(strings.Split(ip.String(),":")) > 1 {
			//fmt.Println("Is ipv6")
			return "ipv6"
		}
	}
	return "unknown"
}

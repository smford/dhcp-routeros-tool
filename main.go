package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/go-routeros/routeros"
)

var (
	command    = flag.String("command", "/ip/dhcp-server/lease/print", "RouterOS command")
	address    = flag.String("address", "172.28.0.1:8729", "RouterOS address and port")
	username   = flag.String("username", "test", "User name")
	password   = flag.String("password", "somepassword", "Password")
	async      = flag.Bool("async", false, "Use async code")
	useTLS     = flag.Bool("tls", true, "Use TLS")
	properties = flag.String("properties", "comment,address,mac-address,client-id,address-lists,server,dhcp-option,status,last-seen,host-name,radius,dynamic,blocked,disabled", "Properties")
)

func dial() (*routeros.Client, error) {
	if *useTLS {
		return routeros.DialTLS(*address, *username, *password, nil)
	}
	return routeros.Dial(*address, *username, *password)
}

func main() {
	flag.Parse()

	c, err := dial()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	if *async {
		c.Async()
	}

	r, err := c.RunArgs(strings.Split(*command, " "))
	if err != nil {
		log.Fatal(err)
	}

	/*
		log.Print(r)
		fmt.Printf("\n======%T\n========\n\n", r)
	*/

	/*
		for k, v := range r.Re {
			fmt.Println("---")
			fmt.Printf("k type and value: %T %d\n", k, k)
			fmt.Printf("v type: %T\n", v)
			fmt.Printf("v.Word: %T\n", v.Word)
			fmt.Printf("v.Word values:: %+v\n", v)
			fmt.Printf("v.Word contents:: %s\n", v)
			fmt.Printf("v.Tag %T\n", v.Tag)
			fmt.Printf("v.Tag contents:: %+v\n", v.Tag)
		}
	*/

	/*
		for _, re := range r.Re {
			for _, p := range strings.Split(*properties, ",") {
				fmt.Print(re.Map[p], "\t")
			}
			fmt.Print("\n")
		}
		fmt.Print("\n")
	*/

	const padding = 1
	w := tabwriter.NewWriter(os.Stdout, 0, 2, padding, ' ', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t\n", "Comment", "Address", "Mac-address", "Client-id", "Address-lists", "Server", "Dhcp-option", "Status", "Last-seen", "Host-name", "Radius", "Dynamic", "Blocked", "Disabled")

	//var printthis string
	for _, re := range r.Re {
		var mytempstring string
		for _, p := range strings.Split(*properties, ",") {
			//fmt.Print(re.Map[p], "\t")
			mytempstring = mytempstring + re.Map[p] + "\t"
			//fmt.Print(printthis + ",")
		}
		fmt.Fprintf(w, "%s\n", mytempstring)
	}

	//fmt.Fprintln(w, printthis)

	w.Flush()
}

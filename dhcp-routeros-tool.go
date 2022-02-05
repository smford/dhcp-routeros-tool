package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/go-routeros/routeros"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const applicationName string = "dhcp-routeros-tool"
const applicationVersion string = "v0.2"
const command string = "/ip/dhcp-server/lease/print"

var (
	columnstodisplay string
)

func init() {
	flag.String("address", "", "The IP address or hostname of mikrotik router")
	flag.Bool("async", true, "Execute commands asyncronously")
	flag.String("config", "config.yaml", "Configuration file: /path/to/file.yaml, default = ./config.yaml")
	flag.String("default", "comment,address,mac-address,client-id,address-lists,server,dhcp-option,status,last-seen,host-name,radius,dynamic,blocked,disabled", "Columns to display by default")
	flag.Bool("displayconfig", false, "Display configuration")
	flag.Bool("help", false, "Display help")
	flag.Int("padding", 2, "Column padding size")
	flag.String("password", "", "Password")
	flag.Bool("simple", false, "Display simple format")
	flag.String("simpledisplay", "address,mac-address,client-id,server,status,last-seen,host-name,disabled", "Columns to display when --simple argument passed")
	flag.String("username", "", "Username")
	flag.Bool("usetls", true, "Use TLS to connect to router")
	flag.Bool("version", false, "Display version information")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	checkErr(err)

	if viper.GetBool("help") {
		displayHelp()
		os.Exit(0)
	}

	if viper.GetBool("version") {
		fmt.Println(applicationName + " " + applicationVersion)
		os.Exit(0)
	}

	configdir, configfile := filepath.Split(viper.GetString("config"))

	// set default configuration directory to current directory
	if configdir == "" {
		configdir = "."
	}

	viper.SetConfigType("yaml")
	viper.AddConfigPath(configdir)

	config := strings.TrimSuffix(configfile, ".yaml")
	config = strings.TrimSuffix(config, ".yml")

	viper.SetConfigName(config)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		} else {
			log.Fatal("Config file was found but another error was discovered: ", err)
		}
	}

	if viper.GetBool("displayconfig") {
		displayConfig()
		os.Exit(0)
	}

	if viper.GetBool("simple") {
		columnstodisplay = viper.GetString("simpledisplay")
	} else {
		columnstodisplay = viper.GetString("defaultdisplay")
	}
}

func main() {

	client, err := dial()
	checkErr(err)

	defer client.Close()

	if viper.GetBool("async") {
		client.Async()
	}

	response, err := client.RunArgs(strings.Split(command, " "))
	checkErr(err)

	w := tabwriter.NewWriter(os.Stdout, 0, 2, viper.GetInt("padding"), ' ', 0)

	var columnheadings string
	for _, heading := range strings.Split(columnstodisplay, ",") {
		columnheadings = columnheadings + strings.ToUpper(heading) + "\t"
	}
	columnheadings = columnheadings + "\n"
	fmt.Fprint(w, columnheadings)

	for _, myreply := range response.Re {
		var mylease string
		for _, p := range strings.Split(columnstodisplay, ",") {
			mylease = mylease + myreply.Map[p] + "\t"
		}
		fmt.Fprintf(w, "%s\n", mylease)
	}

	w.Flush()
}

// displays help information
func displayHelp() {
	message := `
      --config [file]       Configuration file: /path/to/file.yaml (default: "./config.yaml")
      --displayconfig       Display configuration
      --help                Display help
      --simple              Display simple columns
      --version             Display version`
	fmt.Println(applicationName + " " + applicationVersion)
	fmt.Println(message)
}

// display configuration
func displayConfig() {
	allmysettings := viper.AllSettings()
	var keys []string
	for k := range allmysettings {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("CONFIG:", k, ":", allmysettings[k])
	}
}

// checks errors
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dial() (*routeros.Client, error) {
	if viper.GetBool("usetls") {
		return routeros.DialTLS(viper.GetString("address"), viper.GetString("username"), viper.GetString("password"), nil)
	}
	return routeros.Dial(viper.GetString("address"), viper.GetString("username"), viper.GetString("password"))
}

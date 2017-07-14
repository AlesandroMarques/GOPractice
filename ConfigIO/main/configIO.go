package main

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"




)

// Config ...
type Config struct {
	server string
	user   string
	pw     string
}

// Reads info from config file
func ReadConfig() Config {
}

func main() {
	config := ReadConfig()
	fmt.Printf("%s: %s: %s\n", config.server, config.user, config.pw)

}
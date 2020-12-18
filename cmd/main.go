package cmd

import (
	"flag"
	"fmt"
	"time"
)

var c Config
var config string
// Key binding of fishing spell
var Key string
// Debug for better logging
var Debug bool

func init() {
	flag.StringVar(&config, "config", "", "Path to your config.yml")
	flag.BoolVar(&Debug, "debug", false, "Debug option")
	flag.Parse()
	c.GetConfig(config)
	Key = c.Key
	Debug = Debug || c.Debug
}

// Main for fishcraft app
func Main() {
	fmt.Println("🎣 Fishcraft 🎣")
	fmt.Println("Focus on WoW Window")
	fmt.Println("Starting in")
	for i := 3; 0 <= i; i-- {
		fmt.Printf("%v seconds.\n",i)
		time.Sleep(1 * time.Second)
	}

	finder := CreateBobberFinder()
	for {
		Fish(finder)
	}
}

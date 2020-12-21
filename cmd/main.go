package cmd

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var c Config
var config string
var logPath string
// BaitKey where bait is binded
var BaitKey string
// Key binding of fishing spell
var Key string
// Debug for better logging
var Debug bool
// Threshold to detect bobber movement
var Threshold int
// DebugLogger for fishcraft
var DebugLogger *log.Logger
// ErrorLogger for fishcraft
var ErrorLogger *log.Logger
// InfoLogger for fishcraft
var InfoLogger *log.Logger

func init() {
	flag.StringVar(&config, "config", "", "Path to your config.yml")
	flag.BoolVar(&Debug, "debug", false, "Debug option")
	flag.Parse()
	c.GetConfig(config)
	BaitKey = c.Bait
	Key = c.Key
	Debug = Debug || c.Debug
	Threshold = c.Threshold
}

func initLog() {
	c.GetConfig(config)
	if c.LogPath == "" {
		logPath = "./fishcraft.log"
	} else {
		logPath = c.LogPath
	}
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	ErrorLogger = log.New(file, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger.SetOutput(mw)
	InfoLogger = log.New(file, "[INFO] ", log.Ldate|log.Ltime)
	InfoLogger.SetOutput(mw)
	if Debug {
		DebugLogger = log.New(file, "[DEBUG] ", log.Ldate|log.Ltime)
		DebugLogger.SetOutput(mw)
	}
}

// Main for fishcraft app
func Main() {
	initLog()
	fmt.Println("🎣 Fishcraft 🎣")
	fmt.Println("Focus on WoW Window")
	fmt.Println("Starting in")
	for i := 3; 0 <= i; i-- {
		fmt.Printf("%v seconds.\n",i)
		time.Sleep(1 * time.Second)
	}

	finder := CreateBobberFinder()
	ApplyBait()
	for {
		Fish(finder)
	}
}

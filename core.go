package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/syslog"
	"os"
)

const (
	configPath = "/etc/nsstls.json"
	logPrefix  = "NSSTLS "
)

var (
	initSuccessful bool
	initError      error
	isDebugMode    bool
	logger         *log.Logger
	logFile        *os.File
	configuration  config
)

type config struct {
	URL    string `json:"url"`
	Token  string `json:"token"`
	Logger string `json:"logger"`
	Debug  bool   `json:"debug"`

	TLS struct {
		Cert       string `json:"cert"`
		Key        string `json:"key"`
		Root       string `json:"root"`
		VerifyMode string `json:"verify-mode"`
	} `json:"tls"`
}

// LoadConfig ingests and applies the configuration at p.
func LoadConfig(p string) error {
	logger = log.New(os.Stdout, logPrefix, log.Ltime)
	f, err := os.Open(p)
	if err != nil {
		return err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)

	var c config
	err = decoder.Decode(&c)
	if err != nil {
		return err
	}
	configuration = c
	return applyConfig()
}

func applyConfig() error {
	// setup logging
	switch configuration.Logger {
	case "syslog":
		initSyslog()
	default:
		logger = log.New(os.Stdout, logPrefix, log.Ltime)
	}

	isDebugMode = configuration.Debug
	baseURL = configuration.URL
	return tlsInit()
}

func initSyslog() {
	var err error
	logger, err = syslog.NewLogger(syslog.LOG_INFO, log.Ltime)
	if err != nil {
		logger = log.New(os.Stderr, logPrefix, log.Ltime)
		logger.Printf("syslog.Open() err: %v", err)
	}
}

func info(module string, data ...interface{}) {
	s := fmt.Sprintf("[%s] %s", module, fmt.Sprint(data...))
	logger.Println(s)
}

func fatal(module string, err error) {
	s := fmt.Sprintf("[%s] Fatal: %v", module, err)
	logger.Println(s)
}

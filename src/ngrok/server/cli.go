package server

import (
	"flag"
)

type Options struct {
	httpAddr   string
	tunnelAddr string
	domain     string
	logto      string
	loglevel   string
}

func parseArgs() *Options {
	httpAddr := flag.String("httpAddr", ":80", "Public address for HTTP connections, empty string to disable")
	tunnelAddr := flag.String("tunnelAddr", ":4443", "Public address listening for ngrok client")
	domain := flag.String("domain", "ngrok.com", "Domain where the tunnels are hosted")
	logto := flag.String("log", "stdout", "Write log messages to this file. 'stdout' and 'none' have special meanings")
	loglevel := flag.String("log-level", "DEBUG", "The level of messages to log. One of: DEBUG, INFO, WARNING, ERROR")
	flag.Parse()

	return &Options{
		httpAddr:   *httpAddr,
		tunnelAddr: *tunnelAddr,
		domain:     *domain,
		logto:      *logto,
		loglevel:   *loglevel,
	}
}

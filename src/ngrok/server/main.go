package server

import (
	log "ngrok/log"
	"ngrok/conn"
)

// GLOBALS
var (
	opts      *Options
)

func Main() {
	opts = parseArgs()
	// init logging
	log.LogTo(opts.logto, opts.loglevel)
	conn.Listen(":8080");
}

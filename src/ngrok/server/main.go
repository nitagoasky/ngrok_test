package server

import (
	"fmt"
	log "ngrok/log"
)

// GLOBALS
var (
	opts      *Options
)

func Main() {
	opts = parseArgs()

	// init logging
	log.LogTo(opts.logto, opts.loglevel)
	
	fmt.Println("GAOTIAN")
}

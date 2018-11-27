package main

import (
	"flag"
	"log"

	"git.linuxrocker.com/mattburchett/go_tab-magic/pkg/config"
	"git.linuxrocker.com/mattburchett/go_tab-magic/pkg/resolver"
)

func main() {
	var c string
	var debug bool

	flag.StringVar(&c, "config", "", "Configuration to load")
	flag.BoolVar(&debug, "debug", false, "Enables Debugging Mode")
	flag.Parse()

	// Stop the app if they're missing required flags.
	if c == "" {
		log.Fatal("You need to specify a configuration file.")
	}

	cfg, err := config.GetConfig(c, debug)
	if err != nil {
		log.Fatal(err)
	}

	resolver.PerformZoneTransfer(cfg)
}

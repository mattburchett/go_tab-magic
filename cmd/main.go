package main

import (
	"flag"
	"log"

	"git.linuxrocker.com/mattburchett/go_tab-magic/pkg/config"
	"git.linuxrocker.com/mattburchett/go_tab-magic/pkg/resolver"
	"git.linuxrocker.com/mattburchett/go_tab-magic/pkg/shell"
)

func main() {
	var c = flag.String("config", "", "Configuration to load")
	var user = flag.String("user", "", "user for aliases")
	var debug = flag.Bool("debug", false, "Enables Debugging Mode")
	flag.Parse()

	// Stop the app if they're missing required flags.
	if *c == "" {
		log.Fatal("You need to specify a configuration file.")
	}

	if *user == "" {
		log.Fatal("You must specify a username.")
	}

	cfg, err := config.GetConfig(*c, *debug)
	if err != nil {
		log.Fatal(err)
	}

	data := resolver.PerformZoneTransfer(cfg, *debug)
	shell.CreateShellAliases(data, *user, cfg, *debug)
}

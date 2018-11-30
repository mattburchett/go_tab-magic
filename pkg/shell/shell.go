package shell

import (
	"fmt"
	"strings"

	"git.linuxrocker.com/mattburchett/go_tab-magic/pkg/config"
)

func CreateShellAliases(data []string, username string, config config.Config) {
	for _, i := range data {
		splitStrings := strings.Split(i, " ")
		hostname := splitStrings[0]
		txt := splitStrings[2]

		jump := config.JumpHost
		stringSplit := config.SplitString

		remoteUser := username
		sudo := "sudo su -"
		rac := "ssh"
		racOpts := "-AXt -l"
		hop := "ssh -AXt"
		prerac := ""

		host := strings.TrimRight(hostname, stringSplit)
		fqdn := hostname

		greentext := "tput -T xterm setaf 2; "
		// redtext := "tput -T xterm setaf 1; "
		resettext := "tput -T xterm sgr0; "
		message := fmt.Sprintf("%vecho \"Authenticating as: %v\";%v", greentext, remoteUser, resettext)

		if txt == "" {
			fmt.Printf("alias %v=\\'%v%v%v %v@%v \"%v %v %v %v %v\"'\n", host, message, prerac, hop, username, jump, rac, racOpts, remoteUser, fqdn, sudo)
		} else {
			fmt.Printf("alias %v=\\'%v%v%v %v@%v \"%v %v %v %v %v\"'\n", host, message, prerac, hop, username, jump, rac, racOpts, remoteUser, fqdn, sudo)
		}
	}
}

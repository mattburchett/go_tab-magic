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

		useJump := config.UseJump

		jump := config.JumpHost
		stringSplit := config.SplitString

		remoteUser := username
		sudo := "sudo su -"
		rac := "ssh"
		racOpts := "-AXt -l"
		hop := "ssh -AXt"
		prerac := ""
		windowsGeometry := config.WindowsGeometry

		host := strings.TrimSuffix(hostname, stringSplit)
		fqdn := hostname

		greentext := "tput -T xterm setaf 2; "
		redtext := "tput -T xterm setaf 1; "
		resettext := "tput -T xterm sgr0; "
		message := fmt.Sprintf("%vecho \"Authenticating as: %v\";%v", greentext, remoteUser, resettext)

		// TXT Record Parsing

		txtSplit := strings.Split(txt, ";")

		for _, i := range txtSplit {
			if strings.Contains(i, "SSH_PORT") {
				port := strings.TrimLeft(i, "SSH_PORT=")
				racOpts = fmt.Sprintf("-AXt -p %v -l", port)
			} else if strings.Contains(i, "OS_FAMILY") {
				osFamily := strings.Split(i, "=")
				if osFamily[1] == "ESXi" {
					sudo = ""
					remoteUser = "root"
				} else if osFamily[1] == "Ubiquiti" {
					sudo = ""
				} else if osFamily[1] == "Windows" {
					prerac = fmt.Sprintf("%vecho \"Password: \"; %v", redtext, resettext)
					hop = "ssh -XCAT"
					rac = "rdesktop"
					windowsDomain := ""
					racOpts = fmt.Sprintf("-r clipboard:CLIPBOARD -a 16 -k en-us -g %v -p - %v -u", windowsGeometry, windowsDomain)
					sudo = ""
				}
			} else if strings.Contains(i, "REMOTE_USER") {
				user := strings.TrimLeft(i, "REMOTE_USER=")
				remoteUser = user
			}
		}

		if useJump {
			fmt.Printf("alias %v=\\'%v%v%v %v@%v \"%v %v %v %v %v\"'\n", host, message, prerac, hop, username, jump, rac, racOpts, remoteUser, fqdn, sudo)
		} else {
			fmt.Printf("alias %v=\\'%v %v %v %v %v %v\n", host, message, rac, racOpts, remoteUser, fqdn, sudo)
		}
	}
}

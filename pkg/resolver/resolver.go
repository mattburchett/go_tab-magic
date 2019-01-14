package resolver

import (
	"errors"
	"fmt"
	"log"
	"rfi-sower/pkg/utils"
	"strings"

	"git.linuxrocker.com/mattburchett/go_tab-magic/pkg/config"
	"github.com/miekg/dns"
)

// LookupName returns IPv4 address from A record or error.
func lookupName(fqdn, serverAddr string, debug bool) (string, error) {
	if debug {
		defer utils.LogElapsedTime("Lookup Name")
	}
	m := &dns.Msg{}
	m.SetQuestion(dns.Fqdn(fqdn), dns.TypeA)
	in, err := dns.Exchange(m, serverAddr)
	if err != nil {
		return "", err
	}
	if len(in.Answer) < 1 {
		return "", errors.New("no Answer")
	}
	if a, ok := in.Answer[0].(*dns.A); ok {
		ip := a.A.String()
		return ip, nil
	}
	return "", errors.New("no A record returned")
}

// PerformZoneTransfer performs zone transfers and gathers a list from config.Domains
func PerformZoneTransfer(config config.Config, debug bool) []string {
	if debug {
		defer utils.LogElapsedTime("Perform Zone Transfer")
	}
	data := make([]string, 0)

	// Do the transfer
	for _, i := range config.Domains {
		server := fmt.Sprintf("%s:%d", config.Resolver, config.ResolverPort)
		tr := dns.Transfer{}
		m := &dns.Msg{}
		m.SetAxfr(dns.Fqdn(i))
		in, err := tr.In(m, server)
		if err != nil {
			log.Fatal(err)
		}
		for ex := range in {
			for _, a := range ex.RR {
				var ip, hostname, txt string
				switch v := a.(type) {
				case *dns.TXT:
					hostname = v.Hdr.Name
					cip, err := lookupName(strings.TrimRight(v.Hdr.Name, "."), server, debug)
					if err != nil || cip == "" {
						continue
					}
					ip = cip
				case *dns.A:
					ip = v.A.String()
					hostname = v.Hdr.Name
				case *dns.CNAME:
					cip, err := lookupName(v.Target, server, debug)
					if err != nil || cip == "" {
						continue
					}
					ip = cip
					hostname = v.Hdr.Name
				default:
					continue
				}
				data = append(data, fmt.Sprintf("%v %v %v", strings.TrimRight(hostname, "."), ip, txt))
			}
		}
	}

	return data
}

package main

import (
	"flag"
	"log"

	"zdxsv/pkg/config"

	"github.com/miekg/dns"
)

var (
	dnasip    = flag.String("dnasip", "", "Public IP Addr of DNAS server")
	loginsvip = flag.String("loginsvip", "", "Public Addr of Login server")
)

func makeDNSHandler(record string) func(dns.ResponseWriter, *dns.Msg) {
	return func(w dns.ResponseWriter, r *dns.Msg) {
		m := new(dns.Msg)
		m.SetReply(r)
		rr, err := dns.NewRR(record)
		if err != nil {
			log.Println(err)
		}
		m.Answer = append(m.Answer, rr)
		err = w.WriteMsg(m)
		if err != nil {
			log.Println(err)
		}
	}
}

func mainDNS() {
	dnassvIP := *dnasip
	if *dnasip == "" {
		dnassvIP = config.Conf.DNAS.PublicAddr
	}

	loginPublicIP := *loginsvip
	if *loginsvip == "" {
		loginPublicIP = config.Conf.Login.PublicAddr
	}

	log.Println("DNAS server ", dnassvIP)
	log.Println("Login server ", loginPublicIP)
	dns.HandleFunc("kddi-mmbb.jp", makeDNSHandler("www01.kddi-mmbb.jp. 3600 IN A "+loginPublicIP))
	dns.HandleFunc("playstation.org", makeDNSHandler("gate1.jp.dnas.playstation.org. 3600 IN A "+dnassvIP))

	server := &dns.Server{Addr: ":53", Net: "udp"}
	err := server.ListenAndServe()
	log.Fatal(err)
}

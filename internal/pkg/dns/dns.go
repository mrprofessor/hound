package dns

import (
	"fmt"
	"github.com/charmbracelet/glamour"
	"net"
)

func LookUpDnsRecords(url string) {

	var dnsInfo string
	// Write functions to lookup all the possible records.

	dnsInfo = "## A & AAAA record \n" +
		lookUpIpRecords(url) +

		"\n ## CNAME record \n" +
		lookUpCnameRecords(url) +

		"\n ## MX records \n" +
		lookUpMxRecords(url) +

		"\n ## TXT records \n" +
		lookUpTxtRecords(url) +

		"\n ## Name servers \n" +
		lookUpNameServers(url) +

		// PTR records are only for IPs(opposite of A records)
		"\n ## PTR records \n" +
		lookUpPtrRecords(url) +

		"\n ## SRV records \n" +
		lookUpSrvRecords(url)

	Render(dnsInfo)
}

func Render(str string) {
	out, _ := glamour.Render(str, "auto")
	fmt.Print(out)
}

func lookUpIpRecords(url string) string {
	ipRecords, err := net.LookupIP(url)
	if err != nil {
		return "- " + err.Error()
	}
	str := ""
	for _, ip := range ipRecords {
		str += fmt.Sprintln("- ", ip)
	}
	return str
}

func lookUpCnameRecords(url string) string {
	CNAME, err := net.LookupCNAME(url)
	if err != nil {
		return "- " + err.Error()
	}
	return fmt.Sprintln("- ", CNAME)
}

func lookUpMxRecords(url string) string {
	mxRecords, err := net.LookupMX(url)
	if err != nil {
		return "- " + err.Error()
	}
	str := ""
	for _, mx := range mxRecords {
		str += fmt.Sprintln("- ", mx.Host, mx.Pref)
	}
	return str
}

func lookUpTxtRecords(url string) string {
	txtRecords, err := net.LookupTXT(url)
	if err != nil {
		return "- " + err.Error()
	}
	str := ""
	for _, txt := range txtRecords {
		str += fmt.Sprintf("- %s\n", txt)
	}
	return str
}

func lookUpNameServers(url string) string {
	nameServers, err := net.LookupNS(url)
	if err != nil {
		return "- " + err.Error()
	}
	str := ""
	for _, ns := range nameServers {
		str += fmt.Sprintf("- %s \n", ns)
	}
	return str
}

func lookUpPtrRecords(url string) string {
	ptrServers, err := net.LookupAddr(url)
	if err != nil {
		return "- " + err.Error()
	}
	str := ""
	for _, ptr := range ptrServers {
		str += fmt.Sprintf("- %s\n", ptr)
	}
	return str
}

func lookUpSrvRecords(url string) string {
	_, xmppSrvRecords, err := net.LookupSRV("xmpp-server", "tcp", url)
	if err != nil {
		return "- " + err.Error()
	}
	str := ""
	for _, srv := range xmppSrvRecords {
		str += fmt.Sprintf("- %v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
	return str
}

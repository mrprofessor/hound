package whois

import (
	"fmt"

	"github.com/charmbracelet/glamour"
	"github.com/likexian/whois"
	"github.com/likexian/whois-parser"
)

func RunWhoIs(url string) {
	whoIsParsedData, err := WhoIsData(url)
	if err != nil {
		fmt.Println("Error Getting and parsing Whois data, Error: ", err)
		return
	}
	Render(*whoIsParsedData)
}

// WhoIsData accepts the URL and returns a whoisparser.WhoisInfo struct.
func WhoIsData(url string) (*whoisparser.WhoisInfo, error) {
	whoIsRaw, err := whois.Whois(url)
	if err != nil {
		return nil, err
	}

	parsedData, err := whoisparser.Parse(whoIsRaw)
	if err != nil {
		return nil, err
	}

	return &parsedData, nil
}

func Render(resp whoisparser.WhoisInfo) {

	// Create a map of similar types in order to iterate easily.
	whoIsAddrMap := map[string]*whoisparser.Contact{
		"Registrar":      resp.Registrar,
		"Registrant":     resp.Registrant,
		"Administrative": resp.Administrative,
		"Technical":      resp.Technical,
		"Billing":        resp.Billing,
	}

	domainInfo := "## Domain Info \n" +
		fmt.Sprintf("- %-20s: %s\n", "Name", resp.Domain.Name) +
		fmt.Sprintf("- %-20s: %s\n", "Whois Server", resp.Domain.WhoisServer) +
		fmt.Sprintf("- %-20s: %s\n", "Registered On", resp.Domain.CreatedDate) +
		fmt.Sprintf("- %-20s: %s\n", "Expires On", resp.Domain.ExpirationDate) +
		fmt.Sprintf("- %-20s: %s\n", "Domain Status", resp.Domain.Status) +
		fmt.Sprintf("- %-20s: %s\n", "Name Servers", resp.Domain.NameServers)

	// Print WhoIs contact info
	for k, v := range whoIsAddrMap {
		if v != nil {

			str := fmt.Sprintf("## %s \n", k) +
				fmt.Sprintf("- %-20s: %s\n", "Name", v.Name) +
				fmt.Sprintf("- %-20s: %s\n", "Organization", v.Organization) +
				fmt.Sprintf("- %-20s: %s\n", "Street", v.Street) +
				fmt.Sprintf("- %-20s: %s\n", "City", v.City) +
				fmt.Sprintf("- %-20s: %s\n", "State", v.Province) +
				fmt.Sprintf("- %-20s: %s\n", "Country", v.Country) +
				fmt.Sprintf("- %-20s: %s\n", "Phone", v.Phone) +
				fmt.Sprintf("- %-20s: %s\n", "Phone Extension", v.PhoneExt) +
				fmt.Sprintf("- %-20s: %s\n", "Email", v.Email) +
				fmt.Sprintf("- %-20s: %s\n", "Referral URL", v.ReferralURL)

			domainInfo += str
		}
	}

	customRenderer, _ := glamour.NewTermRenderer(
		// detect background color and pick either the default dark or light theme
		glamour.WithAutoStyle(),
		// wrap output at specific width (default is 80)
		glamour.WithWordWrap(100),
	)

	out, _ := customRenderer.Render(domainInfo)
	fmt.Print(out)
}

package ssl

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"net"

	"github.com/charmbracelet/glamour"
)

// TODO
// Error handling needs to improve
func RunSSL(url string) {
	cert, err := SSLCertData(url)
	if err != nil {
		fmt.Errorf("Error getting SSL data, Error: ", err)
		return
	}
	Render(cert)
}

// NOTE
// crypto/tls doesn't return certificates when it has some errors. For example
// with CertificateInvalidError, it returns the certificates, but not for other
// types of errors.
// e.g certificate signed by unknown authority
// test with (self-signed.badssl.com) from badssl.com
func SSLCertData(url string) (*x509.Certificate, error) {
	// By default HTTPS connections use port 443
	conn, err := tls.Dial("tcp", url+":443", nil)
	if err != nil {

		var certErr x509.CertificateInvalidError
		if errors.As(err, &certErr) && certErr.Reason == x509.Expired {
			fmt.Errorf("Certificate has expired or is not yet valid")
			// Certificate exists even though it's expired.
			return certErr.Cert, nil
		} else if errors.As(err, new(*net.DNSError)) {
			fmt.Errorf("Host not Found")
		} else {
			fmt.Errorf("%+v\n", err)
		}
	}

	// Malformed certificates can be present even after an error
	// Only return the first certificate.
	if conn != nil {
		return conn.ConnectionState().PeerCertificates[0], nil
	} else {
		return nil, err
	}
}

func Render(cert *x509.Certificate) {

	sslInfo := "## Certificate Details \n" +
		fmt.Sprintf("- %-20s: %v\n", "Issuer", cert.Issuer.String()) +
		fmt.Sprintf("- %-20s: %v\n", "Subject", cert.Subject.String()) +
		fmt.Sprintf("- %-20s: %v\n", "Version", cert.Version) +
		fmt.Sprintf("- %-20s: %v\n", "Serial Number", cert.SerialNumber) +
		fmt.Sprintf("- %-20s: %v\n", "Start Date", cert.NotBefore) +
		fmt.Sprintf("- %-20s: %v\n", "Expiration Date", cert.NotAfter) +
		fmt.Sprintf("- %-20s: %v\n", "DNS Names", cert.DNSNames) +
		fmt.Sprintf("- %-20s: %v\n", "Email Addresses", cert.EmailAddresses) +
		fmt.Sprintf("- %-20s: %v\n", "IP Addresses", cert.IPAddresses) +
		fmt.Sprintf("- %-20s: %v\n", "URIs", cert.URIs) +
		fmt.Sprintf("- %-20s: %v\n", "Permitted Domains", cert.PermittedDNSDomains)

	out, _ := glamour.Render(sslInfo, "auto")
	fmt.Print(out)
}

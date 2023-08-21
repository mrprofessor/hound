package cmd

import (
	"github.com/mrprofessor/hound/internal/pkg/dns"
	"github.com/spf13/cobra"
)

var DnsCmd = &cobra.Command{
	Use:     "dns",
	Aliases: []string{"dns"},
	Short:   "Domain Name Server details for a URL",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dns.LookUpDnsRecords(args[0])
	},
}

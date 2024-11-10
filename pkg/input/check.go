/*

=======================
Scilla - Information Gathering Tool
=======================

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see http://www.gnu.org/licenses/.

	@Repository:  https://github.com/edoardottt/scilla

	@Author:      edoardottt, https://edoardottt.com

	@License: https://github.com/edoardottt/scilla/blob/main/LICENSE

*/

package input

import (
	"flag"
	"fmt"
	"os"

	ignoreUtils "github.com/edoardottt/scilla/internal/ignore"
	transportUtils "github.com/edoardottt/scilla/internal/transport"
	urlUtils "github.com/edoardottt/scilla/internal/url"
)

// ReportSubcommandCheckFlags performs all the necessary checks on the flags
// for the report subcommand.
func ReportSubcommandCheckFlags(reportCommand flag.FlagSet, reportTargetPtr *string,
	reportPortsPtr *string, reportCommonPtr *bool, reportVirusTotalPtr *bool, reportSubdomainDBPtr *bool,
	startPort int, endPort int, reportIgnoreDirPtr *string,
	reportIgnoreSubPtr *string, reportTimeoutPort *int,
	reportOutputJSON, reportOutputHTML, reportOutputTXT, reportUserAgentPtr *string,
	reportRandomUserAgentPtr *bool) (int, int, []int, bool, []string, []string) {
	// Required Flags
	if *reportTargetPtr == "" {
		reportCommand.PrintDefaults()
		os.Exit(1)
	}

	// Verify good inputs
	if !urlUtils.IsURL(*reportTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}

	// output files all different
	if *reportOutputJSON != "" {
		if *reportOutputJSON == *reportOutputTXT || *reportOutputJSON == *reportOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *reportOutputHTML != "" {
		if *reportOutputHTML == *reportOutputTXT || *reportOutputJSON == *reportOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *reportOutputTXT != "" {
		if *reportOutputJSON == *reportOutputTXT || *reportOutputTXT == *reportOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	// common and p not together
	if *reportPortsPtr != "" && *reportCommonPtr {
		fmt.Println("You can't specify a port range and common option together.")
		os.Exit(1)
	}

	if *reportVirusTotalPtr && !*reportSubdomainDBPtr {
		fmt.Println("You can't specify VirusTotal and not the Open Database option.")
		fmt.Println("If you want to use VirusTotal Api, set also -db option.")
		os.Exit(1)
	}

	var (
		portsArray    []int
		portArrayBool bool
		err           error
	)

	startPort, endPort, portsArray, portArrayBool = transportUtils.PortsInputHelper(reportPortsPtr,
		startPort, endPort, portsArray, portArrayBool)

	var (
		reportIgnoreDir []string
		reportIgnoreSub []string
	)

	if *reportIgnoreDirPtr != "" {
		toBeIgnored := *reportIgnoreDirPtr

		reportIgnoreDir, err = ignoreUtils.CheckIgnore(toBeIgnored)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if *reportIgnoreSubPtr != "" {
		toBeIgnored := *reportIgnoreSubPtr

		reportIgnoreSub, err = ignoreUtils.CheckIgnore(toBeIgnored)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if *reportTimeoutPort < 1 || *reportTimeoutPort > 100 {
		fmt.Println("Port Scan timeout must be an integer between 1 and 100.")
		os.Exit(1)
	}

	if *reportUserAgentPtr != DefaultUserAgent && *reportRandomUserAgentPtr {
		fmt.Println("You cannot specify both ua and rua.")
		os.Exit(1)
	}

	return startPort, endPort, portsArray, portArrayBool, reportIgnoreDir, reportIgnoreSub
}

// DNSSubcommandCheckFlags performs all the necessary checks on the flags
// for the dns subcommand.
func DNSSubcommandCheckFlags(dnsCommand flag.FlagSet, dnsTargetPtr, dnsOutputJSON,
	dnsOutputHTML, dnsOutputTXT *string) {
	// Required Flags
	if *dnsTargetPtr == "" {
		dnsCommand.PrintDefaults()
		os.Exit(1)
	}

	// output files all different
	if *dnsOutputJSON != "" {
		if *dnsOutputJSON == *dnsOutputTXT || *dnsOutputJSON == *dnsOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *dnsOutputHTML != "" {
		if *dnsOutputHTML == *dnsOutputTXT || *dnsOutputJSON == *dnsOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *dnsOutputTXT != "" {
		if *dnsOutputJSON == *dnsOutputTXT || *dnsOutputTXT == *dnsOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	// Verify good inputs
	if !urlUtils.IsURL(*dnsTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}
}

// SubdomainSubcommandCheckFlags performs all the necessary checks on the flags
// for the subdomain subcommand.
func SubdomainSubcommandCheckFlags(subdomainCommand flag.FlagSet, subdomainTargetPtr *string,
	subdomainNoCheckPtr *bool, subdomainDBPtr *bool, subdomainWordlistPtr *string,
	subdomainIgnorePtr *string, subdomainCrawlerPtr *bool, subdomainVirusTotalPtr *bool, subdomainBuiltWithPtr *bool,
	subdomainOutputJSON, subdomainOutputHTML, subdomainOutputTXT, subdomainUserAgentPtr *string,
	subdomainRandomUserAgentPtr *bool, subdomainDNSPtr *string, subdomainAlivePtr *bool) []string {
	// Required Flags
	if *subdomainTargetPtr == "" {
		subdomainCommand.PrintDefaults()
		os.Exit(1)
	}

	// Verify good inputs
	if !urlUtils.IsURL(*subdomainTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}

	// no-check checks
	if *subdomainNoCheckPtr && !*subdomainDBPtr {
		fmt.Println("You can use no-check only with db option.")
		os.Exit(1)
	}

	if *subdomainNoCheckPtr && *subdomainWordlistPtr != "" {
		fmt.Println("You can't use no-check with wordlist option.")
		os.Exit(1)
	}

	if *subdomainNoCheckPtr && *subdomainIgnorePtr != "" {
		fmt.Println("You can't use no-check with ignore option.")
		os.Exit(1)
	}

	if *subdomainNoCheckPtr && *subdomainCrawlerPtr {
		fmt.Println("You can't use no-check with crawler option.")
		os.Exit(1)
	}

	// output files all different
	if *subdomainOutputJSON != "" {
		if *subdomainOutputJSON == *subdomainOutputTXT || *subdomainOutputJSON == *subdomainOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *subdomainOutputHTML != "" {
		if *subdomainOutputHTML == *subdomainOutputTXT || *subdomainOutputJSON == *subdomainOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *subdomainOutputTXT != "" {
		if *subdomainOutputJSON == *subdomainOutputTXT || *subdomainOutputTXT == *subdomainOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *subdomainVirusTotalPtr && !*subdomainDBPtr {
		fmt.Println("You can't specify VirusTotal and not the Open Database option.")
		fmt.Println("If you want to use VirusTotal Api, set also -db option.")
		os.Exit(1)
	}

	if *subdomainBuiltWithPtr && !*subdomainDBPtr {
		fmt.Println("You can't specify BuiltWith and not the Open Database option.")
		fmt.Println("If you want to use BuiltWith Api, set also -db option.")
		os.Exit(1)
	}

	var (
		subdomainIgnore []string
		err             error
	)

	if *subdomainIgnorePtr != "" {
		toBeIgnored := *subdomainIgnorePtr

		subdomainIgnore, err = ignoreUtils.CheckIgnore(toBeIgnored)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if *subdomainUserAgentPtr != DefaultUserAgent && *subdomainRandomUserAgentPtr {
		fmt.Println("You cannot specify both ua and rua.")
		os.Exit(1)
	}

	if *subdomainNoCheckPtr && *subdomainDNSPtr != "" {
		fmt.Println("You can't use no-check with DNS option.")
		os.Exit(1)
	}

	if *subdomainNoCheckPtr && *subdomainAlivePtr {
		fmt.Println("You can't use no-check with alive option.")
		os.Exit(1)
	}

	if *subdomainDNSPtr != "" && *subdomainAlivePtr {
		fmt.Println("You can't use DNS with alive option.")
		os.Exit(1)
	}

	if !*subdomainAlivePtr && (*subdomainUserAgentPtr != DefaultUserAgent || *subdomainRandomUserAgentPtr) {
		fmt.Println("User Agent options are available only with -alive.")
		os.Exit(1)
	}

	return subdomainIgnore
}

// PortSubcommandCheckFlags performs all the necessary checks on the flags
// for the port subcommand.
func PortSubcommandCheckFlags(portCommand flag.FlagSet, portTargetPtr *string, portsPtr *string,
	portCommonPtr *bool, startPort int, endPort int, portTimeout *int,
	portOutputJSON, portOutputHTML, portOutputTXT *string) (int, int, []int, bool) {
	// Required Flags
	if *portTargetPtr == "" {
		portCommand.PrintDefaults()
		os.Exit(1)
	}

	// common and p not together
	if *portsPtr != "" && *portCommonPtr {
		fmt.Println("You can't specify a port range and common option together.")
		os.Exit(1)
	}

	var (
		portArrayBool bool
		portsArray    []int
	)

	startPort, endPort, portsArray, portArrayBool = transportUtils.PortsInputHelper(portsPtr,
		startPort, endPort, portsArray, portArrayBool)

	// output files all different
	if *portOutputJSON != "" {
		if *portOutputJSON == *portOutputTXT || *portOutputJSON == *portOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *portOutputHTML != "" {
		if *portOutputHTML == *portOutputTXT || *portOutputJSON == *portOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *portOutputTXT != "" {
		if *portOutputJSON == *portOutputTXT || *portOutputTXT == *portOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	// Verify good inputs
	if !urlUtils.IsURL(*portTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}

	if *portTimeout < 1 || *portTimeout > 100 {
		fmt.Println("Port Scan timeout must be an integer between 1 and 100.")
		os.Exit(1)
	}

	return startPort, endPort, portsArray, portArrayBool
}

// DirSubcommandCheckFlags performs all the necessary checks on the flags
// for the dir subcommand.
func DirSubcommandCheckFlags(dirCommand flag.FlagSet, dirTargetPtr *string,
	dirIgnorePtr *string, dirOutputJSON, dirOutputHTML, dirOutputTXT, dirUserAgentPtr *string,
	dirRandomUserAgentPtr *bool) []string {
	// Required Flags
	if *dirTargetPtr == "" {
		dirCommand.PrintDefaults()
		os.Exit(1)
	}

	// Verify good inputs
	if !urlUtils.IsURL(*dirTargetPtr) {
		fmt.Println("The inputted target is not valid.")
		os.Exit(1)
	}

	// output files all different
	if *dirOutputJSON != "" {
		if *dirOutputJSON == *dirOutputTXT || *dirOutputJSON == *dirOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *dirOutputHTML != "" {
		if *dirOutputHTML == *dirOutputTXT || *dirOutputJSON == *dirOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	if *dirOutputTXT != "" {
		if *dirOutputJSON == *dirOutputTXT || *dirOutputTXT == *dirOutputHTML {
			fmt.Println("The output paths must be all different.")
			os.Exit(1)
		}
	}

	var (
		dirIgnore []string
		err       error
	)

	if *dirIgnorePtr != "" {
		toBeIgnored := *dirIgnorePtr

		dirIgnore, err = ignoreUtils.CheckIgnore(toBeIgnored)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if *dirUserAgentPtr != DefaultUserAgent && *dirRandomUserAgentPtr {
		fmt.Println("You cannot specify both ua and rua.")
		os.Exit(1)
	}

	return dirIgnore
}

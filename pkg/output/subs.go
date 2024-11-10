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

package output

import (
	"fmt"
	"os"
	"sync"

	urlUtils "github.com/edoardottt/scilla/internal/url"
	"github.com/fatih/color"
)

// PrintSubs prints the results (only the resources not already printed).
// Also performs the checks based on the response status codes.
func PrintSubs(subs map[string]Asset, ignore []string, outputFileJSON, outputFileHTML, outputFileTXT string,
	mutex *sync.Mutex, plain bool) {
	mutex.Lock()
	for domain, asset := range subs {
		if !asset.Printed {
			sub := Asset{
				Value:   asset.Value,
				Printed: true,
			}

			subs[domain] = sub

			var resp = asset.Value

			if !plain {
				fmt.Fprint(os.Stdout, "\r")

				if resp == "" || resp[:3] != "404" {
					subDomainFound := urlUtils.CleanProtocol(domain)
					fmt.Printf("[+]FOUND: %s ", subDomainFound)

					if resp == "" || string(resp[0]) == "2" {
						if outputFileJSON != "" {
							AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "json", outputFileJSON)
						}

						if outputFileHTML != "" {
							AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "html", outputFileHTML)
						}

						if outputFileTXT != "" {
							AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "txt", outputFileTXT)
						}

						color.Green("%s\n", resp)
					} else {
						if outputFileJSON != "" {
							AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "json", outputFileJSON)
						}

						if outputFileHTML != "" {
							AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "html", outputFileHTML)
						}

						if outputFileTXT != "" {
							AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "txt", outputFileTXT)
						}

						color.Red("%s\n", resp)
					}
				}
			} else if resp[:3] != "404" {
				subDomainFound := urlUtils.CleanProtocol(domain)
				fmt.Printf("%s\n", subDomainFound)

				if outputFileJSON != "" {
					AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "json", outputFileJSON)
				}

				if outputFileHTML != "" {
					AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "html", outputFileHTML)
				}

				if outputFileTXT != "" {
					AppendWhere(domain, fmt.Sprint(resp), "SUB", "", "txt", outputFileTXT)
				}
			}
		}
	}
	mutex.Unlock()
}

// AddSubs adds the target found to the subs map.
func AddSubs(target string, value string, subs map[string]Asset, mutex *sync.Mutex) {
	sub := Asset{
		Value:   value,
		Printed: false,
	}

	target = urlUtils.CleanProtocol(target)
	if !PresentSubs(target, subs, mutex) {
		mutex.Lock()
		subs[target] = sub
		mutex.Unlock()
	}
}

// PresentSubs checks if a subdomain is present inside the subs map.
func PresentSubs(input string, subs map[string]Asset, mutex *sync.Mutex) bool {
	mutex.Lock()
	_, ok := subs[input]
	mutex.Unlock()

	return ok
}

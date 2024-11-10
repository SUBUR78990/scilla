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
	"encoding/json"
	"log"
	"os"

	fileUtils "github.com/edoardottt/scilla/internal/file"
)

// File struct helping json output.
type File struct {
	Port      []string            `json:"port,omitempty"`
	DNS       map[string][]string `json:"dns,omitempty"`
	Subdomain []string            `json:"subdomain,omitempty"`
	Dir       []string            `json:"dir,omitempty"`
}

// AppendOutputToJSON appends a (json) row in the JSON output file.
func AppendOutputToJSON(output string, key string, record string, filename string) {
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	if len(file) == 0 {
		file = []byte(`{}`)
	}

	data := File{}

	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	switch {
	case key == "PORT":
		{
			data.Port = append(data.Port, output)
		}
	case key == "SUB":
		{
			data.Subdomain = append(data.Subdomain, output)
		}
	case key == "DIR":
		{
			data.Dir = append(data.Dir, output)
		}
	default:
		{
			if data.DNS == nil {
				data.DNS = make(map[string][]string)
			}

			if _, ok := data.DNS[record]; !ok {
				data.DNS[record] = make([]string, 0)
			}
			data.DNS[record] = append(data.DNS[record], output)
		}
	}

	file, err = json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
	} else {
		err = os.WriteFile(filename, file, fileUtils.Permission0644)
		if err != nil {
			log.Fatal(err)
		}
	}
}

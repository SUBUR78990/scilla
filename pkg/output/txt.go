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
	"log"
	"os"

	fileUtils "github.com/edoardottt/scilla/internal/file"
)

// AppendOutputToTxt appends a (txt) row in the TXT output file.
func AppendOutputToTxt(output string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, fileUtils.Permission0644)
	if err != nil {
		log.Println(err)
	}

	if _, err := file.WriteString(output + "\n"); err != nil {
		log.Fatal(err)
	}

	file.Close()
}

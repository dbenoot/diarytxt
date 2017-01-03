//   Copyright 2016 The diarytxt Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/go-ini/ini"
	"github.com/pmylund/sortutil"
	"os"
	"path/filepath"
	"strings"
)

func pin(a string, r string, l string, i bool, ia bool, sd string, cfgFile string, args []string) {

	var err error

	var cfg, _ = ini.LooseLoad(cfgFile)

	if len(a) > 0 {
		p := cfg.Section("general").Key("pins").String() + ", " + a
		_, _ = cfg.Section("general").NewKey("pins", p)
		err = cfg.SaveTo(cfgFile)
	}

	if len(r) > 0 {
		var p string
		s := cfg.Section("general").Key("pins").Strings(",")
		for i, v := range s {
			if v == r {
				s = append(s[:i], s[i+1:]...)
				break
			}
		}
		for j := range s {
			if j == len(s)-1 {
				p = p + s[j]
			} else {
				p = p + s[j] + ", "
			}
		}
		_, _ = cfg.Section("general").NewKey("pins", p)
		err = cfg.SaveTo(cfgFile)
	}

	if i == true {
		s := cfg.Section("general").Key("pins").Strings(",")
		color.Green("Specified pins:")
		for i := range s {
			fmt.Println(s[i])
		}
	}

	if ia == true {

		s := cfg.Section("general").Key("pins").Strings(",")

		fileList := []string{}
		err = filepath.Walk(sd, func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return nil
		})
		fileList = filterFile(fileList)

		color.Green("Full index of the specified pins and their unique values:")

		for i := range s {

			var index []string

			for _, file := range fileList {
				fo, _ := os.Open(file)

				scanner := bufio.NewScanner(fo)

				for scanner.Scan() {
					if strings.Contains(scanner.Text(), "* "+s[i]+":") == true {
						item := strings.TrimSpace(strings.Replace(scanner.Text(), "* "+s[i]+":", "", -1))
						if len(item) > 0 {
							index = AppendIfMissing(index, item)
						}
					}
				}
			}

			color.Cyan("Entries for pin " + s[i])
			// 			fmt.Println(len(index))
			if len(index) != 0 {
				sortutil.CiAsc(index)
				for j := range index {
					fmt.Printf("%v : %s \n", j+1, index[j])
				}
			}
		}
	}

	if len(l) > 0 {

		fileList := []string{}
		err = filepath.Walk(sd, func(path string, f os.FileInfo, err error) error {
			fileList = append(fileList, path)
			return nil
		})

		fileList = filterFile(fileList) //function used in search.go

		var index []string
		var date []string
		var files []string

		for _, file := range fileList {

			var d string

			fo, _ := os.Open(file)

			scanner := bufio.NewScanner(fo)

			for scanner.Scan() {

				if strings.Contains(scanner.Text(), "* date:") == true {
					d = strings.TrimSpace(strings.Replace(scanner.Text(), "* date:", "", -1))
				}

				if strings.Contains(scanner.Text(), "* "+l+":") == true {
					item := strings.TrimSpace(strings.Replace(scanner.Text(), "* "+l+":", "", 1))
					if len(item) > 0 {
						index = append(index, item)
						date = append(date, d)
						files = append(files, file)
					}
				}
			}

		}

		color.Green("Dated entries for pin " + l)

		if len(index) != 0 {
			sortutil.CiAsc(index)
			for k := range index {
				fmt.Printf("%v \t %s \t %s \n", date[k], index[k], files[k])
			}
		}

	}

	if err != nil {
		fmt.Println(err)
	}
}

func AppendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
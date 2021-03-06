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

// attach
// 1. file
//    - check that diary entry exists
//    - upload file in files folder, rename if necessary
//    - create link to uploaded file in markdown
// 2. image
//    - check that diary entry exists
//    - upload image in files folder, rename if necessary
//    - create img linked to uploaded file in markdown

package main

import (
	// "bufio"
	"fmt"
	// "github.com/pmylund/sortutil"
	// "io/ioutil"
	"os"
	// "path/filepath"
	// "strings"
)

func attach(i string, f string, e string) {
	if len(e) == 0 {
		fmt.Println("Please specify diary entry using the -entry flag.")
	}

	if len(i) > 0 {

		checkfile(i)

		fmt.Println("Image.")
	}

	if len(f) > 0 {

		checkfile(f)

		fmt.Println("File.")
	}
}

func checkfile(name string) {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exists. Please check path.")
			os.Exit(2)
		}
	}
}

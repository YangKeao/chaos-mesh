// Copyright 2020 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package mapreader

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	commonerror "github.com/chaos-mesh/chaos-mesh/pkg/commonerror"
)

// Entry is one line in /proc/pid/maps
type Entry struct {
	StartAddress uint64
	EndAddress   uint64
	Privilege    string
	PaddingSize  uint64
	Path         string
}

// Read parse /proc/[pid]/maps and return a list of entry
// The format of /proc/[pid]/maps can be found in `man proc`.
func Read(pid int) ([]Entry, *Error) {
	data, err := ioutil.ReadFile(fmt.Sprintf("/proc/%d/maps", pid))
	if err != nil {
		return nil, ErrorWrap(&commonerror.IOError{
			Err: err,
		})
	}

	lines := strings.Split(string(data), "\n")

	var entries []Entry
	for _, line := range lines {
		sections := strings.Split(line, " ")
		if len(sections) < 3 {
			continue
		}

		var path string

		if len(sections) > 5 {
			path = sections[len(sections)-1]
		}

		addresses := strings.Split(sections[0], "-")
		startAddress, err := strconv.ParseUint(addresses[0], 16, 64)
		if err != nil {
			return nil, ErrorWrap(&commonerror.ParseIntError{
				S:       addresses[0],
				Base:    16,
				Bitsize: 64,
			})
		}
		endAddresses, err := strconv.ParseUint(addresses[1], 16, 64)
		if err != nil {
			return nil, ErrorWrap(&commonerror.ParseIntError{
				S:       addresses[1],
				Base:    16,
				Bitsize: 64,
			})
		}

		privilege := sections[1]

		paddingSize, err := strconv.ParseUint(sections[2], 16, 64)
		if err != nil {
			return nil, ErrorWrap(&commonerror.ParseIntError{
				S:       addresses[2],
				Base:    16,
				Bitsize: 64,
			})
		}

		entries = append(entries, Entry{
			startAddress,
			endAddresses,
			privilege,
			paddingSize,
			path,
		})
	}

	return entries, nil
}

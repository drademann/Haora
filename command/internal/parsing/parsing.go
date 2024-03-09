//
// Copyright 2024-2024 The Haora Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package parsing

import (
	"errors"
	"fmt"
	"github.com/drademann/haora/app/datetime"
	"regexp"
	"strconv"
	"time"
)

func Time(flag string, args []string) (time.Time, []string, error) {
	if flag != "" {
		t, err := parseTime(flag)
		if err != nil {
			return t, args, err
		}
		return t, args, nil
	}
	if len(args) > 0 {
		t, err := parseTime(args[0])
		if err != nil {
			return t, args, err
		}
		return t, args[1:], nil
	}
	return time.Time{}, args, errors.New("no time found")
}

var timeRE = regexp.MustCompile(`(\d?\d):?(\d\d)`)

func parseTime(timeStr string) (time.Time, error) {
	if timeStr == "now" {
		return datetime.Now(), nil
	}
	groups := timeRE.FindStringSubmatch(timeStr)
	if len(groups) == 0 {
		return time.Time{}, errors.New("invalid time format")
	}
	hour, err := strconv.Atoi(groups[1])
	if err != nil {
		return time.Time{}, err
	}
	if hour > 23 {
		return time.Time{}, fmt.Errorf("invalid hour: %d", hour)
	}
	minute, err := strconv.Atoi(groups[2])
	if err != nil {
		return time.Time{}, err
	}
	if minute > 59 {
		return time.Time{}, fmt.Errorf("invalid minute: %d", minute)
	}
	t := time.Time{}
	return time.Date(t.Year(), t.Month(), t.Day(), hour, minute, 0, 0, t.Location()), nil
}

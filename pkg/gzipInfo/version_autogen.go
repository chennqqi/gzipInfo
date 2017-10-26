//==============================================================================
// This file is part of gzipInfo
// Copyright (c) 2017, FutureQuest, Inc.
//   https://www.FutureQuest.net
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//==============================================================================

// Code generated by go generate; DO NOT EDIT.
// This file was generated by genVersion at
// 2017-10-12 04:21:10.008227115 -0700 PDT m=+0.000181776

package gzipInfo

import (
	"futurequest.net/FQgolibs/FQversion"
)

var (
	PROG    string = "gzipInfo"
	VERSION string = "v3.0.1-1-g8319ee5"
	BUILD   string = "20171012@04:21:09"
)

// +build FQversion

func init() {
	FQversion.Register(PROG, VERSION, BUILD)
}

func Version() string {
	return FQversion.ProgVersion(PROG, VERSION, BUILD)
}

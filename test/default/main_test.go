// Copyright (c) 2020-2022 Doc.ai and/or its affiliates.
//
// Copyright (c) 2023-2024 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package default_test

import (
	"testing"

	"github.com/networkservicemesh/integration-tests/extensions/parallel"
	"github.com/networkservicemesh/integration-tests/suites/basic"
	"github.com/networkservicemesh/integration-tests/suites/ipsec_mechanism"
	"github.com/networkservicemesh/integration-tests/suites/memory"
)

func TestBasic(t *testing.T) {
	parallel.Run(t, new(basic.Suite))
}

func TestMemory(t *testing.T) {
	parallel.Run(t, new(memory.Suite))
}

func TestIPSec(t *testing.T) {
	parallel.Run(t, new(ipsec_mechanism.Suite))
}
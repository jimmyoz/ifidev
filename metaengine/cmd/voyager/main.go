// Copyright 2020 The Infinity Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/klauspost/cpuid"
	"os"

	"github.com/yanhuangpai/voyager/cmd/voyager/cmd"
)

func main() {
	fmt.Println("logicalScore",float64(cpuid.CPU.LogicalCores)*float64(cpuid.CPU.Hz)/1000000000.00)

	if err := cmd.Execute(); err != nil {
		println(os.Stderr, "Error:", err.Error())
		os.Exit(1)
	}
}

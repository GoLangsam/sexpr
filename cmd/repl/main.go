// Copyright 2020 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// REPL for symbolic expressions

package main

import (
	"fmt"
	"os"
	"time"
)

// ===========================================================================

func main() {

	doLoop()

	if true {
		fmt.Println("about to leave ... are goroutines leaking?")
		<-time.After(time.Millisecond * 100)
		os.Exit(1) // to see leaking goroutines, if any
	}

}

// ===========================================================================

// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/GoLangsam/sexpr"
)

// ===========================================================================

func doLoop() {
	fmt.Println("REPL - Please enter some expression and press [Enter]")
	defer fmt.Println("bye ...")

	text := make([]byte, 2048)
	next := func() (int, error) { return os.Stdin.Read(text) }

	for n, err := next(); err == nil; n, err = next() {
		line := string(text[:n-1]) // strip trailing CrLf
		if ast, err := sexpr.Parse(line); err == nil {
			fmt.Println(ast.String())
		} else {
			fmt.Println(err)
		}
	}
}

// ===========================================================================

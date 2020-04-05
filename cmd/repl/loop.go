// Copyright 2017 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	//"strings"

	"github.com/GoLangsam/sexpr"
)

// ===========================================================================

func doLoop() {
	buf := make([]byte,1024) 

	for {
		n, err := os.Stdin.Read(buf)
		if err!=nil {break}
		
		// fmt.Println(by)
		
		//s := strings.Replace(string(buf[:n]), "\n", "", -1) // strip CrLf
		s := string(buf[:n-2])
		
		fmt.Println(n)
		//fmt.Println(s)
		
		ast, err := sexpr.Parse(s)
		//if err == nil {
			fmt.Println(ast.String())
		//} else {
			fmt.Println(err)
		//}
	}
}

// ===========================================================================

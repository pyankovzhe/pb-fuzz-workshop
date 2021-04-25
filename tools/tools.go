// +build tools

package tools

import (
	_ "github.com/dvyukov/go-fuzz/go-fuzz"
	_ "github.com/dvyukov/go-fuzz/go-fuzz-build"
)

//go:generate go build -v -o ../bin/go-fuzz       github.com/dvyukov/go-fuzz/go-fuzz
//go:generate go build -v -o ../bin/go-fuzz-build github.com/dvyukov/go-fuzz/go-fuzz-build

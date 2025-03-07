//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/josephlewis42/skilltreetool/pkg/browser"
)

func main() {
	fmt.Println("started wasm")

	js.Global().Set("svg2yaml", js.FuncOf(func(this js.Value, args []js.Value) any {
		got := browser.SVG2Yaml(args[0].String())
		return js.ValueOf(got)
	}))

	<-make(chan bool) // To use anything from Go WASM, the program may not exit.

}

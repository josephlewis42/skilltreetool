//go:build js && wasm

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/josephlewis42/skilltreetool/pkg/commands"
	"github.com/josephlewis42/skilltreetool/pkg/models"
)

func outputJSON(output string, err error) js.Value {

	type Output struct {
		Data string `json:"data"`
		Err  string `json:"err"`
	}

	var out Output
	if err != nil {
		out.Err = err.Error()
	} else {
		out.Data = output
	}

	data, err := json.Marshal(&out)

	if err != nil {
		fmt.Printf("couldn't encode JSON to return: %w\n", err)
		js.ValueOf(err.Error())
	}

	return js.ValueOf(string(data))
}

func jsonError(err error) js.Value {
	return outputJSON("", err)
}

func main() {
	fmt.Println("started wasm")

	js.Global().Set("svg2yaml", js.FuncOf(func(this js.Value, args []js.Value) any {
		buf := bytes.Buffer{}
		err := commands.SVG2Yaml([]byte(args[0].String()), &buf)

		return outputJSON(buf.String(), err)
	}))

	js.Global().Set("yaml2svg", js.FuncOf(func(this js.Value, args []js.Value) any {
		buf := bytes.Buffer{}
		err := commands.Yaml2SVG([]byte(args[0].String()), &buf)

		return outputJSON(buf.String(), err)
	}))

	js.Global().Set("diff", js.FuncOf(func(this js.Value, args []js.Value) any {

		before, err := models.LoadFromString(args[0].String())
		if err != nil {
			return jsonError(fmt.Errorf("couldn't load before: %w", err))
		}

		after, err := models.LoadFromString(args[1].String())
		if err != nil {
			return jsonError(fmt.Errorf("couldn't read after: %w", err))
		}

		diff := commands.Diff(before, after)

		return outputJSON(diff.ToMarkdown(), nil)
	}))

	<-make(chan bool) // To use anything from Go WASM, the program may not exit.

}

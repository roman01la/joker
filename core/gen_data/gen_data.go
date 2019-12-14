package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strings"

	. "github.com/candid82/joker/core"
	_ "github.com/candid82/joker/std/html"
	_ "github.com/candid82/joker/std/string"
)

/* Note: Rather than depend on initialization order, this code allows
/* any order as long as init() functions are not run at the same time
/* (on different threads). In particular, while both a_*_code.go and
/* a_*_data.go set the .generated field, the former (assuming it
/* exists) will always set it to a valid function, while the latter
/* will avoid resetting it (to nil) by preserving the existing value
/* when it initializes the remainder of the internalNamespaceInfo
/* struct. */

var template string = `// Generated by gen_data. Don't modify manually!

package core

func init() {
	{name}NamespaceInfo = internalNamespaceInfo{init: {name}Init, generated: {name}NamespaceInfo.generated, available: true}
}

func {name}Init() {
	{name}NamespaceInfo.data = []byte("{content}")
}
`

const hextable = "0123456789abcdef"

func main() {
	GLOBAL_ENV.FindNamespace(MakeSymbol("user")).ReferAll(GLOBAL_ENV.CoreNamespace)
	for _, f := range CoreSourceFiles {
		GLOBAL_ENV.SetCurrentNamespace(GLOBAL_ENV.CoreNamespace)
		content, err := ioutil.ReadFile("data/" + f.Filename)
		if err != nil {
			panic(err)
		}
		content, err = PackReader(NewReader(bytes.NewReader(content), f.Name), "")
		PanicOnErr(err)

		name := f.Filename[0 : len(f.Filename)-5] // assumes .joke extension
		if _, err := os.Stat("a_" + name + "_code.go"); err == nil {
			continue // already have generated-code version
		}

		dst := make([]byte, len(content)*4)
		for i, v := range content {
			dst[i*4] = '\\'
			dst[i*4+1] = 'x'
			dst[i*4+2] = hextable[v>>4]
			dst[i*4+3] = hextable[v&0x0f]
		}
		fileContent := strings.ReplaceAll(template, "{name}", name)
		fileContent = strings.Replace(fileContent, "{content}", string(dst), 1)
		ioutil.WriteFile("a_"+name+"_data.go", []byte(fileContent), 0666)
	}
}

package parser

import (
	"flag"
	"fmt"
	"strings"
)

var traceLevel int = 0
var debugSwitch *bool

func init() {
	debugSwitch = flag.Bool("d", false, "debug switch")
}

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(fs string) {
	if !flag.Parsed() {
		flag.Parse()
	}
	if debugSwitch != nil && *debugSwitch {
		fmt.Printf("%s%s\n", identLevel(), fs)
	}
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func untrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}

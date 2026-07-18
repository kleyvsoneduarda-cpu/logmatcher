package main

import (
	"fmt"

	"github.com/digitalocean/captainslog"
	matcherpkg "github.com/digitalocean/logmatcher"
)

func main() {
	// exemplo simples: cria uma mensagem e testa um matcher de programa
	m := captainslog.NewSyslogMsg()
	m.SetProgram("logCatcher_1")

	v := matcherpkg.NewValue(matcherpkg.Program, matcherpkg.PrefixMatch, "logCatcher_")

	fmt.Printf("Matcher: %s\n", v.String())
	fmt.Printf("Matches: %v\n", v.Matches(m))
}

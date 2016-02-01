package cli

import (
	"strings"

	"github.com/alecthomas/kingpin"
)

func Flag(name string, helpOptional ...string) *kingpin.FlagClause {
	help := strings.Join(helpOptional, " ")
	return kingpin.Flag(name, help)
}

func Arg(name string, helpOptional ...string) *kingpin.ArgClause {
	help := strings.Join(helpOptional, " ")
	return kingpin.Arg(name, help)
}

func ParseArgs() string {
	// kingpin.IgnoreUnkownFlags = true
	return kingpin.Parse()
}

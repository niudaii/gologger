package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/formatter"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/gologger/writer"
	"strconv"
)

func main() {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)
	gologger.DefaultLogger.SetFormatter(formatter.NewCLI(true))
	gologger.DefaultLogger.SetWriter(writer.NewLogFile("gologger.txt"))
	gologger.Print().Msgf("\tgologger: sample test\t\n")
	gologger.Info().Str("user", "pdteam").Msg("running simulation program")
	for i := 0; i < 10; i++ {
		gologger.Info().Str("count", strconv.Itoa(i)).Msg("running simulation step...")
	}
	gologger.Debug().Str("state", "running").Msg("planner running")
	gologger.Debug().TimeStamp().Str("state", "running").Msg("with timestamp event")
	gologger.Warning().Str("state", "errored").Str("status", "404").Msg("could not run")
	gologger.Fatal().Msg("bye bye")
}

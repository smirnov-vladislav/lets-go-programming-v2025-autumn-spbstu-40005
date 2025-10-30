package main

import (
	"flag"
	"sort"

	"github.com/smirnov-vladislav/vladislav.smirnov/task-3/internal/jsonwriter"
	"github.com/smirnov-vladislav/vladislav.smirnov/task-3/internal/config"
	"github.com/smirnov-vladislav/vladislav.smirnov/task-3/internal/valute"
	"github.com/smirnov-vladislav/vladislav.smirnov/task-3/internal/xmlparser"
)

func main() {
	configPath := flag.String("config", "config.yaml", "config file path")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	valCurs, err := xmlparser.Read(cfg.Input)
	if err != nil {
		panic(err)
	}

	sort.Slice(valCurs.Valutes, func(a, b int) bool {
		return valCurs.Valutes[i].Value > valCurs.Valutes[j].Value
	})

	if err := jsonwriter.Write(cfg.Output, valCurs.Valutes); err != nil {
		panic(err)
	}
}

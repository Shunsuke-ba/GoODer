package config

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	FavoWord  string
	FavoCount int
	FavoSleep int
}

var Config ConfigList

func init() {
	load, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	Config = ConfigList{
		FavoWord:  load.Section("CreateFavorite").Key("favo_word").String(),
		FavoCount: load.Section("CreateFavorite").Key("favo_count").MustInt(),
		FavoSleep: load.Section("CreateFavorite").Key("favo_sleep").MustInt(),
	}
}

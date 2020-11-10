package main

import (
	"github.com/kimtaek/gamora/pkg/db"
	"github.com/kimtaek/gamora/pkg/helper"
	"os"
)

func init() {
	db.Setup()
}

func main() {
	defer db.CloseDB()

	if err := db.Connection().DB().Ping(); err != nil {
		helper.Error("database connecting ping error!")
		os.Exit(1)
	}

	os.Exit(0)
}

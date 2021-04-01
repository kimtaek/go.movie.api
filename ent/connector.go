package ent

import (
	"context"
	"fmt"
	"github.com/caarlos0/env"
	"github.com/fatih/color"
	"github.com/kimtaek/gamora/pkg/slack"
	"log"
	"os"
	"time"
)

// Configure config for db
type Configure struct {
	Database string `env:"MYSQL_DATABASE" envDefault:"database"`
	Username string `env:"MYSQL_USERNAME" envDefault:"user"`
	Password string `env:"MYSQL_PASSWORD" envDefault:"password"`
	Host     string `env:"MYSQL_HOST" envDefault:"localhost"`
	Port     string `env:"MYSQL_PORT" envDefault:"3306"`
}

// Config global defined db config
var Config Configure

// client global defined db config
var client *Client

// Setup init db
func Setup() {
	_ = env.Parse(&Config)
	c, err := Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Username,
		Config.Password,
		Config.Host,
		Config.Port,
		Config.Database,
	))

	if err != nil {
		slack.SendMessage(slack.Message{
			Text: "Database: " + err.Error(),
		})
		os.Exit(1)
	}

	ctx := context.Background()
	if err := c.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources:%v", err)
	}

	client = c
	_, _ = color.New(color.FgWhite).Println(time.Now().Format(time.RFC3339), "[info]", "[EntORM schema created!]")
	_, _ = color.New(color.FgWhite).Println(time.Now().Format(time.RFC3339), "[info]", "[EntORM database connected!]")
}

// Connection get db connection
func Connection() *Client {
	return client.Debug()
}

// CloseDB close db connection
func CloseDB() {
	defer client.Close()
}

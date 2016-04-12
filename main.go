package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
)

// Config holds the emoji configuration
type Config struct {
	Words map[string]string `toml:"commitKinds"`
}

var pwd string
var hook = "\n# simplifies emoji usage \nemojify -m \"`cat $1`\" > \"$1\""

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

func install() {
	if _, err := os.Stat(pwd + "/.git"); err != nil {
		log.Fatal("fatal: Not a git repository (or any of the parent directories): .git")
	}
	f, err := os.OpenFile(pwd+"/.git/hooks/prepare-commit-msg", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()
	_, werr := f.WriteString(hook)
	if werr != nil || err != nil {
		log.Fatal("fatal: unable to create the prepare-commit-msg hook!", werr, err)
	}
	log.Println("created hook for ", pwd)
}

// Run ...
func Run(c *cli.Context) {
	cfg := loadEmojiMap()
	msg := c.String("message")
	if msg != "" {
		fmt.Println(replace(cfg, msg))
	}
}

func loadEmojiMap() *Config {
	wrds := Config{}
	if _, err := toml.DecodeFile(pwd+"/.emojifyrc", &wrds); err != nil {
		data, err := Asset("config/emoji.toml")
		if err != nil {
			log.Fatal("Fatal: No config file found")
		}
		if _, err := toml.Decode(string(data), &wrds); err != nil {
			log.Fatal(err)
		}
	}
	return &wrds
}

func replace(cfg *Config, msg string) string {
	words := strings.Split(msg, " ")
	for i, w := range words {
		if strings.HasPrefix(w, ":") {
			raw := strings.TrimPrefix(w, ":")
			if cfg.Words[raw] != "" {
				words[i] = cfg.Words[raw]
			}
		}
	}
	return strings.Join(words, " ")
}

func main() {
	// installHook()
	app := cli.NewApp()
	app.Name = "emojify"
	app.Usage = "Make your git commits more expressive"
	app.Action = Run
	app.Version = "0.1.4"
	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "initialize the git hook",
			Action: func(c *cli.Context) {
				install()
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "message, m",
			Usage: "The message to commit",
		},
	}
	app.Run(os.Args)
}

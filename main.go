package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/pborman/getopt"
)

// Config holds the emoji configuration
type Config struct {
	Words map[string]string `toml:"commitKinds"`
}

var pwd string
var hook = "\n# simplifies emoji usage \nlipstick \"`cat $1`\" > \"$1\""

func init() {
	var err error
	pwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

// install adds the hook to this program to the local git repo
func install() {
	if _, err := os.Stat(pwd + "/.git"); err != nil {
		log.Fatal("fatal: Not a git repository (or any of the parent directories): .git")
	}
	f, err := os.OpenFile(pwd+"/.git/hooks/commit-msg", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	defer f.Close()
	_, werr := f.WriteString(hook)
	if werr != nil || err != nil {
		log.Fatal("fatal: unable to create the commit-msg hook!", werr, err)
	}
	log.Println("created hook for ", pwd)
}

// Run is our main function
func Run(c *cli.Context) {
	getopt.Parse()
	args := getopt.Args()
	msg := strings.Join(args, " ")
	cfg, err := loadEmojiMap()
	if err != nil {
		log.Fatal("fatal: could not load config", err)
	}
	if msg != "" {
		fmt.Println(replace(cfg, msg))
	} else {
		log.Fatal("fatal: no message given")
	}
}

// loadEmojiMap loads the config file into our config and fallsback on the
// default built in config
func loadEmojiMap() (*Config, error) {
	cfg := &Config{}
	if _, err := loadLocalConfig(cfg); err != nil {
		if _, err := loadDefaultConfig(cfg); err != nil {
			return nil, err
		}
	}
	return cfg, nil
}

// loadLocalConfig attempts to load the local config file
func loadLocalConfig(cfg *Config) (*Config, error) {
	if _, err := toml.DecodeFile(pwd+"/.lipstickrc", &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// loadDefaultConfig attempts to load the builtin config file from the bindata
// file.
func loadDefaultConfig(cfg *Config) (*Config, error) {
	data, err := Asset("config/lipstickrc.toml")
	if err != nil {
		return nil, err
	}
	if _, err := toml.Decode(string(data), &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// replace finds words that fit our params in the msg and replaces them with
// the words defined in our config file.
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
	app := cli.NewApp()
	app.Name = "lipstick"
	app.Usage = "Make your git commits more expressive"
	app.Action = Run
	app.Version = "0.4.7"
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
	app.Run(os.Args)
}

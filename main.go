package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/codegangsta/cli"
	"github.com/natefinch/atomic"
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

func uninstall() {
	if _, err := os.Stat(pwd + "/.git"); err != nil {
		log.Fatal("fatal: Not a git repository (or any of the parent directories): .git")
	}
	f, err := os.Open(pwd + "/.git/hooks/commit-msg")
	d, rerr := ioutil.ReadAll(f)
	if err != nil || rerr != nil {
		log.Fatal("fatal: unable to remove commit-msg hook", rerr, err)
	}
	f.Close()
	old := string(d)
	new := strings.Replace(old, "\n# simplifies emoji usage", "", -1)
	new = strings.Replace(new, "\nlipstick \"`cat $1`\" > \"$1\"", "", -1)
	out := strings.NewReader(new)
	if err := atomic.WriteFile(pwd+"/.git/hooks/commit-msg", out); err != nil {
		log.Fatal("fatal: unable to remove commit-msg hook", err)
	}
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
	for key, value := range cfg.Words {
		msg = strings.Replace(msg, ":"+key+":", value, -1)
	}
	return msg
}

func createConfig() {
	if _, err := os.Stat(".lipstickrc"); !os.IsNotExist(err) {
		log.Fatal("fatal: .lipstickrc exists")
	}
	data, err := Asset("config/lipstickrc.toml")
	if err != nil {
		log.Fatal("fatal: could not load default .lipstickrc", err)
	}
	r := strings.NewReader(string(data))
	if err := atomic.WriteFile(".lipstickrc", r); err != nil {
		log.Fatal("fatal: could not generate .lipstickrc", err)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "lipstick"
	app.Usage = "Make your git commits more expressive"
	app.Action = Run
	app.Version = "4.2.0"
	app.Commands = []cli.Command{
		{
			Name:    "install",
			Aliases: []string{"i"},
			Usage:   "initialize the git hook",
			Action: func(c *cli.Context) {
				install()
			},
		}, {
			Name:    "uninstall",
			Aliases: []string{"u"},
			Usage:   "remove the git hook",
			Action: func(c *cli.Context) {
				uninstall()
			},
		}, {
			Name:    "initialize",
			Aliases: []string{"init"},
			Usage:   "creates a .lipstickrc file if one does not exist",
			Action: func(c *cli.Context) {
				createConfig()
			},
		},
	}
	app.Run(os.Args)
}

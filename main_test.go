package main

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var cfg *Config

func init() {
	var err error
	cfg, err = loadEmojiMap()
	if err != nil {
		log.Fatal("fatal: could not load map")
	}
}

func TestReplace(t *testing.T) {
	Convey("Given a message with no key words", t, func() {
		msg := "Hello world I am a message with no keywords"
		Convey("The message should be the same as what was put in", func() {
			out := replace(cfg, msg)
			So(msg, ShouldEqual, out)
		})
	})

	Convey("Given a message with keywords", t, func() {
		msg := ":init I am a message with keywords! :bugfix :crucial :docs"
		Convey("The keywords should be replaced with the emoji", func() {
			out := ":tada: I am a message with keywords! :bug: :ambulance: :books:"
			So(replace(cfg, msg), ShouldEqual, out)
		})
	})

	Convey("Given a multiline message with keywords", t, func() {
		msg := `
:init I am a message with keywords! :bugfix :crucial :docs
:init I am a message with keywords! :bugfix :crucial :docs
:init I am a message with keywords! :bugfix
:crucial :docs
`
		Convey("The keywords should be replaced with the emoji", func() {
			out := `
:tada: I am a message with keywords! :bug: :ambulance: :books:
:tada: I am a message with keywords! :bug: :ambulance: :books:
:tada: I am a message with keywords! :bug:
:ambulance: :books:
`
			So(replace(cfg, msg), ShouldEqual, out)
		})
	})
}

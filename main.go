/*
  Copyright (c) 2013 José Carlos Nieto, http://xiam.menteslibres.org/

  Permission is hereby granted, free of charge, to any person obtaining
  a copy of this software and associated documentation files (the
  "Software"), to deal in the Software without restriction, including
  without limitation the rights to use, copy, modify, merge, publish,
  distribute, sublicense, and/or sell copies of the Software, and to
  permit persons to whom the Software is furnished to do so, subject to
  the following conditions:

  The above copyright notice and this permission notice shall be
  included in all copies or substantial portions of the Software.

  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
  NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
  LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
  OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
  WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package cli

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Subcommand.
type Command interface {
	Execute() error
}

// Command entry.
type Entry struct {
	Name        string
	Usage       string
	Description string
	Arguments   []string
	Command     Command
}

// Project name.
var Name string

// Copyright statement.
var Copyright string

// License notice.
var License string

// Version string.
var Version string

// Project's URL.
var Homepage string

// Author's name
var Author string

// Author's e-mail
var AuthorEmail string

// Array of commands.
var commandList []string

// Command names.
var commandNames map[string]Entry

func init() {
	commandList = []string{}
	commandNames = map[string]Entry{}

	Register("help", Entry{
		Description: "Shows information about the given command.",
		Usage:       "help <command>",
		Command:     &helpCommand{},
	})
}

// Registers a subcommand.
func Register(name string, entry Entry) {
	_, exists := commandNames[name]
	if exists == false {
		entry.Name = name
		commandList = append(commandList, name)
		commandNames[name] = entry
	} else {
		panic(fmt.Sprintf("Command \"%s\" was already registered.\n", name))
	}
}

// Displays a banner with available information (if any).
func Banner() {
	banner := []string{}
	if Name != "" {
		var name string
		if Version != "" {
			name = fmt.Sprintf("%s (%s)", Name, Version)
		} else {
			name = Name
		}
		if Homepage != "" {
			name = name + " - " + Homepage
		}
		banner = append(banner, name)
	}
	if Copyright != "" {
		if License != "" {
			banner = append(banner, fmt.Sprintf("%s. %s.", Copyright, License))
		} else {
			banner = append(banner, Copyright)
		}
	}
	if Author != "" {
		if AuthorEmail != "" {
			banner = append(banner, fmt.Sprintf("by %s <%s>", Author, AuthorEmail))
		} else {
			banner = append(banner, fmt.Sprintf("by %s", Author))
		}
	}
	if len(banner) > 0 {
		fmt.Printf("%s\n\n", strings.Join(banner, "\n"))
	}
}

// Shows help for a subcommand.
func Help(name string) error {

	if sort.StringsAreSorted(commandList) == false {
		sort.Strings(commandList)
	}

	if name == "" {
		fmt.Printf("Usage: %s <arguments> <command>\n\n", os.Args[0])
		fmt.Printf("Available commands for %s:\n\n", os.Args[0])
		for _, name := range commandList {
			entry := commandNames[name]
			fmt.Printf("\t%s\t\t%s\n", name, entry.Description)
		}
		fmt.Printf("\nUse \"%s help <command>\" to view more information about a command.\n", os.Args[0])
	} else {
		return Usage(name)
	}
	return nil
}

// Shows usage for a subcommand.
func Usage(name string) error {
	entry, ok := commandNames[name]
	if ok == false {
		return fmt.Errorf(`No such command "%s".`, name)
	}
	if entry.Description != "" {
		fmt.Printf("Command \"%s\": %s\n", name, entry.Description)
	}
	if entry.Usage != "" {
		fmt.Printf("\nUsage: %s %s\n", os.Args[0], entry.Usage)
	}
	if entry.Arguments != nil {
		fmt.Printf("\nArguments for command \"%s\":\n\n", entry.Name)
		for _, argName := range entry.Arguments {
			arg := flag.Lookup(argName)
			if arg == nil {
				panic(fmt.Sprintf("Flag \"-%s\" is expected for command \"%s\" but it's not defined.", argName, entry.Name))
			} else {
				fmt.Printf("\t-%s [%s]: %s\n", arg.Name, arg.DefValue, arg.Usage)
			}
		}
		fmt.Printf("\n")
	}
	return nil
}

// Executes a subcommand.
func Execute(name string) error {
	cmd, ok := commandNames[name]
	if ok == false {
		return fmt.Errorf(`No such command "%s".`, name)
	}
	err := cmd.Command.Execute()
	if err != nil {
		Usage(name)
		fmt.Printf("\n")
	}
	return err
}

// Dispatches the subcommand.
func Dispatch() error {
	flag.Parse()

	if flag.NArg() > 0 {
		name := flag.Arg(0)
		return Execute(name)
	} else {
		return Help("")
	}

	return nil
}

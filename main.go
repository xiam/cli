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
)

type Command interface {
	Help() error
	Usage() error
	Execute() error
}

var Output = os.Stdout

type Entry struct {
	Name    string
	Description string
	Arguments []string
	Command Command
}

var commandList []string
var commandNames map[string]Entry

func init() {
	commandList = []string{}
	commandNames = map[string]Entry{}

	Register("help", Entry{
		Description: "Shows information about an specific command.",
		Command: &helpCommand{},
	})
}

func Register(name string, entry Entry) error {
	entry.Name = name
	commandList = append(commandList, name)
	commandNames[name] = entry
	return nil
}

func Help() error {
	fmt.Printf("Usage: %s <arguments> <command>\n\n", os.Args[0])
	fmt.Printf("Available commands for %s:\n\n", os.Args[0])
	for name, _ := range commandNames {
		entry := commandNames[name]
		fmt.Printf("\t%s\t\t%s\n", name, entry.Description)
		/*
		if entry.Arguments != nil {
			for _, argName := range entry.Arguments {
				arg := flag.Lookup(argName)
				fmt.Printf("\t\t\t-%s [%s]: %s\n", arg.Name, arg.DefValue, arg.Usage)
			}
		}
		*/
	}
	fmt.Printf("\nUse \"%s help <command>\" to view more information about that command.\n", os.Args[0])
	return nil
}

/* Shows command description given its name. */
func Usage(name string) error {
	cmd, ok := commandNames[name]
	if ok == false {
		return fmt.Errorf(`No such command "%s".`, name)
	}
	fmt.Printf("usage: %s %s\n\n", os.Args[0], name)
	return cmd.Command.Usage()
}

/* Executes a command given its name. */
func Execute(name string) error {
	cmd, ok := commandNames[name]
	if ok == false {
		return fmt.Errorf(`No such command "%s".`, name)
	}
	return cmd.Command.Execute()
}

func Dispatch() error {
	flag.Parse()

	if flag.NArg() > 0 {
		name := flag.Arg(0)
		return Execute(name)
	} else {
		return Help()
	}

	return nil
}

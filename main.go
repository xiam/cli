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

type commandEntry struct {
	Name    string
	Command Command
}

var commandList []string
var commandNames map[string]commandEntry

func init() {
	commandList = []string{}
	commandNames = map[string]commandEntry{}
}

func Register(name string, cmd Command) error {
	entry := commandEntry{
		Name:    name,
		Command: cmd,
	}
	commandList = append(commandList, name)
	commandNames[name] = entry
	return nil
}

func Help() {

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
	name := flag.Arg(1)
	if name[0:1] != "-" {
		return Execute(name)
	}
	return nil
}

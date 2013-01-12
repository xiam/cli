/*
  Copyright (c) 2012-2013 José Carlos Nieto, http://xiam.menteslibres.org/

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
	"fmt"
	"flag"
	"os"
)

type helpCommand struct {

}

func (self *helpCommand) Help() error  {
	fmt.Printf("Help on help.\n")
	return nil
}

func (self *helpCommand) Usage() error {
	fmt.Printf("help <command>\n")
	return nil
}

func (self *helpCommand) Execute() error {

	if flag.NArg() > 1 {

		name := flag.Arg(1)

		entry, ok := commandNames[name]

		if ok == false {
			return fmt.Errorf("Cannot help! no such command: %s.", name)
		}

		fmt.Printf("Usage: %s <arguments> %s\n", os.Args[0], entry.Name)

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

	} else {
		fmt.Printf("Need a command.\n")
	}

	return nil
}

// goconc ermöglicht es mehrere Befehle
// nebenläufig auszuführen.
//
// CLI Logik:
// goconc befehl1 arg1 arg2 :: befehl2 -flag arg1 arg2
package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

type cmdArgs struct {
	name string
	args []string
}

const cmdDelimiter = "::"

func parseArgs(args []string) []cmdArgs {
	if len(args) == 0 {
		return nil
	}
	var cmds []cmdArgs
	var cmd cmdArgs
	for _, arg := range args {
		switch {
		case arg == cmdDelimiter:
			cmds = append(cmds, cmd)
			cmd = cmdArgs{}
		case cmd.name == "":
			cmd.name = arg
		default:
			cmd.args = append(cmd.args, arg)
		}
	}
	cmds = append(cmds, cmd)
	return cmds
}

func runCmds(cmds []cmdArgs) {
	wg := &sync.WaitGroup{}
	for i, args := range cmds {
		wg.Add(1)
		go func(nr int, args cmdArgs) {
			log.Printf("Starting cmd %d: %s %s\n",
				nr,
				args.name,
				strings.Join(args.args, " "))
			cmd := exec.Command(args.name, args.args...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				log.Printf("cmd %d with error = %s\n", err)

			}
			wg.Done()
		}(i+1, args)
	}
	wg.Wait()
}

func main() {
	flag.Parse()
	cmds := parseArgs(flag.Args())
	runCmds(cmds)
}

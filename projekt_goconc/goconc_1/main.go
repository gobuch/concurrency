// goconc ermöglicht es mehrere Befehle
// nebenläufig auszuführen.
//
// CLI Logik:
// goconc befehl1 arg1 arg2 :: befehl2 -flag arg1 arg2
package main

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

func main() {

}

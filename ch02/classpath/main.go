package classpath

import "fmt"

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("Version 0.0.1\n")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
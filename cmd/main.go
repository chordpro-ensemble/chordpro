package main

import (
	"fmt"

	"github.com/chris-skud/chordpro2/chordpro"
	"github.com/chris-skud/chordpro2/outputs/pdf"

	"github.com/spf13/cobra"
)

var processName = cmdutil.GetProcessName()
var rootCmd = &cobra.Command{
	Use:     fmt.Sprintf("%v <command> <subcommand> [flags]", processName),
	Short:   "Slack command-line tool",
	Example: `auth login`,
	Long:    `✨ CLI to create, build, and deploy Slack apps`}

func main() {

	cp := chordpro.Processor{
		Formatter: &pdf.Formatter{},
	}
	cp.Parse()
	fmt.Println("hi")
}

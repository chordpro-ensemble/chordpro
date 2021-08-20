package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/chris-skud/chordpro2/chordpro"
	"github.com/chris-skud/chordpro2/outputs/pdf"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "chordpro2",
		Short: "Convert chordpro formatted files to formatted, printable versions",
		Run:   Run,
	}
)

func Run(cmd *cobra.Command, args []string) {
	// cp := chordpro.Processor{
	// 	Formatter: &pdf.Formatter{},
	// }

	// file, err := ioutil.ReadFile("examples/simple.cho")
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// cp.Process(file)

	processor := chordpro.NewProcessor(&pdf.Processor{})
	r := bufio.NewReader(strings.NewReader("[D]This is a [C+] good song\n[C(maj7)]It is"))
	processor.Process(r)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

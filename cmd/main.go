package main

import (
	"fmt"
	"io/ioutil"
	"log"

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
	cp := chordpro.Processor{
		Formatter: &pdf.Formatter{},
	}

	file, err := ioutil.ReadFile("examples/simple.cho")
	if err != nil {
		log.Fatal(err.Error())
	}

	cp.Process(file)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

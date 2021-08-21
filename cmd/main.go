package main

import (
	"bufio"
	"fmt"
	"os"
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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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
	f, err := os.Create("simple.pdf")
	check(err)

	defer f.Close()
	// w := bufio.NewWriter(f)
	err = processor.Process(r, f)
	check(err)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

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

	songFile string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run(cmd *cobra.Command, args []string) {
	processor := chordpro.NewProcessor(&pdf.Processor{})

	b, err := os.ReadFile(songFile)
	check(err)

	r := bufio.NewReader(bytes.NewReader(b))
	f, err := os.Create("simple.pdf")
	check(err)
	defer f.Close()

	// the processor will write to the writer
	err = processor.Process(r, f)
	check(err)
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&songFile, "song", "v", "examples/simple.cho", "path to song file")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

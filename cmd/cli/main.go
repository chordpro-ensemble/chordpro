package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/chris-skud/chordpro2/chordpro"
	"github.com/chris-skud/chordpro2/chordpro/types"
	"github.com/chris-skud/chordpro2/formatters/pdf"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "chordpro2",
		Short: "Convert chordpro formatted files to formatted, printable versions",
		Run:   Run,
	}

	chordproFile string
	outputFile   string
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run(cmd *cobra.Command, args []string) {
	processor := chordpro.NewProcessor(&pdf.Processor{
		Config: types.Config{
			LyricsOnly: false,
		},
	})

	b, err := os.ReadFile(chordproFile)
	check(err)

	r := bufio.NewReader(bytes.NewReader(b))
	f, err := os.Create(outputFile)
	check(err)
	defer f.Close()

	// the processor will write to the writer
	err = processor.Process(r, f)
	check(err)
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&chordproFile, "song", "s", "examples/simple.cho", "--song=~/awesome_song.cho")
	rootCmd.PersistentFlags().StringVarP(&outputFile, "outputFile", "o", "examples/simple.pdf", "--outputFile=~/awesome_song.pdf")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

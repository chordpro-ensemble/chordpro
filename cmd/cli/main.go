package main

import (
	"bufio"
	"bytes"
	"encoding/json"
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
	configFile   string
	lyricsOnly   bool
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Run(cmd *cobra.Command, args []string) {
	cfg := types.Config{}
	if configFile != "" {
		b, err := os.ReadFile(configFile)
		check(err)
		check(json.Unmarshal(b, &cfg))
	}

	processor := chordpro.NewProcessor(&pdf.Processor{
		Config: cfg,
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
	rootCmd.PersistentFlags().StringVarP(&configFile, "configFile", "c", "", "--configFile=~/config.json")
	rootCmd.PersistentFlags().BoolVarP(&lyricsOnly, "lyricsOnly", "l", false, "--lyricsOnly=true")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

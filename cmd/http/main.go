package main

import (
	"bufio"
	"bytes"
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/chordpro-ensemble/formatters/raw"
	"github.com/chordpro-ensemble/pkg/parse"
	"github.com/gorilla/mux"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		p := raw.Processor{}
		cp := parse.NewProcessor(&p)

		reader := bufio.NewReader(bytes.NewReader(
			[]byte("{title: hi}\n[A]This is a song"),
		))
		cp.Process(reader, w)
	})

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // block

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

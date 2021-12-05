package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"github.com/rclancey/podcast"
)

func main() {
	p, err := podcast.NewPodcast(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	enc := json.NewEncoder(os.Stderr)
	enc.SetIndent("", "  ")
	enc.Encode(p.GetFeed())
	os.Stderr.Write([]byte("\n\n"))
	ep := p.Latest()
	enc.Encode(ep)
	os.Stderr.Write([]byte("\n\n"))
	r, err := ep.Reader()
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}
	r.Close()
}

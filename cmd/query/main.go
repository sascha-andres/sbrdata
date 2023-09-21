package main

import (
	"errors"
	"github.com/sascha-andres/sbrdata"
	"log"
	"strings"

	"github.com/sascha-andres/reuse/flag"
)

var (
	collectionFile, number, name string
)

func init() {
	log.SetPrefix("[SBR_QUERY] ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)

	flag.SetEnvPrefix("SBR_QUERY")
	flag.StringVar(&collectionFile, "collection-file", "", "pass name/path of collection file")
	flag.StringVar(&number, "number", "", "pass number to query")
	flag.StringVar(&name, "name", "", "pass name to query")
}

func main() {
	flag.Parse()

	err := run()
	if err != nil {
		log.Fatalf("error running query: %s", err)
	}
}

func run() error {
	if collectionFile == "" {
		return errors.New("you have to provide collection file")
	}
	if number == "" && name == "" {
		return errors.New("you have to provide either number or name")
	}

	c, err := sbrdata.LoadCollection(collectionFile)
	if err != nil {
		return err
	}
	for i := range c.Calls {
		if strings.Contains(c.Calls[i].Number, number) || strings.Contains(c.Calls[i].ContactName, name) {
			log.Printf("Call: %v", c.Calls[i])
		}
	}
	return nil
}

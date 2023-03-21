package main

import (
	"encoding/xml"
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/sascha-andres/flag"
	"github.com/sascha-andres/sbrdata"
)

var (
	collectionFile, callFile, messageFile string
	backup, verbose                       bool
)

func main() {
	log.SetPrefix("[SBR_COLLECTION] ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)

	flag.SetEnvPrefix("SBR_COLLECTION")
	flag.StringVar(&collectionFile, "collection-file", "", "pass name/path of collection file")
	flag.StringVarWithoutEnv(&callFile, "call-file", "", "pass name/path of call file")
	flag.StringVarWithoutEnv(&messageFile, "message-file", "", "pass name/path of message file")
	flag.BoolVarWithoutEnv(&verbose, "verbose", false, "print more information")
	flag.BoolVar(&backup, "backup", false, "do a backup of the file")
	flag.Parse()

	err := run()
	if err != nil {
		log.Fatalf("error running collection: %s", err)
	}
}

func run() error {
	var coll *sbrdata.Collection
	var err error
	if collectionFile == "" {
		return errors.New("you have to provide collection file")
	}
	if _, err = os.Stat(collectionFile); err != nil {
		if !errors.Is(err, fs.ErrNotExist) {
			return err
		}
		coll = &sbrdata.Collection{
			Calls: make([]sbrdata.Call, 0, 0),
			SMS:   make([]sbrdata.SMS, 0, 0),
			MMS:   make([]sbrdata.MMS, 0, 0),
		}
		err = nil
	} else {
		log.Printf("using %q as collection file", collectionFile)
		coll, err = sbrdata.LoadCollection(collectionFile)
	}
	if err != nil {
		return err
	}
	if verbose {
		coll.SetVerbose()
	}
	if backup {
		coll.SetBackup()
	}
	if _, err = os.Stat(messageFile); err == nil {
		log.Printf("using %q as message file", messageFile)
		var messages sbrdata.Messages
		var data []byte
		data, err = os.ReadFile(messageFile)
		if err != nil {
			return err
		}
		err = xml.Unmarshal(data, &messages)
		if err != nil {
			return err
		}
		if err = coll.AddMessages(messages); err != nil {
			return err
		}
	} else {
		log.Printf("no message file or error: %s", err)
	}
	if _, err = os.Stat(callFile); err == nil {
		log.Printf("using %q as call file", callFile)
		var calls sbrdata.Calls
		var data []byte
		data, err = os.ReadFile(callFile)
		if err != nil {
			return err
		}
		err = xml.Unmarshal(data, &calls)
		if err != nil {
			return err
		}
		if err = coll.AddCalls(calls); err != nil {
			return err
		}
	} else {
		log.Printf("no call file or error: %s", err)
	}
	return coll.Save(collectionFile)
}

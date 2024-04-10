package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"log"
	"os"
	"path"

	"github.com/sascha-andres/sbrdata/v2"

	"github.com/sascha-andres/reuse/flag"
)

var (
	baseDirectory, callFile, messageFile string
	backup, verbose, configFile          bool
	groupPeriod                          uint
)

// config is a type representing a configuration struct.
// It contains two fields: GroupPeriod and Backup.
//
// GroupPeriod is an unsigned integer of size 8 bits.
// It is used to represent a group period in the configuration.
//
// Backup is a boolean field indicating whether a backup is enabled in the configuration.
//
// Example usage:
//
//	var cfg config
//	cfg.GroupPeriod = sbrdata.GroupPeriod(5)
//	cfg.Backup = true
type config struct {
	// GroupPeriod is a type declaration for an unsigned integer of size 8 bits.
	// It is used to represent a group period in a configuration.
	GroupPeriod uint
	//
	Backup bool
}

// main is the entry point of the program.
// It sets the prefix and flags for the logger, parses the command-line flags,
// and calls the run function to perform the collection process.
// If there is an error returned by run, it logs a fatal error and terminates the program.
func main() {
	log.SetPrefix("[SBR_COLLECTION_V2] ")
	log.SetFlags(log.LstdFlags | log.LUTC | log.Lshortfile)

	flag.SetEnvPrefix("SBR_COLLECTION_V2")
	flag.StringVar(&baseDirectory, "base-directory", "", "pass path of data directory")
	flag.UintVar(&groupPeriod, "group-period", 99, "use 0 for no grouping, 1 for monthly and 2 for yearly")
	flag.StringVarWithoutEnv(&callFile, "call-file", "", "pass name/path of call file")
	flag.StringVarWithoutEnv(&messageFile, "message-file", "", "pass name/path of message file")
	flag.BoolVarWithoutEnv(&verbose, "verbose", false, "print more information")
	flag.BoolVar(&backup, "backup", false, "do a backup of the file")
	flag.BoolVar(&configFile, "use-config", false, "provide to use config file, named sbr.config located in collection dir")
	flag.Parse()

	err := run()
	if err != nil {
		log.Fatalf("error running collection: %s", err)
	}
}

// run is a function that performs a series of actions to process and save grouped collections.
// It returns an error if any of the actions fail.
// The function checks if the base directory is empty and returns an error if it is.
// It then creates a grouped collection by calling NewGroupedCollection function with the necessary options.
// If a message file exists, it reads the file and unmarshals the data into a Messages struct.
// It then adds the messages to the grouped collection using the AddMessages method.
// If a call file exists, it reads the file and unmarshals the data into a Calls struct.
// It then adds the calls to the grouped collection using the AddCalls method.
// Finally, it saves the grouped collection by calling the Save method.
// The function logs the file being used for messages and calls, or any error related to the files.
// `baseDirectory`, `callFile`, and `messageFile` are package-level variables used in the function.
func run() error {
	if baseDirectory == "" {
		return errors.New("you have to provide collection file")
	}

	if configFile {
		configFilePath := path.Join(baseDirectory, "sbr.config")
		data, err := os.ReadFile(configFilePath)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
		if err != nil && errors.Is(err, os.ErrNotExist) {
			cfg := config{
				GroupPeriod: groupPeriod,
				Backup:      backup,
			}
			data, err := json.MarshalIndent(cfg, "", "  ")
			if err != nil {
				return err
			}
			err = os.WriteFile(configFilePath, data, 0660)
			if err != nil {
				return err
			}
		}
		if err == nil {
			var cfg config
			err = xml.Unmarshal(data, &cfg)
			if err != nil {
				return err
			}
			backup = cfg.Backup
			groupPeriod = cfg.GroupPeriod
		}
	}

	// create grouped collection
	opts := make([]sbrdata.GroupedCollectionOption, 2)
	opts[0] = sbrdata.SetBaseDirectory(baseDirectory)
	opts[1] = sbrdata.SetGroupPeriod(sbrdata.GroupPeriod(groupPeriod))
	if verbose {
		opts = append(opts, sbrdata.SetVerbose())
	}
	if backup {
		opts = append(opts, sbrdata.SetBackup())
	}
	gc, err := sbrdata.NewGroupedCollection(opts...)
	if err != nil {
		return err
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
		if err = gc.AddMessages(messages); err != nil {
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
		if err = gc.AddCalls(calls); err != nil {
			return err
		}
	} else {
		log.Printf("no call file or error: %s", err)
	}
	return gc.Save()
}

package sbrdata

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type GroupPeriod uint8

const (
	// NoGrouping represents a constant value of groupPeriod that indicates there is no grouping.
	// It is defined as GroupPeriod(0) and is used to differentiate between different grouping options.
	NoGrouping = GroupPeriod(0)
	// GroupMonthly represents a constant value of groupPeriod that indicates monthly grouping.
	// It is defined as GroupPeriod(1) and is used to differentiate between different grouping options.
	GroupMonthly = GroupPeriod(1)
	// GroupYearly represents a constant value of groupPeriod that indicates yearly grouping.
	// It is defined as GroupPeriod(2) and is used to differentiate between different grouping options.
	GroupYearly = GroupPeriod(2)
)

const noGroupingMapKey = "collection"

// GroupedCollection saves calls/messages grouped by a period of time
type GroupedCollection struct {
	// groupPeriod is either monthly, yearly or none. If none it behaves exactly as like a single Collection
	// if set it will look for multiple collections
	groupPeriod GroupPeriod
	// baseDirectory is where the collections are persisted. For monthly, it will be yyyy/mm.json, yearly yyyy.json and none collection.json
	baseDirectory string
	// verbose controls verbosity
	verbose bool
	// backup controls whether a backup file is created
	backup bool
	// collections holds possible collections
	collections map[string]*Collection
}

// AddMessages will add all messages (SMS and MMS) to collection which are not yet known
func (gc *GroupedCollection) AddMessages(messages MessageData) error {
	for _, message := range messages.GetMms() {
		var c *Collection
		var err error
		c, err = gc.getCollection(message.GetDate())
		if err != nil {
			return err
		}
		if c == nil {
			log.Printf("could not load collection for grouping %d: %q", gc.groupPeriod, err)
		} else {
			err = c.AddMms(message)
		}
		if err != nil {
			log.Printf("could not add mms: %s", err)
		}
	}
	for _, message := range messages.GetSms() {
		var c *Collection
		var err error
		c, err = gc.getCollection(message.GetDate())
		if c == nil {
			log.Printf("could not load collection for grouping %d: %q", gc.groupPeriod, err)
		} else {
			err = c.AddSms(message)
		}
		if err != nil {
			log.Printf("could not add mms: %s", err)
		}
	}
	return nil
}

// getCollection returns a Collection object for the given date.
// If the GroupedCollection has no grouping period, it retrieves the Collection with an empty key.
// If the GroupedCollection has a grouping period, it creates a key based on the given date
// and retrieves the Collection with that key. If the Collection does not exist, it creates a new one.
// Returns the Collection object and any error occurred during retrieval or creation.
func (gc *GroupedCollection) getCollection(date string) (*Collection, error) {
	var c *Collection
	var err error
	var key string
	if gc.groupPeriod == NoGrouping {
		key = ""
	} else {
		d := createTimeFromStringUnixEpoch(date)
		key, err = gc.createKeyAndDirectoryStructure(d)
		if err != nil {
			return nil, err
		}
	}
	c, err = gc.Get(key)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// createKeyAndDirectoryStructure creates a key for the given time and creates the corresponding directory structure
// for the grouped collection. If groupPeriod is GroupMonthly, it creates a directory with the year and month as the name.
// It returns the generated key and any error encountered during the directory creation.
func (gc *GroupedCollection) createKeyAndDirectoryStructure(d time.Time) (string, error) {
	key := d.Format("2006")
	if gc.groupPeriod == GroupMonthly {
		dir := path.Join(gc.baseDirectory, key)
		if gc.verbose {
			log.Printf("potentially creating %q", dir)
		}
		err := os.MkdirAll(dir, 0770)
		if err != nil {
			return "", err
		}
		key = d.Format("2006/01")
	}
	return key, nil
}

// createTimeFromStringUnixEpoch converts a string representation of a Unix epoch timestamp
// to a time.Time value. It takes a string parameter `date` and returns a time.Time value.
// If the conversion fails, it logs an error message and returns the zero value of time.Time.
// The `date` string should be a valid integer representation of a Unix epoch timestamp.
// The returned time.Time value represents the corresponding time in Unix format.
func createTimeFromStringUnixEpoch(date string) time.Time {
	i, err := strconv.Atoi(date)
	if err != nil {
		log.Printf("could not convert %q to int", date)
	}
	d := time.Unix(int64(i)/1000, 0)
	return d
}

// Save saves all the collections in the GroupedCollection to the file system.
// It iterates over each collection in the collections map and saves it to a JSON file.
// The file path is constructed using the baseDirectory and the key of the collection.
// If an error occurs during the saving process, it is returned.
// If all collections are successfully saved, it returns nil.
func (gc *GroupedCollection) Save() error {
	for key, v := range gc.collections {
		if gc.groupPeriod == GroupMonthly {
			parts := strings.Split(key, "/")
			if len(parts) != 2 {
				log.Printf("expected key of format yyyy/mm, got: %s", key)
				continue
			}
			err := os.MkdirAll(path.Join(gc.baseDirectory, parts[0]), 0700)
			if err != nil {
				return err
			}
		}
		if v != nil {
			err := v.Save(path.Join(gc.baseDirectory, fmt.Sprintf("%s.json", key)))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// AddCalls will add all calls to collection which are not yet known
func (gc *GroupedCollection) AddCalls(calls CallsData) error {
	for _, call := range calls.GetCalls() {
		var c *Collection
		var err error
		c, err = gc.getCollection(call.GetDate())
		if c == nil {
			log.Printf("could not load collection for grouping %d: %q", gc.groupPeriod, err)
		} else {
			err = c.addCalls(call)
		}
		if err != nil {
			log.Printf("could not add call: %s", err)
		}
	}
	return nil
}

// Keys will return a slice of strings containing all the keys in the GroupedCollection's collections map.
func (gc *GroupedCollection) Keys() []string {
	keys := make([]string, 0)
	for k, _ := range gc.collections {
		keys = append(keys, k)
	}
	return keys
}

// AllCalls returns all the calls from the GroupedCollection by iterating
// over the keys and appending calls from each collection to the result slice.
// It returns the result slice of calls and any error encountered during the process.
func (gc *GroupedCollection) AllCalls() ([]Call, error) {
	var result []Call
	for _, key := range gc.Keys() {
		data, err := gc.getCollection(key)
		if err != nil {
			return nil, err
		}
		result = append(result, data.Calls...)
	}
	return result, nil
}

// AllMms returns all MMS messages in the GroupedCollection
func (gc *GroupedCollection) AllMms() ([]MMS, error) {
	var result []MMS
	for _, key := range gc.Keys() {
		data, err := gc.getCollection(key)
		if err != nil {
			return nil, err
		}
		result = append(result, data.Mms...)
	}
	return result, nil
}

// AllSms returns all SMS messages from the GroupedCollection
// by iterating over the collection's keys and retrieving the SMS messages.
// It returns a slice of SMS messages and an error in case of any failure.
func (gc *GroupedCollection) AllSms() ([]SMS, error) {
	var result []SMS
	for _, key := range gc.Keys() {
		data, err := gc.getCollection(key)
		if err != nil {
			return nil, err
		}
		result = append(result, data.Sms...)
	}
	return result, nil
}

// CustomGrouped groups the calls, SMS, and MMS based on the provided key functions and returns a map with the grouped collections.
// Beware: expensive as it iterates over all keys anr returns all data in a map of keys
func (gc *GroupedCollection) CustomGrouped(key KeyFuncs) (map[string]*Collection, error) {
	coll := &Collection{
		Calls: make([]Call, 0),
		Sms:   make([]SMS, 0),
		Mms:   make([]MMS, 0),
	}

	for k, v := range gc.collections {
		err := coll.AddSms(v.Sms...)
		if err != nil {
			return nil, fmt.Errorf("error appending Sms for key %q: %w", k, err)
		}
		err = coll.AddMms(v.Mms...)
		if err != nil {
			return nil, fmt.Errorf("error appending Mms for key %q: %w", k, err)
		}
		err = coll.addCalls(v.Calls...)
		if err != nil {
			return nil, fmt.Errorf("error appending calls for key %q: %w", k, err)
		}
	}

	result := make(map[string]*Collection)
	for _, call := range coll.Calls {
		key, err := key.Call(call)
		if err != nil {
			return nil, err
		}
		if _, ok := result[key]; !ok {
			result[key] = &Collection{
				Calls: make([]Call, 0),
				Sms:   make([]SMS, 0),
				Mms:   make([]MMS, 0),
			}
		}
		result[key].Calls = append(result[key].Calls, call)
	}
	for _, sms := range coll.Sms {
		key, err := key.SMS(sms)
		if err != nil {
			return nil, err
		}
		if _, ok := result[key]; !ok {
			result[key] = &Collection{
				Calls: make([]Call, 0),
				Sms:   make([]SMS, 0),
				Mms:   make([]MMS, 0),
			}
		}
		result[key].Sms = append(result[key].Sms, sms)
	}
	for _, mms := range coll.Mms {
		key, err := key.MMS(mms)
		if err != nil {
			return nil, err
		}
		if _, ok := result[key]; !ok {
			result[key] = &Collection{
				Calls: make([]Call, 0),
				Sms:   make([]SMS, 0),
				Mms:   make([]MMS, 0),
			}
		}
		result[key].Mms = append(result[key].Mms, mms)
	}
	return result, nil
}

// Get retrieves a specific collection based on the provided key.
// If the groupPeriod is NoGrouping, the key must be empty. It returns an error
// if a key is provided.
// If the key is empty, it returns an error.
// If the key exists in the collections map, it checks if the value is nil.
// If the value is nil, it loads the collection from the file system using the
// baseDirectory and key, and stores it in the collections map.
// If the key doesn't exist in the collections map, it returns an error.
// Finally, it returns the collection associated with the key and nil error.
//
// This method returns an error if any file or directory operation fails.
func (gc *GroupedCollection) Get(key string) (*Collection, error) {
	if gc.groupPeriod == NoGrouping {
		if key != "" {
			return nil, errors.New("without grouping no key must be provided")
		}
		key = noGroupingMapKey
	}
	if key == "" {
		return nil, errors.New("key must be provided")
	}
	if v, ok := gc.collections[key]; ok {
		if v == nil {
			coll, err := LoadCollection(path.Join(gc.baseDirectory, fmt.Sprintf("%s.json", key)))
			if err != nil {
				return nil, err
			}
			if gc.verbose {
				coll.SetVerbose()
			}
			if gc.backup {
				coll.SetBackup()
			}
			gc.collections[key] = coll
		}
	} else {
		c := &Collection{
			Key:     key,
			Calls:   make([]Call, 0),
			Sms:     make([]SMS, 0),
			Mms:     make([]MMS, 0),
			verbose: gc.verbose,
			backup:  gc.backup,
		}
		gc.collections[key] = c
	}
	return gc.collections[key], nil
}

// initializeCollections initializes the collections based on the groupPeriod value.
// If the groupPeriod is NoGrouping, it adds a key of "no" to the collections map with a value of nil.
// If the groupPeriod is GroupMonthly, it calls the initializeMonthlyGroupedCollections method.
// If the groupPeriod is GroupYearly, it calls the initializeYearlyGroupedCollections method.
// If the groupPeriod is not any of the defined values, it returns an error with message "no such grouping available".
// This method returns an error if any file or directory operation fails.
func (gc *GroupedCollection) initializeCollections() error {
	if gc.groupPeriod == NoGrouping {
		gc.collections[noGroupingMapKey] = nil
		return nil
	}
	if gc.groupPeriod == GroupMonthly {
		return gc.initializeMonthlyGroupedCollections()
	}
	if gc.groupPeriod == GroupYearly {
		return gc.initializeYearlyGroupedCollections()
	}
	return errors.New("no such grouping available")
}

// initializeYearlyGroupedCollections initializes the yearly grouped collections by reading the items in the base directory.
// It iterates over each item, and if the item is not a directory, has a name of length 9, and has a ".json" extension,
// it adds a key of "yyyy" to the collections map with a value of nil.
// This method returns an error if any file or directory operation fails.
func (gc *GroupedCollection) initializeYearlyGroupedCollections() error {
	items, err := os.ReadDir(gc.baseDirectory)
	if err != nil {
		return err
	}
	for _, item := range items {
		if !item.IsDir() && len(item.Name()) == 9 && strings.HasSuffix(item.Name(), ".json") {
			gc.collections[item.Name()[:4]] = nil
		}
	}
	return nil
}

// initializeMonthlyGroupedCollections initializes the monthly grouped collections by reading the directories and files
// within the base directory. It iterates over the items in the base directory, and if the item is a directory with a
// name of length 4, it reads the sub-items in that directory. For each sub-item that is not a directory, has a name
// of length 7, and has a ".json" extension, it adds a key of "yyyy/mm" to the collections map with a value of nil.
// This method returns an error if any file or directory operation fails.
func (gc *GroupedCollection) initializeMonthlyGroupedCollections() error {
	items, err := os.ReadDir(gc.baseDirectory)
	if err != nil {
		return err
	}
	for _, item := range items {
		if item.IsDir() && len(item.Name()) == 4 { // TODO check for 4 digit number
			// get json files in directory that are mm.json and add key yyyy/mm to dict
			subItems, err := os.ReadDir(path.Join(gc.baseDirectory, item.Name()))
			if err != nil {
				return err
			}
			for _, subItem := range subItems {
				if !subItem.IsDir() && len(subItem.Name()) == 7 && strings.HasSuffix(subItem.Name(), ".json") {
					gc.collections[fmt.Sprintf("%s/%s", item.Name(), subItem.Name()[:3])] = nil
				}
			}
		}
	}
	return nil
}

// GroupedCollectionOption is a function type used to modify the configuration of a GroupedCollection struct.
// It takes a pointer to a GroupedCollection and returns an error if the modification fails.
type GroupedCollectionOption func(gc *GroupedCollection) error

// SetBaseDirectory sets the base directory where the collections are persisted.
// It takes a string parameter `baseDirectory` and returns a `GroupedCollectionOption` function.
// The `baseDirectory` must be non-empty, otherwise an error will be returned.
// The returned `GroupedCollectionOption` function sets the `baseDirectory` field of the `GroupedCollection` struct.
func SetBaseDirectory(baseDirectory string) GroupedCollectionOption {
	return func(gc *GroupedCollection) error {
		if strings.Trim(baseDirectory, " \t") == "" {
			return errors.New("baseDirectory must be non empty")
		}
		gc.baseDirectory = baseDirectory
		return nil
	}
}

// SetGroupPeriod sets the grouping period for the GroupedCollection.
// It takes a parameter period of type GroupPeriod and returns a function of type GroupedCollectionOption.
// The period must be one of 0, 1 or 2, otherwise an error will be returned.
// The returned GroupedCollectionOption function sets the groupPeriod field of the GroupedCollection struct.
func SetGroupPeriod(period GroupPeriod) GroupedCollectionOption {
	return func(gc *GroupedCollection) error {
		if uint8(period) > 2 {
			return errors.New("period must be one of 0, 1 or 2")
		}
		gc.groupPeriod = period
		return nil
	}
}

// SetVerbose sets the `verbose` field of the `GroupedCollection` struct to true.
// It takes no parameters and returns a `GroupedCollectionOption` function.
// The returned `GroupedCollectionOption` function modifies the `verbose` field of the `GroupedCollection` struct.
// The `verbose` field controls the verbosity of the collection.
// If `verbose` is set to true, detailed information will be printed during the collection process.
// If `verbose` is set to false, minimal information will be printed.
// The `verbose` field is initially set to false.
// Example usage:
//
//	option := SetVerbose()
//	err := option(&gc)
//	Note: `gc` refers to an instance of the `GroupedCollection` struct.
func SetVerbose() GroupedCollectionOption {
	return func(gc *GroupedCollection) error {
		gc.verbose = true
		return nil
	}
}

// SetBackup sets the backup flag to true for a GroupedCollection struct.
// It takes no parameters and returns a GroupedCollectionOption function.
// The returned GroupedCollectionOption function sets the backup field of the GroupedCollection struct to true.
func SetBackup() GroupedCollectionOption {
	return func(gc *GroupedCollection) error {
		gc.backup = true
		return nil
	}
}

// NewGroupedCollection creates a new grouped collection
func NewGroupedCollection(opts ...GroupedCollectionOption) (*GroupedCollection, error) {
	gc := &GroupedCollection{}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		err := opt(gc)
		if err != nil {
			return nil, err
		}
	}
	gc.collections = make(map[string]*Collection)
	if _, err := os.Stat(gc.baseDirectory); os.IsNotExist(err) {
		err := os.MkdirAll(gc.baseDirectory, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	return gc, gc.initializeCollections()
}

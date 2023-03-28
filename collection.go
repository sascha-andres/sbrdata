package sbrdata

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/exp/slices"
)

// Collection is a container for all calls/messages that can be filled delta like
type Collection struct {
	// Calls
	Calls []Call
	// SMS data
	SMS []SMS
	// MMS data
	MMS []MMS
	// verbose controls verbosity
	verbose bool
	// backup controls whether a backup file is created
	backup bool
}

// LoadCollection loads a collection of communication data from a file
func LoadCollection(path string) (*Collection, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var coll Collection
	err = json.Unmarshal(data, &coll)
	return &coll, err
}

// Save will store all data in the collection
func (c *Collection) Save(path string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}
	if c.backup {
		err = c.doBackup(path)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(path, data, 0600)
}

// doBackup creates a backup of the collection. file.ext => file.timestamp.ext
func (c *Collection) doBackup(path string) error {
	fs, err := os.Stat(path)
	if err == nil {
		base := filepath.Dir(path)
		ext := filepath.Ext(path)
		name := fs.Name()
		if ext != "" {
			name = name[:len(name)-len(ext)]
		}
		dest := fmt.Sprintf("%s/%s.%d%s", base, name, time.Now().Unix(), ext)
		err = copy(path, dest, 1024)
		if err != nil {
			return err
		}
	}
	return nil
}

// AddCalls will add all calls to collection which are not yet known
func (c *Collection) AddCalls(calls CallsData) error {
	for _, call := range calls.GetCalls() {
		if !c.isKnownCall(call) {
			if c.verbose {
				log.Printf("adding call with %q on %q", call.GetContactName(), call.GetDate())
			}
			c.Calls = append(c.Calls, call)
		}
	}
	return nil
}

// isKnownCall returns true if call is already in collection
func (c *Collection) isKnownCall(call Call) bool {
	return slices.Contains(c.Calls, call)
}

// AddMessages will add SMS/MMS to collection which are not yet known
func (c *Collection) AddMessages(messages MessageData) error {
	for _, s := range messages.GetSms() {
		if !c.isKnownSMS(s) {
			if c.verbose {
				log.Printf("adding sms with %q on %q", s.GetContactName(), s.GetDate())
			}
			c.SMS = append(c.SMS, s)
		}
	}
	for _, s := range messages.GetMms() {
		if !c.isKnownMMS(s) {
			if c.verbose {
				log.Printf("adding sms mms %q on %q", s.GetContactName(), s.GetDate())
			}
			c.MMS = append(c.MMS, s)
		}
	}
	return nil
}

// isKnownSMS scans collection for SMS and returns true if found
func (c *Collection) isKnownSMS(sms SMS) bool {
	return slices.Contains(c.SMS, sms)
}

// isKnownMMS scans collection for MMS and returns true if found
func (c *Collection) isKnownMMS(m MMSData) bool {
	for _, s := range c.MMS {
		if s.Date == m.GetDate() && s.Address == m.GetAddress() {
			return true
		}
	}
	return false
}

// SetVerbose is used to make collection a bit more heavy on informational output
func (c *Collection) SetVerbose() {
	c.verbose = true
}

// SetBackup tells collection to make a backup on save
func (c *Collection) SetBackup() {
	c.backup = true
}

func copy(src, dst string, bufferSize int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, bufferSize)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}

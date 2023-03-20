package sbrdata

import (
	"encoding/json"
	"log"
	"os"

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
	return os.WriteFile(path, data, 0600)
}

// AddCalls will add all calls to collection which are not yet known
func (c *Collection) AddCalls(calls Calls) error {
	for _, call := range calls.Call {
		if !c.isKnownCall(call) {
			log.Printf("adding call with %q on %q", call.ContactName, call.Date)
			c.Calls = append(c.Calls, call)
		}
	}
	return nil
}

// isKnownCall returns true if call is already in collection
func (c *Collection) isKnownCall(call Call) bool {
	return slices.Contains(c.Calls, call)
	//for _, knownCall := range c.Calls {
	//	if knownCall.Date == call.Date && knownCall.ContactName == call.ContactName {
	//		return true
	//	}
	//}
	//return false
}

// AddMessages will add SMS/MMS to collection which are not yet known
func (c *Collection) AddMessages(messages Messages) error {
	for _, s := range messages.Sms {
		if !c.isKnownSMS(s) {
			log.Printf("adding sms with %q on %q", s.ContactName, s.Date)
			c.SMS = append(c.SMS, s)
		}
	}
	for _, s := range messages.Mms {
		if !c.isKnownMMS(s) {
			log.Printf("adding sms mms %q on %q", s.ContactName, s.Date)
			c.MMS = append(c.MMS, s)
		}
	}
	return nil
}

// isKnownSMS scans collection for SMS and returns true if found
func (c *Collection) isKnownSMS(sms SMS) bool {
	return slices.Contains(c.SMS, sms)
	//for _, s := range c.SMS {
	//	if s.Date == sms.Date && s.Address == sms.Address {
	//		return true
	//	}
	//}
	//return false
}

// isKnownMMS scans collection for MMS and returns true if found
func (c *Collection) isKnownMMS(m MMS) bool {
	for _, s := range c.MMS {
		if s.Date == m.Date && s.Address == m.Address {
			return true
		}
	}
	return false
}

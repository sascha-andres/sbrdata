package sbrdata

import "encoding/xml"

type (
	// CallsData allows retrieval of data for all calls
	CallsData interface {
		GetCount() string
		GetBackupSet() string
		GetBackupDate() string
		GetType() string
		GetCalls() []Call
	}

	// CallData is a single call
	CallData interface {
		GetNumber() string
		GetDuration() string
		GetDate() string
		GetType() string
		GetPresentation() string
		GetSubscriptionID() string
		GetPostDialDigits() string
		GetSubscriptionComponentName() string
		GetReadableDate() string
		GetContactName() string
	}
)

// Call is a single call
type Call struct {
	Number                    string `xml:"number,attr"`
	Duration                  string `xml:"duration,attr"`
	Date                      string `xml:"date,attr"`
	Type                      string `xml:"type,attr"`
	Presentation              string `xml:"presentation,attr"`
	SubscriptionID            string `xml:"subscription_id,attr"`
	PostDialDigits            string `xml:"post_dial_digits,attr"`
	SubscriptionComponentName string `xml:"subscription_component_name,attr"`
	ReadableDate              string `xml:"readable_date,attr"`
	ContactName               string `xml:"contact_name,attr"`
}

func (c Call) GetNumber() string {
	return c.Number
}

func (c Call) GetDuration() string {
	return c.Duration
}

func (c Call) GetDate() string {
	return c.Date
}

func (c Call) GetType() string {
	return c.Type
}

func (c Call) GetPresentation() string {
	return c.Presentation
}

func (c Call) GetSubscriptionID() string {
	return c.SubscriptionID
}

func (c Call) GetPostDialDigits() string {
	return c.GetPostDialDigits()
}

func (c Call) GetSubscriptionComponentName() string {
	return c.SubscriptionComponentName
}

func (c Call) GetReadableDate() string {
	return c.ReadableDate
}

func (c Call) GetContactName() string {
	return c.ContactName
}

// Calls contains all
type Calls struct {
	XMLName    xml.Name `xml:"calls"`
	Count      string   `xml:"count,attr"`
	BackupSet  string   `xml:"backup_set,attr"`
	BackupDate string   `xml:"backup_date,attr"`
	Type       string   `xml:"type,attr"`
	Call       []Call   `xml:"call"`
}

func (c Calls) GetCount() string {
	return c.Count
}

func (c Calls) GetBackupSet() string {
	return c.BackupSet
}

func (c Calls) GetBackupDate() string {
	return c.BackupDate
}

func (c Calls) GetType() string {
	return c.Type
}

func (c Calls) GetCalls() []Call {
	if c.Call == nil {
		return make([]Call, 0, 0)
	}
	return c.Call
}

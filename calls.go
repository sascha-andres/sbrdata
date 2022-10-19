package sbrdata

import "encoding/xml"

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

// Calls contains all
type Calls struct {
	XMLName    xml.Name `xml:"calls"`
	Count      string   `xml:"count,attr"`
	BackupSet  string   `xml:"backup_set,attr"`
	BackupDate string   `xml:"backup_date,attr"`
	Type       string   `xml:"type,attr"`
	Call       []Call   `xml:"call"`
}

// Package template provides types the implement the xml.Marshaler interfaces for HC Templates.
package template

import (
	"encoding/xml"
	"strings"
)

// CommaDelimited is a slice of string that marshals into an XML comma-separated string.
type CommaDelimited []string

// SpaceDelimited is a slice of string that marshals into an XML space-separated string.
type SpaceDelimited []string

// Values is a map of strings that marshals into an XML key/value pair.
type Values map[string]string

// MarshalXML implements the xml.Marshaler interface for v. Each key/value pair
// renders an XML element. There is no outer VALUE element.
func (v Values) MarshalXML(e *xml.Encoder, _ xml.StartElement) error {
	for k, value := range v {
		key := strings.ToUpper(k)

		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: key}}); err != nil {
			return err
		}

		if err := e.EncodeToken(xml.CharData(value)); err != nil {
			return err
		}

		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: key}}); err != nil {
			return err
		}
	}

	return e.Flush()
}

// MarshalXML implements the xml.Marshaler interface for s. The slice is joined
// into a single comma-separated string.
func (s CommaDelimited) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strings.Join(s, ","), start)
}

// MarshalXML implements the xml.Marshaler interface for s. The slice is joined
// into a single space-separated string.
func (s SpaceDelimited) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(strings.Join(s, " "), start)
}

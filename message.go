package rfc5424

import "time"

// Message represents a log message as defined by RFC-5424
// (https://tools.ietf.org/html/rfc5424)
type Message struct {
	Priority       int
	Timestamp      time.Time
	Hostname       string
	AppName        string
	ProcessID      string
	MessageID      string
	StructuredData []StructuredData
	Message        []byte
}

// SDParam represents parameters for structured data
type SDParam struct {
	Name  string
	Value string
}

// StructuredData represents structured data within a log message
type StructuredData struct {
	ID         string
	Parameters []SDParam
}

// AddParam adds the given name and value to this structured data's parameters
func (sd *StructuredData) AddParam(name, value string) {
	sd.Parameters = append(sd.Parameters, SDParam{Name: name, Value: value})
}

// AddDatum adds structured data to a log message
func (m *Message) AddDatum(ID string, Name string, Value string) {
	if m.StructuredData == nil {
		m.StructuredData = []StructuredData{}
	}
	for i, sd := range m.StructuredData {
		if sd.ID == ID {
			sd.Parameters = append(sd.Parameters, SDParam{Name: Name, Value: Value})
			m.StructuredData[i] = sd
			return
		}
	}

	m.StructuredData = append(m.StructuredData, StructuredData{
		ID: ID,
		Parameters: []SDParam{
			SDParam{
				Name:  Name,
				Value: Value,
			},
		},
	})
}

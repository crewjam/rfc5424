package GoFuzz

import (
	"bytes"

	"github.com/crewjam/rfc5424"
)

func FuzzReadFrom(data []byte) int {
	m := rfc5424.Message{}
	_, err := m.ReadFrom(bytes.NewReader(data))
	if err != nil {
		return -1
	}
	return 1
}

func FuzzUnmarshalBinary(data []byte) int {
	m := rfc5424.Message{}
	err := m.UnmarshalBinary(data)
	if err != nil {
		return -1
	}
	return 1
}

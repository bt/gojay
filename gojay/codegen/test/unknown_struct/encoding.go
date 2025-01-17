// Code generated by Gojay. DO NOT EDIT.

package unknown_struct

import (
	"github.com/francoispqt/gojay"
)

// MarshalJSONObject implements MarshalerJSONObject
func (m *Message) MarshalJSONObject(enc *gojay.Encoder) {
	enc.IntKey("Id", m.Id)
	enc.StringKey("Name", m.Name)
}

// IsNil checks if instance is nil
func (m *Message) IsNil() bool {
	return m == nil
}

// UnmarshalJSONObject implements gojay's UnmarshalerJSONObject
func (m *Message) UnmarshalJSONObject(dec *gojay.Decoder, k string) error {

	switch k {
	case "Id":
		return dec.Int(&m.Id)

	case "Name":
		return dec.String(&m.Name)

	default:
		return gojay.MakeUnknownFieldErr(m, k)

	}
	return nil
}

// NKeys returns the number of keys to unmarshal
func (m *Message) NKeys() int {
	return 0 // Unmarshal all keys so we can catch trailing unknown keys.
}

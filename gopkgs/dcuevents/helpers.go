package dcuevents

import (
	"encoding/json"
	"fmt"
)

// EncodeEvent takes a given struct, and encodes it to JSON
// returning a byte format
func EncodeEvent(payload interface{}) ([]byte, error) {
	res, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error encoding payload: %v", err)
	}
	return res, nil
}

// DecodeEvent takes a given JSON payload, and a pointer to a struct
// and decodes it
func DecodeEvent(payload []byte, dec interface{}) error {
	err := json.Unmarshal(payload, dec)
	if err != nil {
		return fmt.Errorf("error decoding payload: %v", err)
	}
	return nil
}

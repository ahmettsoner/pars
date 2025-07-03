package base64

import (
	b64 "encoding/base64"
)

func Encode(in []byte) (string, error) {
	return b64.StdEncoding.EncodeToString(in), nil
}

func Decode(in string) ([]byte, error) {
	o, err := b64.StdEncoding.DecodeString(in)
	if err != nil {
		o, err = b64.URLEncoding.DecodeString(in)
		if err != nil {
			return nil, err
		}
	}
	return o, nil
}

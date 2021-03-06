// Copyright 2018-2019 The vogo Authors. All rights reserved.
// author: wongoo
// since: 2018/12/27
//

package zkclient

import (
	"encoding/json"
	"io"
	"reflect"
)

// JSONCodec struct
type JSONCodec struct {
	typ reflect.Type
}

// Encode zk json encode
func (c *JSONCodec) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// Decode zk json decode
func (c *JSONCodec) Decode(data []byte) (interface{}, error) {
	value := reflect.New(c.typ)

	if len(data) == 0 {
		return value.Elem(), io.EOF
	}

	obj := value.Interface()

	if err := json.Unmarshal(data, obj); err != nil {
		return value, err
	}

	return value.Interface(), nil
}

var (
	jsonEncodeCodec = &JSONCodec{}
)

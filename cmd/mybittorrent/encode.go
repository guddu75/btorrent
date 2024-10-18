package main

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
)

type bendcoder struct {
	*bytes.Buffer
}

func (b bendcoder) encode(val interface{}) (err error) {
	switch v := val.(type) {
	case string:
		b.WriteString(fmt.Sprintf("%d:%s", len(v), v))
		return nil
	case int:
		b.WriteString(fmt.Sprintf("i%de", v))
		return nil
	case []interface{}:
		b.WriteByte('l')
		for _, item := range v {
			b.encode(item)
		}
		b.WriteByte('e')
		return nil
	case map[string]interface{}:
		keys := make([]string, 0, len(v))
		for k := range v {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		b.WriteByte('d')

		for _, k := range keys {
			b.encode(k)
			b.encode(v[k])
		}
		b.WriteByte('e')
		return nil
	default:
		return errors.New("unsupported bencode type")
	}
}

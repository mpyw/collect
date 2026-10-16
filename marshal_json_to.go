package collect

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
)

// MarshalJSONTo writes the values to enc as a JSON array, as each is read.
// It stands for Laravel's toJson. json.Marshal calls it, with enc's options, in
// encoding/json/v2 and in encoding/json.
func (c Collection[T]) MarshalJSONTo(enc *jsontext.Encoder) error {
	if err := enc.WriteToken(jsontext.BeginArray); err != nil {
		return err
	}
	for v := range c.Seq() {
		if err := json.MarshalEncode(enc, v); err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndArray)
}

// MarshalJSONTo writes the pairs to enc as a JSON object, in order, as each
// is read. It is an object even when the keys are 0, 1, 2 and so on, where
// PHP would write an array.
//
// A key that encodes as a JSON string, with enc's options, is the name. A
// number or a boolean is quoted, as json.Marshal quotes the keys of a map.
// Any other key is an error that wraps [ErrUnsupportedKey]. A key that
// comes twice is an error in encoding/json/v2, unless enc allows it with
// jsontext.AllowDuplicateNames. encoding/json allows it.
func (c KeyedCollection[K, V]) MarshalJSONTo(enc *jsontext.Encoder) error {
	if err := enc.WriteToken(jsontext.BeginObject); err != nil {
		return err
	}
	for k, v := range c.Seq() {
		name, err := marshalJSONToName(k, enc.Options())
		if err != nil {
			return err
		}
		if err := enc.WriteValue(name); err != nil {
			return err
		}
		if err := json.MarshalEncode(enc, v); err != nil {
			return err
		}
	}
	return enc.WriteToken(jsontext.EndObject)
}

// marshalJSONToName returns k as the name of a JSON object member: a JSON
// string.
func marshalJSONToName[K any](k K, opts json.Options) (jsontext.Value, error) {
	raw, err := json.Marshal(k, opts)
	if err != nil {
		return nil, err
	}
	switch name := jsontext.Value(raw); name.Kind() {
	case '"':
		return name, nil
	case '0', 't', 'f':
		return jsontext.Value(`"` + string(raw) + `"`), nil
	default:
		return nil, fmt.Errorf("%w: a key of type %T is not a JSON string, number or boolean: %s", ErrUnsupportedKey, k, raw)
	}
}

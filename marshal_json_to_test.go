package collect_test

import (
	"bytes"
	"encoding/json"
	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
	"errors"
	"fmt"
	"testing"

	"github.com/mpyw/collect"
)

func ExampleCollection_MarshalJSONTo() {
	b, err := jsonv2.Marshal(collect.Of(user{"Alice", 31}, user{"Bob", 19}))

	fmt.Println(string(b), err)
	// Output: [{"Name":"Alice","Age":31},{"Name":"Bob","Age":19}] <nil>
}

func ExampleKeyedCollection_MarshalJSONTo() {
	b, err := jsonv2.Marshal(collect.Of("x", "y").Keyed())

	fmt.Println(string(b), err)
	// Output: {"0":"x","1":"y"} <nil>
}

func TestMarshalJSONTo(t *testing.T) {
	for _, tt := range []struct {
		v    any
		want string
	}{
		{collect.Of(1, 2), `[1,2]`},
		{collect.Of[int](), `[]`},
		{collect.Of(user{"Alice", 31}), `[{"Name":"Alice","Age":31}]`},
		{keyedOf([]string{"b", "a"}, []int{2, 1}), `{"b":2,"a":1}`},
		{collect.Of("x", "y").Keyed(), `{"0":"x","1":"y"}`},
		{keyedOf([]bool{true}, []int{1}), `{"true":1}`},
	} {
		got, err := jsonv2.Marshal(tt.v)
		if err != nil || string(got) != tt.want {
			t.Errorf("json/v2 Marshal() = %s, %v, want %s", got, err, tt.want)
		}
		got, err = json.Marshal(tt.v)
		if err != nil || string(got) != tt.want {
			t.Errorf("json Marshal() = %s, %v, want %s", got, err, tt.want)
		}
	}
}

func TestMarshalJSONToErrors(t *testing.T) {
	dup := keyedOf([]string{"a", "a"}, []int{1, 2})
	if _, err := jsonv2.Marshal(dup); err == nil {
		t.Error("a key twice is not an error")
	}
	if got, err := jsonv2.Marshal(dup, jsontext.AllowDuplicateNames(true)); err != nil || string(got) != `{"a":1,"a":2}` {
		t.Errorf("AllowDuplicateNames: %s, %v", got, err)
	}
	type point struct{ X int }
	if _, err := jsonv2.Marshal(keyedOf([]point{{1}}, []int{1})); err == nil {
		t.Error("an object key is not an error")
	}
	if _, err := jsonv2.Marshal(collect.Of(func() {})); err == nil {
		t.Error("a func value is not an error")
	}
	if _, err := jsonv2.Marshal(collect.Of(1).Keyed().Map(func(int, int) func() { return nil })); err == nil {
		t.Error("a func value of a pair is not an error")
	}
	if _, err := jsonv2.Marshal(keyedOf([]chan int{nil}, []int{1})); err == nil {
		t.Error("a chan key is not an error")
	}
}

// An encoder that expects an object's name rejects a collection.
func TestMarshalJSONToEncoderErrors(t *testing.T) {
	for name, v := range map[string]any{
		"Collection":      collect.Of(1),
		"KeyedCollection": collect.Of(1).Keyed(),
	} {
		var b bytes.Buffer
		enc := jsontext.NewEncoder(&b)
		if err := enc.WriteToken(jsontext.BeginObject); err != nil {
			t.Fatal(err)
		}
		if err := jsonv2.MarshalEncode(enc, v); err == nil {
			t.Errorf("a %s as a name is not an error", name)
		}
	}
}

// The encoder's options apply to the keys too.
func TestMarshalJSONToKeyOptions(t *testing.T) {
	type id struct{ N int }
	byID := jsonv2.WithMarshalers(jsonv2.MarshalFunc(func(v id) ([]byte, error) {
		return []byte(fmt.Sprintf(`"id-%d"`, v.N)), nil
	}))
	got, err := jsonv2.Marshal(keyedOf([]id{{7}}, []int{1}), byID)
	if err != nil || string(got) != `{"id-7":1}` {
		t.Errorf("Marshal() = %s, %v", got, err)
	}
	type point struct{ X int }
	if _, err := jsonv2.Marshal(keyedOf([]point{{1}}, []int{1})); !errors.Is(err, collect.ErrUnsupportedKey) {
		t.Errorf("an object key: %v", err)
	}
}

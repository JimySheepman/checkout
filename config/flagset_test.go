package config

import (
	"flag"
	"fmt"
	"testing"

	"github.com/magiconair/properties"
	"github.com/stretchr/testify/require"
)

func TestKvValue_newKVValue(t *testing.T) {

	var dictionary map[string]string
	var pDictionary map[string]string
	var expected *kvValue = new(kvValue)

	actual := newKVValue(dictionary, &pDictionary)

	require.Equal(t, expected, actual)
}

func TestKvValue_kvParse(t *testing.T) {

	tests := []struct {
		name     string
		input    string
		expected kvValue
	}{
		{
			name:     "if there is a semicolon at the end, it adds an empty",
			input:    "k1=v1;k2=v2;",
			expected: kvValue{"": "", "k1": "v1", "k2": "v2"},
		},
		{
			name:     "succeed",
			input:    "k1=v1;k2=v2",
			expected: kvValue{"k1": "v1", "k2": "v2"},
		},
		{
			name:     "check the key length",
			input:    "k1;k2=v2",
			expected: kvValue{"k1": "", "k2": "v2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := kvParse(tt.input)

			require.Equal(t, tt.expected, actual)
		})
	}
}

// kvParse to kvString and kvString to kvParse  not equal
func TestKvValue_kvString(t *testing.T) {
	tests := []struct {
		name     string
		input    kvValue
		expected string
	}{
		{
			name:     "succeed",
			input:    kvValue{"k1": "v1", "k2": "v2"},
			expected: "k1=v1;k2=v2",
		},
		{
			name:     "it adds an empty",
			input:    kvValue{"": "", "k1": "v1", "k2": "v2"},
			expected: "=;k1=v1;k2=v2",
		},
		{
			name:     "check the key length",
			input:    kvValue{"k1": "", "k2": "v2"},
			expected: "k1=;k2=v2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := kvString(tt.input)

			require.Equal(t, tt.expected, actual)
		})
	}
}

func TestKvValue_Set(t *testing.T) {
	v := kvValue{}

	input := "k1=v1;k2=v2"
	expected := kvValue{"k1": "v1", "k2": "v2"}

	v.Set(input)

	require.Equal(t, expected, v)
}

func TestKvValue_Get(t *testing.T) {
	v := kvValue{"k1": "v1", "k2": "v2"}

	expected := kvValue{"k1": "v1", "k2": "v2"}

	v.Get()

	require.Equal(t, expected, v)
}

func TestKvValue_String(t *testing.T) {
	v := kvValue{"k1": "v1", "k2": "v2"}

	expected := "k1=v1;k2=v2"

	actual := v.String()

	require.Equal(t, expected, actual)
}

func Test_newKVSliceValue(t *testing.T) {

	var dictionaries []map[string]string
	var pDictionaries []map[string]string
	var expected *kvSliceValue = new(kvSliceValue)

	actual := newKVSliceValue(dictionaries, &pDictionaries)

	require.Equal(t, expected, actual)
}

func TestKvSliceValue_Set(t *testing.T) {

	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "succeed",
			input: "k1=v1,k1=v1;k2=v2,k1=v1;k2=v2;k3=v3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := kvSliceValue{}

			actual := v.Set(tt.input)
			if actual != nil {
				require.Error(t, actual)
			}
		})
	}
}

func TestKvSliceValue_Get(t *testing.T) {
	v := kvSliceValue{{"k1": "v1", "k2": "v2"}}

	expected := kvSliceValue{{"k1": "v1", "k2": "v2"}}

	v.Get()

	require.Equal(t, expected, v)
}

func TestKvSliceValue_String(t *testing.T) {
	v := kvSliceValue{{"k1": "v1", "k2": "v2"}, {"k1": "v1", "k2": "v2"}}

	expected := "k1=v1;k2=v2,k1=v1;k2=v2"

	actual := v.String()

	require.Equal(t, expected, actual)
}

func TestStringSliceValue_newStringSliceValue(t *testing.T) {
	var words []string
	var pWords []string
	var expected *stringSliceValue = new(stringSliceValue)

	actual := newStringSliceValue(words, &pWords)

	require.Equal(t, expected, actual)
}

func TestStringSliceValue_Set(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "succeed",
			input: "k1,v1,k2,v2",
		},
		{
			name:  "succeed",
			input: "k1",
		},
		{
			name:  "succeed",
			input: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			v := stringSliceValue{}

			actual := v.Set(tt.input)

			fmt.Println(v)
			if actual != nil {
				require.Error(t, actual)
			}
		})
	}
}

func TestStringSliceValue_Get(t *testing.T) {
	v := stringSliceValue{"k1", "v1", "k2", "v2"}

	expected := stringSliceValue{"k1", "v1", "k2", "v2"}

	v.Get()

	require.Equal(t, expected, v)
}

func TestStringSliceValue_String(t *testing.T) {
	v := stringSliceValue{"k1", "v1", "k2", "v2"}

	expected := "k1,v1,k2,v2"

	actual := v.String()

	require.Equal(t, expected, actual)
}

func TestFlagSet_NewFlagSet(t *testing.T) {
	var name string
	var errorHandling flag.ErrorHandling

	var expected *FlagSet = &FlagSet{
		set: make(map[string]bool),
	}

	actual := NewFlagSet(name, errorHandling)

	require.Equal(t, expected, actual)
}

func TestFlagSet_IsSet(t *testing.T) {
	f := FlagSet{
		set: map[string]bool{"test": true},
	}
	name := "test"

	expected := true

	actual := f.IsSet(name)

	require.Equal(t, expected, actual)
}

func TestFlagSet_KVVar(t *testing.T) {
	p := &map[string]string{"": "", "k1": "v1", "k2": "v2"}
	name := "test"
	value := map[string]string{"value": "value"}
	usage := "usage"

	f := FlagSet{}
	e := FlagSet{}

	e.KVVar(p, name, value, usage)
	f.KVVar(p, name, value, usage)

	require.Equal(t, e, f)
}

func TestFlagSet_KVSliceVar(t *testing.T) {
	p := &[]map[string]string{{"": "", "k1": "v1", "k2": "v2"}}
	name := "test"
	value := []map[string]string{{"value": "value"}}
	usage := "usage"

	f := FlagSet{}
	e := FlagSet{}

	e.KVSliceVar(p, name, value, usage)
	f.KVSliceVar(p, name, value, usage)

	require.Equal(t, e, f)
}

func TestFlagSet_StringSliceVar(t *testing.T) {
	p := &[]string{"", "k1", "v1", "k2", "v2"}
	name := "test"
	value := []string{"value", "value"}
	usage := "usage"

	f := FlagSet{}
	e := FlagSet{}

	e.StringSliceVar(p, name, value, usage)
	f.StringSliceVar(p, name, value, usage)

	require.Equal(t, e, f)
}

func TestFlagSet_ParseFlags(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		environ  []string
		prefixes []string
		p        *properties.Properties
	}{
		{
			name:     "Parse error",
			args:     []string{"--test1", "=test2"},
			environ:  []string{"test1=test2"},
			prefixes: []string{},
		},
		{
			// TODO: Need help writing tests for f.Visit ,f.VisitAll
			name:     "succeed",
			args:     []string{"test1", "test2"},
			environ:  []string{"test1=test2", "k1=v1;k2=v2"},
			prefixes: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FlagSet{
				FlagSet: flag.FlagSet{},
				set:     make(map[string]bool),
			}

			err := f.ParseFlags(
				tt.args,
				tt.environ,
				tt.prefixes,
				tt.p,
			)

			if err != nil {
				require.Error(t, err)
			}
		})
	}
}

// Package parser reads the webgpu.yml specification and provides
// Go structs representing the full WebGPU C API surface.
package parser

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Spec is the top-level representation of webgpu.yml.
type Spec struct {
	Copyright string     `yaml:"copyright"`
	Name      string     `yaml:"name"`
	Doc       string     `yaml:"doc"`
	Constants []Constant `yaml:"constants"`
	Typedefs  []Typedef  `yaml:"typedefs"`
	Enums     []Enum     `yaml:"enums"`
	Bitflags  []Bitflag  `yaml:"bitflags"`
	Structs   []Struct   `yaml:"structs"`
	Functions []Function `yaml:"functions"`
	Objects   []Object   `yaml:"objects"`
	Callbacks []Callback `yaml:"callbacks"`

	callbackMap map[string]*Callback
}

// Constant represents a named constant value.
type Constant struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
	Doc   string `yaml:"doc"`
}

// Typedef represents a type alias.
type Typedef struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
	Doc  string `yaml:"doc"`
}

// Enum represents a C enum.
type Enum struct {
	Name    string        `yaml:"name"`
	Doc     string        `yaml:"doc"`
	Entries EnumEntryList `yaml:"entries"`
}

// EnumEntryList preserves YAML `null` placeholders so enum values match the C ABI.
type EnumEntryList []EnumEntry

func (l *EnumEntryList) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.SequenceNode {
		return fmt.Errorf("enum entries: expected sequence, got kind %d", value.Kind)
	}
	out := make(EnumEntryList, 0, len(value.Content))
	for _, item := range value.Content {
		var e EnumEntry
		if err := e.UnmarshalYAML(item); err != nil {
			return err
		}
		out = append(out, e)
	}
	*l = out
	return nil
}

// EnumEntry is a single entry in an enum.
type EnumEntry struct {
	Name             string   `yaml:"name"`
	Doc              string   `yaml:"doc"`
	Value            *int     `yaml:"value"`
	ValueCombination []string `yaml:"value_combination"`
	IsNull           bool     `yaml:"-"`
}

func (e *EnumEntry) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind == yaml.ScalarNode && (value.Value == "null" || value.Tag == "!!null") {
		e.IsNull = true
		return nil
	}
	type rawEntry struct {
		Name             string   `yaml:"name"`
		Doc              string   `yaml:"doc"`
		Value            *int     `yaml:"value"`
		ValueCombination []string `yaml:"value_combination"`
	}
	var r rawEntry
	if err := value.Decode(&r); err != nil {
		return err
	}
	e.Name = r.Name
	e.Doc = r.Doc
	e.Value = r.Value
	e.ValueCombination = r.ValueCombination
	return nil
}

// Bitflag represents a set of bitflags.
type Bitflag struct {
	Name    string         `yaml:"name"`
	Doc     string         `yaml:"doc"`
	Entries []BitflagEntry `yaml:"entries"`
}

// BitflagEntry is a single flag value.
type BitflagEntry struct {
	Name             string   `yaml:"name"`
	Doc              string   `yaml:"doc"`
	Value            *int     `yaml:"value"`
	ValueCombination []string `yaml:"value_combination"`
}

// Struct represents a C struct.
type Struct struct {
	Name        string         `yaml:"name"`
	Doc         string         `yaml:"doc"`
	Type        string         `yaml:"type"`
	Extends     []string       `yaml:"extends"`
	FreeMembers bool           `yaml:"free_members"`
	Members     []StructMember `yaml:"members"`
}

// StructMember is a single field of a struct.
type StructMember struct {
	Name     string `yaml:"name"`
	Doc      string `yaml:"doc"`
	Type     string `yaml:"type"`
	Default  any    `yaml:"default"`
	Optional bool   `yaml:"optional"`
	Pointer  string `yaml:"pointer"`
}

// Function represents a C function declaration.
type Function struct {
	Name     string        `yaml:"name"`
	Doc      string        `yaml:"doc"`
	Category string        `yaml:"category"`
	Returns  *ReturnType   `yaml:"returns"` // object with doc, type, passed_with_ownership
	Args     []FunctionArg `yaml:"args"`
	Callback string        `yaml:"callback"` // callback type for async methods
	Object   string        `yaml:"object"`
}

// ReturnType describes the return value of a function.
type ReturnType struct {
	Doc                 string `yaml:"doc"`
	Type                string `yaml:"type"`
	Pointer             string `yaml:"pointer"` // immutable | mutable
	PassedWithOwnership bool   `yaml:"passed_with_ownership"`
	Optional            bool   `yaml:"optional"`
}

// FunctionArg is a single argument of a function.
type FunctionArg struct {
	Name                string `yaml:"name"`
	Doc                 string `yaml:"doc"`
	Type                string `yaml:"type"`
	Optional            bool   `yaml:"optional"`
	Pointer             string `yaml:"pointer"`
	Default             any    `yaml:"default"`
	PassedWithOwnership bool   `yaml:"passed_with_ownership"`
}

// Object represents a WebGPU object (handle type).
type Object struct {
	Name    string     `yaml:"name"`
	Doc     string     `yaml:"doc"`
	Methods []Function `yaml:"methods"`
}

// Callback represents a callback function type.
type Callback struct {
	Name string        `yaml:"name"`
	Doc  string        `yaml:"doc"`
	Args []CallbackArg `yaml:"args"`
}

// CallbackArg is a single argument of a callback.
type CallbackArg struct {
	Name    string `yaml:"name"`
	Doc     string `yaml:"doc"`
	Type    string `yaml:"type"`
	Pointer string `yaml:"pointer"` // immutable | mutable
}

// ParseFile reads and parses a webgpu.yml file.
func ParseFile(path string) (*Spec, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	return Parse(data)
}

// Parse parses webgpu.yml from raw bytes.
func Parse(data []byte) (*Spec, error) {
	var spec Spec
	if err := yaml.Unmarshal(data, &spec); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}
	spec.buildCallbackMap()
	return &spec, nil
}

func (s *Spec) buildCallbackMap() {
	s.callbackMap = make(map[string]*Callback, len(s.Callbacks))
	for i := range s.Callbacks {
		s.callbackMap[s.Callbacks[i].Name] = &s.Callbacks[i]
	}
}

// GetCallback returns the callback definition by name, or nil.
func (s *Spec) GetCallback(name string) *Callback {
	return s.callbackMap[name]
}

// LookupFunction returns a function by name from either top-level functions or object methods.
func (s *Spec) LookupFunction(name string) *Function {
	for i := range s.Functions {
		if s.Functions[i].Name == name {
			return &s.Functions[i]
		}
	}
	for i := range s.Objects {
		for j := range s.Objects[i].Methods {
			if s.Objects[i].Methods[j].Name == name {
				return &s.Objects[i].Methods[j]
			}
		}
	}
	return nil
}

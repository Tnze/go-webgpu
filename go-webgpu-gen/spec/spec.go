package spec

import "strings"

type Spec struct {
	PackageName string `yaml:"-"`
	ImportPath  string `yaml:"-"`

	Copyright  string `yaml:"copyright"`
	Name       string `yaml:"name"`
	Doc        string `yaml:"doc"`
	EnumPrefix uint16 `yaml:"enum_prefix"`

	Constants []Constant `yaml:"constants"`
	Typedefs  []Typedef  `yaml:"typedefs"`
	Enums     []Enum     `yaml:"enums"`
	Bitflags  []Bitflag  `yaml:"bitflags"`
	Structs   []Struct   `yaml:"structs"`
	Callbacks []Callback `yaml:"callbacks"`
	Functions []Function `yaml:"functions"`
	Objects   []Object   `yaml:"objects"`
}

type Base struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
	Doc       string `yaml:"doc"`
	Extended  bool   `yaml:"extended"`
}

type Constant struct {
	Base  `yaml:",inline"`
	Value string `yaml:"value"`
}

type Typedef struct {
	Base `yaml:",inline"`
	Type string `yaml:"type"`
}

type Enum struct {
	Base    `yaml:",inline"`
	Entries []*EnumEntry `yaml:"entries"`
}
type EnumEntry struct {
	Base  `yaml:",inline"`
	Value *uint16 `yaml:"value"`
}

type Bitflag struct {
	Base    `yaml:",inline"`
	Entries []BitflagEntry `yaml:"entries"`
}
type BitflagEntry struct {
	Base             `yaml:",inline"`
	Value            string   `yaml:"value"`
	ValueCombination []string `yaml:"value_combination"`
}

type PointerType string

const (
	PointerTypeMutable   PointerType = "mutable"
	PointerTypeImmutable PointerType = "immutable"
)

type ParameterType struct {
	Name                string      `yaml:"name"`
	Namespace           string      `yaml:"namespace"`
	Doc                 string      `yaml:"doc"`
	Type                string      `yaml:"type"`
	PassedWithOwnership *bool       `yaml:"passed_with_ownership"`
	Pointer             PointerType `yaml:"pointer"`
	Optional            bool        `yaml:"optional"`
	Default             *string     `yaml:"default"`
}

type Callback struct {
	Base  `yaml:",inline"`
	Style string          `yaml:"style"`
	Args  []ParameterType `yaml:"args"`
}

type Function struct {
	Base     `yaml:",inline"`
	Returns  *ParameterType  `yaml:"returns"`
	Callback *string         `yaml:"callback"`
	Args     []ParameterType `yaml:"args"`
}

type Struct struct {
	Base        `yaml:",inline"`
	Type        string          `yaml:"type"`
	FreeMembers bool            `yaml:"free_members"`
	Members     []ParameterType `yaml:"members"`
	Extends     []string        `yaml:"extends"`
}

type Object struct {
	Base    `yaml:",inline"`
	Methods []Function `yaml:"methods"`

	IsStruct bool `yaml:"-"`
}

func (s *Spec) GetDef(typeName string) any {
	i := strings.IndexByte(typeName, '.')
	if i == -1 {
		return nil
	}
	name := typeName[i+1:]

	switch typeName[:i] {
	case "struct":
		for i := range s.Structs {
			if s.Structs[i].Name == name {
				return &s.Structs[i]
			}
		}
	case "callback":
		for i := range s.Callbacks {
			if s.Callbacks[i].Name == name {
				return &s.Callbacks[i]
			}
		}
	}
	return nil
}

func (s *Spec) IsPureData(input *Struct) bool {
	for _, field := range input.Members {
		if strings.HasPrefix(field.Type, "array<") ||
			strings.HasPrefix(field.Type, "object.") {
			return false
		}
		if strings.HasPrefix(field.Type, "struct.") {
			if !s.IsPureData(s.GetDef(field.Type).(*Struct)) {
				return false
			}
		}
		switch field.Type {
		case "nullable_string", "string_with_default_empty", "out_string":
			return false
		}
	}
	return true
}

package main

import "github.com/Tnze/go-webgpu/gen/parser"

// ---------------------------------------------------------------------------
// Template data model for the low-level webgpu package.
// ---------------------------------------------------------------------------

// TemplateData is the top-level data passed to every low-level template.
type TemplateData struct {
	Spec          *parser.Spec
	Enums         []EnumData
	Bitflags      []BitflagData
	Structs       []StructData
	Funcs         []FuncData
	Handles       []HandleData
	Consts        []ConstData
	CallbackInfos []CallbackInfoData
	ErrorTypes    []ErrorTypeEntryData
}

// ConstData describes a named constant.
type ConstData struct {
	Name  string
	Value string
	Doc   string
}

// HandleData describes a WebGPU opaque object type.
type HandleData struct {
	Name  string
	CName string
	Doc   string
}

// EnumData describes a Go enum type.
type EnumData struct {
	Name    string
	CName   string
	Doc     string
	Entries []EnumEntryData
}

// EnumEntryData describes a single enum value.
type EnumEntryData struct {
	Name  string
	CName string
	Value string
	Doc   string
}

// BitflagData describes a Go bitflag type.
type BitflagData struct {
	Name    string
	CName   string
	Doc     string
	Entries []BitflagEntryData
}

// BitflagEntryData is a single flag value.
type BitflagEntryData struct {
	Name  string
	CName string
	Value string
	Doc   string
}

// StructData describes a Go struct (C layout).
type StructData struct {
	Name        string
	CName       string
	Type        string // "extensible" | "extensible_callback_arg" | "extension" | "standalone"
	Doc         string
	Members     []StructMemberData
	HasDefault  bool
	FreeMembers bool
}

// StructMemberData is a single struct field.
type StructMemberData struct {
	Name       string // Go PascalCase field name
	GoType     string
	TypeRef    string // original YAML type reference
	CName      string
	Doc        string
	HasDefault bool
	DefaultVal string
	IsPtr      bool // pointer: immutable/mutable
	IsOptional bool
	IsSlice    bool // array<T> data pointer field
	IsArray    bool
	ArrayLen   int
}

// FuncData describes one generated function.
//
// The low-level package mirrors the C API one-to-one: async calls take a raw
// by-value callback-info record and return a Future; status enums are returned
// as-is. Blocking wrappers and error folding live in the gpu package.
type FuncData struct {
	Name         string
	CName        string
	Ident        string // unique Go identifier for internal symbols (proc vars)
	GoArgs       []FuncArgData
	GoReturn     string
	ReturnRef    string // original YAML return type reference
	RetKind      string // handle | status | pointer | wait | future | bool | value | none
	OpName       string
	Doc          string
	IsMethod     bool
	IsRelease    bool
	ObjName      string // receiver type name, e.g. "Device"
	HasCallback  bool   // takes a raw callback-info and returns Future
	CallbackFn   string // e.g. RequestAdapterFn (consumed by the gpu wrapper)
	CallbackName string // e.g. request_adapter
	// Precomputed function bodies, one line each (no leading tab).
	CgoBody    []string
	SyscallBody []string
}

// FuncArgData is a function argument.
type FuncArgData struct {
	Name        string
	GoType      string
	TypeRef     string
	CName       string
	IsPtr       bool
	IsOptional  bool
	IsSlice     bool
	IsStr       bool
	IsCBInfo    bool     // by-value WGPUNnnCallbackInfo parameter (async APIs)
	HasChain    bool     // struct has a NextInChain field
	SliceFields []string // slice members needing pin
	StrFields   []string // string members (pin unsafe.StringData)
	PtrFields   []string // pointer members (pin pointee)
	SliceCPtr   string   // cgo pointer type for slice data
}

// CallbackInfoData describes a WGPUNnnCallbackInfo record and its trampoline.
type CallbackInfoData struct {
	Name            string // DeviceLostCallbackInfo
	CName           string // WGPUDeviceLostCallbackInfo
	Doc             string
	FnType          string // DeviceLostFn
	HasMode         bool   // true except uncaptured_error
	BaseName        string // RequestAdapter
	ExportName      string // goRequestAdapterCB
	CFnType         string // WGPURequestAdapterCallback
	OpName          string
	Args            []CallbackArgInfo
	HasTrampoline   bool
	OneShot         bool     // true for async ops; false for persistent hooks
	FnParams        []string // pure Go callback signature, C argument order
	StatusArg       *CallbackArgInfo
	MessageArg      *CallbackArgInfo
	ErrorTypeArg    *CallbackArgInfo
	ErrorMessageArg *CallbackArgInfo
	SuccessConst    string
	CallArgs        []string
	ErrConvert      string
}

// CallbackArgInfo describes one C parameter of a callback trampoline.
type CallbackArgInfo struct {
	YAMLName   string
	Name       string // Go parameter name for the C value (c_type)
	GoName     string // converted Go value name (typeVal)
	CGoType    string // cgo type in //export
	ExternType string // C type in extern decl
	GoType     string
	PureType   string // webgpu-level type (no error folding): ErrorType, string, Adapter, …
	Kind       string // status | message | error_type | error_msg | object | enum | struct_ptr | string | value
	IsStatus   bool
	IsMessage  bool
}

// CallbackArgData is one parameter of a Go callback func type.
type CallbackArgData struct {
	Name   string
	GoType string
}

// ErrorTypeEntryData is one distinct Go error type from WGPUErrorType.
type ErrorTypeEntryData struct {
	Name    string // ValidationError
	Value   string // ErrorTypeValidation
	Display string // "validation"
}

package main

import (
	"regexp"
	"strings"
	"unicode"
)

// ---------------------------------------------------------------------------
// Identifier and documentation helpers.
// ---------------------------------------------------------------------------

var goKeywords = map[string]bool{
	"break": true, "case": true, "chan": true, "const": true, "continue": true,
	"default": true, "defer": true, "else": true, "fallthrough": true, "for": true,
	"func": true, "go": true, "goto": true, "if": true, "import": true,
	"interface": true, "map": true, "package": true, "range": true, "return": true,
	"select": true, "struct": true, "switch": true, "type": true, "var": true,
}

// pascalCase converts a YAML identifier to PascalCase, dropping the WGPU
// prefix and expanding known acronyms (GPU, D3D12, RGBA, …).
func pascalCase(s string) string {
	if s == "" {
		return s
	}
	s = strings.TrimPrefix(s, "WGPU")

	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		switch strings.ToLower(p) {
		case "gpu":
			b.WriteString("GPU")
		case "d3d11", "d3d12":
			b.WriteString(strings.ToUpper(p))
		case "rgb", "rgba", "bgra", "rg":
			b.WriteString(strings.ToUpper(p))
		default:
			runes := []rune(p)
			runes[0] = unicode.ToUpper(runes[0])
			b.WriteString(string(runes))
		}
	}
	return b.String()
}

// camelCase is pascalCase with the first letter lowercased, escaping Go keywords.
func camelCase(s string) string {
	if s == "" {
		return s
	}
	result := lowerFirst(pascalCase(s))
	if goKeywords[result] {
		result += "Val"
	}
	return result
}

// lowerFirst lowercases the first rune of s.
func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

// enumValue is the C-style EnumName_Value form (template helper).
func enumValue(enumName, entryName string) string {
	return pascalCase(enumName) + "_" + pascalCase(entryName)
}

// receiverName is the conventional single-letter Go receiver name.
func receiverName(typeName string) string {
	if typeName == "" {
		return "x"
	}
	return strings.ToLower(typeName[:1])
}

var commentRe = regexp.MustCompile(`(?m)^`)

// cleanDoc trims a YAML doc comment and drops empty / TODO placeholders.
func cleanDoc(s string) string {
	s = strings.TrimSpace(s)
	if s == "" || strings.EqualFold(s, "TODO") {
		return ""
	}
	return s
}

// commentLine prefixes every line of s with "// " (no leading indent).
func commentLine(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	return commentRe.ReplaceAllString(s, "// ")
}

// commentIndent is commentLine with a leading tab, for use inside blocks.
func commentIndent(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return ""
	}
	return commentRe.ReplaceAllString(s, "\t// ")
}

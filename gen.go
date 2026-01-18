// Package gen provides a simple and powerful code generation toolkit.
package gen

import (
	"fmt"
	"strings"
)

// Generator provides a fluent interface for generating code in any programming language.
//
// It handles indentation and line breaks automatically
type Generator struct {
	sb            strings.Builder
	indentLevel   int
	indentString  string
	atStartOfLine bool // Tracks if the next write should be indented
}

// New creates a new Generator instance with default settings (2 spaces indentation).
// The generator starts positioned at the beginning of a line.
func New() *Generator {
	return &Generator{
		sb:            strings.Builder{},
		indentLevel:   0,
		indentString:  "  ",
		atStartOfLine: true,
	}
}

// WithSpaces configures the generator to use the specified number of spaces for each indentation level.
func (g *Generator) WithSpaces(spaces int) *Generator {
	g.indentString = strings.Repeat(" ", spaces)
	return g
}

// WithTabs configures the generator to use tabs for indentation.
func (g *Generator) WithTabs() *Generator {
	g.indentString = "\t"
	return g
}

// Indent increases the current indentation level by one.
// This affects subsequent writes that require indentation.
func (g *Generator) Indent() *Generator {
	g.indentLevel++
	return g
}

// Dedent decreases the current indentation level by one.
// This affects subsequent writes that require indentation.
func (g *Generator) Dedent() *Generator {
	if g.indentLevel > 0 {
		g.indentLevel--
	}
	return g
}

// Raw writes the given content directly to the output, bypassing
// indentation rules and automatic newline handling.
//
// If the last character of the content is a newline, the subsequent
// write will be indented if needed.
//
// Use this for pre-formatted text or specific formatting needs.
func (g *Generator) Raw(content string) *Generator {
	if content == "" {
		return g
	}
	g.sb.WriteString(content)
	if strings.HasSuffix(content, "\n") {
		g.atStartOfLine = true
	} else {
		g.atStartOfLine = false
	}
	return g
}

// Rawf formats the arguments according to the format specifier and writes
// the result directly using Raw's logic (bypassing indentation/newlines).
//
// See Raw for details on behavior and state updates.
func (g *Generator) Rawf(format string, args ...any) *Generator {
	return g.Raw(fmt.Sprintf(format, args...))
}

// Break writes a single newline character to the output.
//
// It ensures the generator is positioned at the start of a new line for subsequent writes.
func (g *Generator) Break() *Generator {
	g.sb.WriteString("\n")
	g.atStartOfLine = true
	return g
}

// Inline writes content, adding the current indentation *only if* the generator
// is currently positioned at the start of a line.
//
// It does not add a trailing newline itself.
//
// If the content string contains newlines, lines following the newline character
// within the content string will be correctly indented.
func (g *Generator) Inline(content string) *Generator {
	if content == "" {
		return g
	}

	sublines := strings.Split(content, "\n")
	for idx, subline := range sublines {
		if idx > 0 {
			g.sb.WriteString("\n")
			g.atStartOfLine = true
		}

		if subline != "" {
			if g.atStartOfLine {
				g.sb.WriteString(strings.Repeat(g.indentString, g.indentLevel))
			}
			g.sb.WriteString(subline)
			g.atStartOfLine = false
		}
	}

	// If the original input string itself ended with a newline,
	// the next write should start on a new line.
	if strings.HasSuffix(content, "\n") {
		g.atStartOfLine = true
	}

	return g
}

// Inlinef formats the arguments and writes the result using Inline's logic.
//
// See Inline for details on behavior and state updates.
func (g *Generator) Inlinef(format string, args ...any) *Generator {
	return g.Inline(fmt.Sprintf(format, args...))
}

// Line writes content with appropriate indentation (if needed) and concludes
// with a newline, ensuring the next write starts on a fresh, indented line.
//
// It's a convenient combination of Inline followed by Break.
func (g *Generator) Line(content string) *Generator {
	g.Inline(content) // Write the content, indenting if at start of line
	g.Break()         // Add the newline and set state for the next line
	return g
}

// Linef formats the arguments and writes the result using Line's logic.
//
// See Line for details on behavior and state updates.
func (g *Generator) Linef(format string, args ...any) *Generator {
	return g.Line(fmt.Sprintf(format, args...))
}

// Block executes the provided function `fn` within a temporarily increased
// indentation level. Indentation is increased before calling `fn` and
// decreased afterward, regardless of what happens inside `fn`.
//
// Think of it as a way to temporarily increase the indentation level for a
// group of related lines of code written inside the `fn` argument.
func (g *Generator) Block(fn func()) *Generator {
	g.Indent()
	fn()
	g.Dedent()
	return g
}

// String returns the final generated code as a single string.
func (g *Generator) String() string {
	return g.sb.String()
}

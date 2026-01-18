# Varavel Gen

<img src="./gopher.png" height="250px" width="auto">

<br/>

<p>
  <a href="https://github.com/varavelio/gen/actions">
    <img src="https://github.com/varavelio/gen/actions/workflows/ci.yaml/badge.svg" alt="CI status"/>
  </a>
  <a href="https://pkg.go.dev/github.com/varavelio/gen">
    <img src="https://pkg.go.dev/badge/github.com/varavelio/gen" alt="Go Reference"/>
  </a>
  <a href="https://goreportcard.com/report/varavelio/gen">
    <img src="https://goreportcard.com/badge/varavelio/gen" alt="Go Report Card"/>
  </a>
  <a href="https://github.com/varavelio/gen/releases/latest">
    <img src="https://img.shields.io/github/release/varavelio/gen.svg" alt="Release Version"/>
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/varavelio/gen.svg" alt="License"/>
  </a>
  <a href="https://github.com/varavelio/gen">
    <img src="https://img.shields.io/github/stars/varavelio/gen?style=flat&label=github+stars"/>
  </a>
</p>

Gen is a simple and powerful code generation toolkit that provides a fluent interface for generating code in any programming language.

It handles indentation and line breaks automatically, making it easy to generate clean and properly formatted code.

## Features

- Zero dependencies
- Fluent interface for easy code generation
- Customizable indentation (spaces or tabs)
- Block-based code generation
- String formatting support
- Language-agnostic design
- Raw content insertion
- Manual line breaks

## Installation

```bash
go get github.com/varavelio/gen
```

## Usage

### Basic Example

```go
g := gen.New()

g.Line("function greet(name) {")

// Functions can be optionally chained for fluent code generation
g.Indent().
    Line("console.log('Hello, ' + name)").
    Dedent().
    Line("}")

fmt.Println(g.String())
```

Output:

```javascript
function greet(name) {
  console.log("Hello, " + name);
}
```

### Custom Indentation

```go
// Use 4 spaces
g := gen.New().WithSpaces(4)

// Use tabs
g := gen.New().WithTabs()
```

### Block-based Generation

With blocks, you can automatically manage indentation.

```go
g := gen.New()
g.Line("class User {")
g.Block(func() {
    g.Line("constructor(name) {")
    g.Block(func() {
        g.Line("this.name = name")
    })
    g.Line("}")
})
g.Line("}")
```

Output:

```javascript
class User {
  constructor(name) {
    this.name = name;
  }
}
```

### Conditional Lines

```go
isAngry := true
g := gen.New()

g.Line("func saySomething() string {")
g.Block(func() {
	if isAngry {
	    g.Line(`return "I'm angry"`)
    } else {
        g.Line(`return "I'm not angry"`)
    }
})
g.Line("}")
```

Output:

```go
func saySomething() string {
    return "I'm angry"
}
```

### Formatted Lines

```go
g := gen.New()
g.Linef("const %s = %q", "greeting", "Hello, World!")
```

Output:

```javascript
const greeting = "Hello, World!";
```

### Raw Content

```go
g := gen.New()
g.Raw("console.log('Hello, World!');")
```

Output:

```javascript
console.log("Hello, World!");
```

### Manual Line Breaks

```go
g := gen.New()
g.Line("console.log('Hello')").Break().Line("console.log('World')")
```

Output:

```javascript
console.log("Hello");

console.log("World");
```

### Inline Content

```go
g := gen.New()
g.Inline("console.log('Hello')").Inline("console.log('World')")
```

Output:

```javascript
console.log('Hello')console.log('World')
```

### Generating the Code

```go
g := gen.New()
g.Line("package main")

result := g.String()
```

## Examples

### TypeScript Interface

```go
g := gen.New()
g.Line("interface User {").
    Indent().
    Line("id: string").
    Line("name: string").
    Line("age: number").
    Dedent().
    Line("}")
```

### Python Class

```go
g := gen.New().WithSpaces(4)
g.Line("class User:").
    Indent().
    Line("def __init__(self, name, age):").
    Indent().
    Line("self.name = name").
    Line("self.age = age").
    Dedent().
    Line("").
    Line("def greet(self):").
    Indent().
    Line("return f\"Hello, {self.name}!\"").
    Dedent()
```

### Go Struct

```go
g := gen.New()
g.Line("type User struct {").
    Indent().
    Line("ID   int").
    Line("Name string").
    Dedent().
    Line("}")
```

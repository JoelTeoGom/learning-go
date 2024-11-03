# Naming Conventions in Go: Writing Clean, Readable Code

🚀 Want to write clear, maintainable Go code? Follow these naming conventions to make your code not just work, but shine! 💎

## 1. Capitalization for Visibility 🔍
- **Uppercase Initial (ExportedName):** Makes functions or variables exported and accessible outside the package, like `public` in other languages.
- **Lowercase Initial (unexportedName):** Keeps them unexported, accessible only within the package.

**Example:**

```go
// Exported: visible outside the package
func PublicFunction() {}

// Unexported: only visible within the package
func privateFunction() {}
```

## 2. Naming Variables and Functions 📜
- **Short, descriptive names:** Go prefers concise, meaningful names, like n, err, or buf.
- **Avoid abbreviations:** Full words increase clarity (userCount > usrCnt).
- **CamelCase for functions and variables:** Keep it camelCase, no underscores.

**Example:**

```go
var userCount int // Good
func calculateTotal() {}
```

## 3. Package Naming 📦
- Package names are short, lowercase, and singular (e.g., fmt, strings). 
- They shouldn’t need prefixes when used, as in fmt.Println.


## 4. Naming Interfaces 🎭
- End in -er for action-based interfaces (Reader, Writer).
- Use broader names for more complex interfaces.

**Example:**
```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

## 5. Constants Naming 🔢

- Use `SNAKE_CASE` for global constants, fully uppercase.
- Local constants can be lowercase.

**Example:**

```go
const PI = 3.14159
const maxUsers = 100
```

## 6. Error Naming ⚠️

- Use `err` for error variables and an `Err` prefix for error constants to make them easily identifiable.

**Example:**

```go
var ErrNotFound = errors.New("not found")
```


## 📝 Summary of Conventions

- **Capitalization**: Use uppercase initials for exported names and lowercase initials for unexported names.
- **Short Names**: Keep names descriptive and concise.
- **Packages**: Use short, lowercase names.
- **Interfaces**: Use `-er` suffix for action-based interfaces (e.g., `Reader`, `Writer`).
- **Constants**: Use `SNAKE_CASE` for global constants and lowercase for local constants.
- **Errors**: Use `err` for error variables and `Err` prefix for error constants.


Try it out! 🚀 These conventions make Go code cleaner, simpler, and easier to work with. Let’s make Go code beautiful, one convention at a time! ✨
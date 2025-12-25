Got it 👍
Here are **ONLY the function-related notes** — **clean, minimal, code-first**, no extra topics.

---

# Go Functions (Video 10–13)

---

## Function Syntax

```go
func name(parameters) returnType {
	// body
}
```

---

## Function with Return Type

```go
func jog(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}
```

| Part              | Meaning          |
| ----------------- | ---------------- |
| `num1 int`        | parameter        |
| `int` (after `)`) | return type      |
| `return`          | sends value back |

---

## Function with No Return Type

```go
func print() {
	fmt.Println("Howitzer Impact")
}
```

| Feature          | Value        |
| ---------------- | ------------ |
| Return type      | none         |
| `return` keyword | not required |

---

## Function with Parameters

```go
func saySmtng(name string) {
	fmt.Println("Welcome to the real world,", name)
}
```

Call:

```go
saySmtng("HSEVEN")
```

---

## Function with Multiple Return Values

```go
func maxmin(num1 int, num2 int) (int, int) {
	max := num1 + num2
	min := num1 - num2
	return max, min
}
```

---

## Receiving Multiple Return Values

```go
p, q := maxmin(a, b)
```

| Variable | Receives              |
| -------- | --------------------- |
| `p`      | first returned value  |
| `q`      | second returned value |

---

## Function Call Flow

```go
func main() {
	result := jog(10, 5)
	fmt.Println(result)
}
```

---

## Naming Rules

| Rule                               | Allowed           |
| ---------------------------------- | ----------------- |
| Same function name in same package | ❌                 |
| Lowercase function                 | same package only |
| Uppercase function                 | exported          |

---

## Key Recall

| Concept              | Rule             |
| -------------------- | ---------------- |
| Return type location | After `()`       |
| Multiple return      | `(type, type)`   |
| Parameters           | `name type`      |
| No return function   | omit return type |

---



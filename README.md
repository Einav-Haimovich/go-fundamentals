# Go Fundamentals

Course code from **Go - The Complete Guide** by Maximilian Schwarzmüller. Every lesson is kept as its own snapshot folder, so each section reads as a progression from the first naive version to the final refactor.

## Sections

### 01. Go essentials

Language basics: variables, types, operators, control flow, functions, and the first contact with files and user input.

**What I learned:**
- Go's type system is static but inference-friendly — `:=` covers most declarations, explicit types only where they carry meaning.
- Errors are values returned alongside results, not thrown; the `if err != nil` check is the control flow, not an afterthought.
- Multiple return values remove the need for out-params and let a function report a result and a failure at once.

---

### 02. Packages

Splitting a program across files and packages, controlling visibility, and pulling in a third-party module.

**What I learned:**
- Exporting is decided by capitalization alone — an uppercase identifier is public, lowercase is package-private.
- Files in the same package share a namespace with no imports between them; the package, not the file, is the unit of encapsulation.
- `go mod init` / `go get` make a module explicit, and the module path becomes the import prefix.

---

### 03. Pointers

Value semantics versus reference semantics, and passing pointers to mutate data in place.

**What I learned:**
- Go passes everything by value, so a function mutates the caller's data only through a pointer.
- Pointers here are about avoiding copies and enabling mutation, not about manual memory management — the GC still owns lifetime.
- `*T` and `&x` are the whole surface; there is no pointer arithmetic to reason about.

---

### 04. Structs and custom types

Defining structs, attaching methods, constructor functions, embedding, and struct tags.

**What I learned:**
- Go has no constructors, so a `New...` function returning `(*T, error)` is the idiom for enforcing validation at creation.
- Pointer receivers mutate, value receivers copy — the receiver choice is part of the method's contract.
- Embedding composes behavior without inheritance: the outer type promotes the inner type's fields and methods.
- Struct tags are metadata read by libraries at runtime, which is how a Go struct maps onto a JSON shape.

---

### 05. Interfaces and generics

Defining interfaces, satisfying them implicitly, type switches, the empty interface, and an intro to type parameters.

**What I learned:**
- Interface satisfaction is structural and implicit — a type never declares which interfaces it implements, which lets the consumer define the abstraction.
- Small interfaces defined at the point of use keep implementations decoupled; embedded interfaces compose larger contracts from them.
- `interface{}` / `any` gives up type safety, so a type switch is needed to get it back at runtime.
- Generics solve the same problem at compile time: constrained type parameters instead of runtime assertions.

---

### 06. Arrays, slices, and maps

Fixed arrays, slices as views over them, `make`, map operations, and choosing between maps and structs.

**What I learned:**
- A slice is a window onto a backing array — reslicing shares memory, so mutations are visible to every slice over the same array.
- `make` preallocates length and capacity, which matters once `append` starts triggering reallocations.
- Maps are for dynamic, homogeneous key/value data; structs are for a fixed, heterogeneous, compile-time-known shape.

---

### 07. Functions deep dive

Functions as values, anonymous functions, closures, variadic parameters, and recursion.

**What I learned:**
- Functions are first-class values, so behavior can be passed into a transform instead of hardcoded inside it.
- A closure captures its enclosing variables by reference, which is how a factory function produces preconfigured behavior.
- Variadic parameters and the `slice...` spread are two sides of the same call convention.

---

### 08. Practice project — price calculator

A CLI tool that reads prices, applies tax rates, and writes JSON — refactored step by step from a single file into packages behind an interface.

**What I learned:**
- Extracting an `IOManager` interface let the same job run against a file or the command line without the job knowing which.
- Wrapping low-level errors at the boundary keeps the caller's error handling readable and hides the source of the failure.
- The refactor sequence itself is the lesson: get it working, find the seam, then invert the dependency.

---

### 09. Concurrency

Goroutines, channels for signalling completion, coordinating multiple channels with `select`, and `defer`.

**What I learned:**
- `go` starts a goroutine but the program exits when `main` returns, so completion has to be communicated explicitly.
- Channels are the synchronization primitive, not just a queue — a receive blocks until a send happens.
- `select` waits on several channels at once, which is what makes fan-out with independent success and error channels workable.
- `defer` guarantees cleanup runs on every exit path, including early returns.

---

### 10. REST API

Building a JSON event-booking API with Gin, SQLite, JWT auth, and route middleware.

**What I learned:**
- Routing, model, and persistence layers separate cleanly: handlers parse and respond, models own their own SQL.
- Route groups plus a middleware make authentication a property of a group of routes rather than a check repeated in every handler.
- Passwords are hashed at rest and identity travels in a signed JWT — the token carries the user id the handlers authorize against.
- Authorization is distinct from authentication: knowing who the caller is does not decide whether they own the event they're editing.

Note: this section's code under `10-rest-api/author-source/` is the author's original source, kept so the course reference is complete. Everything else in this repo is my own work from following along.

---

## How to Run

Each lesson folder is a standalone snapshot. From inside one:

```bash
cd 09-concurrency/practice-project
go mod init example.com/price-calculator   # only where no go.mod is present
go mod tidy
go run .
```

For the REST API section:

```bash
cd 10-rest-api/author-source/16-cancelling-registrations
go mod tidy
go run .
# serves on http://localhost:8080
```

## Folder Structure

```
01-go-essentials/
02-packages/
03-pointers/
04-structs/
05-interfaces/
06-arrays-slices-maps/
07-functions-deep-dive/
08-practice-project/
09-concurrency/
10-rest-api/author-source/
certificate/
```

Each section contains one folder per lesson, named after the topic it introduces.

---

Thanks to Maximilian Schwarzmüller for the course — [Go - The Complete Guide](https://www.udemy.com/course/go-the-complete-guide/).

[Certificate of completion](certificate/Go%20Fundamentals%20-%20Einav%20Haimovich.pdf)

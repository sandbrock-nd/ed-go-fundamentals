# ed-go-fundamentals
Exercise files for Go fundamentals programming

# 1. Typical directory structure of a Go program
```text
my-go-app/
|-- cmd/
|-- internal/
|-- pkg/
|-- test/
|-- vendor/
|-- go.mod
|-- README.md
```
**Root Directory**
The root of your Go application is typically named after your project. It contains the go.mod file.

**`cmd` Directory**
The cmd directory contains the main applications for your project. Each subdirectory inside cmd is named for each executable your project will generate.

**`pkg` Directory**
The pkg directory holds library code that's ok to use by external applications. The code is organized into packages.

**`internal` Directory**
The internal directory is for private application and library code. This code isn't intended to be imported by other applications.

**`test` Directory**
While Go tests typically reside alongside the code they test, some projects may choose to have a separate test directory for system tests or test data.

## 1.1. Exercise - Creating a New Project
1. Clone the repository.
2. Navigate to the directory in the command line.
3. Create a new go project: `go mod init ed-go-fundamentals`
4. Create a `cmd` directory parented to the root directory.
5. Create an `app` directory parented to the `cmd` directory.
6. Create a `main.go` file in the `cmd` directory. Populate the file with the following code:
```go
package main

func main() {
	println("Hello, world!")
}
```
7. Navigate to the app directory in the command line.
8. Run the application: `go run main.go`
9. Commit your changese:
```bash
git add .
git commit -m "Completed exercise 1.1"
```

# 2) Packages
In Go, a package is a fundamental concept that helps organize code and promote reusability. It's a collection of Go source files in the same directory that are compiled together.

## 2.1) Characteristics of Packages
1. Directory-based: A package corresponds to a single directory in the Go workspace. All Go files in the same directory are part of the same package.
2. Declaration: Each Go file in a package starts with the package keyword, followed by the package name. For instance, package math declares that the file belongs to the math package.
```go
package math
```
3. Naming: The package name is usually the same as the last segment of the import path (e.g., the package in src/myapp/utils would be utils). However, it's not a strict requirement.
4. Scope: Identifiers (like types, variables, functions) defined in a package are accessible from any file within the same package.

## 2.2) Types of Packages
1. Executable Packages: If a package is named main, it's an executable package. This means that when you build the package, it will generate an executable file.
2. Library Packages: Any package that is not main is a library package. These packages are used to provide reusable code, such as utility functions, types, constants, etc.

## 2.3) Importing Packages
You can import other packages into your Go file using the import statement. Imported packages can be standard library packages or custom packages you've created.

## 2.4) Best Practices
- Organize related code into packages for better code management.
- Keep your package's purpose focused; avoid creating "do-it-all" packages.
- Use clear, concise, and descriptive names for your packages.
- Minimize the number of exported identifiers to maintain a clean and usable package interface.

## 2.5) Gotchas
- Packages cannot share a bidirectional relationship.

## 2.6) Exercise - Determine if a year is leap year
1. Navigate to the project root directory in the command line.
2. Add a reference to the external date package: `go get github.com/rickb777/date`. Notice that new entries were added to `go.mod`` in the root directory.
3. In main.go, add the following function:
```go
func isLeapYear(year int) bool {
	return gregorian.IsLeap(year)
}
```
4. Update the `main` function to call `isLeapYear`:
```go
func main() {
	println("Hello, world!")
	println("Is 2020 a leap year?", isLeapYear(2020))
}
```
5. Run the application. Play with it by changing the year and running it.
6. Commit your changese:
```bash
git add .
git commit -m "Completed exercise 2.6"
```

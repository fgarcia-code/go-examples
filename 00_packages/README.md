# Code organization
## Repository
* A repository contains one or more modules (a go repository can be seen exactly the same thing as a git repository)

## Packages
* Go programs are organized into packages
* A *Package* is a collection of source files in the same directory that are compiled together. 
  * In the current folder we have `pkg1`, `pkg2` and `main`
* Functions, types, variables, and constants defined in one source file are visible to all other source files within same packages.

## Modules
* Using modules, developers ar no longer confined to working inside `GOPATH`, version dependency information is explicit yet lightweight, and builds are more reliable and reproducible.
* A module is a collection of related Go packages.
* Modules are the unit of source code interchange and versioning.
* Module-aware mode is active by default whenever a go.mod file is found in the current directory or in any parent directory.
* A module is defined by a tree of Go sources files with go.mod file in the three's directory. The directory containing the go.mod file is called the module root.
* Typically the module root will also correspond to a **source code repository root**
* The module is the set of all **Go packages** in the module root and its subdirectories, **but excluding subtrees with their own go.mod files.
* The **module path** is the import path prefix corresponding to the module root.
* The *go.mod* file defines the module path and lists the specific versions of other modules.
* Start a new module:
  ```bash
  go mod init example.com/m
  ```

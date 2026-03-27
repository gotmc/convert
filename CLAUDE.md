# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

`convert` is a Go utility library (`github.com/gotmc/convert`) providing functions to convert between types — primarily string-to-float parsing and string sanitization. It is a single-package library with no external dependencies.

## Build & Test Commands

```bash
# Format, vet, and test (primary check before PRs)
make check        # or: just check

# Run a single test
go test -run TestStripDoubleQuotes ./...

# Lint with staticcheck
make lint         # or: just lint

# HTML coverage report
make cover        # or: just cover
```

Both `Makefile` and `Justfile` are maintained with equivalent recipes.

## Code Structure

Single package `convert` with two source files:
- `parse.go` — package doc comment only (formerly the main file before rename from `parse` to `convert`)
- `strings.go` — all exported functions: `StripDoubleQuotes`, `StringToNFloats`, `StringToFloats`
- `strings_test.go` — table-driven tests for all functions

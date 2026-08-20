# Infrastructure As Code

![CI](https://github.com/Qyroxen/Infrastructure-As-Code/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Infrastructure-As-Code/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Infrastructure-As-Code?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Infrastructure-As-Code)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Infrastructure-As-Code)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Infrastructure-As-Code?style=social)](https://github.com/Qyroxen/Infrastructure-As-Code/stargazers)

## What is it?

Infrastructure As Code is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Infrastructure-As-Code.git
cd Infrastructure-As-Code
go build -o infrastructureascode .

# Run
./infrastructureascode --help
```

## CLI Usage

```bash
# Basic usage
./infrastructureascode

# With flags
./infrastructureascode --verbose --output json

# Get help
./infrastructureascode --help
```

## Examples

```bash
# Example 1
./infrastructureascode example1

# Example 2
./infrastructureascode example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o infrastructureascode .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Infrastructure-As-Code/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Infrastructure-As-Code?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Infrastructure-As-Code/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Infrastructure-As-Code?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Infrastructure-As-Code/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Infrastructure-As-Code" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Infrastructure-As-Code/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Infrastructure-As-Code" alt="Pull Requests">
  </a>
</p>

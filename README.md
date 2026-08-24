# ⚙️ Infrastructure As Code

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v2.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Management tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`management` `operations` `cli` `golang` `io`

---

## What is Infrastructure-As-Code?

**Infrastructure-As-Code** is an operations management tool for automating, tracking, and coordinating development workflows.

## Features

- ✅ Streaming file processing
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Infrastructure-As-Code.git
cd Infrastructure-As-Code

# Build
go build -o infrastructure-as-code .

# Run
./infrastructure-as-code Usage: iac-validator <config-file>
```

### Or directly with `go run`:
```bash
go run main.go Usage: iac-validator <config-file>
```

## Usage

```bash
# Basic usage
./infrastructure-as-code Usage: iac-validator <config-file>

# With flags
./infrastructure-as-code Usage: iac-validator <config-file> value Usage: iac-validator <config-file>
```

### Example Output

```
$ ./infrastructure-as-code Usage: iac-validator <config-file>
Usage: iac-validator <config-file>
Error:
  [WARN] Hardcoded credential at: %s\n
```

## Project Structure

```
Infrastructure-As-Code/
  main.go          # Entry point (41 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)

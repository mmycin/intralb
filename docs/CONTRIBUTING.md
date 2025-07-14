# 🤝 Contributing to IntraLB

Thank you for your interest in contributing to **IntraLB** — a blazing-fast, intelligent load balancer for Go HTTP services.  
Whether it’s a bug fix, feature request, or a new idea, we welcome all contributions!

---

## 🧾 Table of Contents

- [Before You Start](#-before-you-start)
- [How to Contribute](#-how-to-contribute)
  - [Bug Reports](#-bug-reports)
  - [Feature Requests](#-feature-requests)
  - [Pull Requests](#-pull-requests)
- [Code Style & Standards](#-code-style--standards)
- [Development Setup](#-development-setup)
- [Community Expectations](#-community-expectations)

---

## ✅ Before You Start

- Make sure to read our [README](./README.md) and [ISSUES.md](../.github/ISSUES.md)
- Check the [open issues](https://github.com/mmycin/intralb/issues) to avoid duplication
- Read the [Code of Conduct](./CODE_OF_CONDUCT.md)

---

## 🚀 How to Contribute

### 🐛 Bug Reports

Found a bug? Great! Help us squash it:

1. Search [existing issues](https://github.com/mmycin/intralb/issues)
2. If not already reported, open a new issue using the **Bug** template
3. Include logs, steps to reproduce, and environment details

See [`ISSUES.md`](../.github/ISSUES.md) for the full template.

---

### 💡 Feature Requests

Got an idea to improve IntraLB?

1. Check if it's already discussed in [Issues](https://github.com/mmycin/intralb/issues)
2. Open a new issue using the **Feature Request** template
3. Provide use cases, examples, or mockups if possible

---

### 📦 Pull Requests

We ❤️ PRs! Here's how to get started:

1. Fork the repository
2. Create a new branch: `git checkout -b feature/my-new-feature`
3. Make your changes (with comments & tests if needed)
4. Run tests: `go test ./...`
5. Commit and push:  
   `git commit -m "feat: add my-new-feature"`  
   `git push origin feature/my-new-feature`
6. Open a [pull request](https://github.com/mmycin/intralb/pulls) on the `main` branch
7. Include:
   - Description of changes
   - Linked issue (if applicable)
   - Screenshot or logs (if UI or performance-related)

---

## 🧑‍💻 Code Style & Standards

- Follow [Go idioms](https://golang.org/doc/effective_go)
- Keep functions small, composable, and well-documented
- Use `gofmt`, `go vet`, and `golangci-lint` before submitting
- Test coverage is **not optional** — write tests for new logic
- Panics should be handled with care (e.g., recovery middleware)
- Prefer composition over inheritance

---

## ⚙ Development Setup

```bash
git clone https://github.com/mmycin/intralb.git
cd intralb
go mod tidy
go run example/main.go
````

---

## 🌍 Community Expectations

By contributing, you agree to follow our [Code of Conduct](https://chatgpt.com/c/CODE_OF_CONDUCT.md).  
We strive to foster an open, inclusive, and respectful environment for everyone.

---

## 🙏 Thank You

Thank you for being part of the **IntraLB** journey.  
Your contributions help make this tool faster, safer, and more reliable for the Go ecosystem.

If you’re ever unsure, feel free to [open an issue](https://github.com/mmycin/intralb/issues) or [start a discussion](https://github.com/mmycin/intralb/discussions).

Happy coding! 💻⚡

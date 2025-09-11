# Contributing to NACHA Generator

Thank you for your interest in contributing to NACHA Generator! This document provides guidelines and instructions for
contributing to the project.

## Development Setup

1. Fork and clone the repository
2. Install Go 1.25 or later
3. Run `go mod download` to install dependencies
4. Run tests using `go test ./...`

## Code Style Guidelines

- Follow standard Go code formatting
- Run `go fmt` before committing any changes
- Write clear, descriptive variable and function names
- Add comments for exported functions and types
- Include proper error handling
- Follow Go's standard naming conventions
- Use meaningful package names

## Testing

- Write unit tests for new functionality
- Ensure test coverage for bug fixes
- Run all tests locally before submitting PR
- Include both positive and negative test cases
- Test edge cases and error conditions
- Use table-driven tests when appropriate

## Pull Request Process

1. Create a new branch from `main` for your changes
2. Make your changes following the code style guidelines
3. Write or update tests as needed
4. Run all tests locally
5. Update documentation if necessary
6. Submit a pull request with a clear description of changes
7. Address any review comments

## Reporting Issues

When reporting issues, please include:

- Go version and system information
- Clear steps to reproduce the issue
- Expected vs actual behavior
- Any relevant error messages
- Sample code demonstrating the issue (if possible)

## Communication

- Use GitHub Issues for bug reports and feature requests
- Be respectful and constructive in discussions
- Provide context and examples when discussing changes
- Ask questions if something is unclear

## License

By contributing to NACHA Generator, you agree that your contributions will be licensed under the Apache-2.0 License.

## Questions?

If you have any questions about contributing, feel free to open an issue for discussion.

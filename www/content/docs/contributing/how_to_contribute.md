---
weight: 100
title: "How to contribute"
description: "A quick guide on how to contribute to cligen"
icon: "volunteer_activism"
date: "2025-10-20T11:01:25+11:00"
lastmod: "2026-02-15T14:44:00+00:00"
draft: false
toc: true
---

Thank you for your interest in contributing to **cligen**! We welcome all contributions, from bug reports and documentation updates to new features and scaffolds.

## Getting Started

### Prerequisites

To build and test cligen locally, you'll need:

* **Go**: version 1.25 or higher.
* **Make**: for running build and test commands.
* **golangci-lint**: for linting (optional but recommended).

### Fork and Clone

1.  **Fork** the repository on GitHub: [twinsnes/cligen](https://github.com/twinsnes/cligen).
2.  **Clone** your fork locally:
    ```bash
    git clone https://github.com/YOUR_USERNAME/cligen.git
    cd cligen
    ```

## Development Workflow

### Creating a Branch

Always create a new branch for your changes:

```bash
git checkout -b feature/your-feature-name
# or
git checkout -b fix/bug-description
```

### Making Changes

cligen is structured as follows:
- `cmd/`: CLI command definitions.
- `internal/`: Core logic for generation, prompting, and configuration.
- `templates/`: The actual scaffolds used to generate new projects.

If you are modifying templates, refer to the [Templates Overview](../templates) for more details on the structure.

### Running Tests and Linting

Before submitting your changes, ensure everything is working correctly and follows the project's style.

*   **Linting**: Run the linter to check for style issues.
    ```bash
    make lint
    ```
*   **Running Tests**: Execute the test suite.
    ```bash
    make test
    ```
*   **Building**: Build the binary locally to test the CLI manually.
    ```bash
    make build
    ./dist/cligen --help
    ```

## Submitting a Pull Request

1.  **Commit** your changes using clear and descriptive commit messages. We encourage [Conventional Commits](https://www.conventionalcommits.org/).
2.  **Push** your branch to your fork on GitHub.
3.  **Open a Pull Request** against the `main` branch of the original repository.
4.  Provide a clear description of your changes and link to any related issues.

## Code of Conduct

Please note that this project is released with a [Contributor Code of Conduct](code_of_conduct). By participating in this project you agree to abide by its terms.
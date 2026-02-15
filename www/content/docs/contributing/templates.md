---
weight: 200
title: "Templates Overview"
description: "Overview of the template structure"
icon: "article"
date: "2026-02-15T14:00:00+00:00"
lastmod: "2026-02-15T14:00:00+00:00"
draft: false
---

The templates directory is the core of the generator project, it contains all the template files used to generate the cli scaffolding.

## Structure

The directory is organized into three main categories:

1.  **`common/`**: Contains files that are always included in every generated project, regardless of the selected type or features (e.g., `LICENSE`, `Makefile`, `README.md`).
2.  **`types/`**: Contains different project "blueprints". These are mutually exclusive. For example:
    -   `basic`: A simple CLI structure.
    -   `interactive`: A CLI with interactive prompts.
3.  **`features/`**: Contains optional components that can be added to a project. A common example is the `pages/www` feature which adds a documentation website.

## Template Rendering

Files within these directories can be either static files or dynamic templates:

-   **Go Templates (`.tmpl`)**: Files ending with the `.tmpl` extension are processed using Go's `text/template` package. The extension is removed during generation.
-   **Static Files**: Files without the `.tmpl` extension are copied directly to the output directory without any modifications.

## Embedding

The entire `templates/` directory is embedded into the Go binary using the `embed` package, specifically in `templates/main.go`. This ensures that the generator is a single, self-contained binary that doesn't depend on external files at runtime.

```go
package templates

import "embed"

//go:embed all:*
var FolderFS embed.FS
```

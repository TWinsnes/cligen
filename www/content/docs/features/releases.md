---
weight: 400
title: "Automated Releases"
description: "Automatically create a new release for your application"
icon: "release_alert"
date: "2026-01-31T13:45:00+00:00"
lastmod: "2026-01-31T13:45:00+00:00"
draft: false
toc: true
---

The generated application includes a fully automated release pipeline using [GoReleaser](https://goreleaser.com/) and GitHub Actions. This pipeline handles building binaries for multiple platforms, creating a GitHub Release, generating an SBOM, and optionally updating a Homebrew Tap.

Default platforms configured for release:

* Linux amd64
* Linux arm64
* Linux i386
* macOS arm64
* macOS amd64
* Windows amd64
* Windows arm64
* Windows i386

## Prerequisites

Before creating your first release, ensure you have:

1.  **GITHUB_TOKEN**: This is automatically provided by GitHub Actions and has permissions to create releases.
2.  **TAP_GITHUB_TOKEN** (Optional): If you enabled Homebrew integration, you must add this secret to your repository. See the [Homebrew Integration](./homebrew.md) for details.

## Creating a Release

The release pipeline is triggered by pushing a git tag that starts with `v`.

### 1. Tag your commit

Navigate to your project root and create a new semantic version tag:

```bash
git tag -a v1.0.0 -m "Release v1.0.0"
```

### 2. Push the tag

Push the tag to your GitHub repository:

```bash
git push origin v1.0.0
```

### 3. Monitor the Release

Once the tag is pushed, a new workflow run will start in the **Actions** tab of your GitHub repository. The `Release` workflow will:

- Run linting checks.
- Build binaries for Linux, macOS, and Windows.
- Generate an SBOM (Software Bill of Materials).
- Create a new GitHub Release with the generated [Release Notes](./release_notes.md).
- Upload all artifacts to the release.
- Update your Homebrew Tap (if configured).

## Local Testing

You can test the release process locally using the GoReleaser CLI without publishing:

```bash
goreleaser release --snapshot --clean
```

This will build the binaries and create the archives in the `dist/` directory, allowing you to verify the output before pushing a tag.

---
weight: 500
title: "Configuration"
description: ""
icon: "settings"
date: "2025-10-20T10:50:10+11:00"
lastmod: "2025-10-20T10:50:10+11:00"
draft: false
toc: true
---

You can change the default configuration for the generated applications using a configuration file. This is especially useful for features that use the same settings across multiple applications, like homebrew, where the homebrew tap repository is the same for all applications.

The configuration file is located under `~/.config/.cligen.yaml` under macOS and linux.

You can modify the file directly, or use the `cligen configure` command to follow a wizard taking you through each of the configuration options.

{{< alert context="info" >}}
The configuration file is not encrypted.
{{< /alert >}}

## Configuration options

| Option                   | Description                                                                                                                                       |
|--------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------|
| homebrew.enabled         | If the homebrew feature will automatically be enabled in the prompt                                                                               |
| homebrew.repo            | The repository containing the homebrew tap, e.g. `homebrew-tap`. Needs to follow the naming convention `homebrew-*` for homebrew commands to work |
| homebrew.github_username | The username or organisation name for the github repository                                                                                       |

## Example configuration file

```yaml
homebrew:
    enabled: true
    repo: homebrew-tap
    github_username: twinsnes
```
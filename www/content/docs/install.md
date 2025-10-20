---
weight: 200
title: "Install"
description: "Guide to setting up and installing Cligen"
icon: "system_update_alt"
date: "2025-10-16T20:15:04+11:00"
lastmod: "2025-10-16T20:15:04+11:00"
draft: false
toc: true
---

Cligen is available for installation on MacOS, Linux, and Windows as a stand alone binary or built from source. The 
recommended installation method for macOS and linux is via Homebrew.

## MacOS / Linux with Homebrew

Install the latest version of Cligen with Homebrew from the tap.

```shell
brew install twinsnes/tap/cligen
```

{{< alert context="info" text="Installing with homebrew will work around the Signing and Notarizing requirement on macOS by removing the quarantine bit form the binary post install" />}}

## MacOS / Linux / Windows with Binary Download

Download the latest release from the [releases page](https://github.com/twinsnes/cligen/releases) and add it to your path.

## From source

To build the binary from source, clone the repository, run `go mod tidy` to download the dependencies, and run the `make build` command.

```shell
git clone git@github.com:TWinsnes/cligen.git
cd cligen
go mod tidy
make build
```

This will output the binary to `dist/cligen`, which can be copied to your path or called directly from the command line.

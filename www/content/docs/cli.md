---
weight: 400
title: "Cli"
description: "An overview of the CLI and the available commands"
icon: "terminal"
date: "2025-10-20T10:52:04+11:00"
lastmod: "2025-10-20T10:52:04+11:00"
draft: false
toc: true
---

The CLI comes with a few commands that can be used to generate your application scaffold, as well as manage the configuration and default values.

## Commands

The following commands are used by passing them to the cli binary as arguments. e.g. `cligen new` or `cligen help configure`.

| Command   | Description                                                                                                                       |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------|
| help, h   | Displays the help page for cligen. If called with a command (e.g. new), it will display the help text for that particular command |
| new       | Starts the scaffold generator wizard in the current directory to generate a new cli application                                   |
| configure | Starts the configuration wizard to step through setting up or updating an existing configuration file                             |

## New Command

The new command is used to generate a new cli application scaffold in the current directory. It will prompt you for the name of the application and the language to generate it in.

The following options are available for the new command:

| Option     | Description                                                                           |
|------------|---------------------------------------------------------------------------------------|
| --dry-mode | Run the generator without writing any files to disk and instead log what would happen |

## Global options

The following options are globally available to all commands.

| Option        | Description                                                                                           |
|---------------|-------------------------------------------------------------------------------------------------------|
| -h, --help    | Displays the help page for cligen, call it on a specific command to display the help for that command |
| -v, --version | Display the version of the local binary                                                               |

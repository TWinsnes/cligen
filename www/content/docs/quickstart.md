+++
weight = 100
date = "2025-10-11T17:00:00+10:00"
draft = false
author = "Thomas Winsnes"
title = "Quickstart"
icon = "rocket_launch"
toc = true
description = "A quickstart guide to install and use cligen"
publishdate = "2025-10-11T17:00:00+10:00"
tags = ["Beginners"]
+++
## Requirements

* Git
* Golang ≥ 1.23

### Optional

* Docker

## Install

The simplest way to install cligen is with Homebrew:

```shell
brew install twinsnes/tap/cligen
```

This makes this tool available in your terminal.

## Usage

Creat a new directory for your CLI app, and navigate to it.

```shell 
mkdir my-cli-app
cd my-cli-app
```

Optionally, set up a git repository and configure a remote. This will be picked up by the generator and used as your project's module name.

```shell 
git init
git remote add origin git@github.com:my-username/my-cli-app.git
```

To generate a new app scaffold, run the following command in the directory you navigated to:

```shell
cligen new
```

This will prompt you to fill in some information about you app, and select which features you would like to include. Depending on the features you select, you may be prompted with additional questions to ensure the app is configured correctly. Some of the features also requires additional dependencies to work properly. Please see the [features](docs/features) page for more information.

Follow the prompts to generate the app, and then open the generated README file for instructions on the next steps.

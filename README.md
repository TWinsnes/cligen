# cligen 

![GitHub Release](https://img.shields.io/github/v/release/twinsnes/cligen)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/twinsnes/cligen/ci.yml)
![GitHub License](https://img.shields.io/github/license/twinsnes/cligen)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/twinsnes/cligen)
[![Static Badge](https://img.shields.io/badge/docs-github_pages-blue)](https://twinsnes.github.io/cligen/docs)


Cligen is a command line tool to generate go cli applications with a set of sane defaults, including ci/cd and release management. It takes a fair bit of time setting up the basics when building a new cli application, and this tool aims to make that process easier.

![Demo](demo.gif)

Check out the quickstart to get started: https://twinsnes.github.io/cligen/docs/quickstart/

## Why did I make this?
I found myself writing quite a few cli tools, and it was always the same boilerplate over and over again, so I decided to create a tool to generate it for me. I could use an LLM to generate a lot of the structure, but I also wanted to make sure I had an opinionated and tested scaffold around building and releasing the app. So I've put together a tool that does it all for me. 

This way I have a well-structured base to build on, and I can focus on building the actual app.

I hope that you find it useful!

## Installation

Simplest way to install is by using homebrew

```shell
brew install twinsnes/tap/cligen
```

This makes cliegen available on our path and works around the MacOS Notarization requirement.

Other installation methods can be found in the docs: https://twinsnes.github.io/cligen/docs/install/

## Usage

Call `cligen new` in the folder you would like to create a new cli application.

Details on usage can be found in the docs: https://twinsnes.github.io/cligen/docs/quickstart/#usage
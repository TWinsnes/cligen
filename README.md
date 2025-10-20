# cligen 

![GitHub Release](https://img.shields.io/github/v/release/twinsnes/cligen)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/twinsnes/cligen/ci.yml)
![GitHub License](https://img.shields.io/github/license/twinsnes/cligen)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/twinsnes/cligen)

Cligen is a command line tool to generate go cli applications with a set of sane defaults, including ci/cd and release management. It takes a fair bit of time setting up the basics when building a new cli application, and this tool aims to make that process easier.

![Demo](demo.gif)

Check out the quickstart to get started: https://twinsnes.github.io/cligen/docs/quickstart/

## Why did I make this?
I found myself writing quite a few cli tools, and it was always the same boilerplate over and over again, so I decided to create a tool to generate it for me. I could use an LLM to generate a lot of the golang code, but I also wanted to make sure I had an opinionated and tested scaffold around building and releasing the app. So I've put together a tool that does it all for me.

With this I can skip all the repetitive stuff and get straight into building the app.

I hope that you find it useful!

## Included

- [x] Opinionated CLI app scaffold
- [x] Github Actions
- [x] Makefile
- [x] Golangci-lint
- [x] Homebrew tap
- [x] Github release
- [x] Github release notes
- [x] Gitleaks secret scanning
- [x] Dependency license review
- [x] Demo generator using VHS 
- [x] Github pages


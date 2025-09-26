# cligen 

![GitHub Release](https://img.shields.io/github/v/release/twinsnes/cligen)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/twinsnes/cligen/ci.yml)
![GitHub License](https://img.shields.io/github/license/twinsnes/cligen)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/twinsnes/cligen)


A CLI tool to generate a CLI app scaffold with sane defaults and complete github CI/CD so you can focus on building the app rather than the boilerplate.

![Demo](demo.gif)

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

# Quickstart

## Install
### MacOS / Linux with Homebrew

```shell
brew install twinsnes/tap/cligen
```

### MacOS / Linux / Windows with Binary Download

Download the latest release from the [releases page](https://github.com/twinsnes/cligen/releases) and add it to your path.

### From source with Go

Clone the repository and run `make build` to build the binary.

## Usage

To generate a new CLI app, run the following command in the location you want to generate the app:

```shell
cligen new
```

Follow the prompts to generate the app, and then open the generated README file for instructions on what to do next.

# Docs

## Config

Cligen supports a yaml configuration file to preset values for generated applications. It's stored 
in your home directory (`~/.config/.cligen.yaml`)

Use the `configure` command to set up and save the configuration file to the correct location.

```yaml
homebrew:
    enabled: true
    repo: <insert repo name>
    github_username: <insert repo owner's github username, e.g. twinsnes>
```


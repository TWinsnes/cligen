---
weight: 50
title: "Overview"
description: "A quick overview of cligen and how it works"
icon: "dashboard"
date: "2025-10-20T10:49:22+11:00"
lastmod: "2025-10-20T10:49:22+11:00"
draft: false
toc: true
---

Cligen is a command line tool to generate go cli applications with a set of sane defaults, including ci/cd and release management. It takes a fair bit of time setting up the basics when building a new cli application, and this tool aims to make that process easier. 

The tool itself is a single binary file that can be installed with no external dependencies or shared libraries. This makes it easy to distribute and run on any system. The goal is that this translates to the scaffold as well, meaning tools that you build using this generator will be as portable and easy to adopt as possible. 

## How it works

When you run the generator, it will ask you a few questions about your project, and then it will generate a new project for you in your current directory. It will also look at the local system for information that can help with setting up defaults, like the git repository info and remotes. 

The project will be a go module, and it will contain a few basic files and directories. In addition to this it will also generate all the required scripts and configuration files to get you started.

Cligen contains quite a few different features that can be used to build your app, please check out the [features page](docs/features) for more details on the individual features and how to use them.
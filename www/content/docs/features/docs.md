---
weight: 600
title: "Docs on GitHub Pages"
description: "Overview of the docs on GitHub Pages feature"
icon: "article"
date: "2025-10-21T15:53:06+11:00"
lastmod: "2025-10-21T15:53:06+11:00"
draft: false
toc: true
---

The docs on GitHub Pages feature allows you to host your documentation on GitHub Pages using a static site generator to convert markdown files to html and javascript. It is powered by [Hugo](https://gohugo.io/) and uses the [Lotus Docs](https://lotusdocs.dev/) theme. This gives a good balance between a simple and clean look and feel, while still being able to keep your documentation repository lightweight.

{{<alert context="info" title="Note">}}
The site is located in the `www` folder in the root of the repository. All commands for generating site content and testing changes should be run from this folder.
{{</alert>}}

## Overview of files and folders

- **assets/images/** - Images used in the documentation reside here
- **content/docs/** - All your documentation markdown files reside here
- **data/** - This is mainly used to store the landing page configuration
- **.gitignore** - The .gitignore file which prevents generated files produced during testing from being added to git
- **go.mod** - The go.mod file which contains the dependencies for the project
- **go.sum** - The go.sum file which contains the checksum for the project
- **hugo.toml** - The configuration file for the Hugo static site generator

## Adding Documentation

To add new pages to the generated site, create a new Markdown file in the `contents/docs` folder. To set up the front matter automatically you can use the `hugo new` command.

For a more indepth guide on how to set up and manage the documentation, please see the [Lotus Docs documentation](https://lotusdocs.dev/docs/quickstart/)

## Updating the Landing page

The landing page is a special page that is configured in the `data/landing.yml` file. 

For details on how to configure the landing page, please see the [Lotus Docs documentation](https://lotusdocs.dev/docs/guides/landing-page/overview/)

## Release Pipeline

The Pages release pipeline is triggered by a version tag being pushed to the repository and will build the docs and deploy them to the GitHub Pages environment. It can also be triggered manually, and in this case will build from the latest commit on the `main` branch by default.

The workflow is defined in the `.github/workflows/pages.yml` file.

## Troubleshooting

#### Tag "vX.X.X" is not allowed to deploy to github-pages due to environment protection rules.

If you run into this issue when the release pipeline runs, this is likely because the github pages environment has been limited to only allow deployments from the `main` branch, and you will need to change environment to allow deployments from tags matching the pattern: `v*`

#### My pages are not generating in Github Pages

This could be due to the draft flag being set to true in the front matter of the markdown file.
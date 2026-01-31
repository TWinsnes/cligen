---
weight: 500
title: "Automated Release Notes"
description: "Overview of the Automated Release Notes feature"
icon: "article"
date: "2025-10-21T15:37:57+11:00"
lastmod: "2025-10-21T15:37:57+11:00"
draft: false
toc: true
---

When a new release is created, the scaffolded release pipeline will automatically generate release notes based on the commits since the last release. Each commit will be added to the release notes as a bullet point. This process expects that the commit messages follow the [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) standard, and will organize the commits into categories based on the commit type.

By default, a few tags will be ignored: 

* `chore:`
* `docs:`
* `test:` 

These tags are intended to be used for internal changes that don't affect the end user.

To configure the tags that will be ignored, edit the .goreleaser.yaml file in the root of the project.
Details on how to configure the release notes can be found in the [goreleaser documentation](https://goreleaser.com/customization/release/#github).

For more information on how to trigger a release, see the [Automated Releases](./releases.md) guide.
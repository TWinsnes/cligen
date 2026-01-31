---
weight: 200
title: "Homebrew Integration"
description: "Overview of the Homebrew feature"
icon: "article"
date: "2025-10-21T15:36:28+11:00"
lastmod: "2025-10-21T15:36:28+11:00"
draft: false
toc: true
---

Releasing your CLI app as a Homebrew formula is an optional feature that can be enabled when running the `cligen new` command. It requires a couple of extra settings to be configured correctly:
|setting|description|example|
|---|---|---|
|`owner`|The owner of the tap|twinsnes|
|`repo`|The name of the repo, the name needs to start with `homebrew-` to comply with the homebrew standard |homebrew-tap|

The above settings will be converted to a github url ( e.g. `https://github.com/twinsnes/homebrew-tap`) and used by the release pipeline to publish the formula.

## Set up access for the github actions to publish the formula

By default, the release pipeline will not have access to updating the tap repo, so you will need to set up a token with the correct permissions.

The built in actions will look for a secret called `TAP_GITHUB_TOKEN` to use for publishing the formula to the tap repo.

To generate the token, go to your github account settings and create a new personal access token:
https://github.com/settings/personal-access-tokens

Make sure scope is limited to `contents` with `read and write` permissions.
It is also recommended to limit the token to only the tap repo.

## Adding the token secret to the repo

Once you have the token, you need to add it to the repo as a secret called `TAP_GITHUB_TOKEN`.

Under the repo settings, go to the `Secrets and variables` menu item and select `actions` from the dropdown. Then press the `New repository secret` button to add the secret with the name `TAP_GITHUB_TOKEN` and the value of the token you generated.

Once that is done it will be available to the release pipeline and you can publish the formula.

To test this, create a new release by following the [Automated Releases](./releases.md) guide and watch the logs to see if the formula is published correctly.

---
weight: 300
title: "Features"
description: "Breakdown of features and how to use them"
icon: "auto_awesome"
date: "2025-10-16T22:02:06+11:00"
lastmod: "2025-10-16T22:02:06+11:00"
draft: false
toc: true
---

{{<alert context="info" title="Note">}}
This page is a work in progress.
{{</alert>}}

## Automatic demo video using VHS

Cligen comes with the required tools and scripts to automatically generate a demo video using VHS. This makes it easy to show off you CLI app as part of the readme or documentation. It also makes it very easy to update this as you're making changes to your app.

In the root directory of your project there is a `demo.tape` file that contains the steps of the demo. This is a simple step file that emulates the steps of a user interacting with your app using keyboard presses. It is quite permissive and allows you to do anything you want.

We run the generation of the video / gif using a docker container, so there is no need to install anything locally, e.g. a headless chrome.

To update your video, simply run the `make vhs` command from the root of your project and wait for the output. You should see a `demo.gif` file in the root of your project that you can link directly in your README.md file.

More details on the VHS syntax and examples can be found in the (VHS repository)[https://github.com/charmbracelet/vhs]

## Homebrew

Releasing your CLI app as a Homebrew formula is an optional feature that can be enabled when running the `cligen new` command. It requires a couple of extra settings to be configured correctly:
|setting|description|example|
|---|---|---|
|`owner`|The owner of the tap|twinsnes|
|`repo`|The name of the repo, the name needs to start with `homebrew-` to comply with the homebrew standard |homebrew-tap|

The above settings will be converted to a github url ( e.g. `https://github.com/twinsnes/homebrew-tap`) and used by the release pipeline to publish the formula. 

### Set up access for the github actions to publish the formula

By default, the release pipeline will not have access to updating the tap repo, so you will need to set up a token with the correct permissions.

The built in actions will look for a secret called `TAP_GITHUB_TOKEN` to use for publishing the formula to the tap repo.

To generate the token, go to your github account settings and create a new personal access token:
https://github.com/settings/personal-access-tokens

Make sure scope is limited to `contents` with `read and write` permissions.
It is also recommended to limit the token to only the tap repo.

### Adding the token secret to the repo

Once you have the token, you need to add it to the repo as a secret called `TAP_GITHUB_TOKEN`.

Under the repo settings, go to the `Secrets and variables` menu item and select `actions` from the dropdown. Then press the `New repository secret` button to add the secret with the name `TAP_GITHUB_TOKEN` and the value of the token you generated.

Once that is done it will be available to the release pipeline and you can publish the formula.

To test this, create a new release and watch the logs to see if the formula is published correctly.

## Docs site on Github pages
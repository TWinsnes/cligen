---
weight: 300
title: "Demo Video Generation"
description: "Overview of the Demo Video Generation feature"
icon: "video_library"
date: "2025-10-16T22:02:06+11:00"
lastmod: "2025-10-16T22:02:06+11:00"
draft: false
toc: true
---

Cligen comes with the required tools and scripts to automatically generate a demo video using VHS. This makes it easy to show off you CLI app as part of the readme or documentation. It also makes it very easy to update this as you're making changes to your app.

In the root directory of your project there is a `demo.tape` file that contains the steps of the demo. This is a simple step file that emulates the steps of a user interacting with your app using keyboard presses. It is quite permissive and allows you to do anything you want.

We run the generation of the video / gif using a docker container, so there is no need to install anything locally, e.g. a headless chrome.

To update your video, simply run the `make vhs` command from the root of your project and wait for the output. You should see a `demo.gif` file in the root of your project that you can link directly in your README.md file.

More details on the VHS syntax and examples can be found in the (VHS repository)[https://github.com/charmbracelet/vhs]

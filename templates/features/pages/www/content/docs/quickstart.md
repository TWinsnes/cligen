+++
weight = 100
date = "2025-02-10T19:00:00-10:00"
draft = true
author = "Thomas Winsnes"
title = "Quickstart"
icon = "rocket_launch"
toc = true
description = "A quickstart guide to creating new content in Lotus Docs"
publishdate = "2025-02-10T19:00:00-10:00"
tags = ["Beginners"]
+++

## Create New Content

Navigate to the www folder in your repo and use the `hugo new` command to create a file in the `content/docs` directory:

```shell
hugo new docs/examplepage.md
```

## Preview your docs site 

Run the following command in the `www` folder to start the server with drafts.

```shell
hugo server -D
```

Your docs server should now be running on http://localhost:1313/docs
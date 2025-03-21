# Split file to chunks

Problem: I have a file like `out` with many 11-lines-chunk like this

```
---
layout: post
date: 2025-03-21
title: Dotfiles
details: Linux, Windows, Bash, Vim
hascontent: false
repourl: https://github.com/minhhoccode111/dotfiles
demourl:
noteurl:
---

```

- I want to split it into many chunks, each chunk has 11 lines
- Put each chunk to a separate file
- Each new file's name create by:
  - Taking the 4th line of the chunk
  - Remove the prefix `title: `
  - Lowercase the string
  - Replace all ` ` and `'` by `-`

Conclusion: I have this problem when trying to convert the markdown table format of my projects to separate Front Matter format file for each project to use with Liquid Template.

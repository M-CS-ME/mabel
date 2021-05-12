# The Mabel Literate Programming Tool

## Introduction

I wanted to take notes with executable code snippets inside, and I didn't want it to be glued to my text editor. This little project is supposed to (eventually) do that.

I chose markdown because it's simple, light weight, and has many implementations. A markdown engine or `pandoc` will be better than any weaver that I could write. A pdf sample using pandoc is provided (`mabel.pdf`) which was produced by `pandoc mabel.md -o mabel.pdf`

The source in `src/mabel.go` is generated through `make tangle` or `mabel mabel.md > src/mabel.go`

## Installation

```bash
$ git clone https://github.com/M-CS-ME/mabel mabel
$ cd mabel
$ make tangle
$ sudo make install
```

## Usage

`mabel` prints the source blocks into stdout, therefore to write to a file just do `mabel src.md > src`.

To specify a specific source block `n` just do `mabel src.md n`

The syntax is the same as markdown, only code blocks are counted and in-line code is ignored.

## What's next?

- [X] Ability to print to stdout a selected group of code blocks (like `org-babel`)
- [ ] Add concurency for speed 
- [ ] Ability to run a specific code block and write to output to a file (either through pipe or buf)
- [ ] Configuring specific compilers and interpreters.

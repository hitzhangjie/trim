# README

## Description

`trim` is used for trim leading spaces of each line read from stdin (you can use redirect from file),

it works as following:
- scans each lines of input
- determine the mininum number of spaces to trim for all lines
- trim leading spaces and output

## Installation

```bash
go get github.com/hitzhangjie/trim
```

## Usage

You can use it like this way, `echo .....DATA..... | trim`.

Here's an example, the data is output by `git log`:

```bash
$ echo "
    fix/stream: aaaaaaaaaaaaaaaaa
    
    This CL does .........：
    - aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
      bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
    - cccccccccccccccccccccccccccccccccccccccccccccccccccccccccc
      ddddddddddddddddddddddddddddddddddddddddddddddddddddddd
      
    close #551
    
    --story=xxxxxxxxx eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee
" | trim
```

The output will look like this:

```
fix/stream: aaaaaaaaaaaaaaaaa

This CL does .........：
- aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
  bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb
- cccccccccccccccccccccccccccccccccccccccccccccccccccccccccc
  ddddddddddddddddddddddddddddddddddddddddddddddddddddddd
  
close #551

--story=xxxxxxxxx eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee
```

You can use `echo "..." | trim | pbcopy` pipeline to copy the output to 
clipboard on macOS. So we can paste the pretty formatted data to fill 
the PR message or paste it anywhere we want.

Use it at your will.
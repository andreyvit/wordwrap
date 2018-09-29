wordwrap (Golang)
=================

[![GoDoc](https://godoc.org/github.com/andreyvit/wordwrap?status.svg)](https://godoc.org/github.com/andreyvit/wordwrap)

Splits a bunch of text into lines.


Installation
------------

```
import (
    "github.com/andreyvit/wordwrap"
)
```


Example
-------

Enumerate lines:

```go
wordwrap.Wrap(lorem, 40, wordwrap.Options{}, func(line string) {
    fmt.Println(line)
})
```

or get a string:

```go
fmt.Print(wordwrap.WrapString(lorem, 40, wordwrap.Options{}))
```

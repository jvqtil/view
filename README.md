# view [![Go Report Card](https://goreportcard.com/badge/github.com/jvqtil/view)](https://goreportcard.com/report/github.com/jvqtil/view) [![Go Reference](https://pkg.go.dev/badge/github.com/jvqtil/view.svg)](https://pkg.go.dev/github.com/jvqtil/view)
Simple Go package for displaying any text via [bat](https://github.com/sharkdp/bat), [glow](https://github.com/charmbracelet/glow), [less](https://www.greenwoodsoftware.com/less), or your `$PAGER`, colored and without an infinite `if {} else {}` in your code.

Pretty good way for showing help messages in CLIs / TUIs

## Install
```go
go get github.com/jvqtil/view
```

## Functions
`func Show(text string, paintFunc ...func(a ...any) string)`

## Examples
- [`showhelp.go`](https://github.com/jvqtil/view/blob/master/examples/showhelp.go)

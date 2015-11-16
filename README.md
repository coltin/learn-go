# learn-go

Following the directory structure under https://golang.org/doc/code.html, let's learn Go-lang!

I `git clone` this repo under my `go_workspace/src/` repo. Which is strange, but works for go.

Resources I picked up along the way:
	Official docs: https://golang.org/doc/
	Very basic walkthrough, good for getting set up: https://golang.org/doc/code.html
	Slides on static analysis with Go: https://talks.golang.org/2014/static-analysis.slide#1
	Go trainning: https://github.com/ardanlabs/gotraining

Interesting Tidbits:

1. `go test ./...` will run all tests in the current directory and all subdirectories.
1. `go test foo/...` likewise will run everything under `foo`.
1. You can import sources directly by supplying the url `import "github.com/golang/example/stringutil"`.
1. Package "reflect" has a method `reflect.TypeOf(varaible)` which will tell you what the type of a varaible/interface is.
1. `len(varaible)` is a way of accessing the lenght of something like a string or array.
1. Go talks about why it chose to do `x int` instead of `int x`: https://blog.golang.org/gos-declaration-syntax
1. How slices work internally: https://blog.golang.org/go-slices-usage-and-internals
1. The end of the slides with more links: https://tour.golang.org/concurrency/10
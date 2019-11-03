# Readahead
This library adds support for
[`readahead`](http://man7.org/linux/man-pages/man2/readahead.2.html) for Linux.
[![GoDoc](https://godoc.org/github.com/hodgesds/goreadahead?status.svg)](https://godoc.org/github.com/hodgesds/goreadahead)


# Example
To use with an open file call the `Fd` method and cast to `int`. Readahead can be validated with [`mincore`](http://man7.org/linux/man-pages/man2/mincore.2.html).

```
if err := readahead.Readahead(int(tmpfile.Fd()), 0, 2<<10); err != nil {
	t.Fatal(err)
}
```

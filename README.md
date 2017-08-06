# Stdlib

[![Build Status](https://img.shields.io/travis/goph/stdlib.svg?style=flat-square)](https://travis-ci.org/goph/stdlib)
[![Go Report Card](https://goreportcard.com/badge/github.com/goph/stdlib?style=flat-square)](https://goreportcard.com/report/github.com/goph/stdlib)
[![GoDoc](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=flat-square)](https://godoc.org/github.com/goph/stdlib)

**All kinds of utilities and extensions of the standard library.**


## Documentation

The project closely follows the package structure in the [standard library](https://golang.org/pkg/#stdlib)
and aims to provide high quality extensions and utilities.


## Development

The project requires [Glide](https://glide.sh/) to install dependencies.

``` bash
$ make dep
```

When all coding and testing is done, please run the test suite:

``` bash
$ make check
```


### Experimental features

When a feature is not mature or stable enough for general availability, it can be marked as *experimental*.

This means that those features can only be used with the `experimental` build tag.

Being experimental does not provide any BC promise.


## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.

# Change Log


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- **os:** `Env` function to return the complete environment as a map of strings
- **time:** `Abs` function to time package for durations

### Changed

- **os:** Moved `MustEnv` to separate file
- **os:** Moved `DefaultEnv` to separate file

### Removed

- **time:** `Clock` interface: it should be defined by consumers, though it is considered to be part of the public API


## 0.14.0 - 2017-09-06

### Changed

- Improved trace package

### Removed

- `log` package


## 0.13.0 - 2017-08-30

### Removed

- **ext:** `errors` package usage from Closer implementations
- `errors` package moved to [github.com/goph/emperror](https://github.com/goph/emperror)


## 0.12.0 - 2017-08-30

### Added

- **errors:** Composite handler from Emperror
- **errors:** Test handler from Emperror
- **errors:** Utils from Emperror
- **errors:** Handler function type for log-like functions (like `log.Print`, `testing.T.Error`)
- **errors:** Handler function type


## 0.11.0 - 2017-08-28

### Changed

- **ext/types:** *(experimental)* Moved `Arguments` to types package
- **ext/types:** Removed *experimental* flag from `Arguments`


## 0.10.0 - 2017-08-27

### Added

- **errors:** Stack trace to recovered errors
- **log:** Structured logger interface
- **log:** Default nop logger implementation
- **errors:** Default nop error handler implementation


## 0.9.0 - 2017-08-22

### Added

- **ext:** *(experimental)* An argument list with type assertions
- **errors:** `Cause` method to contextual error
- **errors:** `Causer` interface
- **errors:** `Handler` interface
- **errors:** `github.com/pkg/errors` as a dependency
- **errors:** `StackTracer` interface
- **errors:** Aliases for `github.com/pkg/errors` package

### Changed

- **errors:** Renamed `ContextualError` to `Contextor`


## 0.8.0 - 2017-08-06

### Added

- **os:** `DefaultEnv` to fallback to a default value when a key is not found
- **net/http:** `HandlerAcceptor` for accepting HTTP handlers
- **net/http/pprof:** `RegisterRoutes` to allow registering routes in other routers than the `http.DefaultServeMux`
- **x/net/trace:** `RegisterRoutes` to allow registering routes in other routers than the `http.DefaultServeMux`
- **x/net/trace:** `NoAuth` to allow remote tracing
- **expvar:** `RegisterRoutes` to allow registering routes in other routers than the `http.DefaultServeMux`
- **archive/zip:** *(experimental)* A Reader which reads a certain file from a ZIP archive


## 0.7.0 - 2017-07-27

### Added

- **internal/testing:** `Equal` assertion
- **internal/testing:** Panic assertions
- **os:** `MustEnv` to make sure an environment variable exists

### Changed

- Replaced internal testing code with [testify](https://github.com/stretchr/testify)


## 0.6.0 - 2017-07-06

### Added

- **errors:** `ContextualError` provides context to errors


## 0.5.0 - 2017-07-06

### Added

- **errors:** Optional message in `MultiErrorBuilder`
- **ext:** `Close` function to try calling Close on interfaces

### Changed

- **errors:** `MultiError` is not exported anymore


## 0.4.1 - 2017-06-24

### Added

- **ext:** Dummy `ext.Closer` implementation


## 0.4.0 - 2017-06-22

### Added

- **errors:** `ErrorCollection` serving as an interface for a multi-error struct
- **errors:** `MultiError` immutable structure aggregating multiple errors into a single value
- **errors:** `MultiErrorBuilder` collecting and aggregating multiple errors into a single value (`MultiError`)
- **ext:** Package ext contains some extra code which does not fit into the categorization of the stdlib
- **ext:** Closing tools (similar to `io.Closer`)


## 0.3.0 - 2017-06-19

### Added

- **time:** `ClockFunc` for making functions Clocks
- **time:** global `SystemClock` to avoid creating new instances of Clocks
- **time:** `StoppedAt` clock type

### Changed

- Moved back to `_test` test packages

### Removed

- `ShutdownManager` since it's not really an extension of the stdlib.
- Error handling code (moved to [github.ibm.com/goph/emperror](https://github.ibm.com/goph/emperror))
- **time:** `StoppedClock` (use StoppedAt instead)


## 0.2.1 - 2017-05-26

### Added

- **time:** `MySQLDateTime` format string


## 0.2.0 - 2017-05-23

### Added

- **net:** `NewVirtualAddress` to create virtual `net.Addr` instances
- **net:** Pipe Listener-Dialer pair for testing client-server applications (eg. gRPC)


## 0.1.0 - 2017-05-16

### Added

- **errors:** `Recover` function to create an error from a recovered panic
- **errors:** `Handler` interface to handle errors
- **errors:** `LogHandler` implementation (tested with [sirupsen/logrus](https://github.com/sirupsen/logrus))
- **errors:** `TestHandler` implementation for testing purposes
- **errors:** `NullHandler` implementation as a fallback
- **util:** `ShutdownManager` to register and execute shutdown handlers when an application exits


## 0.0.5 - 2017-05-10

### Added

- `Must` func to panic when error is passed
- **time:** `Clock` interface and implementations to make testing with time easy


## 0.0.4 - 2017-04-11

### Added

- **strings:** `ToSpinal` util to convert a string to *spinal-case*
- **strings:** `ToTrain` util to convert a string to *Train-Case*
- **strings:** `ToCamel` util to convert a string to *CamelCase*


## 0.0.3 - 2017-04-09

### Added

- **archive/tar:** A Reader which reads a certain file from a TAR archive, optionally decompressing it
- **archive/tar:** `NewTarGzFileReader` (returns a Reader which decompresses and unarchives a .tar.gz stream and returns a file from it)


## 0.0.2 - 2017-04-03

### Changed

- Matched `str` package with `strings` in the stdlib


## 0.0.1 - 2017-03-11

### Added

- **strings:** `ToSnake` util to convert a string to *snake_case*


[Unreleased]: https://github.com/goph/stdlib/compare/v0.14.0...HEAD

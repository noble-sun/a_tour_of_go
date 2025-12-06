# A Tour of Go

This repository is a tag-along of the official [A Tour of Go](https://go.dev/tour/list) guide.

## Running a Go Program

### Prerequisites

Make sure you have Go installed. You can download it from [go.dev/dl](https://go.dev/dl/).

Verify your installation:

```bash
go version
```

### Running a Program

To run a Go file directly:

```bash
go run filename.go
```

To build an executable and then run it:

```bash
go build filename.go
./filename
```

### Running with a Module

If the project uses Go modules (has a `go.mod` file):

```bash
go run .
```

Or build and run:

```bash
go build
./a_tour_of_go
```

# Mower 🚜

Mower simulates a group of automatic mowers that mow a virtual lawn simultaneously without colliding.

## Get Started

### Requirements

Make sure the following dependencies are installed:

- [Go](https://golang.org/dl/)

### Running

Create a configuration file and add it as an argument:

```bash
$ go mod download
$ go run . test/sample.txt
```

### Building

```bash
$ go build -o build .
```

### Testing

```bash
$ go test ./...
```

---

## In Retrospect

### Code

Since this project does not belong to any organization, I decided to follow the conventions and idioms advised by the Go team.

I chose Go because it is efficient while having a simple syntax. Afterwards, I think it would have been interesting to use a more classical object-oriented language such as Java because the problem lends itself well to it.

### Testing

For quality and productivity purposes, I decided to test all the features I was adding in order to guarantee a quality standard and to be able to move on to the next problem.

### Split by problems

I divided the project into several sub-problems in order to better understand and isolate them. This also has the advantage of naturally separating concerns.

#### The mowers

A lawnmower is an object that moves in 2D space and responds to basic instructions. My first task was to bring them to life, regardless of their ability to mow a lawn. The implementation was quite fast and the tests were quickly conclusive.

#### The lawn

The lawn needs to have a dimension, and to know which parcels of land are occupied. The issue of acquiring a shared ressource introduced me to the concept of Semaphores which I used to restrict access to the plots.

#### The parser

I really enjoyed working on the parser because it allowed me to continue my research on compiler design. For this parser, I got inspiration from the Go compiler and a lisp built by Rob Pike using Go.
The parser code could be simpler, but the advantage of using a scanner/parser architecture is that the syntax is easily extensible.

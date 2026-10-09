# learning_go

My Go learning repo. I'm a backend engineer (Node, TypeScript, Postgres) who wants to build systems in Go:
a Raft key-value store, a toy database with a write-ahead log, a load balancer, a job scheduler. I started over
from "what does `package main` mean?" so the basics would be solid before any of that.

I typed every line of Go here. I used an AI assistant (Claude Code) as teacher and reviewer. The section
"How the AI was used" says what that did and didn't mean.

## How it started

On 2026-10-06 I started with Docker and Kubernetes around a tiny Go HTTP server (`http-server/`, not in this
repo). I learned that a container is a normal Linux process with namespaces and cgroups, not a VM. I watched a
Deployment, a ReplicaSet and pods on minikube, then did a rolling update and a rollback. That evening I decided my
Go was too shaky to build on, threw away my earlier Go lessons, and restarted from the basics. This repo is that
restart. I started it on 2026-10-07.

## How I work through it

Each step is small: a short explanation of how something works, then I write 10 to 30 lines, write down what I
expect it to print, run it, and compare. At the end I break it on purpose and read the error.

A prediction I got wrong is a good example. I guessed that after `b := a` and `b[0] = 99`, `a[0]` would still be 1.
It printed 99. A slice is a small header (pointer, length, capacity), so `b := a` copies the header and both
point at one array. Later I found that maps behave the same way (a variable holds a handle to a shared table), and
that is why a method with a value receiver can still change a map but not a plain `int`. One wrong guess explained
three later surprises.

Each experiment lives in its own folder, so nothing gets overwritten.

## How the AI was used

It did not write my code. It explained concepts, read and ran my programs to check them, and told me what was
wrong and why. It also ran small throwaway programs in a scratch directory to test its own claims, then deleted
them, and it created one file, `stage-01-basics/testdata/sample.txt`.

It made mistakes, and I kept them in my notes. Early on it wrote an entire key-value server for me to read. I
objected ("if you write all this how would I learn?"), and it deleted the server. It also told me a `go doc`
command that fails on my Go version. I tested several of its claims myself by breaking the code, for example by
changing a tie-break in `Top` and watching which tests failed.

I wrote the first README draft with the AI from my notes. I still edit it as I go.

## What is here

`stage-01-basics/` is the only stage so far.

| Folder | What I learned there |
|---|---|
| `slices/` | the slice header, shared arrays, `append` and `copy` |
| `strings/` | `strings.Fields` vs `strings.Split`, and the empty strings `Split` can return |
| `wordcount/` | counting words in a map and sorting by count with a tie-break, so the output is the same every run |
| `wordfile/` | errors as return values, `os.ReadFile`, a function that returns `([]string, error)` |
| `readfile/` | my first own package: an exported `ReadWords`, imported by `counter/` (a library, so `go run` doesn't apply) |
| `structs/` | value vs pointer receivers |
| `counter/` | a `TopK` type with `Add`, `Top` and `Count`, and its tests |

Next: interfaces, then a small in-memory key-value store over HTTP, then concurrency and the race detector.

## Run it

```bash
cd stage-01-basics
go run ./wordcount
go test ./...
go vet ./...
```

Each folder except `readfile/` is its own small `package main`. Run one at a time with `go run ./<folder>`. The programs that read a
file expect to run from `stage-01-basics/`, because the paths are relative.

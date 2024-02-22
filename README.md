# Ultimate Go ∴ February 2024

Miki Tebeka
📬 [miki@ardanlabs.com](mailto:miki@ardanlabs.com), 𝕏 [@tebeka](https://twitter.com/tebeka), 👨 [mikitebeka](https://www.linkedin.com/in/mikitebeka/), ✒️[blog](https://www.ardanlabs.com/blog/)

#### Shameless Plugs

- [Ardan Labs Courses](https://www.ardanlabs.com/training/)
- [LinkedIn Learning Classes](https://www.linkedin.com/learning/instructors/miki-tebeka)
- [Books](https://pragprog.com/search/?q=miki+tebeka)

---


## Day 1

### Agenda

- Programming vs engineering
- Understanding memory allocations:
    - Value semantics
    - Pointer semantics & ownership rules
    - Stack vs heap allocations and how they affect your code
        - Detecting and understanding escape analysis
- How the garbage collector (GC) works and its effect on performance
    - Common causes for memory leaks
- Writing CPU cache friendly code
- Working with slices in memory efficient way
    - Range loop semantics

### Code

- [inc.go](inc/inc.go) - Values & pointers
- [user.go](user/user.go) - Escape analysis
- [units.go](units/units.go) - Value vs pointer semantics
- [marshal.go](marshal/marshal.go) - Pointer semantics in marshaling
- [tokenizer.go](tokenizer/tokenizer.go) - A look at the GC
- [author.go](author/author.go) - Leaking memory
- [matrix.go](matrix/matrix.go) - CPU cache friendly code
- [bank.go](bank/bank.go) - Slices "for" loop semantics

### Links

- [Effective Go](https://go.dev/doc/effective_go) - Keep this under your pillow
- [Go proverbs](https://go-proverbs.github.io/) - Reflect about them
- [gotraining](https://github.com/ardanlabs/gotraining) GitHub repository
    - [Profiling a Larger Web Service](https://github.com/ardanlabs/gotraining/blob/master/topics/go/profiling/project/README.md)
- [What is Software Engineering?](https://research.swtch.com/vgo-eng) by Russ Cox
- [Garbage Collection in Go](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)
- [A Guide to the Go Garbage Collector](https://go.dev/doc/gc-guide)
- [gctrace output](https://pkg.go.dev/github.com/chanshik/gotraining/topics/profiling/godebug/gctrace#section-readme)
- [slice implementation](https://github.com/golang/go/blob/master/src/runtime/slice.go#L15)
- [Tri color marking](https://en.wikipedia.org/wiki/Tracing_garbage_collection#Tri-color_marking)
- [Organizing a Go Module](https://go.dev/doc/modules/layout)

### Data & Other

- [inc by value](_extra/inc-value.png)
- [inc by ptr](_extra/inc-ptr.png)
- [By value & by pointer](_extra/value-ptr.png)
- [CPU cache](_extra/cpu-cache.png)
    - [Matrix Sum](_extra/matrix-sum.png)
- [Slices](_extra/slice.png)

---

## Day 2

### Agenda

- Data oriented design
    - Choosing value vs pointer semantics
    - Built-in types and user defined types
- Polymorphism with interfaces
    - Modeling interfaces
    - How interfaces are implemented - understanding cost
- Using embedding
    - Difference from traditional OO
- Using generics
    - Generics functions
    - Generic data structures
- Method set rules

### Code

- [types.go](types/types.go) - Basic types (mostly slices)
- [freq.go](freq/freq.go) - Working with maps
- [mana.go](mana/mana.go) - Iterating over maps
- [etl.go](etl/etl.go) - Discovering interfaces
- [log.go](log/log.go) - Interfaces are on the heap
- [client.go](client/client.go) - Using interfaces for mocking errors
- [flag.go](flag/flag.go) - Implementing `fmt.Stringer`
- [auth.go](auth/auth.go) - Custom errors (and a bug)
- [admin.go](admin/admin.go) - Struct embedding
- [driver.go](driver/driver.go) - Embedding is not inheritance
- [deep.go](deep/deep.go) - Using generics


### Links

- [Go Slices: usage and internals](https://go.dev/blog/slices-intro)
- [Methods, interfaces & embedded types in Go](https://www.ardanlabs.com/blog/2014/05/methods-interfaces-and-embedded-types.html)
- [Methods & Interfaces](https://go.dev/tour/methods/1) in the Go tour
- [sort examples](https://pkg.go.dev/sort/#pkg-examples) - Read and try to understand
- [Generics can make your Go code slower](https://planetscale.com/blog/generics-can-make-your-go-code-slower)
- [When to use generics](https://go.dev/blog/when-generics)
- [Generics tutorial](https://go.dev/doc/tutorial/generics)
- [Method Sets](https://www.youtube.com/watch?v=Z5cvLOrWlLM)
    - [Formal definition](https://go.dev/ref/spec#Method_sets)
- [Domain Driven, Data Oriented Architecture with Bill Kennedy](https://www.youtube.com/watch?v=bQgNYK1Z5ho)

### Data & Other

- [Error cost](_extra/cost.png)

---

## Day 3

### Agenda
- The OS scheduler
    - Design and trade-offs, how they affect you
- The Go scheduler
- CPU vs I/O bound jobs
- Synchronization and orchestration
    - Understanding channels
    - Using sync and sync/atomic
- Concurrency patterns
    - Wait for results
    - Fan out
    - Polling
    - …

### Code

- [stack.go](stack/stack.go) - Generic data structures
- [sched.go](sched/sched.go) - Go scheduler example
- [chan.go](chan/chan.go) - Goroutines, channels and patterns
- [counter.go](counter/counter.go) - Race conditions, mutex and atomic
- [db.go](db/db.go) - RWMutex

### Links

- [Kubernetes CPU Limits and Go](https://www.ardanlabs.com/blog/2024/02/kubernetes-cpu-limits-go.html)
- [automaxprocs](https://pkg.go.dev/go.uber.org/automaxprocs)
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576/photo/1)
- [Scheduling in Go](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part2.html) by Bill Kennedy
- [The race detector](https://go.dev/doc/articles/race_detector)
- [errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup)
- [Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
- [Go Concurrency Patterns: Context](https://go.dev/blog/context)
- [Curious Channels](https://dave.cheney.net/2013/04/30/curious-channels)
- [The Behavior of Channels](https://www.ardanlabs.com/blog/2017/10/the-behavior-of-channels.html)
- [Channel Semantics](https://www.353solutions.com/channel-semantics)
- [Why are there nil channels in Go?](https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308)
- [Amdahl's Law](https://en.wikipedia.org/wiki/Amdahl%27s_law) - Limits of concurrency
- [Concurrency is not Parallelism](https://www.youtube.com/watch?v=cN_DpYBzKso) by Rob Pike
- [ArdanLabs service repo](https://github.com/ardanlabs/service/)
- [conc](https://github.com/sourcegraph/conc) - Concurrency package

### Data & Other

- [taxi.tar](https://storage.googleapis.com/353solutions/c/data/taxi.tar)

---

## Day 4

### Agenda

- Before you start optimizing
- CPU profiling
    - Writing benchmarks
    - Using the profiler
    - Using the execution traces
- Memory profiling
    - Benchmarking memory
    - Using the GC tracer
    - Using the execution tracer

### Code


- [token.go](token/token.go) - Refreshing a token while clients wait
- [cache.go](cache/cache.go) - Benchmarking right data
- [tokenizer](tokenizer) - CPU optimization
- [search](search) - Memory optimization

### Links

- [The Rules of Optimization Club](https://wiki.c2.com/?RulesOfOptimizationClub)
- [Rob Pike's 5 Rules of Programming](https://users.ece.utexas.edu/~adnan/pike.html)
- [Computer Latency at Human Scale](https://twitter.com/jordancurve/status/1108475342468120576)
- [High Performance Go Workshop](https://dave.cheney.net/high-performance-go-workshop/gophercon-2019.html)
- [Profile-guided optimization](https://go.dev/doc/pgo)
- [So you wanna go fast](https://www.slideshare.net/TylerTreat/so-you-wanna-go-fast-80300458)
- [Miki's Optimization Overview](_extra/optimize.md)
- [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) - Compare benchmarks
- [Graphviz](https://graphviz.org/) - Generate graphs (used by `pprof`)
- [Open Telemetry (otel)](https://github.com/open-telemetry/opentelemetry-go)
- [Profile-guided optimization](https://go.dev/doc/pgo)
- [Uber Go Style Guide](https://github.com/uber-go/guide/)

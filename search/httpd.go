package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"runtime/trace"
)

type Server struct {
	db  *DB
	log *log.Logger
}

func (s *Server) searchHandler(w http.ResponseWriter, r *http.Request) {
	ctx, task := trace.NewTask(r.Context(), "search")
	defer task.End()

	if r.Method != http.MethodGet {
		s.log.Printf("ERROR: bad method - %s (%s)", r.Method, r.RemoteAddr)
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	reg := trace.StartRegion(ctx, "query")
	query := r.URL.Query().Get("q")
	if query == "" {
		s.log.Printf("ERROR: missing query (%s)", r.RemoteAddr)
		http.Error(w, "missing query", http.StatusBadRequest)
		return
	}
	reg.End()

	reg = trace.StartRegion(ctx, "search")
	matches := s.db.Search(query)
	if len(matches) == 0 {
		s.log.Printf("ERROR: no matches for %q", query)
		http.Error(w, "no matches", http.StatusNotFound)
		return
	}
	reg.End()

	reg = trace.StartRegion(ctx, "json")
	s.log.Printf("INFO: %q - %d matches", query, len(matches))
	reply := make([]map[string]any, len(matches))
	for i, item := range matches {
		reply[i] = map[string]any{
			"title":   item.Title,
			"content": item.Content,
			"link":    item.Link,
		}
	}

	if err := json.NewEncoder(w).Encode(reply); err != nil {
		s.log.Printf("ERROR: can't encode - %s", err)
	}
	reg.End()
}

func main() {
	s := Server{
		db:  NewDB(),
		log: log.New(os.Stdout, "[search] ", log.LstdFlags|log.Lshortfile),
	}

	s.log.Printf("INFO: %d entries in the database", len(dbData))

	mux := http.NewServeMux()
	mux.HandleFunc("/search", s.searchHandler)
	// TODO: Guard with build-tag or user/password or ...
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	const addr = ":8080"
	srv := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	s.log.Printf("INFO: starting on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}

// CPU
// curl http://localhost:8080/search?q=president
// hey -c 100 -z 15s http://localhost:8080/search?q=president
// go tool pprof -http :9999 http://localhost:8080/debug/pprof?seconds=10

// Memory
// go test -run NONE -bench . -benchmem
// go test -run NONE -bench . -memprofile mem.pprof
// go tool pprof -http :9999 mem.pprof

// Trace
// curl -o trace.out http://localhost:8080/debug/trace?seconds=10
// go tool trace -http :9999 trace.out

/* Journal

Initial:
	RPS: 3991
	GC mark: 5ms

Lower case
	RPS: 16098
	GC mark: 4ms

Index
	RPS: 17868
	GC mark: 3.3ms


go test -bench . -benchmem (initial, lower, index)
BenchmarkQuery-12           7575            151284 ns/op           75176 B/op        505 allocs/op
BenchmarkQuery-12          75604             14630 ns/op            7504 B/op          7 allocs/op
BenchmarkQuery-12        2394890               459.7 ns/op          2304 B/op          1 allocs/op
*/

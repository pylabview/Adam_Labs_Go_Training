package tokenizer

import "testing"

var text = `
TO THE RED-HEADED LEAGUE: On account of the bequest of the late
Ezekiah Hopkins, of Lebanon, Pennsylvania, U.S.A., there is now another
vacancy open which entitles a member of the League to a salary of £ 4 a
week for purely nominal services. All red-headed men who are sound in
body and mind and above the age of twenty-one years, are eligible.
Apply in person on Monday, at eleven o’clock, to Duncan Ross, at the
offices of the League, 7 Pope’s Court, Fleet Street.
`

func BenchmarkTokenize(b *testing.B) {
	// for i := 0; i < b.N; i++ {
	for range b.N { // go1.22+
		tokens := Tokenize(text)
		if len(tokens) != 86 {
			b.Fatal(len(tokens))
		}
	}
}

/* Speed
- Initial:      	20000
- Manual split:  	 5700
- Pre-allocate fs:   4900
- Pre-allocated tok: 3800
*/

// Profile
// go test -run NONE -bench . -count 5 -cpuprofile cpu.pprof
// View
// go tool pprof -http :8080 cpu.pprof

// --- Looking at memory ---
// go test -bench . -benchmem -count 5

// GODEBUG=gctrace=1 go test -bench .
// GODEBUG=gctrace=1,gcpacertrace=1 go test -bench .

// go test -bench . -trace trace.out
// go tool trace -http :8080 trace.out

# benchdiff

A simple tool that computes diff between two different benchmarks.

Here is a simple benchmark function, we want to compare the performance between `bytes.Buffer` and `strings.Builder`:
```go
package main_test

import (
	"bytes"
	"strings"
	"testing"
)

var result string

func BenchmarkExample(b *testing.B) {
	b.Run(
		"strings.Builder", func(b *testing.B) {
			buf := strings.Builder{}
			buf.Write([]byte("Hello world!"))
			for i := 0; i < b.N; i++ {
				result = buf.String()
			}
		},
	)
	b.Run(
		"bytes.Buffer", func(b *testing.B) {
			buf := bytes.Buffer{}
			buf.Write([]byte("Hello world!"))
			for i := 0; i < b.N; i++ {
				result = buf.String()
			}
		},
	)
}

```

We run the test and analyze it using the `bencstat` tool:
```shell
$ go test -run='^$' -bench=. -benchmem -count 5 | tee result.txt
$ benchstat result.txt
```

But `bencstat` shows each summary only:
```
name                        time/op
Example/strings.Builder-16  0.71ns ± 1%
Example/bytes.Buffer-16     21.9ns ± 0%

name                        alloc/op
Example/strings.Builder-16   0.00B
Example/bytes.Buffer-16      16.0B ± 0%

name                        allocs/op
Example/strings.Builder-16    0.00
Example/bytes.Buffer-16       1.00 ± 0%

```

Wants to get a diff? try `benchdiff`:

```shell
$ go install github.com/j178/benchdiff@latest
$ bechdiff -name 'Buffer-vs-Builder' -pat 'bytes.Buffer,strings.Builder' result.txt
```

Here is the nice diff result:
```
name               old time/op    new time/op    delta
Buffer-vs-Builder    21.9ns ± 0%     0.7ns ± 1%   -96.78%  (p=0.016 n=4+5)
Buffer-vs-Builder    21.9ns ± 0%     0.7ns ± 1%   -96.78%  (p=0.016 n=4+5)
Buffer-vs-Builder    21.9ns ± 0%     0.7ns ± 1%   -96.78%  (p=0.016 n=4+5)
Buffer-vs-Builder    21.9ns ± 0%     0.7ns ± 1%   -96.78%  (p=0.016 n=4+5)
Buffer-vs-Builder    21.9ns ± 0%     0.7ns ± 1%   -96.78%  (p=0.016 n=4+5)
Buffer-vs-Builder    21.9ns ± 0%     0.7ns ± 1%   -96.78%  (p=0.016 n=4+5)

name               old alloc/op   new alloc/op   delta
Buffer-vs-Builder     16.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder     16.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder     16.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder     16.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder     16.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder     16.0B ± 0%      0.0B       -100.00%  (p=0.008 n=5+5)

name               old allocs/op  new allocs/op  delta
Buffer-vs-Builder      1.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder      1.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder      1.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder      1.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder      1.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
Buffer-vs-Builder      1.00 ± 0%      0.00       -100.00%  (p=0.008 n=5+5)
```
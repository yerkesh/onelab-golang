package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		n, dup, prev, i int
	)

	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Fscan(reader, &n)
	writer := bufio.NewWriter(os.Stdout)
	for i = 0; i < n; i++ {
		_, _ = fmt.Fscan(reader, &dup)

		if prev != dup {
			prev = dup
			_, _ = fmt.Fprintf(writer, "%d", dup)
			_ = writer.WriteByte('\n')
		} else if i == 0 && dup == 0 {
			_, _ = fmt.Fprintf(writer, "%d", dup)
			_ = writer.WriteByte('\n')
		}
	}
	_ = writer.Flush()
}


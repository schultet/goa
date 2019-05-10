package util

// Buffer provides an unbounded buffer between in and out. Buffer exits when in
// is closed and all items in the buffer have been sent to out, at which point
// it closes out.
func Buffer(in <-chan int, out chan<- int) {
	var buf []int
	for in != nil || len(buf) > 0 {
		var i int
		var c chan<- int
		if len(buf) > 0 {
			i = buf[0]
			c = out // enable send case
		}
		select {
		case n, ok := <-in:
			if ok {
				buf = append(buf, n)
			} else {
				in = nil // disable receive case
			}
		case c <- i:
			buf = buf[1:]
		}
	}
	close(out)
}

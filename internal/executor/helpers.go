package executor

import "io"

// coalescReader returns the first non-nil io.Reader
func coalescReader(vals ...io.Reader) io.Reader {
	for _, v := range vals {
		if v != nil {
			return v
		}
	}
	return nil
}

// coalescWriter returns the first non-nil io.Writer
func coalescWriter(vals ...io.Writer) io.Writer {
	for _, v := range vals {
		if v != nil {
			return v
		}
	}
	return nil
}

// closeAll closes all closers, ignoring nil values
func closeAll[T io.Closer](closers []T) {
	for _, c := range closers {
		closer := any(c).(io.Closer)
		if closer != nil {
			closer.Close()
		}
	}
}

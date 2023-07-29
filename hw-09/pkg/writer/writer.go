package writer

import "io"

func Write(w io.Writer, a ...any) {
	for _, v := range a {
		if val, ok := v.(string); ok {
			w.Write([]byte(val))
		}
	}
}

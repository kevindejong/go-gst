package app

// #include <gst/gst.h>
import "C"

import (
	"github.com/tinyzimmer/go-gst/gst"
)

func wrapAppSink(elem *gst.Element) *Sink  { return &Sink{elem} }
func wrapAppSrc(elem *gst.Element) *Source { return &Source{elem} }

// gobool provides an easy type conversion between C.gboolean and a go bool.
func gobool(b C.gboolean) bool { return int(b) > 0 }

// gboolean converts a go bool to a C.gboolean.
func gboolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

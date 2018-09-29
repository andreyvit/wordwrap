// Package wordwrap splits a bunch of text into lines.
package wordwrap

import (
	"io"
	"strings"
)

// Options adjust word-wrapping behavior.
type Options struct {
	// NoWrap disables word wrapping, so that only the existing line breaks are used.
	NoWrap bool

	// BreakWords allows to break the line mid-word if absolutely necessary.
	BreakWords bool

	// BreakMarker is appended to a line that is broken mid-word.
	// It counts as a single character for the purposes of width computation.
	BreakMarker string
}

var eol = []byte("\n")

/*
WrapString splits the text into lines up to the given width, and returns the
result, including the trailing end-of-line character.

Any pre-existing line breaks are preserved. If opt.BreakWords is true,
the resulting lines are guaranteed to fit into the given width. If the width
is zero or opt.NoWrap is true, no new line breaks will be added.
*/
func WrapString(text string, width int, opt Options) string {
	var buf strings.Builder
	_, _ = WrapTo(&buf, text, width, opt)
	return buf.String()
}

/*
WrapTo splits the text into lines up to the given width, and writes the resulting
lines into the given io.Writer, including the trailing end-of-line character.

Any pre-existing line breaks are preserved. If opt.BreakWords is true,
the resulting lines are guaranteed to fit into the given width. If the width
is zero or opt.NoWrap is true, no new line breaks will be added.
*/
func WrapTo(w io.Writer, text string, width int, opt Options) (n int, err error) {
	Wrap(text, width, opt, func(line string) {
		if err != nil {
			return
		}

		var k int
		k, err = io.WriteString(w, line)
		n += k
		if err != nil {
			return
		}

		k, err = w.Write(eol)
		n += k
	})
	return
}

/*
Estimate predicts the number of lines in the given text for memory allocation
purposes.
*/
func Estimate(text string, width int) int {
	explicit := strings.Count(text, "\n") + 1

	if width == 0 {
		return explicit
	} else {
		return explicit + (len(text)+width-1)/width
	}
}

/*
WrapSlice splits the text into lines up to the given width, and returns the lines
as a slice of strings.

Any pre-existing line breaks are preserved. If opt.BreakWords is true,
the resulting lines are guaranteed to fit into the given width. If the width
is zero or opt.NoWrap is true, no new line breaks will be added.
*/
func WrapSlice(text string, width int, opt Options) []string {
	lines := make([]string, 0, Estimate(text, width))
	Wrap(text, width, opt, func(line string) {
		lines = append(lines, line)
	})
	return lines
}

/*
Wrap splits the text into lines up to the given width, and calls linef for each
line.

Any pre-existing line breaks are preserved. If opt.BreakWords is true,
the resulting lines are guaranteed to fit into the given width. If the width
is zero or opt.NoWrap is true, no new line breaks will be added.
*/
func Wrap(text string, width int, opt Options, linef func(line string)) {
	for {
		i := strings.IndexByte(text, byte('\n'))
		if i < 0 {
			wrapLine(text, width, opt, linef)
			return
		}
		wrapLine(text[:i], width, opt, linef)
		text = text[i+1:]
	}
}

func wrapLine(text string, w int, opt Options, linef func(line string)) {
	if len(text) == 0 || w == 0 || opt.NoWrap {
		linef(text)
		return
	}

	for {
		n := len(text)
		if n == 0 {
			break
		}

		var b int
		var marker bool
		if n <= w {
			b = n
		} else if i := strings.LastIndexByte(text[:min(w+1, len(text))], ' '); i >= 0 {
			b = i
		} else if opt.BreakWords {
			if len(opt.BreakMarker) > 0 && w > 1 {
				b = w - 1
			} else {
				b = w
			}
			marker = true
		} else {
			i := strings.IndexByte(text[min(w, len(text)):], ' ')
			if i < 0 {
				b = n
			} else {
				b = w + i
			}
		}

		s := rtrim(text, b)
		if s > 0 {
			line := text[:s]
			if marker {
				linef(line + opt.BreakMarker)
			} else {
				linef(line)
			}
		}

		text = text[ltrim(text, b):]
	}
}

func ltrim(s string, p int) int {
	n := len(s)
	for p < n && s[p] == ' ' {
		p++
	}
	return p
}

func rtrim(s string, p int) int {
	for p > 0 && s[p-1] == ' ' {
		p--
	}
	return p
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

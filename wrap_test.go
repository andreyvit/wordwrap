package wordwrap

import (
	"testing"
)

var (
	nobr = Options{}
	br   = Options{BreakWords: true}
	dash = Options{BreakWords: true, BreakMarker: "-"}
)

func TestWrap(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		w        int
		opt      Options
	}{
		{"abc def ghi", "abc def ghi\n", 40, nobr},
		{"abc def ghi", "abc def ghi\n", 11, nobr},
		{"abc def ghi", "abc def\nghi\n", 10, nobr},
		{"abc def ghi", "abc def\nghi\n", 9, nobr},
		{"abc def ghi", "abc def\nghi\n", 8, nobr},
		{"abc def ghi", "abc def\nghi\n", 7, nobr},
		{"abc def ghi", "abc\ndef\nghi\n", 6, nobr},
		{"abc def ghi", "abc\ndef\nghi\n", 5, nobr},
		{"abc def ghi", "abc\ndef\nghi\n", 4, nobr},
		{"abc def ghi", "abc\ndef\nghi\n", 3, nobr},

		// space handling
		{" abc defg hij", " abc\ndefg\nhij\n", 4, nobr},
		{"abc defg hij ", "abc\ndefg\nhij\n", 4, nobr},
		{"abc defg hij  ", "abc\ndefg\nhij\n", 4, nobr},
		{"abc defg hij        ", "abc\ndefg\nhij\n", 4, nobr},
		{"abc  def  ghi", "abc\ndef\nghi\n", 4, nobr},
		{"abc  def  ghi", "abc\ndef\nghi\n", 3, nobr},
		{"abcdef  ghi", "abcdef\nghi\n", 5, nobr},

		// BreakWords == false
		{"abcdef ghi", "abcdef ghi\n", 10, nobr},
		{"abcdef ghi", "abcdef\nghi\n", 9, nobr},
		{"abcdef ghi", "abcdef\nghi\n", 8, nobr},
		{"abcdef ghi", "abcdef\nghi\n", 7, nobr},
		{"abcdef ghi", "abcdef\nghi\n", 6, nobr},
		{"abcdef ghi", "abcdef\nghi\n", 5, nobr},
		{"abcdef ghi", "abcdef\nghi\n", 4, nobr},
		{"abcdef ghi", "abcdef\nghi\n", 3, nobr},

		// BreakWords
		{"abcdefg", "abcdefg\n", 40, br},
		{"abcdefg", "abcdefg\n", 7, br},
		{"abcdefg", "abcdef\ng\n", 6, br},
		{"abcdefg", "abcde\nfg\n", 5, br},
		{"abcdefg", "abcd\nefg\n", 4, br},
		{"abcdefg", "abc\ndef\ng\n", 3, br},
		{"abcdefg", "ab\ncd\nef\ng\n", 2, br},
		{"abcdefg", "a\nb\nc\nd\ne\nf\ng\n", 1, br},

		// BreakWords + BreakMarker
		{"abcdefg", "abcdefg\n", 7, dash},
		{"abcdefg", "abcde-\nfg\n", 6, dash},
		{"abcdefg", "abcd-\nefg\n", 5, dash},
		{"abcdefg", "abc-\ndefg\n", 4, dash},
		{"abcdefg", "ab-\ncd-\nefg\n", 3, dash},
		{"abcdefg", "a-\nb-\nc-\nd-\ne-\nfg\n", 2, dash},
		{"abcdefg", "a-\nb-\nc-\nd-\ne-\nf-\ng\n", 1, dash},

		{"abc\ndef\nghi", "abc\ndef\nghi\n", 40, br},
		{"abc def\n\nghi", "abc\ndef\n\nghi\n", 4, nobr},
	}
	for _, tt := range tests {
		actual := WrapString(tt.input, tt.w, tt.opt)
		if actual != tt.expected {
			t.Errorf("** Wrap(%q, %d) == %q, expected %q", tt.input, tt.w, actual, tt.expected)
		} else {
			t.Logf("✓ Wrap(%q, %d) == %q", tt.input, tt.w, actual)
		}
	}
}

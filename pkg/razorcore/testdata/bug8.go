// This file is generated by gorazor 1.2.2
// DON'T modified manually
// Should edit source file and re-generate: cases/bug8.gohtml

package cases

import (
	"github.com/sipin/gorazor/gorazor"
	"io"
	"strings"
)

// Bug8 generates cases/bug8.gohtml
func Bug8(l *Locale) string {
	var _b strings.Builder
	RenderBug8(&_b, l)
	return _b.String()
}

// RenderBug8 render cases/bug8.gohtml
func RenderBug8(_buffer io.StringWriter, l *Locale) {
	// Line: 3
	_buffer.WriteString("\n<span>")
	// Line: 4
	_buffer.WriteString(gorazor.HTMLEscape(l.T("for")))
	// Line: 4
	_buffer.WriteString("</span>")

}

// This file is generated by gorazor 2.0.1
// DON'T modified manually
// Should edit source file and re-generate: cases/comment.gohtml

package cases

import (
	"io"
	"strings"
)

func Comment() string {
	var _b strings.Builder
	RenderComment(&_b)
	return _b.String()
}

func RenderComment(_buffer io.StringWriter) {
	_buffer.WriteString("\n\n\n\n<p>hello </p>")

	hello

}
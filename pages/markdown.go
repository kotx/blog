package pages

import (
	"bytes"

	"github.com/a-h/templ"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/parser"
	"go.abhg.dev/goldmark/anchor"
)

func markdown(source []byte) templ.Component {
	var buf bytes.Buffer
	err := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithExtensions(
			&anchor.Extender{},
			highlighting.NewHighlighting(
				highlighting.WithStyle("xcode-dark"),
			),
		),
	).Convert(source, &buf)

	return templ.Raw(buf.String(), err)
}

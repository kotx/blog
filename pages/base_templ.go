// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"strconv"
	"time"
)

func base(title string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/base.templ`, Line: 13, Col: 17}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><link rel=\"stylesheet\" href=\"/site.css\"><link rel=\"preconnect\" href=\"https://rsms.me/\"><link rel=\"stylesheet\" href=\"https://rsms.me/inter/inter.css\"></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<details id=\"comments\"><summary id=\"commentsSummary\">Comments</summary><noscript>JavaScript must be enabled to view comments here.</noscript><noscript>Alternatively, view on <a href=\"https://github.com/kotx/blog/discussions\">GitHub Discussions</a>!</noscript><script>\n\t\t\t\t\tlet commentsLoaded = false;\n\t\t\t\t\tconst comments = document.getElementById(\"comments\");\n\t\t\t\t\tconst commentsSummary = document.getElementById(\"commentsSummary\");\n\n\t\t\t\t\tcomments.addEventListener(\"toggle\", (event) => {\n\t\t\t\t\t\tif (comments.open && !commentsLoaded) {\n\t\t\t\t\t\t\tcommentsSummary.innerText = \"Comments (loading...)\";\n\t\t\t\t\t\t}\n\t\t\t\t\t});\n\n\t\t\t\t\tfunction handleGiscusEvent(event) {\n\t\t\t\t\t\tif (event.origin !== 'https://giscus.app') return;\n\t\t\t\t\t\tif (!(typeof event.data === 'object' && event.data.giscus)) return;\n\n\t\t\t\t\t\tif (event.data.giscus.resizeHeight > 0) {\n\t\t\t\t\t\t\tcommentsSummary.innerText = \"Comments\";\n\t\t\t\t\t\t\tcommentsLoaded = true;\n\t\t\t\t\t\t}\n\t\t\t\t\t}\n\t\t\t\t\n\t\t\t\t\twindow.addEventListener(\"message\", handleGiscusEvent);\n\t\t\t\t</script><script src=\"https://giscus.app/client.js\" data-repo=\"kotx/blog\" data-repo-id=\"R_kgDOMYvuow\" data-category=\"Announcements\" data-category-id=\"DIC_kwDOMYvuo84ChCB-\" data-mapping=\"pathname\" data-strict=\"1\" data-reactions-enabled=\"1\" data-emit-metadata=\"0\" data-input-position=\"top\" data-theme=\"noborder_gray\" data-lang=\"en\" data-loading=\"lazy\" crossorigin=\"anonymous\" async>\n\t\t\t\t</script></details><footer>&copy; Kot ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(time.Now().Year()))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `pages/base.templ`, Line: 70, Col: 48}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" - <a href=\"/rss.xml\">rss</a> - <a href=\"https://github.com/kotx/blog\" target=\"_blank\">source</a></footer></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

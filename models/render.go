package models

import (
	"net/http"
	"text/template"
)


func (g *Game) RenderBoardHTML(w http.ResponseWriter, r *http.Request) {
    board := g.GetPrintableBoard()

    tmpl := template.Must(template.New("board").Parse(`
        <html><body>
        <table>
            {{range .}}
            <tr>{{range .}}<td>{{.}}</td>{{end}}</tr>
            {{end}}
        </table>
        </body></html>
    `))

    tmpl.Execute(w, board)
}

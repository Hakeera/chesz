
package main

import (
	"chesz/routes"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer implementa echo.Renderer
type TemplateRenderer struct {
	templates *template.Template
}

// Render renderiza um template HTML
func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	// Você pode injetar dados globais aqui se quiser, por exemplo, user logado etc.
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	// Compila todos os templates HTML
	tmpl := template.Must(template.ParseGlob("view/**/*.html")) 

	renderer := &TemplateRenderer{
		templates: tmpl,
	}

	// Instância do Echo
	e := echo.New()

	// Define pasta de arquivos estáticos (CSS, JS, imagens)

	// Define o renderer customizado
	e.Renderer = renderer

	// Define rotas
	routes.SetUpRoutes(e)

	// Starta o servidor
	e.Logger.Fatal(e.Start(":1323"))
}

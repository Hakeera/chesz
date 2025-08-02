package main

import (
	"chesz/routes"
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer implementa echo.Renderer
type TemplateRenderer struct {
	templates *template.Template
}

// Render renderiza um template HTML
func (t *TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	err := t.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("Erro ao renderizar template %q: %v\n", name, err)
	}
	return err
}

func main() {

	// Compila todos os templates HTML e adiciona funções auxiliares
	renderer := &TemplateRenderer{
		templates: template.Must(template.New("").Funcs(template.FuncMap{
			"add": func(a, b int) int { return a + b },
			"sub": func(a, b int) int { return a - b },
			"neg": func(a int) int { return -a },
		}).ParseGlob("view/**/*.html")),
	}

	// Instância do Echo
	e := echo.New()

	// Define pasta de arquivos estáticos (CSS, JS, imagens)
	e.Static("/static", "view/static")

	// Define o renderer customizado
	e.Renderer = renderer

	// Define rotas
	routes.SetUpRoutes(e)

	// Starta o servidor
	e.Logger.Fatal(e.Start(":1323"))
}

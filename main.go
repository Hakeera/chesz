package main

import (
	"html/template"
	"io"
	"log"

	"chesz/models"
	"chesz/routes"

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
			    "mod": func(a, b int) int { return a % b },
			    "eq":  func(a, b any) bool { return a == b },
			    "ne":  func(a, b any) bool { return a != b },
			    "seq": func(start, end int) []int {
				    var result []int
				    if start <= end {
					    for i := start; i <= end; i++ {
						    result = append(result, i)
					    }
				    } else {
					    for i := start; i >= end; i-- {
						    result = append(result, i)
					    }
				    }
				    return result
			    },
		    }).ParseGlob("view/**/*.html")),
	    }

	    // Instância do Echo
	    e := echo.New()

	// Define pasta de arquivos estáticos (CSS, JS, imagens)
	e.Static("/static", "view/static")

	// Define o renderer customizado
	e.Renderer = renderer

	// Inicia o Jogo
	models.CurrentGame = models.NewGame()
	go models.CurrentGame.PlayLoop()

	// Define rotas
	routes.SetUpRoutes(e)

	// Starta o servidor
	e.Logger.Fatal(e.Start(":1323"))
}

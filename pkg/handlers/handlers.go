package handlers

import (
	"github.com/GuskkovD/go-course/pkg/config"
	"github.com/GuskkovD/go-course/pkg/models"
	"github.com/GuskkovD/go-course/pkg/render"
	"net/http"
)

//Repo is the repository used
var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a repos
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the new handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	//perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again "

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

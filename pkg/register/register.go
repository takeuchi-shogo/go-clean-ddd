package register

import (
	"log"

	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/config"
	"github.com/takeuchi-shogo/go-clean-ddd/pkg/infra/database"
)

type Registry struct {
	repo *Repository
}

func New() *Registry {
	c := config.New("LOCAL")
	db, err := database.NewDB(*c)
	if err != nil {
		log.Fatalln("failed database connection")
	}
	repo, err := NewRepository(db)
	if err != nil {
		log.Fatalln("failed create repository")
	}
	return &Registry{
		repo: repo,
	}
}

func (r *Registry) Repository() *Repository {
	return r.repo
}

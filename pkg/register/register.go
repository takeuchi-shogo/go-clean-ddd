package register

type Registry struct {
	repo Repository
}

func New() *Registry {
	// repo := NewRepository(db)
	return &Registry{
		// repo: repo,
	}
}

func (r *Registry) Repository() Repository {
	return r.repo
}

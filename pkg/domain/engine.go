package domain

type Engine struct {
	providers []Provider
}

func NewEngine(providers []Provider) Engine {
	return Engine{providers: providers}
}

func (e Engine) Search(query string) string {
	return e.providers[0].Search(query)
}

package domain

type Provider interface {
	Search(q string) string
}

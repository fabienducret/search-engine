package infra

type Google struct {
}

func (p Google) Search(q string) string {
	return q
}

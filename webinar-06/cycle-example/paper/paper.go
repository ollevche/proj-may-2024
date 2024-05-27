package paper

type Paper string

type Storer interface {
	Store(Paper)
	// Store(p Paper)
}

func (p Paper) StoreIn(pr Storer) {
	pr.Store(p)
}

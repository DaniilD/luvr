package category

type Category struct {
	Id         int
	Name       string
	SlugEn     string
	Categories *CategoryCollection
}

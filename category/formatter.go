package category

type CategoryFormatter struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func FormatCategory(category Category) CategoryFormatter {
	categoryFormatter := CategoryFormatter{}
	categoryFormatter.ID = category.ID
	categoryFormatter.Title = category.Title
	return categoryFormatter
}

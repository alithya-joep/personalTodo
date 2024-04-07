package dict

type Task struct {
	Status      Status
	Title       string
	Description string
}

func (t Task) GetTitle() string {
	return t.Title
}

func (t Task) GetDescription() string {
	return t.Description
}
func (t Task) FilterValue() string {
	return t.Title
}

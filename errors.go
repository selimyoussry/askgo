package askgo

// AddError adds an error to the errors list
func (t *Trv) AddError(err error) {
	t.Errors = append(t.Errors, err)
}

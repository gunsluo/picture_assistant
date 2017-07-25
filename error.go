package assistant

type Errors []error

// Errors returns itself.
func (es Errors) Errors() []error {
	return es
}

func (es Errors) Error() string {
	var err string
	for _, e := range es {
		err += e.Error() + ";"
	}
	return err
}

func (es Errors) add(err error) {
	es = append(es, err)
}

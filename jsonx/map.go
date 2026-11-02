package jsonx

func ToMap(v any) (map[string]any, error) {
	if v == nil {
		return nil, nil
	}

	m := make(map[string]any)
	err := Copy(&m, v)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MustToMap(v any) map[string]any {
	m, err := ToMap(v)
	if err != nil {
		panic(err)
	}
	return m
}

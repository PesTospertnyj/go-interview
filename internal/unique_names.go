package internal

func UniqueNames(a, b []string) []string {
	var result []string
	uniqueEntries := make(map[string]struct{})
	for _, aVal := range a {
		uniqueEntries[aVal] = struct{}{}
	}

	for _, bVal := range b {
		uniqueEntries[bVal] = struct{}{}
	}

	for key := range uniqueEntries {
		result = append(result, key)
	}

	return result
}

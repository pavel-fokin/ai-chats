package sqlite

import "strings"

func QueryIn[T any](query string, items []T) (string, []any) {
	if len(items) == 0 {
		return "", nil
	}

	placeholders := make([]string, 0, len(items))
	args := make([]any, 0, len(items))
	for _, item := range items {
		placeholders = append(placeholders, "?")
		args = append(args, item)
	}

	return query + " IN (" + strings.Join(placeholders, ", ") + ")", args
}

func isUniqueConstraintViolation(err error) bool {
	return strings.Contains(err.Error(), "UNIQUE constraint failed")
}

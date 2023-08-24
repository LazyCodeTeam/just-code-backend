package util

import "log/slog"

func MapSlice[TIn any, TOut any](items []TIn, mapper func(TIn) TOut) []TOut {
	result := make([]TOut, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}

	return result
}

func TryMapSlice[TIn any, TOut any](items []TIn, mapper func(TIn) (TOut, error)) ([]TOut, error) {
	result := make([]TOut, len(items))
	for i, item := range items {
		r, err := mapper(item)
		if err != nil {
			slog.Warn("Failed to map slice", "err", err, "input", item, "index", i)
			return nil, err
		}
		result[i] = r
	}

	return result, nil
}

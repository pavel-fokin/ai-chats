package json

import (
	"context"
	"encoding/json"
	"log/slog"
)

func MustMarshal(ctx context.Context, v interface{}) []byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		slog.ErrorContext(ctx, "failed to marshal json", "err", err)
		panic(err)
	}
	return bytes
}

func MustUnmarshal(ctx context.Context, data []byte, v interface{}) {
	if err := json.Unmarshal(data, v); err != nil {
		slog.ErrorContext(ctx, "failed to unmarshal json", "err", err)
		panic(err)
	}
}

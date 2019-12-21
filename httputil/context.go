package httputil

import "context"

// MustGetClient ...
// It panics if value was not found.
func MustGetClient(ctx context.Context) string {
	fp, ok := ctx.Value(ClientContextKey).(string)
	if !ok {
		panic("client not found in context")
	}
	return fp
}

// MustGetProduct ...
// It panics if value was not found.
func MustGetProduct(ctx context.Context) string {
	fp, ok := ctx.Value(ClientContextKey).(string)
	if !ok {
		panic("product not found in context")
	}
	return fp
}

// GetAcceptedLanguages extracts accepted languages from context
func GetAcceptedLanguages(ctx context.Context) []string {
	alangs, _ := ctx.Value(AcceptLanguageContextKey).([]string)
	return alangs
}

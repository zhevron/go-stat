// Package stat provides system statistics collection functions.
package stat

// MemInfo holds current memory usage information.
// Note that .Buffers and .Cached will always be 0 on Windows.
type MemInfo struct {
	Total   int64 // Total available system memory (in bytes).
	Free    int64 // Total free system memory (in bytes).
	Buffers int64 // Total system memory used for buffers (in bytes).
	Cached  int64 // Total system memory used for caching (in bytes).
}

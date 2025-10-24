package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func MapTryGet[K comparable, V any](m map[K]V, k K) (V, error) {
	if item, ok := m[k]; ok {
		return item, nil
	}
	var zeroValue V
	return zeroValue, fmt.Errorf("key not found: %v", k)
}

func MapTryGetDefault[K comparable, V any](m map[K]V, k K, defaultValue V) V {
	if item, ok := m[k]; ok {
		return item
	}
	return defaultValue
}

type LengthHaver interface {
	Len() int
}

func IsEmpty[T any](value T) bool {
	switch v := any(value).(type) {
	case string:
		return !IsEmptyString(v)
	case LengthHaver:
		return true
	case []T:
		return !IsListEmpty(v)
	default:
		return false
	}
}

func IsEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func IsListEmpty[T any](slice []T) bool {
	return len(slice) == 0
}

func IsEmptyLength(l LengthHaver) bool {
	return l.Len() == 0
}

func JoinArgs(args []string) string {
	return strings.Join(args, " ")
}

func HasMoreArgs(i int, args []string) bool {
	return i+1 < len(args)
}

func TryParseInt(s string) (int, error) {
	if s, err := strconv.Atoi(s); err == nil {
		return s, nil
	}
	return 0, fmt.Errorf("failed to parse int: %s", s)
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandInt(n int) int {
	return random.Intn(n)
}

func RandIntRange(start, end int) int {
	if start > end {
		start, end = end, start
	}

	return start + random.Intn(end-start+1)
}

func Capitalise(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}

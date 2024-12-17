package utils

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"path/filepath"
	"runtime"
)

func PrintMsg(msg string) {
	fmt.Println(msg)
}

type WithErr[T any] struct {
    Value  T
    Err error
}


func Must[T any](v T, err error, message ...string) T {
	if (err != nil) {
		if len(message) > 0 {
			panic(fmt.Errorf("%s: %v", message[0], err))
		}
		panic(err)
	}
	return v
}

func MustHaveKey[K comparable, V any](m map[K]V, key K) V {
	value, exists := m[key]
	if !exists {
		panic(fmt.Errorf("key %v not found", key))
	}
	return value
}

func MustHaveFile(filepath string) []string {
	file := Must(os.Open(filepath))
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return Must(lines, scanner.Err())
}

func GetInputFilePath() string {
	// NOTE: assumption that it is called from part1/2.go files
	_, currFile, _, ok := runtime.Caller(1)
	if !ok {
		panic("couldnt get current filename")
	}
	currDir := filepath.Dir(currFile)

	return filepath.Join(currDir, "input.txt")
}

func Map[T, U any](seq iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for a := range seq {
			mapped := f(a)
			if !yield(mapped) {
				return
			}
		}
	}
}

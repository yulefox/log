package core

import "strings"

func TrimPath(file string) string {
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return file
	}

	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return file
	}

	return file[idx+1:]
}

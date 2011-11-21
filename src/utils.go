package config

import "strings"

func ToString(name interface{}) string {
	return name.(string)
}

func TrimSpacesFromArray(val []string) []string {
    for i, v := range val {
        val[i] = strings.TrimSpace(v)
    }
    return val
}

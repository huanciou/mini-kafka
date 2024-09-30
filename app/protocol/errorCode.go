package protocol

import "fmt"

type Set map[uint16]struct{}

func (s Set) Add(element uint16) {
	s[element] = struct{}{}
}

func (s Set) Contains(element uint16) bool {
	_, exists := s[element]
	return exists
}

func NewSet(elements ...uint16) Set {
	set := Set{}
	for _, elem := range elements {
		set.Add(elem)
	}

	return set
}

type APIRule struct {
	ValidVersions Set
	DefaultCode   uint16
	ValidCode     uint16
}

var apiRules = map[uint16]APIRule{
	1: {
		ValidVersions: NewSet(0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16),
		DefaultCode:   35,
		ValidCode:     0,
	},
	18: {
		ValidVersions: NewSet(0, 1, 2, 3, 4),
		DefaultCode:   35,
		ValidCode:     0,
	},
}

func ErrorCodeChecker(apiKey, apiVersion uint16) uint16 {

	fmt.Println(apiKey, apiVersion)

	if rule, exists := apiRules[apiKey]; exists {
		if rule.ValidVersions.Contains(apiVersion) {
			return rule.ValidCode
		}

		return rule.DefaultCode
	}

	return uint16(35)
}

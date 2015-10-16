package location

import (
	"errors"
	"regexp"
)

type Location struct {
	rule *regexp.Regexp
	Root string
}

func InitLocation(rule string, root string) (*Location, error) {
	reg, err := regexp.Compile(rule)

	if err != nil {
		return nil, err
	}

	return &Location{reg, root}, nil
}

type LocationManager struct {
	locations []*Location
}

func (l *LocationManager) Match(path string) (*Location, error) {
	for _, loc := range l.locations {
		if loc.rule.MatchString(path) {
			return loc, nil
		}
	}

	return nil, errors.New("Location not found")
}

func InitLocationManager(locations []*Location) LocationManager {
	return LocationManager{locations}
}

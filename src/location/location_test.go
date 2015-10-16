package location

import (
	"testing"
)

func TestMatching(t *testing.T) {
	location1, _ := InitLocation("\\.(img|jpg)$", "path1")
	location2, _ := InitLocation("/lol/", "path2")
	location3, _ := InitLocation("/", "path3")

	locationManager := InitLocationManager([]*Location{location1, location2,
		location3})

	if match, _ := locationManager.Match("/lol.jpg"); match.Root != "path1" {
		t.Error("Error matching location: path1")
	}

	if match, _ := locationManager.Match("/lol/dimorinny"); match.Root != "path2" {
		t.Error("Error matching location: path2")
	}

	if match, _ := locationManager.Match("/abc/"); match.Root != "path3" {
		t.Error("Error matching location: path3")
	}

	if _, err := locationManager.Match("error"); err == nil {
		t.Error("Error matching location: error location")
	}
}

package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	yearSeconds := 31557600.0
	switch planet {
	case "Mercury":
		yearSeconds *= 0.2408467
	case "Venus":
		yearSeconds *= 0.61519726
	case "Mars":
		yearSeconds *= 1.8808158
	case "Jupiter":
		yearSeconds *= 11.862615
	case "Saturn":
		yearSeconds *= 29.447498
	case "Uranus":
		yearSeconds *= 84.016846
	case "Neptune":
		yearSeconds *= 164.79132
	case "Earth":
		yearSeconds *= 1
	default:
		return -1.0
	}
	return seconds / yearSeconds
}

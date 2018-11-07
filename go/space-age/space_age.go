package space

// Planet - The planet name
type Planet = string

var (
	// Planets orbital durations ralative to earth orbital duration
	relativeOrbitalPeriod = map[string]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	// Earth orbital duration in seconds
	earthOrbitDuration = 31557600.0
)

// Age - Get the age of a person on a planet
func Age(seconds float64, planet Planet) float64 {
	return seconds / earthOrbitDuration / relativeOrbitalPeriod[planet]
}

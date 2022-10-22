package scoring

import (
	"github.com/samber/lo"
)

type FeatureRegistration struct {
	Feature
	Title string
	Tags  []string
}

var registrations []FeatureRegistration

func RegisterScoringFeature(registration FeatureRegistration) {
	for _, it := range registrations {
		if it.Title == registration.Title {
			return
		}
	}

	registrations = append(registrations, registration)
}

func GetFeatures(includeTags ...string) FeatureSet {
	filteredRegistrations := lo.Filter[FeatureRegistration](registrations, func(it FeatureRegistration, _ int) bool {
		return len(includeTags) == 0 || lo.IsNotEmpty(len(lo.Intersect[string](includeTags, it.Tags)))
	})

	return lo.Map[FeatureRegistration, Feature](filteredRegistrations, func(it FeatureRegistration, _ int) Feature {
		return it.Feature
	})
}

package scoring

import (
	"github.com/samber/lo"
)

type FeatureRegistration struct {
	Feature
	Title string
	Tags  []string
}

type FeatureRegistry struct {
	registrations []FeatureRegistration
}

func NewDefaultRegistry() (r FeatureRegistry) {
	r.Register(affiliateLinkCountRegistration(), wordCountRegistration())
	return
}

func (f *FeatureRegistry) Register(registrations ...FeatureRegistration) {
	for _, registration := range registrations {
		f.registrations = append(f.registrations, registration)
	}
}

func (f *FeatureRegistry) GetFeatures(includeTags ...string) FeatureSet {
	filteredRegistrations := lo.Filter[FeatureRegistration](f.registrations, func(it FeatureRegistration, _ int) bool {
		return len(includeTags) == 0 || lo.Some[string](includeTags, it.Tags)
	})

	return lo.Map[FeatureRegistration, Feature](filteredRegistrations, func(it FeatureRegistration, _ int) Feature {
		return it.Feature
	})
}

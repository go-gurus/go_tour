## Website Scoring 

Let's build a simple website parser with Go!

----
### Objectives

Compute a credibility score for a website, taking into account

- Links. Affiliate Links might mean the website is bad
- Word Count. Much text might mean that the site has real content
----

### Strategy

- Implement a set of Features
- Build a feature registry
- Let the features compute an individual score
- Combine scores into one total score

----
### Types
#### Features

A feature computes a score from a document.
```golang
type Feature func(document *goquery.Document) float64
```

A set of features:
```golang
type FeatureSet []Feature
```
----
### Types
#### Feature Registry

A feature registry contains a set of registrations
```golang
var registrations []FeatureRegistration
```

A registration consists of:
```golang
type FeatureRegistration struct {
    Feature
    Title string
    Tags  []string
}
```
----
### Feature Registry
#### Registration and Deduplication

```golang
func RegisterScoringFeature(registration FeatureRegistration) {
	for _, it := range registrations {
		if it.Title == registration.Title {
			return
		}
	}
	registrations = append(registrations, registration)
}
```

----
### Feature Registry
#### Register Features

```golang
func MyFeature(doc *goquery.Document) float64 {
}
```

```golang
func init() {
    RegisterScoringFeature(FeatureRegistration{
        Feature: MyFeature,
        Title:   "My new scoring feature",
        Tags:    []string{"MARKETING", "AFFILIATE", "UNTRUSTWORTHY"},
    })
}
```

----
### Feature Registry
#### Retrieve Features

```golang
func GetFeatures(includeTags ...string) FeatureSet {
    filteredRegistrations := lo.Filter[FeatureRegistration]
	(registrations, func(it FeatureRegistration, _ int) bool {
        return len(includeTags) == 0 
		    || lo.IsNotEmpty(len(lo.Intersect[string](includeTags, it.Tags)))
    })
    
    return lo.Map[FeatureRegistration, Feature](filteredRegistrations, 
		func(it FeatureRegistration, _ int) Feature {
        return it.Feature
    })
}
```
----
### Putting it all together
#### Commandline Interface

```go
func main() {
	var targetUrl = flag.String("url", "", "URL of the site to be parsed")
	flag.Parse()
	score := scoring.Score(*targetUrl)
	fmt.Printf("Website=%s Score=%f", *targetUrl, score)
}
```
----
### Putting it all together
#### Feature Runner and Combiner

```go
func computeScore(features FeatureSet, document *goquery.Document) (score float64) {
	for _, feature := range features {
		score += feature(document)
	}
	return
}

func Score(url string, featureTags ...string) (score float64) {
	features := GetFeatures(featureTags...)
	document, _ := download.DownloadWebsite(url)
	return computeScore(features, document)
}
```
----

### What we have learned

- Types can also be functions
- Functions can be values
- The ... operator allows readable and flexible argument list and substitutes optional arguments
- The Init function can perform runtime initializion tasks
---

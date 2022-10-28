<!-- .slide: data-background="img/WEBSITE_SCORING/00.jpg" data-background-size="100%" data-background-position="50% 50%" -->
----

## Website Scoring 

Let's build a simple CLI tool and learn something about parsing commandlines, first-class functions and function types in Go along with a little generics. 

----
### Objectives

Compute a credibility score for a website, taking into account

- Links. Affiliate Links might mean the website is bad
- Word Count. Much text might mean that the site has real content
----

### Strategy

- Download the website to a queryable document
- Build a set of features the document is evaluated against
- Build a tiny registry to manage those features
- Let the features compute an individual score
- Combine scores into one total score


----
### Tooling

We won't reinvent the wheel, so we will use:

- [goquery](https://github.com/PuerkitoBio/goquery) for parsing HTML documents
- [lo](https://github.com/samber/lo") for Java Streams/C# LINQ like expressions

----
### Types
#### Features

A feature computes a score from a document.
```golang
type Feature func(document *goquery.Document) float64
```
<small>pkg/scoring/scoring.go</small>

A set of features:
```golang
type FeatureSet []Feature
```
<small>pkg/scoring/scoring.go</small>
----
### Types
#### Feature Registry

A feature registry contains a set of registrations
```golang
type FeatureRegistry struct {
    registrations []FeatureRegistration
}
```
<small>pkg/scoring/scoringFeatureRegistry.go</small>

A registration consists of:
```golang
type FeatureRegistration struct {
    Feature
    Title string
    Tags  []string
}
```
<small>pkg/scoring/scoringFeatureRegistry.go</small>

----
### Feature Registry
#### Registering new Features
Lets start with describing our expectations in a test

```golang
func TestRegisterScoringFeature(t *testing.T) {
    fakeFeature := func(_ *goquery.Document) float64 {
        return 1
    }   

    fakeRegistration := FeatureRegistration{
        Feature: fakeFeature,
        Title:   "FAKE_REGISTRATION",
        Tags:    []string{"FAKE"},
    }
    
    t.Run("should return the previously registered feature when only one registered", func(t *testing.T) {
        registry := FeatureRegistry{}
        registry.Register(fakeRegistration)
        features := registry.GetFeatures()
        assert.Len(t, features, 1)
    })
```
<small>pkg/scoring/scoringFeatureRegistry_test.go</small>

----
### Feature Registry
#### Registering new Features
```
func (f *FeatureRegistry) Register(registrations ...FeatureRegistration) {
	for _, registration := range registrations {
		f.registrations = append(f.registrations, registration)
	}
}

func (f *FeatureRegistry) GetFeatures() FeatureSet {
	return f.registrations
}
```
<small>pkg/scoring/scoringFeatureRegistry.go</small>


----
### Feature Registry
#### Filter support

```
func TestRegisterScoringFeature(t *testing.T) {
	// ...

	fakeRegistration := FeatureRegistration{
		Feature: fakeFeature,
		Title:   "FAKE_REGISTRATION",
		Tags:    []string{"FAKE"},
	}

	fakeRegistration2 := FeatureRegistration{
		Feature: fakeFeature,
		Title:   "FAKE_REGISTRATION2",
		Tags:    []string{"FAKE2"},
	}
	
	t.Run("should return all registered features when multiple registered and no filter provided", func(t *testing.T) {
		registry := FeatureRegistry{}
		registry.Register(fakeRegistration, fakeRegistration2)
		features := registry.GetFeatures()
		assert.Len(t, features, 2)
	})

	t.Run("should return features matching filter", func(t *testing.T) {
		registry := FeatureRegistry{}
		registry.Register(fakeRegistration, fakeRegistration2)
		features := registry.GetFeatures("FAKE2")
		assert.Len(t, features, 1)
	})
}
```
<small>pkg/scoring/scoringFeatureRegistry_test.go</small>

----
### Feature Registry
#### Filter support
Lets use *lo* to filter our subscriptions.
```
func (f *FeatureRegistry) GetFeatures(includeTags ...string) FeatureSet {
	filteredRegistrations := lo.Filter[FeatureRegistration](f.registrations, 
	    func(it FeatureRegistration, _ int) bool {
		    return len(includeTags) == 0 || lo.Some[string](includeTags, it.Tags)
	})

	return lo.Map[FeatureRegistration, Feature](filteredRegistrations, 
	    func(it FeatureRegistration, _ int) Feature {
		    return it.Feature
	})
}
```
<small>pkg/scoring/scoringFeatureRegistry.go</small>
----
### Features
#### Word Count
The result datatype
```
type WordFrequencyResult struct {
	TotalWords   int            
	CountedWords int            
	WordCounts   map[string]int 
}
```
<small>pkg/scoring/wordCount.go</small>

----
### Features
#### Word Count

```
func WordCount(doc *goquery.Document) WordFrequencyResult {
	result := WordFrequencyResult{
		WordCounts: make(map[string]int),
	}

	text := doc.Find("p").Contents().Text()
	words := strings.Split(text, " ")
	result.TotalWords = len(words)

	for i := 0; i < len(words); i++ {
		wordSanitized := strings.ToUpper(sanitizeString.ReplaceAllString(words[i], ""))
		if len(wordSanitized) > 4 {
			result.WordCounts[wordSanitized]++
		}
	}

	for _, v := range result.WordCounts {
		result.CountedWords += v
	}

	return result
}
```
<small>The actual Word Count in pkg/scoring/wordCount.go</small>
----
### Putting it all together
#### Registering the Word Counter

```
func wordCountRegistration() FeatureRegistration {
	return FeatureRegistration{
		Feature: ScoreWrapper[WordFrequencyResult](WordCount),
		Title:   "Word Count",
		Tags:    []string{"CONTENT"},
	}
}
```
<small>pkg/scoring/wordCount.go</small>

```
func NewDefaultRegistry() (r FeatureRegistry) {
	r.Register(affiliateLinkCountRegistration())
	return
}
```
<small>pkg/scoring/scoringFeatureRegistry.go</small>

----
### Putting it all together
#### The score wrapper

```
// Scoreable constraints to types providing a method to be scored
type Scoreable interface {
	Score() float64
}

// ScoreWrapper returns a function converting its result to a score
func ScoreWrapper[T Scoreable](scoreFunction func(doc *goquery.Document) T) Feature {
	return func(document *goquery.Document) float64 {
		result := scoreFunction(document)
		return result.Score()
	}
}
```

<small>pkg/scoring/scoreWrapper.go</small>
----
### Putting it all together
#### Commandline Interface

```
func main() {
	var targetUrl = flag.String("url", "", "URL of the site to be parsed")
	flag.Parse()
	if len(*targetUrl) == 0 {
		flag.Usage()
	}
	score := scoring.Score(*targetUrl)
	fmt.Printf("Website=%s Score=%f", *targetUrl, score)
}
```
<small>cmd/score/score.go</small>
----
### Putting it all together
#### Feature Runner and Combiner

```
func computeScore(features FeatureSet, document *goquery.Document) (score float64) {
	for _, feature := range features {
		score += feature(document)
	}
	return
}

func Score(url string, featureTags ...string) (score float64) {
	registry := NewDefaultRegistry()
	features := registry.GetFeatures(featureTags...)
	document, _ := download.DownloadWebsite(url)
	return computeScore(features, document)
}
```
<small>pkg/scoring/scoring.go</small>
----
### Run it

```
$ website_score --url https://codecentric.de
Website=https://codecentric.de Score=244.200000
```

----
## Lets add a second feature!
----
### Affiliate Link Count
#### Type definition

```
type LinkCountResult struct {
	TotalLinks           int `json:"TotalLinks"`
	LocalLinks           int `json:"LocalLinks"`
	AffiliateLinks       int `json:"AffiliateLinks"`
	MaskedAffiliateLinks int `json:"MaskedAffiliateLinks"`
	ShortendedUrls       int `json:"ShortendedUrls"`
}

func (l LinkCountResult) Score() float64 {
	return float64(l.TotalLinks + l.LocalLinks - l.AffiliateLinks*2 - l.MaskedAffiliateLinks*4 - l.ShortendedUrls)
}
```
<small>pkg/scoring/affiliateLinkCount.go</small>
----


### Affiliate Link Count
#### The actual link count

```
var affiliateLinksExpression, _ = regexp.Compile("(\\.amazon\\.)")
var maskedAffiliateLinksExpression, _ = regexp.Compile("(amzn\\.to)")
var shortenedUrlExpression, _ = regexp.Compile("(bit\\.ly|tinyurl\\.com)")

func AffiliateLinkCount(doc *goquery.Document) LinkCountResult {
	result := LinkCountResult{}
	localdomain := doc.Url.Host

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		result.TotalLinks++
		link, exists := s.Attr("href")

		if exists {
			if strings.Contains(link, localdomain) {
				result.LocalLinks++
			}

			if affiliateLinksExpression.MatchString(link) {
				result.AffiliateLinks++
			}

			if maskedAffiliateLinksExpression.MatchString(link) {
				result.MaskedAffiliateLinks++
			}

			if shortenedUrlExpression.MatchString(link) {
				result.ShortendedUrls++
			}
		}

	})
	return result
}
```
<small>pkg/scoring/affiliateLinkCount.go</small>
----

### Affiliate Link Count
#### The registration

```
func affiliateLinkCountRegistration() FeatureRegistration {
	return FeatureRegistration{
		Feature: ScoreWrapper[LinkCountResult](AffiliateLinkCount),
		Title:   "Affiliate Link Count",
		Tags:    []string{"MARKETING", "AFFILIATE", "UNTRUSTWORTHY"},
	}
}
```
<small>pkg/scoring/affiliateLinkCount.go</small>

```
func NewDefaultRegistry() (r FeatureRegistry) {
	r.Register(affiliateLinkCountRegistration(), wordCountRegistration())
	return
}
```
<small>pkg/scoring/scoringFeatureRegistry.go</small>
----

### What we have learned

- Types can also be functions
- Functions can be values
- Functions can return functions
- The ... operator allows readable and flexible argument lists and substitutes optional arguments
----
# Bonus

Implement a CLI parameter to specify the feature filter query 
----

### Further readings
* [Programming with go / first class functions](https://livebook.manning.com/book/get-programming-with-go/chapter-14/31)
* [variadic function parameters](https://go.dev/ref/spec#Passing_arguments_to_..._parameters)



---
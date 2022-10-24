<!-- .slide: data-background="img/ERROR_HANDLING/00.jpg" data-background-size="60%" data-background-position="50% 50%" -->
----

## Error Handling

Let's talk about errors and error handling in Go.

----

### Golang Error Key Principles

- Errors are values
- Errors don't directly affect control flow
- Developer is encouraged to explicitly handle errors

----

### Errors are Values

Function returning an error:

```golang
func fail() error {
    return fmt.Errorf("This did not work out.")
}
```
----
### Errors don't directly affect control flow
Check if an error occurred:

```golang
err := fail()
if err != nil {
	panic(err)
}
```
> An error has no effect unless actively handled 

----

#### Distinguishing between multiple errors

Use-cases for distinguishable errors:
- Systematic logging
- Sophisticated remedy strategies

----

##### Distinguishing via Sentinel Values

```golang
var noStageNameProvided = errors.New("No Stage name provided")
var invalidStageNameProvidedError = errors.New("Invalid stage name provided")

func resolveService() (string, error) {
	stageName := os.Getenv(stageEnvironmentKey)

	switch stageName {
	case "dev":
		return "https://dev.fake", nil
	case "staging":
		return "https://stage.my.cloud", nil
	case "":
		return invalidService, noStageNameProvided
	default:
		return invalidService, invalidStageNameProvidedError
	}
}
````
----
##### Distinguishing via Sentinel Values

```golang
serviceUrl, err := resolveService()
if err != nil {
    switch {
        case errors.Is(err, noStageNameProvided):
            fmt.Println("No stage name provided.")
            serviceUrl = remedyForMissingStageName()

        case errors.Is(err, invalidStageNameProvidedError):
            panic(err)
    }
}
````
----

#### Wrapping errors
Use-Cases:
- Adding context to errors, such as custom exceptions
- Supporting an inner error construct

----

#### Wrapping errors
The previous function with wrapped errors:

```golang
func resolveService() (serviceUrl string, err error) {
	stageName := os.Getenv(stageEnvironmentKey)

	serviceUrl = invalidService
	switch stageName {
	case "dev":
		serviceUrl = "https://dev.fake"
	case "staging":
		serviceUrl = "https://stage.my.cloud"
	case "":
		err = noStageNameProvided
	default:
		err = fmt.Errorf("%w . %s is not a known stageName", 
			invalidStageNameProvidedError, stageName)
	}
	return
}
````
----

#### Wrapping errors
Handling an invalid input

```golang
serviceUrl, err := resolveService()
	if err != nil {
		switch {
		case errors.Is(err, noStageNameProvided):
			fmt.Println("No stage name provided.")
			serviceUrl = remedyForMissingStageName()

		case errors.Is(err, invalidStageNameProvidedError):
			panic(err)
		}
	}
````

Output:
```
panic: Invalid stage name provided . Gophers! is not a known stageName
```
----

### What we have learned
* Errors are values in Go
* Errors have to be handled explicitly
* How errors are thrown and detected
* How to add context to errors

----

### Further readings
* Working with Errors in Go 1.13
  * [go.dev/blog/go1.13-errors](https://go.dev/blog/go1.13-errors)
* Errors are Values
  * [go.dev/blog/errors-are-values](https://go.dev/blog/errors-are-values)

---

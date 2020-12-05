Time span
======================================================================

The Go package timespan provides the types represent the span of time.

### example

```go
// the span represents 2 hours from 23:00Z to 01:00Z(next day)
span := timespan.MustParse("23:00Z/2h")
epoch := time.Unix(0, 0)
fmt.Println(span.Contains(epoch.Add(time.Hour - 1)))    //  true: 00:59:59.999...
fmt.Println(span.Contains(epoch.Add(time.Hour)))        // false: 01:00:00.000...
fmt.Println(span.Contains(epoch.Add(23*time.Hour - 1))) // false: 22:59:59.999...
fmt.Println(span.Contains(epoch.Add(23 * time.Hour)))   //  true: 23:00:00.000...
// Output:
// true
// false
// false
// true
```

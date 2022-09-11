// Go offers extensive support for times and durations; here are some examples.
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// We’ll start by getting the current time.
	now := time.Now()
	p(now) // -> 2022-09-11 10:54:31.7455537 +0800 +08 m=+0.001607801

	// You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then) // -> 2009-11-17 20:34:58.651387237 +0000 UTC

	// You can extract the various components of the time value as expected.
	p(then.Year())       // -> 2009
	p(then.Month())      // -> November
	p(then.Day())        // -> 17
	p(then.Hour())       /// -> 20
	p(then.Minute())     // -> 34
	p(then.Second())     // -> 58
	p(then.Nanosecond()) // -> 651387237
	p(then.Location())   //-> UTC

	// The Monday-Sunday Weekday is also available.
	p(then.Weekday()) // -> Tuesday

	// These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.
	p(then.Before(now)) // -> true
	p(then.After(now))  // -> false
	p(then.Equal(now))  // -> false

	// The Sub methods returns a Duration representing the interval between two times.
	diff := now.Sub(then)
	p(diff) // -> 112326h19m33.094166463s

	// We can compute the length of the duration in various units.
	p(diff.Hours())       // -> 112326.32585949068
	p(diff.Minutes())     // -> 6.739579551569441e+06
	p(diff.Seconds())     // -> 4.0437477309416646e+08
	p(diff.Nanoseconds()) // -> 404374773094166463

	// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.
	p(then.Add(diff))  // -> 2022-09-11 02:54:31.7455537 +0000 UTC
	p(then.Add(-diff)) // -> 1997-01-24 14:15:25.557220774 +0000 UTC
}

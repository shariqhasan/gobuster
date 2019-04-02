package libgobuster

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type intSet struct {
	Set map[int]bool
}

var mu sync.Mutex

type stringSet struct {
	Set map[string]bool
}

func newStringSet() stringSet {
	return stringSet{Set: map[string]bool{}}
}

// Add an element to a set
func (set *stringSet) Add(s string) bool {
	_, found := set.Set[s]
	set.Set[s] = true
	return !found
}

// Add a list of elements to a set
func (set *stringSet) AddRange(ss []string) {
	for _, s := range ss {
		set.Set[s] = true
	}
}

// Test if an element is in a set
func (set *stringSet) Contains(s string) bool {
	_, found := set.Set[s]
	return found
}

// Check if any of the elements exist
func (set *stringSet) ContainsAny(ss []string) bool {
	for _, s := range ss {
		if set.Set[s] {
			return true
		}
	}
	return false
}

// Stringify the set
func (set *stringSet) Stringify() string {
	values := []string{}
	for s := range set.Set {
		values = append(values, s)
	}
	return strings.Join(values, ",")
}

func newIntSet() intSet {
	return intSet{Set: map[int]bool{}}
}

// Add an element to a set
func (set *intSet) Add(i int) bool {
	mu.Lock()
	defer mu.Unlock()
	_, found := set.Set[i]
	set.Set[i] = true
	return !found
}

// Test if an element is in a set
func (set *intSet) Contains(i int) bool {
	_, found := set.Set[i]
	return found
}

// Stringify the set
func (set *intSet) Stringify() string {
	values := []int{}
	for s := range set.Set {
		values = append(values, s)
	}
	sort.Ints(values)

	delim := ","
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(values)), delim), "[]")
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 1
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func FixUrl(url *string) error {
	if !strings.HasSuffix(*url, "/") {
		*url = fmt.Sprintf("%s/", *url)
	}

	if !strings.HasPrefix(*url, "http") {
		// check to see if a port was specified
		re := regexp.MustCompile(`^[^/]+:(\d+)`)
		match := re.FindStringSubmatch(*url)

		if len(match) < 2 {
			// no port, default to http on 80
			*url = fmt.Sprintf("http://%s", *url)
		} else {
			port, err := strconv.Atoi(match[1])
			if err != nil || (port != 80 && port != 443) {
				return fmt.Errorf("url scheme not specified")
			} else if port == 80 {
				*url = fmt.Sprintf("http://%s", *url)
			} else {
				*url = fmt.Sprintf("https://%s", *url)
			}
		}
	}

	return nil
}
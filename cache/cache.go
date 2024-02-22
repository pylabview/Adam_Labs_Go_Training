package cache

// User is a user in the system.
type User struct {
	Login string
	// TODO: More fields
}

// SliceCache is a user cash using a slice.
type SliceCache []User

// Get gets a user from the cache, return false if not found.
func (c SliceCache) Get(login string) (User, bool) {
	for _, u := range c {
		if u.Login == login {
			return u, true
		}
	}

	return User{}, false
}

// MapCache is a user cash using a map.
type MapCache map[string]User // login -> User

// Get gets a user from the cache, return false if not found.
func (c MapCache) Get(login string) (User, bool) {
	u, ok := c[login]
	if ok {
		return u, true
	}

	return User{}, false
}

// The two hardest problems in computer science are naming things and cache invalidation
// and off-by one errors.

package day21

type Set struct {
	items map[string]struct{}
}

func NewSet(items ...string) *Set {
	set := &Set{}
	set.items = make(map[string]struct{})
	set.Add(items...)
	return set
}

func (s *Set) Add(items ...string) {
	for _, item := range items {
		s.items[item] = struct{}{}
	}
}

func (s *Set) Contains(item string) bool {
	_, ok := s.items[item]
	return ok
}

func (s *Set) Intersect(as *Set) []string {
	var result []string
	for item := range s.items {
		if as.Contains(item) {
			result = append(result, item)
		}
	}
	return result
}

func (s *Set) Items() []string {
	i := 0
	items := make([]string, len(s.items))
	for item := range s.items {
		items[i] = item
		i++
	}
	return items
}

func (s *Set) Size() int {
	return len(s.items)
}

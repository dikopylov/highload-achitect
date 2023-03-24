package users

type SearchUserSpec struct {
	FirstName string
	LastName  string
}

func (s *SearchUserSpec) IsValid() bool {
	return s != nil && s.FirstName != "" && s.LastName != ""
}

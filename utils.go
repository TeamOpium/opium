package opium

func IsInList(users []*User, it *User) bool{
	for _, u := range users{
		if u == it{
			return true
		}
	}
	return false
}

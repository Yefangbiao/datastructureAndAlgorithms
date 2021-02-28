package main

type userIterator struct {
	users []*user
	index int
}

func (u *userIterator) hasNext() bool {
	return u.index < len(u.users)
}

func (u *userIterator) getNext() *user {
	res := u.users[u.index]
	u.index++
	return res
}

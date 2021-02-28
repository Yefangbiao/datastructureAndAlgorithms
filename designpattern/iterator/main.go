package main

import "fmt"

func main() {
	user1 := &user{name: "zhangsan", age: 23}
	user2 := &user{name: "lisi", age: 43}
	user3 := &user{name: "wangwu", age: 13}
	user4 := &user{name: "shuailiu", age: 76}
	collection := userCollection{users: []*user{user1, user2, user3, user4}}
	iter := collection.createIterator()
	for iter.hasNext() {
		u := iter.getNext()
		fmt.Printf("user: %v, age:%v\n", u.name, u.age)
	}
}

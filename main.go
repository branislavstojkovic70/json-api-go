package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, firstname, lastname, pw string) *Account {
	acc, err := NewAccount(firstname, lastname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "branislav", "stojkovic", "bane331431")
}

func main() {
	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()
	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}
	if err := store.Init(); err != nil {
		log.Fatal(err)
	}
	if *seed {
		fmt.Println("seeding the db")
		seedAccounts(store)
	}
	fmt.Printf("%+v\n", store)
	server := NewAPIServer(":3000", store)
	server.Run()
	fmt.Println("des")
}

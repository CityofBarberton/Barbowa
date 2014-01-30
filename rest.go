package main

import (
    "net/url"
    "github.com/dougblack/sleepy"
)

type Contact struct { }

func (contact Contact) Get(values url.Values) (int, interface{}) {
    contacts := []string{"contact1", "contact2"}
    data := map[string][]string{"contacts": contacts}
    return 200, data
}

func main() {
    contact := new(Contact)

    var api = sleepy.NewAPI()
    api.AddResource(contact, "/contacts")
    api.Start(8080)
}

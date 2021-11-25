package main

import (
	"fmt"
	"strings"
)

func main() {
	usernames := map[string]string{"ikehakinyemi": "IkehAkinyemi", "fortune": "priscilla"}

	for {
		var name string

		fmt.Println("Find your username online")
		_, err := fmt.Scanln(&name)

		if err != nil {
			panic(err)
		}

		name = strings.TrimSpace(name)

		if u, ok := usernames[name]; ok {
			fmt.Printf("%q is the username of the name, %q\n", u, name)
			continue
		} else {
			fmt.Printf("Can't find the username for %q\n", name)
		}

		fmt.Println("What's the username you're looking for?", name)

		var username string
		_, err = fmt.Scanln(&username)

		if err != nil {
			usernames[name] = username
		}

		fmt.Printf("Your username has been updatedt to %q", username)
		break
	}
}
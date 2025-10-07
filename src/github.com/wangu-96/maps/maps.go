package main

import (
	"fmt"
)

func main() {

	// email := make(map[string]string)

	// email["wangugalati@gmail.com"] = "Wangu-Email"

	// email["percynguni@gmail.com"] = "Percy-Email"

	// delete(email, "percynguni@gmail.com")

	email := map[string]string{"wangugalati@gmail.com": "Wangu-Email", "percynguni@gmail.com": "Percy-Email"}

	delete(email, "percynguni@gmail.com")

	fmt.Println(email)

}

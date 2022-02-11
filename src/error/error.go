package error

import "fmt"

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

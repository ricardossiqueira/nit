package output

import "fmt"

func PrintCommit(msg string) error {
	fmt.Println(msg)
	return nil
}

func PrintPR(title string, body string) error {
	fmt.Println(title)
	fmt.Println()
	fmt.Println(body)
	return nil
}

func PrintTitle(title string) error {
	fmt.Println(title)
	return nil
}

func PrintBody(body string) error {
	fmt.Println(body)
	return nil
}

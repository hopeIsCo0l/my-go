package main

import (
	"fmt"
	"net/http"
	_ "net/http"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Go on Localhost")

	for i := 10; i > 0; i-- {
		fmt.Fprintf(w, "eye is: ", i)
	}

}

func main() {

	http.HandleFunc("/", helloHandler)

	fmt.Println("server is running on port 8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}

	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		fmt.Println("i =", 100/i)
	}

}

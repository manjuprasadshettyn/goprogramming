// string_to_html project main.go
package main

import "fmt"

func main() {
	name := "Manjuprasad"
	fmt.Println(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Go Rocks!</title>
	</head>
	<body>
	<h1>` +
		name +
		`</h1>
	</body>
	</html>
	`)
}

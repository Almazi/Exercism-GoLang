package hello

const testVersion = 2

// Here I am writing a greeting message to the user who will run this code.
// User will get a message to input name, then it will print "Hello, $NAME",
// if there is no input from user it will simply print "Hello, world!"
func HelloWorld(name string) string {
	if name != ""{
		return "Hello, " + name + "!"
	} else {
		return "Hello, World!"
	}

}

package main

type OS interface {
	ChangeDirectory(args []string)
	Pwd() string
	Echo(args []string)
	Kill(args []string)
	ps(args []string)
}

func main() {

}

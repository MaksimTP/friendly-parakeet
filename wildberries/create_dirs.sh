
for (( a = $1; a < $2 + 1; a++ ))
do
    mkdir $a
    cd $a 
    touch main.go
    echo "package main

import \"fmt\"

func main() {
	fmt.Println(\"Hello World!\")
}
" > main.go
    cd ..
done
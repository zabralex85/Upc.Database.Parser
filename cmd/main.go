package main
import "../parser"
import "os"

func main(){
	barcode:= "025192960024"
	argsWithProg := os.Args[1:]

	if(len(argsWithProg) > 0){
		barcode = argsWithProg[0]
	}

	data := parser.GetDataJSON(barcode)
	println(string(data))
}
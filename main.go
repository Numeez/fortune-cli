package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
var files []string

func visit(path string , f os.FileInfo,err error) error{
	if err!=nil{
		log.Fatal(err)
	}
	if strings.Contains(path,"/off/") {
		return nil
	}
	if filepath.Ext(path)==".dat"{
		return nil
	}
	if f.IsDir(){
		return nil
	}
	files = append(files,path)
	return nil
}
func randomInt(min,max int ) int {
	return min+rand.Intn(max-min)
}

func openFile(filePath string) error {
	file,err:= os.Open(filePath)
	if err!=nil{
		return err
	}
	defer file.Close()
	b,err:=io.ReadAll(file)
	if err!=nil{
		return err
	}
	quotes:=string(b)
	quoteSlice:=strings.Split(quotes,"%")
	randomNum:=randomInt(1,len(quoteSlice))
	fmt.Println(quoteSlice[randomNum])
	return nil
}

func main (){
	fortuneCommand:=exec.Command("fortune","-f")
	pipe,err:=fortuneCommand.StderrPipe()
	if err !=nil{
		panic(err)
	}
	fortuneCommand.Start()
	outputStream:=bufio.NewScanner(pipe)
	outputStream.Scan()
	line:=outputStream.Text()
	root:=line[strings.Index(line,"/"):]
	err=filepath.Walk(root,visit)
	if err!=nil{
		panic(err)
	}
	randomNumber:= randomInt(1,len(files))
	randomFile:= files[randomNumber]
	err=openFile(randomFile)
	if err!=nil{
		panic(err)
	}
}
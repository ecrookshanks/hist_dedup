package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/user"
	"strings"
)

const bash_file = "/.bash_history"

func readFileLineByLine(file string) error {
	f, err := os.Open(file)
	if err != nil{
		return err
	}
	fWrite, err := os.Create("dup_test.txt")
	if err != nil{
		return err
	}

	defer f.Close()
	defer fWrite.Close()

	counts := make(map[string]int)
	r := bufio.NewReader(f)
	
	num_lines := 0
	for{
		line, err := r.ReadString('\n')
		if err == io.EOF{
			break
		} else if err != nil{
			fmt.Printf("Error reading file: %s\n", err)
		}
		counts[line]++
		num_lines++
	}
	fmt.Printf("Total number of lines: %d\n", num_lines)
	fmt.Printf("Total unique lines: %d\n", len(counts) )

	fmt.Printf("\n****\nSummary of repeated items\n****\n")
	for line, n := range counts {
		if n > 1{
			fmt.Printf("\"%s\" appears %d times.\n", strings.Trim(line, "\n"), n)
		}
		// write out each line - will be unique lines
		fWrite.WriteString(line)
	}
	return nil
}

func constructCompleteFileName() (string, error){
	usr, err := user.Current();
	if err != nil{
		return "", err
	}
	dir := usr.HomeDir
	file := dir + bash_file

	return file, nil
}

func main(){
	args := os.Args
	if len(args) > 1 {
		fmt.Printf("process options here \n")
	}

	file, err := constructCompleteFileName()
	if err != nil{
		fmt.Println(err)
		return
	}
	
	err = readFileLineByLine(file)
	if (err != nil){
		fmt.Println(err)
	}

}
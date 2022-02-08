package main

import (
	"fmt"
	"log"
	"os/exec"
)

func SendCode() {
	python := "/usr/bin/python3"
	script := "main.py"
	args_phone := "-P 79963231252"
	args_hash := "-H b2f40d97db0a4b4cdf6a4c33d180e927"
	args_id := "-I 8162123"
	out, err := exec.Command(python, script, args_phone, args_hash, args_id).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func Verify() {
	python := "/usr/bin/python3"
	script := "verify.py"
	args_phone := "-P 79963231252"
	args_hash := "-H b2f40d97db0a4b4cdf6a4c33d180e927"
	args_id := "-I 8162123"
	args_code := "-C 17447"
	args_codeHash := "-G 0255ac92bb56131a7e"
	out, err := exec.Command(python, script, args_phone, args_hash, args_id, args_code, args_codeHash).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func main() {
	//SendCode()
	Verify()
}

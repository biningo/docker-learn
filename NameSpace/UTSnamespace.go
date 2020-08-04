package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

/**
*@Author bingo
*@Date 2020/8/4
*
**/

func main() {

	cmd := exec.Command("sh") //开启一个shell环境
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS, //在不同的UTFnamespace
	}

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err:=cmd.Run();err!=nil{
		log.Println(err)
	}

}

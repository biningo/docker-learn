package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

/**
*@Author icepan
*@Date 2020/8/4
*
**/

func main() {





	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}

	log.Println("PID：", cmd.Process.Pid)

	cpath := path.Join("/sys/fs/cgroup/memory", "icepan")
	//创建一个cgroup
	os.Mkdir(cpath, 0755)
	//加入cgroup
	ioutil.WriteFile(path.Join(cpath, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
	//限制内存资源
	ioutil.WriteFile(path.Join(cpath, "memory.limit_in_bytes"), []byte("100m"), 0644)
	cmd.Process.Wait()
}

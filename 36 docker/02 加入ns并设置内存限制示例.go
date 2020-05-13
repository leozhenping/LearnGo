package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"

func main() {

	// 当调用cmd.Process.Wait()时，才回调执行此处的if分支代码
	if os.Args[0] == "/proc/self/exe" { // 此处为容器内部执行的代码
		// 容器进程
		fmt.Printf("current pid %d\n", syscall.Getpid())
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}

		os.Exit(1)
	}

	// 1. 先执行此处的cmd
	cmd := exec.Command("/proc/self/exe") // /proc/self/exe, 表示当前正在运行的进程, 或者可以正解为/proc/PROCESS_NUM/exe.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	} else {
		// 得到fork出来进程映射在外部命名空间的Pid
		fmt.Printf("%v\n", cmd.Process.Pid)
		// 在系统默认创建挂载了memory subsystem的hierarchy上创建group
		os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)
		//将容器进程加入到这个cgroup中
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "memory.limit_in_bytes"), []byte("100m"), 0644)
	}
	cmd.Process.Wait()
}

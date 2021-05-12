package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

// 顶层处理
func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	err := runJob("1")
	if err != nil {
		// 打印堆栈
		log.Printf("%+v", err)
		// 打印用户需要的信息
		fmt.Println(err)
	}
}

// 底层处理
type LowLevelErr struct {
	error
}

func isGloballyExec(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, errors.Wrapf(LowLevelErr{err}, "isGloballyExec: ")
	}
	return true, nil
}

// 中层处理
func runJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := isGloballyExec(jobBinPath)
	if err != nil {
		return errors.WithMessage(err, "runJob: ")
	} else if isExecutable == false {
		return errors.WithMessagef(err, "runJob: can not run job %q: requisite binaries are not executable", id)
	}
	return nil
}

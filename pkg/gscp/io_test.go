package gscp

import (
	"fmt"
	"log"
	"os/exec"
	"testing"
)

func execute(args ...string) string {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.Output()

	fmt.Println("Execute: ", args[0])
	if err != nil {
		log.Fatalln("failed to execute: ", args[0])
		return ""
	} else {
		fmt.Println(string(out))
		return string(out)
	}
}

func setup_io_test() {
	execute("rm", "-rf", "/tmp/test")
	execute("mkdir", "-p", "/tmp/test")
}

func teardown_io_test() {
	execute("rm", "-rf", "/tmp/test")
}

func TestReadFile(t *testing.T) {
	setup_io_test()

	tf := TestFrame{Test: t}

	tf.Run("read empty", func(tf TestFrame) {
		execute("touch", "/tmp/test/a.txt")
		f, err := ReadFile("/tmp/test/a.txt")
		tf.AssertNil(err)
		tf.AssertEquals(string(f), "")
	})

	tf.Run("read filled", func(tf TestFrame) {
		execute("/bin/bash", "-c", "echo hello > /tmp/test/a.txt")
		f, err := ReadFile("/tmp/test/a.txt")
		tf.AssertNil(err)
		tf.AssertEquals(string(f), "hello\n")
	})

	teardown_io_test()
}

func TestWriteFile(t *testing.T) {
	setup_io_test()

	tf := TestFrame{Test: t}

	tf.Run("write empty", func(tf TestFrame) {
		WriteFile("/tmp/test/a.txt", []byte(""))
		v := execute("cat", "/tmp/test/a.txt")
		tf.AssertEquals(v, "")
	})

	tf.Run("write filled", func(tf TestFrame) {
		WriteFile("/tmp/test/a.txt", []byte("hello"))
		v := execute("cat", "/tmp/test/a.txt")
		tf.AssertEquals(v, "hello")
	})

	teardown_io_test()
}

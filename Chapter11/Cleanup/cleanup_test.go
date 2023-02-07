package cleanup

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func myCleanUp() func() {
	return func() {
		fmt.Println("Cleaning up!")
	}
}

func TestFoo(t *testing.T) {
	t1 := path.Join(os.TempDir(), "test01")
	t2 := path.Join(os.TempDir(), "test02")
	err := os.Mkdir(t1, 0755)
	if err != nil {
		t.Error("os.Mkdir() failed:", err)
		return
	}
	// 테스트 파일에 또는 프로파일링을 할 때 TempDir에 파일을 저장해야할 때 다시 Cleanup을 위하여 t.Cleanup(func(직접 삭제 구현))
	defer t.Cleanup(func() {
		err = os.Remove(t1)
		if err != nil {
			t.Error("os.Mkdir() failed:", err)
		}
	})

	err = os.Mkdir(t2, 0755)
	if err != nil {
		t.Error("os.Mkdir() failed:", err)
		return
	}
}

func TestBar(t *testing.T) {
	t1 := t.TempDir()
	fmt.Println(t1)
	t.Cleanup(myCleanUp())
}

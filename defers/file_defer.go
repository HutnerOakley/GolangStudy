package defers

import "os"

func ReadFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	// ✅ 注册defer：函数无论从哪个出口退出，都会执行Close
	defer f.Close()

	buf := make([]byte, 1024)
	n, err := f.Read(buf)
	if err != nil {
		return "", err
	}

	return string(buf[:n]), nil
}

package writer

import (
	"fmt"
	"github.com/projectdiscovery/gologger/levels"
	"os"
	"sync"
)

type LogFile struct {
	mutex *sync.Mutex
	file  string
}

func NewLogFile(file string) *LogFile {
	logFile := &LogFile{
		file:  file,
		mutex: &sync.Mutex{},
	}
	return logFile
}

func (w *LogFile) Write(data []byte, level levels.Level) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	switch level {
	case levels.LevelSilent:
		os.Stdout.Write(data)
		os.Stdout.Write([]byte("\n"))
	default:
		os.Stderr.Write(data)
		os.Stderr.Write([]byte("\n"))
	}
	fl, err := os.OpenFile(w.file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, _ = fl.Write(data)
	_, _ = fl.Write([]byte("\n"))
	fl.Close()
}

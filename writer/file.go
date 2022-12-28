package writer

import (
	"github.com/projectdiscovery/gologger/levels"
	"io"
	"os"
	"sync"
)

type LogFile struct {
	mutex  *sync.Mutex
	writer io.Writer
}

func NewLogFile(writer io.Writer) *LogFile {
	logFile := &LogFile{
		mutex:  &sync.Mutex{},
		writer: writer,
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
	_, _ = w.writer.Write(data)
	_, _ = w.writer.Write([]byte("\n"))
}

package impl

import (
	"github.com/pkg/errors"
	"os"
	"path"
	"time"
)

type FileNameGenerator interface {
	Gen(app string) string
}

type fileLogger struct {
	dir             string
	currentFileName string
	currentFile     *os.File
	buffer          *LogBuffer
	generator       FileNameGenerator
}

func NewDefaultFileLogger(dir string, cacheSize, flushDurationSecond int) Logger {
	return NewFileLogger(dir, cacheSize, flushDurationSecond, NewHourFileNameGenerator(dir))
}

func NewFileLogger(dir string, cacheSize int, flushDuration int, generator FileNameGenerator) Logger {
	logger := fileLogger{
		dir:       dir,
		generator: generator,
	}
	logger.buffer = NewLogBuffer(cacheSize, time.Duration(flushDuration)*time.Second, logger.onFlush)
	return &logger
}

func (t *fileLogger) Flush() error {
	return t.buffer.Flush()
}

func (t *fileLogger) Info(app string, msg string) error {
	if len(msg) > 0 && msg[len(msg)-1] != '\n' {
		msg += "\n"
	}
	var err error
	genFileName := t.generateFileName(app)
	if t.currentFile == nil {
		t.currentFile, err = os.OpenFile(genFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		t.currentFileName = genFileName
		if err != nil {
			return errors.Wrap(err, "create log file failed")
		}
	}
	if genFileName != t.currentFileName && t.currentFile != nil {
		if err := t.buffer.Flush(); err != nil {
			return err
		}
		_ = t.currentFile.Close()
		t.currentFile, err = os.OpenFile(genFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		t.currentFileName = genFileName
		if err != nil {
			return errors.Wrap(err, "create log file failed")
		}
	}
	return t.buffer.Append(msg)
}

func (t *fileLogger) onFlush(data []byte) error {
	if t.currentFile == nil {
		return nil
	}
	if _, err := t.currentFile.Write(data); err != nil {
		return errors.Wrap(err, "flush data failed")
	}
	return nil
}

func (t *fileLogger) generateFileName(app string) string {
	return t.generator.Gen(app)
}

type HourFileNameGenerator struct {
	dir string
}

func NewHourFileNameGenerator(dir string) FileNameGenerator {
	ret := &HourFileNameGenerator{
		dir: dir,
	}
	return ret
}

func (t *HourFileNameGenerator) Gen(app string) string {
	return path.Join(t.dir, app+"_"+time.Now().Format("20060102_15")+".log")
}

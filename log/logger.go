package log

import (
	"fmt"
	"log"
	"runtime"
)

func Verbosef(msg string, args ...interface{}) {
	// log.Printf(msg, args...)
}
func Verbose(v ...interface{}) {
	// log.Println(v...)
}

func Debugf(msg string, args ...interface{}) {
	log.Printf("[Debug] "+msg, args...)
}

func Debug(v ...interface{}) {
	log.Println(v...)
}

func Infof(msg string, args ...interface{}) {
	log.Printf("[Info]: "+detectFunc()+msg, args...)
}

func Info(v ...interface{}) {
	args := []interface{}{"[Info]: " + detectFunc()}
	args = append(args, v...)
	log.Println(args...)
}

func Errorf(msg string, args ...interface{}) {
	log.Printf("[Error]: "+detectFunc()+msg, args...)
}

func Error(v ...interface{}) {
	args := []interface{}{"[Error]: " + detectFunc()}
	args = append(args, v...)
	log.Println(args)
}

func Fatalf(msg string, args ...interface{}) {
	log.Fatalf("[Fatal]: "+detectFunc()+msg, args...)
}

func Fatal(v ...interface{}) {
	args := []interface{}{"[Fatal]: " + detectFunc()}
	args = append(args, v...)
	log.Fatal(args...)
}

func detectFunc() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf(" %s:%d ", file, line)
}

package log

import (
	"log"
	"runtime"
	"fmt"
)

func  Verbosef(msg string, args ...interface{}) {
	// log.Printf(msg, args...)
}
func  Verbose(v ...interface{}) {
	// log.Println(v...)
}

func  Debugf(msg string, args ...interface{}) {
	log.Printf("[Debug] "+msg, args...)
}

func  Debug(v ...interface{}) {
	log.Println(v...)
}


func  Infof(msg string, args ...interface{}) {
	log.Printf(detectFunc()+" [Info]: "+msg, args...)
}

func  Info(v ...interface{}) {
	log.Println(detectFunc()+" [Info]: ", v)
}

func  Errorf(msg string, args ...interface{}) {
	log.Printf(detectFunc()+" [Error]: "+msg, args...)
}

func  Error(v ...interface{}) {
	log.Println(detectFunc()+" [Error]: ", v)
}

func  Fatalf(msg string, args ...interface{}) {
	log.Fatalf(detectFunc()+" [Fatal]: "+msg, args...)
}

func  Fatal(v ...interface{}) {
	log.Fatal(detectFunc()+" [Fatal]: ", v)
}

func detectFunc() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf("%s at %d", file, line)
}
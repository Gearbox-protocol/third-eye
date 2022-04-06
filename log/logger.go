package log

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"runtime"
	"testing"
)

var testLogModule *testing.T

func SetTestLogging(t *testing.T) {
	testLogModule = t
}
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

func Warnf(msg string, args ...interface{}) {
	msgFormat := "[Warn] " + DetectFunc() + msg
	amqpSendf(msgFormat, args)
	if testLogModule == nil {
		log.Printf(msgFormat, args...)
	} else {
		testLogModule.Logf(msgFormat, args...)
	}
}

func Warn(v ...interface{}) {
	args := []interface{}{"[Warn]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(args)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Infof(msg string, args ...interface{}) {
	msg = "[Info]: " + DetectFunc() + msg
	if testLogModule == nil {
		log.Printf(msg, args...)
	} else {
		testLogModule.Logf(msg, args...)
	}

}

func Info(v ...interface{}) {
	args := []interface{}{"[Info]: " + DetectFunc()}
	args = append(args, v...)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}
func InfoStack3(v ...interface{}) {
	args := []interface{}{"[Info]: " + DetectFuncAtStack3()}
	args = append(args, v...)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Errorf(msg string, args ...interface{}) {
	msgFormat := "[Error]: " + DetectFunc() + msg
	amqpSendf(msgFormat, args)
	if testLogModule == nil {
		log.Printf(msgFormat, args...)
	} else {
		testLogModule.Logf(msgFormat, args...)
	}
}

func Error(v ...interface{}) {
	args := []interface{}{"[Error]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(args)
	if testLogModule == nil {
		log.Println(args...)
	} else {
		testLogModule.Log(args...)
	}
}

func Msgf(msg string, args ...interface{}) {
	amqpSendf(msg, args)
	msgFormat := DetectFunc() + msg
	log.Printf("[AMPQ]"+msgFormat, args...)
}

func Msg(v ...interface{}) {
	amqpSend(v)
	args := []interface{}{"[AMPQ]" + DetectFunc()}
	args = append(args, v...)
	log.Println(args...)
}

func Fatalf(msg string, args ...interface{}) {
	msgFormat := "[Fatal]: " + DetectFunc() + msg
	amqpSendf(msgFormat, args)
	if testLogModule == nil {
		log.Fatalf(msgFormat, args...)
	} else {
		testLogModule.Fatalf(msgFormat, args...)
	}
}

func Fatal(v ...interface{}) {
	args := []interface{}{"[Fatal]: " + DetectFunc()}
	args = append(args, v...)
	amqpSend(args)
	if testLogModule == nil {
		log.Fatal(args...)
	} else {
		testLogModule.Fatal(args...)
	}
}

func DetectFunc() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf(" %s:%d ", file, line)
}
func DetectFuncAtStack3() string {
	_, file, line, _ := runtime.Caller(3)
	return fmt.Sprintf(" %s:%d ", file, line)
}

func CheckFatal(err error) {
	args := []interface{}{"[Fatal]: " + DetectFunc(), err}
	if err != nil {
		amqpSend(args)
		if testLogModule == nil {
			log.Fatal(args...)
		} else {
			testLogModule.Fatal(args...)
		}
	}
}

var ch *amqp.Channel
var netName string

func SetAMQP(_ch *amqp.Channel, name string) {
	ch = _ch
	netName = name
}
func amqpSend(v []interface{}) {
	alert := fmt.Sprint(v...)
	send(alert)
}
func amqpSendf(msg string, args []interface{}) {
	alert := fmt.Sprintf(msg, args...)
	send(alert)
}
func send(message string) {
	if ch == nil {
		return
	}
	err := ch.Publish(
		"TelegramBot", // exchange
		netName,       // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("[%s]Third eye:", netName) + message),
		})
	if err != nil {
		log.Println("Cant sent notification", err)
	}
}

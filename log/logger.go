package log

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"runtime"
)

func Verbosef(msg string, args ...interface{}) {
	log.Printf(msg, args...)
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
	msgFormat := "[Warn] " + detectFunc() + msg
	log.Printf(msgFormat, args...)
	amqpSendf(msgFormat, args)
}

func Warn(v ...interface{}) {
	args := []interface{}{"[Warn]: " + detectFunc()}
	args = append(args, v...)
	amqpSend(args)
	log.Println(args...)
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
	msgFormat := "[Error]: " + detectFunc() + msg
	amqpSendf(msgFormat, args)
	log.Printf(msgFormat, args...)
}

func Error(v ...interface{}) {
	args := []interface{}{"[Error]: " + detectFunc()}
	args = append(args, v...)
	amqpSend(args)
	log.Println(args...)
}

func Msgf(msg string, args ...interface{}) {
	msgFormat := detectFunc() + msg
	amqpSendf(msg, args)
	log.Printf(msgFormat, args...)
}

func Msg(v ...interface{}) {
	amqpSend(v)
	args := []interface{}{detectFunc()}
	args = append(args, v...)
	log.Println(args...)
}

func Fatalf(msg string, args ...interface{}) {
	msgFormat := "[Fatal]: " + detectFunc() + msg
	amqpSendf(msgFormat, args)
	log.Fatalf(msgFormat, args...)
}

func Fatal(v ...interface{}) {
	args := []interface{}{"[Fatal]: " + detectFunc()}
	args = append(args, v...)
	amqpSend(args)
	log.Fatal(args...)
}

func detectFunc() string {
	_, file, line, _ := runtime.Caller(2)
	return fmt.Sprintf(" %s:%d ", file, line)
}

func CheckFatal(err error) {
	if err != nil {
		args := []interface{}{"[Fatal]: " + detectFunc(), err}
		amqpSend(args)
		log.Fatal(args)
	}
}

var ch *amqp.Channel

func SetAMQP(_ch *amqp.Channel) {
	ch = _ch
}
func amqpSend(v []interface{}) {
	alert := fmt.Sprint(v...)
	send(alert, "")
}
func amqpSendf(msg string, args []interface{}) {
	alert := fmt.Sprintf(msg, args...)
	send(alert, "")
}
func send(message string, routingKey string) {
	if ch == nil {
		return
	}
	err := ch.Publish(
		"TelegramBot", // exchange
		routingKey,    // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Support server:" + message),
		})
	if err != nil {
		log.Println("Cant sent notification", err)
	}
}

package logger

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Debugw(msg string, keysValues ...interface{})

	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Infow(msg string, keysValues ...interface{})

	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Warnw(msg string, keysValues ...interface{})

	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Errorw(msg string, keysValues ...interface{})

	Panic(args ...interface{})
	Panicf(template string, args ...interface{})
	Panicw(msg string, keysValues ...interface{})

	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Fatalw(msg string, keysValues ...interface{})
}

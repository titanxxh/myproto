package trait

import "fmt"

type (
	loggerItf interface {
		Log(s string)
	}
	logger struct {
		cellBean
		loggerItf
	}

	FmtLogger struct {
	}
)

func (t *logger) init(l loggerItf) {
	t.loggerItf = l
}

func (l *FmtLogger) Log(s string) {
	fmt.Println(s)
}

func FmtLoggerProducer() *FmtLogger {
	return &FmtLogger{}
}

func FmtLoggerInitMiddleware(producer beanProducer) beanProducer {
	return func() cellBean {
		bean := producer().(loggerItf).(*logger)
		bean.init(FmtLoggerProducer())
		return bean
	}
}

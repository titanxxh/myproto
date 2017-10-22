package trait

import "log"

type (
	messagingItf interface {
		SendMessage(s string)
	}
	messaging struct {
		cellBean
		messagingItf
	}

	ESMessaging struct {
	}
)

func (t *messaging) init(m messagingItf) {
	t.messagingItf = m
}

func (l *ESMessaging) SendMessage(s string) {
	log.Println(s)
}

func ESMessagingProducer() *ESMessaging {
	return &ESMessaging{}
}

func ESMessagingInitMiddleware(producer beanProducer) beanProducer {
	return func() cellBean {
		bean := producer().(messagingItf).(*messaging)
		bean.init(ESMessagingProducer())
		return bean
	}
}

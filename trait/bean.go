package trait

type (
	cellBean interface{}

	beanProducer func() cellBean

	beanInitMiddleware func(beanProducer) beanProducer
)

package trait

import (
	"fmt"
	"testing"
)

type myClass struct {
	messaging
	logger
	privateData string
}

func (x *myClass) Shout() {
	fmt.Println(x.privateData)
}

func myClassFactory(d string) *myClass {
	return FmtLoggerInitMiddleware(ESMessagingInitMiddleware(func() cellBean {
		return &myClass{
			privateData: d,
		}
	}))().(*myClass)
}

func TestChainProducer(t *testing.T) {
	c := myClassFactory("hello")
	c.Log(c.privateData)
	c.SendMessage(c.privateData)
	c.Shout()
}

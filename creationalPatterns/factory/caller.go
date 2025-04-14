package factory

func Caller() {
	factory := NewFactory()
	//
	sms := NewSMS()
	//
	factory.SetNotify("sms", sms)
	//
	factory.GetNotify("sms")
	factory.GetNotify("sms").Send()
}

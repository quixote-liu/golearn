// 假设我现在要做一个短信推送的系统，那么需要
// 1、检查短信字数是否超过限制
// 2、检查手机号是否正确
// 3、发送短信
// 4、返回状态
// 我们可以发现，在发送短信的时候由于不同的供应商调用的接口不同，
// 所以会有一些实现上的差异，但是他的算法（业务逻辑）是固定的

package template

import "fmt"

// ISMS ISMS
type ISMS interface {
	send(content string, phone int) error
}

// SMS 短信发送基类
type sms struct {
	ISMS
}

// Valid 校验短信字数
func (s *sms) Valid(content string) error {
	if len(content) > 63 {
		return fmt.Errorf("content is too long")
	}
	return nil
}

// Send 发送短信
func (s *sms) Send(content string, phone int) error {
	if err := s.Valid(content); err != nil {
		return err
	}

	// 调用子类的方法发送短信
	return s.send(content, phone)
}

// TelecomSms 走电信通道
type TelecomSms struct {
	*sms
}

// NewTelecomSms NewTelecomSms
func NewTelecomSms() *TelecomSms {
	tel := &TelecomSms{}
	// 这里有点绕，是因为 go 没有继承，用嵌套结构体的方法进行模拟
	// 这里将子类作为接口嵌入父类，就可以让父类的模板方法 Send 调用到子类的函数
	// 实际使用中，我们并不会这么写，都是采用组合+接口的方式完成类似的功能
	tel.sms = &sms{ISMS: tel}
	return tel
}

func (tel *TelecomSms) send(content string, phone int) error {
	fmt.Println("send by telecom success")
	return nil
}

/*
	总结：sms所使用的组合逻辑就是一个模板，其中的send方法是基础接口，不同的运营商会
	调用不同的接口去实现这个接口，所以变的send方法里面的内容，而不变的模板定义的整个逻辑。
*/
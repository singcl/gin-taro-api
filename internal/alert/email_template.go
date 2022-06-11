package alert

import (
	"bytes"
	"fmt"
	"html/template"
	"time"
)

// NewHTMLEmail 告警邮件模板
func newHTMLEmail(method, host, uri, id string, msg interface{}, stack string) (subject string, body string, err error) {
	mailData := &struct {
		URL   string
		ID    string
		Msg   string
		Stack string
		Year  int
	}{
		URL:   fmt.Sprintf("%s %s%s", method, host, uri),
		ID:    id,
		Msg:   fmt.Sprintf("%+v", msg),
		Stack: stack,
		Year:  time.Now().Year(),
	}

	// subject 邮件主题
	subject = fmt.Sprintf("[系统告警]-%s", uri)

	// body 邮件内容
	body, err = getEmailHTMLContent(mailTemplate, mailData)

	return
}

// getEmailHTMLContent 获取邮件模板
func getEmailHTMLContent(mailTpl string, mailData interface{}) (string, error) {
	tpl, err := template.New("email tpl").Parse(mailTpl)
	if err != nil {
		return "", err
	}

	buffer := new(bytes.Buffer)
	err = tpl.Execute(buffer, mailData)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}

package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/xhyonline/xutil/helper"
	"github.com/xhyonline/xutil/mail"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	MailerConfig mail.Config `toml:"-"`
	Content      struct {
		Subject string   `toml:"subject"`
		From    string   `toml:"from"`
		ToUser  []string `toml:"to_user"`
		Body    string   `toml:"body"`
	} `toml:"content"`
}

// SendByConfig 通过配置文件发送
func SendByConfig(path, template, bodyPath string) error {
	exists, err := helper.PathExists(path)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("配置文件不存在")
	}
	cfg := new(Config)
	if _, err = toml.DecodeFile(path, cfg); err != nil {
		return fmt.Errorf("配置文件不存在")
	}
	cfg.MailerConfig.User = os.Getenv("MAILER_USER")
	if cfg.MailerConfig.User == "" {
		return fmt.Errorf("环境变量 MAILER_USER 不能为空")
	}

	cfg.MailerConfig.Pass = os.Getenv("MAILER_PASS")
	if cfg.MailerConfig.Pass == "" {
		return fmt.Errorf("环境变量 MAILER_PASS 不能为空")
	}

	cfg.MailerConfig.Host = os.Getenv("MAILER_HOST")
	if cfg.MailerConfig.Host == "" {
		return fmt.Errorf("环境变量 MAILER_PASS 不能为空")
	}

	if cfg.MailerConfig.Port, err = strconv.Atoi(os.Getenv("MAILER_PORT")); err != nil {
		return err
	}
	if cfg.MailerConfig.Port == 0 {
		return fmt.Errorf("环境变量 MAILER_PORT 不能为空")
	}
	if bodyPath!="" {
		if  exists, err = helper.PathExists(bodyPath);!exists||err!=nil{
			return fmt.Errorf("bodyPath 路径不存在")
		}
		body,err:=ioutil.ReadFile(bodyPath)
		if err!=nil {
			return err
		}
		cfg.Content.Body=string(body)
	}

	if template != "" {
		if cfg.Content.Body, err = TemplateReplace(template, cfg.Content.Body); err != nil {
			return err
		}
	}
	return mail.NewMail(cfg.MailerConfig).Send(cfg.Content.Subject, cfg.Content.From, cfg.Content.Body, cfg.Content.ToUser...)
}

// TemplateReplace 模板替换
func TemplateReplace(template, body string) (string, error) {
	leftIndex := strings.Index(template, "{{")
	rightIndex := strings.Index(template, "}}")
	if leftIndex == -1 || rightIndex == -1 {
		return body, nil
	}
	tag := template[leftIndex : rightIndex+2]
	if !strings.Contains(body, tag) {
		return "", fmt.Errorf("模板语法错误,不存在标签 %s 的替换", tag)
	}

	// 开始找该标签对应的内容
	template = strings.Replace(template, tag, "", 1)
	leftIndex = strings.Index(template, "{{")
	rightIndex = strings.Index(template, "}}")
	var info string
	// 意味着到底了
	if leftIndex == -1 || rightIndex == -1 {
		info = template[strings.Index(template, "=")+1:]
	} else {
		info = template[strings.Index(template, "=")+1 : leftIndex]
	}
	body = strings.Replace(body, tag, info, 1)
	body = strings.TrimRight(body,",")
	fmt.Println(body)
	template = strings.Replace(template, info, "", 1)
	template = strings.Replace(template, "=", "", 1)
	return TemplateReplace(template, body)
}

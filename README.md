# mailer
## 一、简介
mailer 是一个助手工具,提供命令行发送邮件的功能

## 二、安装方式

`go install github.com/xhyonline/mailer@latest`

## 三、配置方式
### (一)、配置环境变量
请在环境变量中配置邮件发送的基本配置
由以下几个环境变量(必须配置)

示例如下
```
MAILER_USER=xxxx.qq.com; 邮件账号   
MAILER_PASS=授权码;
MAILER_HOST=smtp.qq.com;
MAILER_PORT=465
```


`mailer.toml` 文件内容示例
```
[content]
subject="一封测试主题"
from="来自线上告警"
to_user=["xxxxxxx@qq.com"]
body="测试模板变量 {{xgo}} {{server}}模板变量"
```

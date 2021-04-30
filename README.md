# mailer
## 一、简介
mailer 是一个助手工具,提供命令行发送邮件的功能

## 二、安装方式

`go install github.com/xhyonline/mailer@latest`

## 三、配置方式
### (一)、配置环境变量
请在环境变量中配置邮件发送的基本配置
由以下几个环境变量组成(必须配置)

示例如下
```
MAILER_USER=xxxx.qq.com 邮件账号   
MAILER_PASS=授权码
MAILER_HOST=smtp.qq.com
MAILER_PORT=465
```

## (二)、你还需要配置收件人
你可以创建任意一个 `.toml` 结尾的文件,定义邮件主题与发送人等信息

例如 `/etc/mailer.toml` 文件内容示例
```
[content]
subject="一封测试主题"
from="来自线上告警"
# to_user 是一个数组,你可以在此添加收件人
to_user=["xxxxxxx@qq.com"]
body="你好我是 {{name}} 今年 {{age}}岁了"
```
通过命令即可发送邮件
```
mailer -c /etc/mailer.toml -template={{name}}=小明,{{age}}=24  
```
或者你可以自定义模板文件,例如创建 `/root/index.html`内容如下
```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
<h2>这是一封测试邮件 </h2>

<p>
    你好我是 {{name}} ,今年 {{age}} 岁了
</p>
</body>
</html>
```
然后定义 `/root/mailer.toml` 文件
```
[content]
subject="一封测试主题"
from="来自线上告警"
# to_user 是一个数组,你可以在此添加收件人
to_user=["xxxxxxx@qq.com"]
```
最终通过命令
```
mailer -c  /root/mailer.toml -template={{name}}=小明,{{age}}=24  -bodyPath /root/index.html  
```
来发送邮件
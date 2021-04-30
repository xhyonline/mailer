package main

import (
	"flag"
	"github.com/xhyonline/xutil/xlog"
	"os"
)




var logger=xlog.Get(true)

func main() {

	var (
		configPath=flag.String("c","","配置文件路径")	// like nginx -c nginx.conf

		template=flag.String("template","","模板变量:示例 {{xgo}}=xgo服务,{{server}}=后端服务")

		bodyPath=flag.String("bodyPath","","body是邮件正文的配置文件路径,你可以用 -bodyPath 指向某个 html 文件,并且也可以在其中使用模板变量,它将作为邮件正文发送,例如" +
			"-bodyPath=/var/template.html")

	)
	flag.Parse()

	if *configPath=="" {
		logger.Errorf("请设置参数完整的发送参数")
		os.Exit(1)
	}
	err:=SendByConfig(*configPath,*template,*bodyPath)
	if err!=nil {
		logger.Errorf("%s",err)
		os.Exit(1)
	}

}

---
title: "Systemctl 服务配置环境变量"
date: 2020-08-14T16:27:16+08:00
draft: false
---

今天在开发过程中发现一个接口报错，究其原因是调用了go的os.Getenv()方法时，没有获取到相应的环境变量。在.zshrc中export KEY=VALUE 之后，systemctl restart xx.service仍然不生效。
直接执行服务对应的可执行文件则不会出错。  
查资料发现，要在systemd处该服务的配置文件中加上Environment="key = value"。对于多个环境变量，要注意双引号，以免前面的环境变量把后面的key也包含进去。  
之后systemctl daemon reload 再用systemctl重启服务即可生效。
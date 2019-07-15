# gofun娱乐版

## 发送钉钉消息
- 上下班打卡
- 随机点店

# 使用方法
```
#下载
git clone https://github.com/yezihack/gofun.git
cd gofun
#编译
make dev
#添写有效的钉钉机器人token
vim run/gofun.toml

#运行
make run

#如果常驻内存
make deam
```

# 第三方包
- 读toml文件配置 github.com/BurntSushi/toml
- 简单的cache github.com/ThreeKing2018/goCache
- 调试利器spew github.com/davecgh/go-spew/spew
- 文件监控 github.com/fsnotify/fsnotify
- curl工具包 github.com/mikemintang/go-curl
- 计划任务工具包 github.com/robfig/cron
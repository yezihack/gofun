package server

import (
	"github.com/ThreeKing2018/k3log"
	"github.com/fsnotify/fsnotify"
	"time"
)

func Watcher() (err error) {
	go func() {
		watch, err := fsnotify.NewWatcher()
		if err != nil {
			k3log.Error("Watcher", err)
			return
		}
		//defer watch.Close()
		//监控文件
		err = watch.Add(ConfigFileAll)
		if err != nil {
			k3log.Error(err)
		}
		k3log.Info("对文件进行监控...", ConfigFileAll)
		for {
			select {
			case event := <-watch.Events:
				{
					k3log.Info("aaaa", event.Op)
					//判断事件发生的类型，如下5种
					// Create 创建
					// Write 写入
					// Remove 删除
					// Rename 重命名
					// Chmod 修改权限
					if event.Op&fsnotify.Remove == fsnotify.Remove ||
						event.Op&fsnotify.Rename == fsnotify.Rename ||
						event.Op&fsnotify.Write == fsnotify.Write ||
						event.Op&fsnotify.Create == fsnotify.Create {
						k3log.Info(event.Name)
						k3log.Info(event.String())
					}
				}
			default:
				time.Sleep(1 * time.Second)
			}
		}
	}()
	return
}

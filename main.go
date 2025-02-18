package main

import (
	"fmt"
	"log"
	"os"
	"translate/apiservice"

	"github.com/go-vgo/robotgo/clipboard"
	hook "github.com/robotn/gohook"
)

func main() {
	// 打开日志文件
	logFile, err := os.OpenFile("monitor.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	defer logFile.Close()

	// 设置日志输出
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)

	// 注册快捷键 Ctrl+T
	hook.Register(hook.KeyDown, []string{"ctrl", "t"}, func(e hook.Event) {
		// 读取剪贴板内容
		content, err := clipboard.ReadAll()
		if err != nil {
			log.Println("read clipboard failed, err:", err)
			return
		}
		// 输出剪贴板内容
		fmt.Println("Clipboard content:", content)
		log.Println("Clipboard content:", content)

		err = clipboard.WriteAll("正在翻译，请稍后")
		if err != nil {
			log.Println("write clipboard failed, err:", err)
			return
		}

		sk := apiservice.NewDeepSeekService()
		resp, err := sk.SendDeepSeekRequest(content)

		if err != nil {
			log.Println("send deepseek request failed, err:", err)
			return
		}
		fmt.Println("deepseek response:", resp)
		err = clipboard.WriteAll(resp.Choices[0].Message.Content)
		if err != nil {
			log.Println("write clipboard failed, err:", err)
			return
		}
	})

	// 注册快捷键 Esc 用于退出程序
	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		fmt.Println("Exiting program...")
		hook.End()
	})

	// 启动钩子
	s := hook.Start()
	<-hook.Process(s)
	fmt.Println("The program is over!")
}

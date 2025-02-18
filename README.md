# KeyTranslate

> 该项目帮助你通过快捷键一键翻译内容为英文。

## 安装与配置

1. 在项目的根目录下创建一个 `.env` 文件。
2. 在 `.env` 文件中写入以下内容：

   ```env
   KEY=xxx
   ```

   其中，`xxx` 为你从 SiliconCloud 获取的 API key。

## 使用方法

1. 执行以下命令来运行项目：

   ```bash
   go run main.go
   ```

2. 复制你想要翻译的内容。
3. 按下快捷键 `Ctrl + T`（Windows/Linux）或 `Cmd + T`（Mac）。
   - 该操作将会把你复制的内容传递到大模型进行翻译，翻译结果会返回给你。
4. 接下来，你可以按 `Ctrl + V` 或 `Cmd + V`（根据操作系统不同）将翻译后的内容粘贴到目标位置。

## 注意事项

- 确保你已经设置了正确的 API key。
- 确保在使用过程中网络连接正常，以便获取翻译结果。

## License

此项目遵循 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件。

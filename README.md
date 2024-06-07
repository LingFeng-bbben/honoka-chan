# honoka-chan

LoveLive! 学园偶像祭、学园偶像季: 群星闪耀 自用私服。

## 文档列表（更新中）

### 主程序部署

首先确认已经安装了 Go 环境以及 Git 工具，然后clone本项目到本地。

进入 `honoka-chan` 文件夹，执行 `go build` 编译，会自动下载依赖。

编译完成运行 `honoka-chan.exe` 即可，Linux 上可执行文件为 `honoka-chan`。

端口为8080，若被占用请解决占用问题，端口号已经写死惹

用电脑进入`http://127.0.0.1:8080/`，看到hello world就说明可以了

### 下载数据包（仅有安卓）

解压android.zip到static文件夹。解压后确保文件结构是像static/Android/archives/9243个数据包这样

### 启动hosts

下载[personal DNS filter](https://github.com/IngoZenz/personaldnsfilter)装进板子里

高级设置>配置其他主机>最下面加入> `>www.llsif.sb 你的服务器ip`

用平板启动浏览器看`http://www.llsif.sb:8080/`，有hello world说明可以了。

## 启动游戏

用密码登陆 随便输
下载数据的时候建议直接选全部数据
（记得调速）

## 特别感谢（下面是一些很牛逼的人）

 - 虹原翼 (yazawa@niconi.co.ni) @ [LLSIF 查卡器
](https://card.niconi.co.ni/) 的技术支持。
 - tungnotpunk @ LL Hax 在 iOS 客户端修改方面的支持。
 - reversing-sifas (https://github.com/Francesco149/reversing-sifas) 项目的灵感和技术细节。
 - LLAS WiKi (https://wiki.loveliv.es/) 提供的全量 AS Live Stage 信息。
 - 其他可能一时没想起来的大佬们和项目。
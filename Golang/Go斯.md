[toc]

## Go语言介绍

### 一、GO语言简介

Go（又称`Golang`）是`Google`开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。

`2009年11月10号`Go语言正式成为了我们开源编程语言中的一员，同时Go语言又被称为云时代的C语言，其地位可以想象，此时肯定会有同学会问：世界上这么多编程语言难道还不够吗？我们程序员就要掌握几十种编程语言或者上百种吗？
其实不是的，Go语言的产生主要是为了能让我们程序员再编程上有更高的一个效率，举例：

如果我们现在要开发一个爬虫，但是我们用c语言的话代码可能要写几百行，这样的话代码量太多，用python的话虽然代码量少但是执行速度慢，我们执行这个程序，c语言执行可能只要0.5秒，python的话可能要用到5秒或者10秒，所以就有了这个Go语言。
Go语言拥有了c语言的执行速度又拥有了python的快速开发。

### 二、作者介绍

相信大家肯定很好奇到底是什么样的人才能开发的出来Go语言，接来下我就来给大家介绍一下。

![99999.jpg](https://www.zlkt.net/media/course/y2snviYUp6grnjX2KiAMnQ.jpg)

从左到右分别是 ：

1. **罗伯特·格瑞史莫（Robert Griesemer）：** 曾为谷歌的V8 JavaScript引擎和Chubby开发代码。现任Go作者之一一职。作为行业享有盛名的大咖，Robert Griesemer行事低调，对工作热情饱满，多次受邀作为嘉宾出席各类大会，并发表了精彩演讲。
   2015年04月25日,Robert Griesemer受邀参加了在浦东新区博云路2号浦东软件园的《GOPHER CHINA 2015上海大会》,并发表了精彩演讲
2. **罗勃·派克（Rob Pike）：** 罗布·派克是Unix的先驱，是贝尔实验室最早和Ken Thompson以及 Dennis M. Ritche开发Unix的猛人，UTF-8的设计人。他还在美国名嘴David Letterman的晚间节目上露了一小脸，一脸憨厚地帮一胖子吹牛搞怪。让偶佩服不已的是，罗伯伯还是1980年奥运会射箭的银牌得主。他还是个颇为厉害的业余天文学家，设计的珈玛射线望远镜差点被NASA用在航天飞机上。Rob Pike AT&T Bell Lab前Member of Technical Staff ，现在google研究操作系统。
3. **汤普逊（Ken Thompson）：** Ken Thompson图灵奖得主，C语言前身B语言的作者，Unix的发明人之一， 操作系统Plan 9的主要作者。 共同开发了UTF-8。

### 三、为什么设计GO语言？

既然都已经有了这么多编程语言为什么还要开发一个Go语言呢？

当时设计Go的目标是为了消除程序各种缓慢和笨重、改进各种低效和扩展性，解决这些当时`Google`开发遇到的问题。

- 大量的`C++`代码，同时又引入了`Java`和`Python`。
- 成千上万的工程师。
- 数以万计行的代码。
- 分布式的编译系统。
- 数百万的服务器。

### 四、为什么选择使用Go语言,或者说Go语言有哪些特点？

在编程界，编程语言可谓是百花争鸣，追求运行速度可以用`C`，追求开发速度可以用`Python`。作为一个新起之秀，`Go`语言究竟有哪些优势，能让吸引这么多忠诚的追随者？以下从几个部分来谈谈`Go`语言的优势。

#### 1. 上手难度低：

`Go`语言的语法有`C`和`Python`的影子，如果你之前学习过其中一种语言，你会发现他们是非常相像的，所以只要有过编程基础的朋友，学习`Go`语言基本上没有什么技术难点。

#### 2. 快：

快是指执行效率和开发效率。作者在开发`Go`语言时候，就有意把`Go`的执行效率和`C`语言进行比较，因此执行效率是除`C`以外其他语言难以媲美的。另外，在语法设计方面，也尽量设计得和`Python`是非常相像，所以开发效率也是极高。在国内，已经有很多中小型公司的技术栈就是基于`Go`语言实现的，因为他的开发效率可以满足创业公司产品快速迭代的需求，并且性能也是很高的。

#### 3. 出身名门：

`Go`语言是由谷歌公司负责牵头，召集了三位人类巅峰的计算机大牛负责开发，其中Robert Griesemer参与开发`Java HotSpot`虚拟机，Rob Pike贝尔实验室`Unix`团队成员，Ken Thompson贝尔实验室`Unix`团队成员，`C`语言创始人之一。阵容豪华，背景强大，可以大胆放心的在公司中使用。

#### 4. 组合：

支持当前所有的编程方式，过程式编程，面向对象编程，函数式编程，开发起来自由高效。任何语言的程序员转`Go`，都可以找到自己喜欢的编程方式。

#### 5. 工具箱：

拥有超多超好用的第三方库，如果你用Go语言写爬虫，能拥有`C`语言的爬行速度又能拥有`Python`那样的少量代码。也是借鉴了`Python`的那句：人生苦短，快用`Go`。

#### 6. 技能释放：

在辛辛苦苦写完一套项目后，如何部署又是一个大问题。而我们的`Go`语言项目在编译后会生成一个可执行的文件，当我部署时只要把这个文件放上去即可，极为方便，大大的提高了运维的效率，技能释放更加迅速。

#### 7. 天生神力：

天生就支持并发，相比于`Python`、`PHP`等编程语言来说，`Go`天然的就更适应大型项目的开发。

### 五、GO语言的应用范围

1. 服务器编程，以前你如果使用`C`或者`C++`做的那些事情，用`Go`来做很合适，例如处理日志、数据打包、虚拟机处理、文件系统等。
2. 分布式系统、数据库代理器等，例如`Etcd`。**为分布式系统而生**
   这一点，在那些用`Golang`打造的工具和软件中显而易见。`Docker`，作为一种微服务的容器，也是用 `Golang`打造。
3. 网络编程，这一块目前应用最广，包括`Web`应用、`API`应用、下载应用，而且`Go`内置的`net/http`包基本上把我们平常用到的网络功能都实现了。
4. 内存数据库。
5. 云平台，目前国外很多云平台在采用Go开发。
6. 区块链，区块链是分布式数据存储、点对点传输、共识机制、加密算法等计算机技术的新型应用。(区块链的服务器云存储那部分就是使用go开发的)
   ![201803131537439596278181.png](https://www.zlkt.net/media/course/YRveYXMngZdm6VHZSiAdoJ.png)

### 六、哪些公司在使用GO语言？

随着Go语言的发展趋势越来越好现在国内国外很多知名公司都有使用Go语言，那到底有哪些公司呢？以下做一个简短描述：

1. 国外的有：
   - Google
   - Docker
   - Facebook
2. 国内的有：
   - 腾讯
   - 百度
   - 阿里云CDN
   - 京东
   - 小米
   - 360

现在很多公司都开始尝试`Golang`，除了上面提到的，还有美团、滴滴、新浪以及七牛等。一般的选择，都是选择用于自己公司合适的产品系统来做，比如消息推送的、监控的、容器的等，`Golang`特别适合做网络并发的服务，这是他的强项，所以也是被优先用于这些项目。

# 区块链介绍

区块链是分布式数据存储、点对点传输、共识机制、加密算法等计算机技术的新型应用模式。

区块链（Blockchain）是比特币的一个重要概念，它本质上是一个去中心化的数据库，同时作为比特币的底层技术，是一串使用密码学方法相关联产生的数据块，每一个数据块中包含了一次比特币网络交易的信息，用于验证其信息的有效性（防伪）和生成下一个区块。这里以借钱为例说明一下。

## 一、没有用区块链借钱：

小明找老王借了1000块钱，过了一段时间之后老王去找小明拿钱，小明却说根本没有这个事情，你又没有证据，有谁能够证明我欠了你钱。
![111.png](https://www.zlkt.net/media/course/QzTRSAEoNMAcK2b7N6ViH.png)

### 中间人：

小明又找老王借2000块钱，老王又借给他了，但是这一次让支付宝做中间人，可能前面小小的借钱没有什么关系，但是如果是一千万呢？借出去之后小明拿着一千万雇了个人让支付宝消失了，接下来老王要小明还钱的时候支付宝都没了，你又没有证据了。所以才有了这个区域链的产生。
![2626.png](https://www.zlkt.net/media/course/x2SgiZySQPGoGo5m5n56XH.png)

## 二、用区块链借钱：

简单的讲，区块链就是去中心化的分布式账本。怎么样去中心化，就是没有中心，或者说每个人都可以是中心，这是和传统的中心化方式不同的。分布式账本，意味着数据的存储不只是在每一个节点上，而是每一个节点会复制并共享整个账本的数据。

还是小明和老王的例子：
小明又找老王借一千万，这次走了个某借钱App的程序，然后某借钱App把钱再给小明，下面步骤解释。

1. 老王把钱传入某借钱App
2. 老王用笔记本记起来我把钱给了
3. 然后传播开来，我借给了小明一千万
4. 某借钱App就把钱发给小明的这条数据保存起来，也保存到各个地方
5. 小明也用笔记本记起来我欠老王一千万
6. 然后传播开来
7. 自此流程完成

这样的话就算其中有人不承认，但是笔记本上面都有这个记录，而且这件事情也传播开来了，很多人都知道，所以想耍赖也是不行的。

# Go语言环境搭建

## 一、安装

相信大家很多人都认为安装这一块是最麻烦的也是最难的，其实只要你跟着我一步一步走就不会有问题。

### 1. 下载程序安装包：

首选不管我们是什么编程语言都得先有安装包，go语言就有go语言的安装包，c语言就有c语言的安装包，他们的安装包不是一样的，所以我们先去官网下载go语言的一个安装包

Go安装包下载网址：https://www.golangtc.com/download

![15446166551.jpg](https://www.zlkt.net/media/course/nbou2T5SvwyYUwtSuVViJn.jpg)

这里面有zip压缩版和msi安装版两个版本下载我们选择msi版本的。（这里使用msi安装版，比较方便）。

### 2. 安装及环境配置

然后我们运行msi安装文件，记住大家安装的时候系统会让你保存的，不要保存到中文的文件夹里面就行，千万不要在安装路径中出现中文不然会安装失败，然后就是傻瓜式安装一路Next。

![JAAT0PC8NWBTGL_VT.png](https://www.zlkt.net/media/course/FmCGveja44VV3DyGV5tf4a.png)

安装过程简单，一路“next”即可，也可以定义 安装目录
**由于使用msi安装文件，所以Go语言的环境变量已经自动设置好了。**

#### 2.1. 检查是否安装成功（进入黑窗口的方式：win+r）

或者你直接点击你左下角的“开始”然后输入cmd就会显示一个命令之令符，进去之后我们输入`go version`查看GO版本，如下图一样就是安装成功：
![15450241351.jpg](https://www.zlkt.net/media/course/G8nFfycBAJMepUGeFrevyX.jpg)

#### 2.2. 查看环境变量:

输入`go env`。查看`GOROOT`，因为`GOROOT=C:\GO`这是我的安装路径，你确保你的安装路径和这个是一样的就行下图所示：
![0K1R9Z7J9DUHA0V0PAGO.png](https://www.zlkt.net/media/course/kTqDtAmqiNDBdf2TpnJyK5.png)

#### 2.3. 设置工作空间`GOPATH`目录(Go语言开发的项目路径)

首先进入我的C盘（你放到其他盘也行），新建一个文件夹，名字叫做`mygo`（这个就是你的工作目录），然后再进入这个`mygo`文件夹，到里面再新建一个文件夹，名字叫做`src`（用来存放代码）。

然后把我们这个工作路径复制下来，我的是`C:\mygo`。

然后我们去设置环境变量，右击一下“我的电脑”，最下面有个属性点击一下，然后我们点击高级系统设置，最后再最下面有个环境变量。

我们点击系统变量的“新建”，变量名设置为：`GOPATH`，变量值就是工作路径（我的是`C:\mygo`），然后点确定就行了。
![Go环境变量](https://www.zlkt.net/media/course/BRh6gJBCbfsJnU6KDoP6LL.png)

## 二、 LiteIDE安装：

`LiteIDE`是一个非常方便的开发工具。`LiteIDE`上手简单，而且是开源的，也是跨平台的。

1. 官网：http://liteide.org/cn/
2. 发行版下载地址：https://sourceforge.net/projects/liteide/files

### 1. 下载和安装LiteIDE：

![kaifagongju.png](https://www.zlkt.net/media/course/Z8BA5Z6xMLSQHPMn8YjZV7.png)

点击这个`x35.2`，点击进去之后就会出现下图：
![kaifagongju2.png](https://www.zlkt.net/media/course/TfTGESGc6QvR29H4DLr4iA.png)

然后选择最新的一个版本，选择`5.9.5.zip`，点击下载
下载完成之后就把这个zip解压，解压完成就会有一个文件夹，点击文件夹就会看到一个bin点击进入bin。

![_4UPB5QRGE89I4RGR.png](https://www.zlkt.net/media/course/KBe8TvevU7DEFnwovREGCT.png)

你可以看到我箭头所示有一个`liteide`，然后我们双击打开就能使用了。

### 2. 新建一个小项目：

我们点击新建，然后选择`empty File`，然后选择最下方的一个位置，这是我们保存代码的位置，我们点击浏览选择我们前面设置好的`c:\mygo\src`，你选择你前面设置好的就行。再给他起个名字叫做`test.go`,然后点击ok，立即加载。
![Snipaste_20181218_113558.png](https://www.zlkt.net/media/course/R5avUG8q5bcmj4iaaJ3q2T.png)

在`test.go`中输入下面代码：

```
package main
import "fmt"
func main(){
  fmt.Println("hello")
}
```

然后直接按`Ctrl+R`就能运行。
![Snipaste_20181218_113747.png](https://www.zlkt.net/media/course/DJoC4Z8BbR2y4nMzRdU2mk.png)
可以看到运行成功了。

**注意：如果你执行不出来点击上方工具栏，查看编辑当前环境，看看`GOROOT`是否和你的安装目录一致！**
![Snipaste_20181218_112953.png](https://www.zlkt.net/media/course/fUZ2nEsdj6USU3DSyEXkHQ.png)

## 三、Goland安装

1. 官网下载地址：https://www.jetbrains.com/go/
2. 选择电脑系统对应的版本
   ![1544845656.jpg](https://www.zlkt.net/media/course/EX52RJcvrkbV2v9SAaYaKF.jpg)
3. 打开安装
   ![15448459761.jpg](https://www.zlkt.net/media/course/kaGCujvsSEqnKt4rhXzAwS.jpg)
4. 点击“next”按钮，选择要安装的路径，然后点击“next”，会出现安装选项。根据你自己电脑的型号，选择合适的版本后点击“next”按钮。
   ![15448464061.jpg](https://www.zlkt.net/media/course/eMy2bbNKtk5q22ePr4pjGg.jpg)
5. 点击安装：
   ![1544846615.jpg](https://www.zlkt.net/media/course/XppWkyBB6ycFpDS4mtpAQT.jpg)
6. 点击完成：
   ![999.jpg](https://www.zlkt.net/media/course/xMe2HSoRtacVXDPWnUAM6m.jpg)
7. 打开Goland激活，输入激活码（激活码自己想办法解决），即可使用。

## 四、标准命令概述：

Go语言中包含了大量用于处理Go语言代码的命令和工具。其中，go命令就是最常用的一个，它有许多子命令。这些子命令都拥有不同的功能，如下所示。

- `build`：用于编译给定的代码包或`Go`语言源码文件及其依赖包。
- `clean`：用于清除执行其他`go`命令后遗留的目录和文件。
- `doc`：用于执行`godoc`命令以打印指定代码包。
- `env`：用于打印`Go`语言环境信息。
- `fix`：用于执行`go tool fix`命令以修正给定代码包的源码文件中包含的过时语法和代码调用。
- `fmt`：用于执行`gofmt`命令以格式化给定代码包中的源码文件。
- `get`：用于下载和安装给定代码包及其依赖包(提前安装git或hg)。
- `list`：用于显示给定代码包的信息。
- `run`：用于编译并运行给定的命令源码文件。
- `install`：编译包文件并编译整个程序。
- `test`：用于测试给定的代码包。
- `tool`：用于运行Go语言的特殊工具。
- `version`：用于显示当前安装的Go语言的版本信息。

## 五、学习资料

1. Go语言官网(需要翻墙)：https://golang.org/
2. 配套课程：https://study.163.com/course/introduction.htm?share=2&shareId=1025897964&courseId=1006508031

# 几种Go开发工具介绍

## 一、LiteIDE

LiteIDE是一款专门为Go语言开发的跨平台轻量级集成开发环境（IDE），由`QT`编写。
![b5ace4f69028404995b7a17d0e0e831d.jpg](https://www.zlkt.net/media/course/pRL3j5TfoCYC46zz756DYa.jpg)

由于它是为`Golang`直接设计的，LiteIDE为开发人员提供了许多有用的功能，包括可配置的构建命令，高级代码编辑器和广泛的`Golang`支持。其他功能包括代码管理，`gdb`和`Delve`调试器，自动完成和使用`WordApi`的主题，基于MIME类型的系统等等。

## 二、Goland

JetBrains公司出品，Goland就是专门针对Go语言而打造的一款编程工具，未付费版本只能试用30天。但是专业做`Go`开发还是推荐使用这款开发工具。
![004ee08b.jpg](https://www.zlkt.net/media/course/aJsd6PvA4LonnSa3tYX3MY.jpg)

## 三、VS Code

它是微软开发的广受欢迎的开源IDE，有一个开箱即用的Go扩展可供VS Code使用。`vscode-go`插件为开发人员提供了更多功能，包括与许多`Go`工具集成。
![bfb4455f22e04ede9ebc4042d1184850.jpg](https://www.zlkt.net/media/course/jdXzryogkndUgg6SkRDEek.jpg)

VS Code通过`IntelliSense`，内置`Git`集成，直接从编辑器调试代码等功能提供智能完成功能。`VS Code`具有高度可扩展性，并通过其许多扩展提供了许多自定义选项。它还提供了几十种语言的支持，这使得它成为了受开发者欢迎的工具。

## 四、Sublime text

这个文本编辑器在开发者中较为普及，应该说sublime并非一个完全成熟的IDE，但是它具备很多语言的扩展插件，比如Python、Lua等，其中有一个插件`GoSublime`专门针对go语言，`GoSublime`提供了语法高亮、自动补全等功能，这些功能使得Sublime Text成为一很实用的Go IDE。
![sublimegolang.png](https://www.zlkt.net/media/course/UKL8St5koDi5AxHWxDrrtk.png)

## 五、Eclipse

GoClipse是Eclipse的插件。使用GoClipse插件，开发人员可以利用流行的Eclipse IDE进行编程。Eclipse IDE和GoClipse插件都是免费和开源的。GoClipse编辑器为开发人员提供了广泛的功能，包括源代码编辑器，项目向导和构建器，以帮助报告编辑器内构建的错误，以及功能齐全的GDB调试器支持。

![035bc2d15d6d44b8bd512a7c73981b6e.jpg](https://www.zlkt.net/media/course/GsDwpJVDYkz444UHo4mxye.jpg)

## 六、Vim

Vim有许多插件可以帮助开发人员更轻松地编辑他们的Go代码。vim-go插件自动安装所有必要的东西，为Vim中的Go开发人员提供更平滑的集成。

Vim-go具有许多有用的功能，包括编译器，改进的语法高亮和折叠，完成支持以及一系列具有集成支持的调试程序。还有一些使用的高级源分析工具，包括：`GoImplements`，`GoCallees`和`GoReferrers`。

![37acd69e32b24f6db5d31c26278ebd9b.jpg](https://www.zlkt.net/media/course/kviSXtAabPRJ8BbRnhW2Bc.jpg)

其他vim插件包括用于反馈编译器错误的`Syntastic`插件，用于`Gotags`的`tagbar`插件，用于语法检查的vim编译器插件，甚至用于生成`.virmrc`配置的`vim-bootstrap`。

# 第一个Go程序

## 一、Hello Go

```
// hello.go
package main
import "fmt"
func main() {
  fmt.Println("Hello Go!")
}
```

## 二、代码分析

1. 每个`Go`源代码文件的开头都是一个`package`声明，表示该`Go`代码所属的包。包是`Go`语言里最基本的分发单位，也是工程管理中依赖关系的体现。

2. 要想成功运行一个程序，必须建立一个名字为`main`的包，并且在该包中包含一个叫`main()`的函数（该函数是`Go`可执行程序的执行起点）。

3. `Go`语言的`main()`函数不能带参数，也不能定义返回值。

4. 在包声明之后，是一系列的`import`语句，用于导入该程序所依赖的包。由于本示例程序用到了`Println()`函数，所以需要导入该函数所属的`fmt`包。

5. 所有`Go`函数以关键字`func`开头。一个常规的函数定义包含以下部分：

   ```
   func函数名(参数列表)(返回值列表) {
     // 函数体
   }
   ```

6. `Go`程序的代码注释与`C++`保持一致，即同时支持以下两种用法：

   - /* 块注释这个中间是可以注释的内容 */
   - // 行注释

7. `Go`程序并不要求开发者在每个语句后面加上分号表示语句结束，这是与`C`和`C++`的一个明显不同之处。

8. **注意：** 强制左花括号`{`的放置位置（不能单独起一行），如果把左花括号另起一行放置，`Go`编译器会编译错误。

## 三、 命令行运行程序

![15448644451.jpg](https://www.zlkt.net/media/course/wX9ZfXfdSCET9eYQHawFMn.jpg)

# 变量命名

`Go`语言中的函数名、变量名、常量名、类型名、语句标号和包名等所有的命名，都遵循一个简单的命名规则：一个名字必须以一个字母（`Unicod`e字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。并且变量命名是大小写敏感的，A和a是两个不同的名字。

```
a="老王"
A="老王"
```

还有一些关键字也不能使用，那么关键字有哪些呢？

Go语言中类似`if`和`switch`的关键字有25个(均为小写)。关键字不能用于自定义名字，只能在特定语法结构中使用。

| break    | default     | func   | interface | select |
| :------- | :---------- | :----- | :-------- | :----- |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

此外，还有大约30多个预定义的名字，比如`int`和`true`等，主要对应内建的常量、类型和函数。

1. 内建常量:`true false iota nil`。

2. 内建类型:

   ```go
       int int8 int16 int32 int64
   uint uint8 uint16 uint32 uint64 uintptr
   float32 float64 complex128 complex64
   bool byte rune string error
   ```

3. 内建函数:

   ```go
   make len cap new append copy close delete
   complex real imag
   panic recover
   ```

# 变量

变量是几乎所有编程语言中最基本的组成元素，变量是程序运行期间可以改变的量，举例：

比如你生了个娃，然后起了个名字入了户口`var name string="小明"`，但是这个娃长大之后就不喜欢这个名字了，他就决定改一个名字改成老王`name="老王"`，以后这个娃的名字就变成了老王，这就是变量。

从根本上说，变量相当于是对一块数据存储空间的命名，程序可以通过定义一个变量来申请一块数据存储空间，之后可以通过引用变量名来使用这块存储空间。

## 一、变量声明

`Go`语言的变量声明方式与`C`和`C++`语言有明显的不同。对于纯粹的变量声明，`Go`语言引入了关键字`var`，而类型信息放在变量名之后，示例如下：

```
var x1 int
var x2 int

//一次定义多个变量
var x3, x4 int

var (
  x5 int
  x6 int
)
```

`var`就是说你这是声明的一个变量，x1是自己起的变量名字，`int`是你给他指定的数据类型，像`var x1 int`,你打印出来这个代码的话结果为`0`就像孩子刚生出来就是`0`岁一样。

## 二、变量初始化

对于声明变量时需要进行初始化的场景，`var`关键字可以保留，但不再是必要的元素，如下所示：

```
var x1 int = 10 // 方式1
var x2 = 10 // 方式2，编译器自动推导出x2的类型
x3 := 10 // 方式3，编译器自动推导出x3的类型
fmt.Println("x3 type is ", reflect.TypeOf(x3)) //x3 type is int

//出现在 := 左侧的变量不应该是已经被声明过，:=定义时必须初始化
var x4 int
x4 := 2 //err
```

初始化的意思就是说别的孩子刚生出来都是0岁，我让我孩子刚生出来就是10岁，也就是`var x1 int =10` 这样就行了。

## 三、变量赋值

```
var x1 int
x1 = 123

var x2, x3, x4 int
x2, x3, x4 = 1, 2, 3 //多重赋值

i := 10
j := 20
i, j = j, i //多重赋值
```

你要是想要你的孩子长快一点你就可以自己给他定义岁数，比如我孩子是 `var x1 int` 默认是`0`岁，我可以`x1=100`岁这样子我孩子就`100`岁了，这就是简单的赋值。

## 四、 匿名变量

_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃：

```
_, i, _, j := 1, 2, 3, 4

func test() (int, string) {
  return 88, "old"
}

_, str := test()
```

# 常量

在Go语言中，常量是指编译期间就已知且不可改变的值。常量可以是数值类型（包括整型、浮点型和复数类型）、布尔类型、字符串类型等。

## 一、字面常量(常量值)

所谓字面常量（literal），是指程序中硬编码的常量，如：

```
123
3.1415 // 浮点类型的常量
3.2+12i // 复数类型的常量
true // 布尔类型的常量
"foo" // 字符串常量
```

也就是说印刷厂刷钱，那个钱一刷出来就是一百块或者20块，它后面是改变不了这个值的，就算你使用了这个20块那么他还是20块，只是不在你这里了而已。常量的意思是不可改变的。

## 二、常量定义

```
const Pi float64 = 3.14
const zero = 0.0 // 浮点常量, 自动推导类型

const ( // 常量组
  size int64 = 1024
  eof = -1 // 整型常量, 自动推导类型
)
const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
const a, b, c = 3, 4, "foo"
// a = 3, b = 4, c = "foo" //err, 常量不能修改
```

只要是使用`const`修饰的变量都是不能修改的。

## 三、iota枚举

常量声明可以使用`iota`常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。**在一个`const`声明语句中，在第一个声明的常量所在的行，iota将会被置为0，后面每加新增一行，常量值会自动加1**。

```
const (
  x = iota // x == 0
  y = iota // y == 1
  z = iota // z == 2
  w // 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

const (
  h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
  a = iota //a=0
  b = "B"
  c = iota //c=2
  d, e, f = iota, iota, iota //d=3,e=3,f=3
  g = iota //g = 4
)

const (
  x1 = iota * 10 // x1 == 0
  y1 = iota * 10 // y1 == 10
  z1 = iota * 10 // z1 == 20
)
```

也就是说：就像生孩子一样，跨一行就是比你早生一年，那你还是0岁，别人在你下面那行的就1岁，在你下面两行的就2岁，如果和你同一行的那就是和你同岁，他0岁你也0岁。

```
//大一年的例子：
const (
  x = iota // x == 0
  y = iota // y == 1
  z = iota // z == 2
  w // 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)
//同年的例子：
const (
  h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)
```

# 基础数据类型

## 一、分类

Go语言内置以下这些基础类型：

| **类型**      | **名称** | **长度** | **零值** | **说明**                                      |
| :------------ | :------- | :------- | :------- | :-------------------------------------------- |
| bool          | 布尔类型 | 1        | false    | 其值不为真即为假，不可以用数字代表true或false |
| byte          | 字节型   | 1        | 0        | uint8别名                                     |
| rune          | 字符类型 | 4        | 0        | 专用于存储unicode编码，等价于uint32           |
| int, uint     | 整型     | 4或8     | 0        | 32位或64位                                    |
| int8, uint8   | 整型     | 1        | 0        | -128 ~ 127, 0 ~ 255                           |
| int16, uint16 | 整型     | 2        | 0        | -32768 ~ 32767, 0 ~ 65535                     |
| int32, uint32 | 整型     | 4        | 0        | -21亿~ 21亿, 0 ~ 42亿                         |
| int64, uint64 | 整型     | 8        | 0        |                                               |
| float32       | 浮点型   | 4        | 0.0      | 小数位精确到7位                               |
| float64       | 浮点型   | 8        | 0.0      | 小数位精确到15位                              |
| complex64     | 复数类型 | 8        |          |                                               |
| complex128    | 复数类型 | 16       |          |                                               |
| uintptr       | 整型     | 4或8     |          | ⾜以存储指针的uint32或uint64整数              |
| string        | 字符串   |          | “”       | utf-8字符串                                   |

## 二、布尔类型

```
var v1 bool
v1 = true
v2 := (1 == 2) // v2也会被推导为bool类型

//布尔类型不能接受其他类型的赋值，不支持自动或强制的类型转换
var b bool
b = 1 // err, 编译错误
b = bool(1) // err, 编译错误
```

布尔类型只有`true`和`false`两个值。

## 三、整型

```
var v1 int32
v1 = 123
v2 := 64 //v1将会被自动推导为int类型
```

就以我们国家的货币，100，50，20，10，5，2，1，0.5其中除了0.5其他的都是整型，只要是数字而且后面没有小数点的就是整型。

## 四、浮点型

```
 var f1 float32
 f1 = 12
 f2 := 12.0 // 如果不加小数点， f2会被推导为整型而不是浮点型，float64
```

浮点型刚好和整型相反，像0.1，0.5这种就是浮点型，只要数字后面带了小数点的都是浮点型。

## 五、字符类型

在Go语言中支持两个字符类型，一个是`byte`（实际上是`uint8`的别名），代表`utf-8`字符串的单个字节的值。另一个是`rune`，代表单个`unicode`字符。

```
package main

import (
    "fmt"
)

func main() {
    var ch1, ch2, ch3 byte
    ch1 = 'a' //字符赋值
    ch2 = 97 //字符的ascii码赋值
    ch3 = '\n' //转义字符
    fmt.Printf("ch1 = %c, ch2 = %c, %c", ch1, ch2, ch3)
}
```

单引号的就是字符，双引号的就是字符串，
如果你`var a byte='0'`输出的结果并不是0，而是48,`var a byte='a'`输出结果为97,`var a byte='A'`输出结果为`65`，单引号里面放`0-9`输出的结果为`48-57`。单引号里面放`a-z`输出结果为`65-90`。单引号里面放`A-Z`输出结果为`97-117`，所有的符号在单引号里面都会有结果，有兴趣的同学可以去看字符的`ASCII`对照表。

## 6、 字符串

在Go语言中，字符串也是一种基本类型：

```
    var str string // 声明一个字符串变量
    str = "abc" // 字符串赋值
    ch := str[0] // 取字符串的第一个字符
    fmt.Printf("str = %s, len = %d\n", str, len(str)) //内置的函数len()来取字符串的长度
    fmt.Printf("str[0] = %c, ch = %c\n", str[0], ch)

    //`(反引号)括起的字符串为Raw字符串，即字符串在代码中的形式就是打印时的形式，它没有字符转义，换行也将原样输出。
    str2 := `hello
    mike \n \r测试
    `
    fmt.Println("str2 = ", str2)
    /*
        str2 = hello
             mike \n \r测试
    */
```

你在字符串里面赋值的任何东西它都是原样的，不会改变，`var a string="小明"`，你打印`a`得出的结果还是小明，不管你放什么进去都是一样的结果.

## 7、 复数类型

复数实际上由两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag）。

```
    var v1 complex64 // 由2个float32构成的复数类型
    v1 = 3.2 + 12i
    v2 := 3.2 + 12i // v2是complex128类型
    v3 := complex(3.2, 12) // v3结果同v2

    fmt.Println(v1, v2, v3)
    //内置函数real(v1)获得该复数的实部
    //通过imag(v1)获得该复数的虚部
    fmt.Println(real(v1), imag(v1))
```

# fmt包的格式化

## 一、格式化说明

| **格式** | **含义**                                                     |
| :------- | :----------------------------------------------------------- |
| %%       | 一个%字面量                                                  |
| %b       | 一个二进制整数值(基数为2)，或者是一个(高级的)用科学计数法表示的指数为2的浮点数 |
| %c       | 字符型。可以把输入的数字按照ASCII码相应转换为对应的字符      |
| %d       | 一个十进制数值(基数为10)                                     |
| %e       | 以科学记数法e表示的浮点数或者复数值                          |
| %E       | 以科学记数法E表示的浮点数或者复数值                          |
| %f       | 以标准记数法表示的浮点数或者复数值                           |
| %g       | 以%e或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出 |
| %G       | 以%E或者%f表示的浮点数或者复数，任何一个都以最为紧凑的方式输出 |
| %o       | 一个以八进制表示的数字(基数为8)                              |
| %p       | 以十六进制(基数为16)表示的一个值的地址，前缀为0x,字母使用小写的a-f表示 |
| %q       | 使用Go语法以及必须时使用转义，以双引号括起来的字符串或者字节切片[]byte，或者是以单引号括起来的数字 |
| %s       | 字符串。输出字符串中的字符直至字符串中的空字符（字符串以’\0‘结尾，这个’\0’即空字符） |
| %t       | 以true或者false输出的布尔值                                  |
| %T       | 使用Go语法输出的值的类型                                     |
| %U       | 一个用Unicode表示法表示的整型码点，默认值为4个数字字符       |
| %v       | 使用默认格式输出的内置或者自定义类型的值，或者是使用其类型的String()方式输出的自定义值，如果该方法存在的话 |
| %x       | 以十六进制表示的整型值(基数为十六)，数字a-f使用小写表示      |
| %X       | 以十六进制表示的整型值(基数为十六)，数字A-F使用小写表示      |

## 二、输出

```
//整型
a := 15
fmt.Printf("a = %b\n", a) //a = 1111
fmt.Printf("%%\n") //只输出一个%

//字符
ch := 'a'
fmt.Printf("ch = %c, %c\n", ch, 97) //a, a

//浮点型
f := 3.14
fmt.Printf("f = %f, %g\n", f, f) //f = 3.140000, 3.14
fmt.Printf("f type = %T\n", f) //f type = float64

//复数类型
v := complex(3.2, 12)
fmt.Printf("v = %f, %g\n", v, v) //v = (3.200000+12.000000i), (3.2+12i)
fmt.Printf("v type = %T\n", v) //v type = complex128

//布尔类型
fmt.Printf("%t, %t\n", true, false) //true, false

//字符串
str := "hello go"
fmt.Printf("str = %s\n", str) //str = hello go
```

输出就是把一个结果上面加一些格式，以各种样式展现在我们面前，比如：
`g:'a'`然后我们打印的时候给他加上一些格式，`fmt.Printf("g=%c",g)`那么输出得结果就是97，更多得格式可以查看上面得格式化说明。

## 三、输入

```
var v int
fmt.Println("请输入一个整型：")
fmt.Scanf("%d", &v)
//fmt.Scan(&v)
fmt.Println("v = ", v)
```

这里在使用`Scanf`输入赋值的时候，注意需要使用取地址符`&`，否则并不能直接修改`v`的值。

# 类型转换和别名

## 一、类型转换：

Go语言中不允许隐式转换，所有类型转换必须显式声明，而且转换只能发生在两种相互兼容的类型之间。

```
var ch byte = 97
//var a int = ch //err, cannot use ch (type byte) as type int in assignment
var a int = int(ch)
```

类型装换顾名思义就是把一个类型变成另一种类型，比如我现在手上有100块钱`a := 100`（默认是int类型）然后我们有三个人平分这100块钱，是不是`fmt.Println(a/3)`但是你这样打印的结果你会发现每人只能分到33块钱，那还有一块钱去哪里了，因为a默认是int类型后面是不会带小数点的，那么忽然少了一块钱我也不开心，所以我们要把这个a转换为浮点类型`fmt.Println(float64(a)/3)`这样每个人就都能分到33.33333了。

## 二、类型别名：

```
type bigint int64 //int64类型改名为bigint
var x bigint = 100

type (
  myint int //int改名为myint
  mystr string //string改名为mystr
)
```

也就是说我给你起个外号，例如：

我本名叫张三 （`int64`）我朋友都喜欢叫我老王，我也默认这个名称，他们叫我我也回答他们，这样和叫我张三效果是一样的，那我们怎么实现呢，`type laowang(也就是起的外号老王) int64(也就是我得本名字张三)`然后别人叫我老王也是一样得效果`var a laowang` ，然后我们打印这个a得到的结果也是`0`，类型也是`int64`。

# 算术运算符

| **运算符** | **术语**           | **示例** | **结果** |
| :--------- | :----------------- | :------- | :------- |
| +          | 加                 | 10 + 5   | 15       |
| -          | 减                 | 10 - 5   | 5        |
| *          | 乘                 | 10 * 5   | 50       |
| /          | 除                 | 10 / 5   | 2        |
| %          | 取模(取余)         | 10 % 3   | 1        |
| ++         | 后自增，没有前自增 | a=0; a++ | a=1      |
| –          | 后自减，没有前自减 | a=2; a–  | a=1      |

加减乘除相信大家都会用，我们要来说的是这个`++`以及`--`

1. `++`：代表在当前变量上加1。比如`a := 1;a++`，那么`a`会变成2。
2. `--`：跟`++`相反，`--`会给当前变量减1。

# 关系运算符

| **运算符** | **术语** | **示例** | **结果** |
| :--------- | :------- | :------- | :------- |
| ==         | 相等于   | 4 == 3   | false    |
| !=         | 不等于   | 4 != 3   | true     |
| <          | 小于     | 4 < 3    | false    |
| >          | 大于     | 4 > 3    | true     |
| < =        | 小于等于 | 4 < = 3  | false    |
| >=         | 大于等于 | 4 >= 1   | true     |

关系运算符常用于判断语句当中

## 一、==：

判断两个值是否一样，例如：

```
a:=100
b:=99,
a==b//返回值肯定是一个false
```

## 二、!=:

还是用上方密码进行案例

```
var v string //用于接收用户密码的变量
var t string=123456 //自己定义的密码
fmt.Scan(&v) //用户传入一个密码。
if(v != t){
  fmt.Println("您输入的密码错误");
}
```

## 三、< :

```
var a int =18//a 的年龄18岁
var b int =19//b 的年龄19岁
if(a<b){
  fmt.Println("a小于b")
}
```

## 五、>:

```
var a int =18//a 的年龄18岁
var b int =19//b 的年龄19岁
if(a<b){
  fmt.Println("a大于b")
}
```

## 五、<=:小于等于

```
var a int =18//a 的年龄18岁
var b int =19//b 的年龄19岁
if(a<=b){
  fmt.Println("a小于等于b")
}
```

## 六、>=：大于等于

```
var a int =18//a 的年龄18岁
var b int =19//b 的年龄19岁
if(a>=b){
  fmt.Println("a大于等于b")
}
```

# 逻辑运算符

| **运算符** | **术语** | **示例** | **结果**                                                 |
| :--------- | :------- | :------- | :------------------------------------------------------- |
| !          | 非       | !a       | 如果a为假，则!a为真；如果a为真，则!a为假。               |
| &&         | 与       | a && b   | 如果a和b都为真，则结果为真，否则为假。                   |
| \|\|       | 或       | a \|\| b | 如果a和b有一个为真，则结果为真，二者都为假时，结果为假。 |

## 一、&&:

```
//用登录作为案例：
//登录条件需要，一个用户名，还有一个用户密码
//我首先自己定义一下用户名和密码

var username int=777 //自己定义的用户名
var userpwd int=123 //自己定义的密码

var a int=777 //假设用户输入的用户名
var b int=456 //假设用户输入的密码

//我们先拿用户名作为比较，username==a,这里是不是一样的值，所以这个返回是true
//然后我再来对比密码,userpwd==b,这里的值是不一样的，所以这个返回是false
//那么我们把他连起来呢？

username==a&&userpwd==b

//其中是不是有一个true和一个false
//&&的作用就是只要这两个判断中有一个为false就会返回false
//必须两个都为true才回返回true
//也就是说你用户名虽然对了，但是你密码不对你还是登不上去
//反之只有你用户名和密码都对了才能登陆
```

## 二、||

拿身份证取票乘车作为比较，我们坐车要有车票，网上买的车票是不是要取票，取票就需要把身份证刷一下，如果你忘记带身份证了那就可以去办个临时身份证，也是一样可以取票的 例如：

```
var a int=2 //身份证等于2的原因是假设我没带身份证
var b int=1 //临时身份证

//假设取票必须等于1
var c int=1 //取票必须值一样
//首先我们a==c吗？很明显是不等于的，那么返回值就是false
//再来b==c吗？验证结果是相等的，办的临时身份证也是可以取票的，那么返回值就是true
//那么我们连起来

a==c||b==c
//"||"的作用就是只要你这两个条件之中有一个为true他就给你返回true
//就算你有一个条件是false他也是不予理睬的还是会算你通过
//也就是说，你没带身份证没关系办个临时的也能用
```

# 赋值运算符

| **运算符** | **说明**       | **示例**                              |
| :--------- | :------------- | :------------------------------------ |
| =          | 普通赋值       | c = a + b 将 a + b 表达式结果赋值给 c |
| +=         | 相加后再赋值   | c += a 等价于 c = c + a               |
| -=         | 相减后再赋值   | c -= a 等价于 c = c - a               |
| *=         | 相乘后再赋值   | c *= a 等价于 c = c * a               |
| /=         | 相除后再赋值   | c /= a 等价于 c = c / a               |
| %=         | 求余后再赋值   | c %= a 等价于 c = c % a               |
| <<=        | 左移后赋值     | c <<= 2 等价于 c = c << 2             |
| >>=        | 右移后赋值     | c >>= 2 等价于 c = c >> 2             |
| &=         | 按位与后赋值   | c &= 2 等价于 c = c & 2               |
| ^=         | 按位异或后赋值 | c ^= 2 等价于 c = c ^ 2               |
| \|=        | 按位或后赋值   | c \|= 2 等价于 c = c \| 2             |

我们就讲讲我们常用的一些运算符

## 一、=

```
var a int =1
var b int
b = a+1//这就是赋值
```

也就是说我用`a+1`赋值给了b，b中就拥有了值，这就是赋值。

## 二、+=

```
var a int 1
var b int
b += a
```

可以理解为`b=b+a`也就是`b=0+1`

## 三、-=

```
var a int 1
var b int
b -= a
```

可以理解为`b=b-a`也就是`b=0-1`。

## 四、*=

```
var a int 1
var b int
b *= a
```

可以理解为`b=b*a`也就是`b=0*1`。

## 五、/=

```
var a int 1
var b int
b /= a
```

可以理解为`b=b/a`也就是`b=0/1`。

# 运算符补充

## 一、其他运算符

| **运算符** | **术语**     | **示例** | **说明**                |
| :--------- | :----------- | :------- | :---------------------- |
| &          | 取地址运算符 | &a       | 变量a的地址             |
| *          | 取值运算符   | *a       | 指针变量a所指向内存的值 |

### 1. &

取址符，就像变量刚创建出来就会放在内存里面，我们用&就能查到它在内存中的位置例如:

```
var a int =5
fmt.printf("%p\n",&a)//就能打印出内存地址了
```

### 2. *

就是到内存中把这个变量的值给取出来。

```
var a int =5
fmt.printf("%d\n",*&a) //为什么要加&号呢？首先我们要找到a变量在内存中的地址
```

## 二、运算符优先级

在Go语言中，一元运算符拥有最高的优先级，二元运算符的运算方向均是从左至右。

下表列出了所有运算符以及它们的优先级，由上至下代表优先级由高到低：

| **优先级** | **运算符**           |
| :--------- | :------------------- |
| **7**      | **^ ! **             |
| **6**      | *** / % <<>>& &^**   |
| **5**      | **+ - \| ^**         |
| **4**      | **== != < < = >= >** |
| **3**      | **< -**              |
| **2**      | **&&**               |
| **1**      | **\|\|**             |

当然，你可以通过使用括号来临时提升某个表达式的整体运算优先级。

# 选择结构

## 一、 if语句

### 1. if

```
var a int=3
if a == 3{//条件表达式没有括号
  fmt.Println("a==3")
}
//支持一个初始化表达式,初始化字句和条件表达式直接需要用分号分隔
if b:= 3;b == 3{
  fmt.Println("b==3")
}



//拿登录密码做一个举例：
var userpwd int =123 //我自己定义的密码
var upd int =123 //假设用户自己输入的密码

if userpwd == upd {
  fmt.Println("我密码正确了!!!!!我成功了")
  //反之如果不相等就不会进入花括号里面来，也就不会运行这里面的代码
}
```

### 2. if … else

```
if a:=3; a==4{
  fmt.Println("a==4")
}else{
  //左大括号必须和条件语句或else在同一行
  fmt.Println("a!=4")
}


//拿登录密码做一个举例：
var userpwd int =123  //我自己定义的密码
var upd int =123456  //假设用户自己输入的密码

if upd==userpwd{
  //上面这一行的意思（如果用户输入的密码等于了我自己设定的密码那就让他进来）
  //那么我就可以进入到这个花括号里面做点事情
  //比如打印

  fmt.Println("我密码正确了!!!!!我成功了")
  //反之如果不相等就不会进入花括号里面来，也就不会运行这里面的代码
}else{
  //如果用户输入的密码和我定义的密码不一样，他就不会进入上面那个花括号
   //就会进入到我这里然后我就打印

  fmt.Println("就这？密码都是错的)

}
```

### 3. if … else if … else

```
if a:=3; a>3 {
  fmt.Println("a>3")
} else if a<3 {
  fmt.Println("a<3")
} else if a==3 {
  fmt.Println("a==3")
} else {
  fmt.Println("error")
}

//上面代码的意思就是，我先定义一个a：=3的变量，然后去判断，从第一个开始，
//a大于3吗？很明显不大于,然后继续往下面走
//a小于3吗？很明显也不小于,然后继续往下走
//a等于3吗？对了a等于3然后他就会输出"a==3"
//我们假设定义的变量a:=4那么上面的条件都不满足他就会进入最后一个else
//就会打印"error"
```

## 二、 switch语句

Go里面`switch`默认相当于每个`case`最后带有`break`，匹配成功后不会自动向下执行其他`case`，而是跳出整个`switch`，但是可以使用`fallthrough`强制执行后面的`case`代码：

```
var score int = 90

switch score {
  case 90:
    fmt.Println("A")
    //fallthrough
  case 80:
    fmt.Println("B")
    //fallthrough
  case 50,60,70:
    fmt.Println("C")
    //fallthrough
  default:
    fmt.Println("D")
}


//其实和if语句是差不多的道理，switch 后面就写你的条件比如上面的分数
//然后进入带着分数进入第一个case,如果分数等于90那我就打印A如果不等于
//我就进入下一条case。如果全部都没有匹配的那我们就进入default
```

可以使用任何类型或表达式作为条件语句：

```
//1
switch s1:=90 ; s1{//初始化语句;条件
  case 90:
    fmt.Println("A")
  case 80:
    fmt.Println("B")
  default:
    fmt.Println("C")
}


//这里是一样的只是在switch 后面定义了一个初始化变量,switch s1:=90 其实就是
//不用到上面定义var s1 int =90,switch s1和switch s1:=90 是一样的道理

//2
var s2 int=90
switch{ //这里没有写条件
  case s2 >= 90: //这里写判断语句
    fmt.Println("A")
  cases2 >= 80:
    fmt.Println("B")
  default:
    fmt.Println("C")
}

//这里的话就用上了条件了，，switch后面就不用写变量了,直接写花括号就行，
//我们在case 后面去判断，就和if 后面加条件一抹一样

//3
switch s3 := 90;{ //只有初始化语句，没有条件
  case s3 >= 90: //这里写判断语句
    fmt.Println("A")
  case s3 >= 80:
    fmt.Println("B")
  default:
    fmt.Println("C")
}

//这里就是结合了初始化变量和条件语句
```

## 三、 区别

1. `if-else`语句更适合于对区间（范围）的判断，而`switch`语句更适合于对离散值的判断。
2. `switch`语句只支持常量值相等的分支判断，而`if`语句支持更为灵活，任意布尔表达式均可。
3. `switch`语句通常比一系列嵌套`if`语句效率更高；逻辑更加清晰。

# 循环语句

## 一、for

我们为什么要使用循环语句呢？如果我要大家打印一个"今天我吃饭了"那是不是`fmt.println("今天我吃饭了")`就行了，但是我要大家打印100次呢，那要复制100次，循环语句就是重复执行某个事情，直到执行到了你相对应的那个条件就会结束。如下：

```
var i,sum int
for i = 1; i <= 100; i++ {
  sum+=i
}
fmt.Println("sum=",sum)
```

## 二、range

关键字 range 会返回两个值，第一个返回值是元素的数组下标，第二个返回值是元素的值：

```
s := "abc"
for i := range s{//支持string/array/slice/map。
  fmt.Printf("%c\n",s[i])
}

for _, c := range s{//忽略index
  fmt.Printf("%c\n",c)
}

for i, c := range s{
  fmt.Printf("%d,%c\n",i,c)
}
```

`for i,c := range s`这里面的`i`,`c`是自己定义的名字，你可以随便起，`i`是下标，`c`代表循环到的值。

## 三、跳出循环（break和continue）：

在循环里面有两个关键操作break和continue，break操作是跳出当前循环，continue是跳过本次循环。

```
for i := 0; i < 5; i++ {
if 2 == i {
  //break操作是跳出当前循环
  continue //continue是跳过本次循环
}
  fmt.Println(i)
}
```

`break`是跳出整个循环，剩下还没执行的循环也一并退出。

`continue`跳出当前循环，会跳过当前还没有执行完的剩余代码。但是剩余的循环还是会执行。

# 函数基本使用

## 一、定义格式

函数构成代码执行的逻辑结构。在Go语言中，函数的基本组成为：关键字`func`、函数名、参数列表、返回值、函数体和返回语句。

Go 语言函数定义格式如下：

```
func FuncName(/*参数列表*/) (o1type1,o2type2/*返回类型*/) {
  //函数体
  return v1, v2//返回多个值
}
```

使用函数主要是为了方便我们以后一些大量代码的调用。

`func FuncName`就是定义一个函数给他起个名字，`func FuncName (a int, b int)`在这个里面可以传两个int值，`func FuncName (a int, b int)(va int,vb int){}`中的`va`和`vb`是函数返回值。

## 二、无参无返回值

```
func Test() {//无参无返回值函数定义
  fmt.Println("this is a test func")
}

func main() {
  Test() //无参无返回值函数调用
}
```

这个是最简单的函数了，只要在`main`里面调用就行，`Test()`然后就会打印,`this is a test func`

## 二、 有参无返回值

### 1. 普通参数列表

```
func Test01(v1int,v2int) {//方式1
  fmt.Printf("v1=%d,v2=%d\n",v1,v2)
}
func Test02(v1,v2int) {//方式2, v1, v2都是int类型
  fmt.Printf("v1=%d,v2=%d\n",v1,v2)
}

func main() {
  Test01(10,20) //函数调用
  Test02(11,22) //函数调用
}
```

意思就是我`Test01`明确表明了我的参数是两个`int`类型，`Test02`里面的`v1`就没有明确表明是什么类型，其实都是一样的，然后在`main`里面调用展示一下。

### 2. 不定参数列表

#### 2.1 不定参数类型

不定参数是指函数传入的参数个数为不定数量。为了做到这点，首先需要将函数定义为接受不定参数类型：

```
//形如...type格式的类型只能作为函数的参数类型存在，并且必须是最后一个参数
func Test(args...int) {
  for_,n := range args{//遍历参数列表
    fmt.Println(n)
  }
}
func main() {
  //函数调用，可传0到多个参数
  Test()
  Test(1)
  Test(1,2,3)
}
```

`Test(args...int)`代表不确定有多少个int类型的参数要放到里面。

#### 2.2 不定参数的传递

```
func MyFunc01(args...int) {
  fmt.Println("MyFunc01")
  for_,n := range args{//遍历参数列表
    fmt.Println(n)
  }
}

func MyFunc02(args...int) {
  fmt.Println("MyFunc02")
  for_,n:=range args{//遍历参数列表
    fmt.Println(n)
  }
}

func Test(args...int) {
  MyFunc01(args...)//按原样传递,Test()的参数原封不动传递给MyFunc01
  MyFunc02(args[1:]...)//Test()参数列表中，第1个参数及以后的参数传递给MyFunc02
}

func main() {
  Test(1,2,3)//函数调用
}
```

`Test`里面调用`MyFunc01`和`MyFunc02`，并且把所有参数都传给了`MyFunc01`，把除了第0个参数外的其他参数传给了`MyFunc02`。

## 三、 无参有返回值

有返回值的函数，必须有明确的终止语句，否则会引发编译错误。

### 1. 一个返回值

```
//方式1
func Test01() int { 
  return 250
}
// 这个int就是定义的返回值类型，也就是说你的返回值只能是int类型，不能回其他类型
// 官方建议：最好命名返回值`func Test01()(a int){ a=250 return a}`
// 因为不命名返回值，虽然使得代码更加简洁了，但是会造成生成的文档可读性差

//方式2,给返回值命名
func Test02() (value int) { 
  value=250
  return value
}

//方式3,给返回值命名
func Test03() (value int) { 
  value = 250
  return
}

func main(){
  v1 := Test01()//函数调用
  v2 := Test02()//函数调用
  v3 := Test03()//函数调用
  fmt.Printf("v1=%d,v2=%d,v3=%d\n", v1, v2, v3)
}
```

### 2. 多个返回值

```
//方式1
func Test01() (int,string) {
  return 88,"zhiliao"
}

//方式2,给返回值命名
func Test02() (a int,str string) {
  a = 88
  str = "zhiliao"
  return a,str
}

func main(){
  v1,v2 := Test01() // 函数调用
  _,v3 := Test02() // 函数调用，第一个返回值丢弃
  v4,_ := Test02() // 函数调用，第二个返回值丢弃
  fmt.Printf("v1=%d,v2=%s,v3=%s,v4=%d\n",v1,v2,v3,v4)
}
```

## 四、有参有返回值

```
//求2个数的最小值和最大值
func MinAndMax(num1 int, num2 int) (min int, max int) {
  if num1 > num2{//如果num1大于num2
    min = num2
    max = num1
  } else {
    max = num2
    min = num1
  }
  return
}
func main() {
  min,max := MinAndMax(33,22)
  fmt.Printf("min=%d,max=%d\n", min, max) //min=22,max=33
}
```

用两个参数和两个返回值，在我们掉用方法的时候就会传两个参数`32`，`22`然后我们进去判断一下如果`32`比`22`大那么我们就把`22`赋值给返回值`min`，然后把`32`赋值给返回值`max`，然后`return`出去就好了。

# 递归函数

递归指函数可以直接或间接的调用自身。

递归函数通常有相同的结构：**一个跳出条件和一个递归体**。所谓**跳出条件**就是根据传入的参数判断是否需要停止递归，而**递归体**则是函数自身所做的一些处理。

```
// 通过循环实现1+2+3……+100
func Test01() int {
    i := 1
    sum := 0
    for i = 1; i <= 100; i++ {
        sum += i
    }
    return sum
}

// 通过递归实现1+2+3……+100
func Test02(num int) int {
    if num == 1 {
        return1
    }
    return num + Test02(num-1) // 函数调用本身
}
// 通过递归实现1+2+3……+100
func Test03(num int) int {
    if num == 100 {
        return 100
    }
    return num + Test03(num+1)  // 函数调用本身
}

func main() {
    fmt.Println(Test01())    // 5050
    fmt.Println(Test02(100))    // 5050
    fmt.Println(Test03(1))    // 5050
}
```

# 函数类型

在`Go`语言中，函数也是一种数据类型，我们可以通过`type`来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型。

```
type FuncType func(int,int) int //声明一个函数类型,func后面没有函数名

// 函数中有一个参数类型为函数类型：fFuncType
func Calc(a, b int,f FuncType) (result int) {
    result = f(a, b) //通过调用f()实现任务
    return
}

func Add(a, b int) int {
    return a + b
}

func Minus(a, b int) int {
    return a - b
}

func main() {
    //函数调用，第三个参数为函数名字，此函数的参数，返回值必须和FuncType类型一致
    result := Calc(1, 1, Add)
    fmt.Println(result)    //2

    var f FuncType = Minus
    fmt.Println("result=",f(10, 2))    //result=8
}
```

`type FuncType func(int, int)int`这里就是声明一个函数类型，并且返回值为`int`类型。后面就可以把`FuncType`当做一个变量类型来使用了。

# 匿名函数与闭包

所谓闭包就是一个函数“捕获”了和它在同一作用域的其它常量和变量。这就意味着当闭包被调用的时候，不管在程序什么地方调用，闭包能够使用这些常量或者变量。它不关心这些捕获了的变量和常量是否已经超出了作用域，所以只有闭包还在使用它，这些变量就还会存在。

在Go语言里，所有的匿名函数(Go语言规范中称之为函数字面量)都是闭包。匿名函数是指不需要定义函数名的一种函数实现方式，它并不是一个新概念，最早可以回溯到1958年的Lisp语言。

```
func main() {
    i := 0
    str := "mike"

    // 方式1
    f1 := func(){ //匿名函数，无参无返回值
        //引用到函数外的变量
        fmt.Printf("方式1：i=%d,str=%s\n",i,str)
    }

    f1() // 函数调用

    // 方式1的另一种方式
    type FuncType func()    // 声明函数类型,无参无返回值
    var f2 FuncType = f1
    f2()    // 函数调用

    //方式2
    var f3 FuncType = func(){
        fmt.Printf("方式2：i=%d,str=%s\n",i,str)
    }
    f3() // 函数调用

    //方式3
    func(){  // 匿名函数，无参无返回值
        fmt.Printf("方式3：i=%d,str=%s\n",i,str)
	}()    // 别忘了后面的(),()的作用是，此处直接调用此匿名函数

    //方式4,匿名函数，有参有返回值
    v := func(a,b int) (result int) {
        result=a+b
        return
	}(1,1) // 别忘了后面的(1,1),(1,1)的作用是，此处直接调用此匿名函数，并传参
    fmt.Println("v=",v)
}
```

## 1. 闭包捕获外部变量特点：

```
func main() {
    i := 10
    str := "mike"
    func(){
        i = 100
        str = "go"
        // 内部：i=100,str=go
        fmt.Printf("内部：i=%d,str=%s\n",i,str)
    }()    // 别忘了后面的(),()的作用是，此处直接调用此匿名函数

    // 外部：i=100,str=go
    fmt.Printf("外部：i=%d,str=%s\n",i,str)
}
```

## 2. 函数返回值为匿名函数：

```
//squares返回一个匿名函数，func()int
//该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
    var x int
    return func() int { // 匿名函数
        x++ // 捕获外部变量
        return x*x
    }
}

func main() {
    f:=squares()
    fmt.Println(f())    //"1"
    fmt.Println(f())    //"4"
    fmt.Println(f())    //"9"
    fmt.Println(f())    //"16"
}
```

函数`squares`返回另一个类型为`func() int`的函数。对`squares`的一次调用会生成一个局部变量`x`并返回一个匿名函数。每次调用时匿名函数时，该函数都会先使`x`的值加1，再返回`x`的平方。第二次调用`squares`时，会生成第二个`x`变量，并返回一个新的匿名函数。新匿名函数操作的是第二个`x`变量。

通过这个例子，我们看到变量的生命周期不由它的作用域决定：`squares`返回后，变量`x`仍然隐式的存在于`f`中。

# 函数作用域

## 一、局部变量

前面我们定义的函数中，都经常使用变量。那么我们看一下如下程序的输出结果：

```
func Test() {
    a := 5
    a += 1
}

func main() {
    a := 9
    Test()
    fmt.Println(a)
}
```

最终的输出结果是9，为什么呢？在执行`fmt.Println(a)`语句之前，我们已经调用了函数`Test()`，并在该函数中我们已经重新给变量`a`赋值了。但是为什么结果没有发生变化呢？这就是变量的作用范围（作用域）的问题。在`Test()`函数中定义的变量`a`，它的作用范围只在改函数中有效，当`Test()`函数执行完成后，在该函数中定义的变量也就无效了。也就是说，当`Test()`函数执行完以后，定义在改函数中所有的变量，所占有的内存空间都会被回收。

所以，我们把定义在函数内部的变量称为局部变量。

局部变量的作用，为了临时保存数据需要在函数中定义变量来进行存储，这就是它的作用。

并且，通过上面的案例我们发现：不同的函数，可以定义相同的名字的局部变量，但是各用个的不会产生影响。例如：我们在`main()`函数中定义变量`a`,在`Test()`函数中也定义了变量`a`，但是两者之间互不影响，就是因为它们属于不同的函数，作用范围不一样，在内存中是两个存储区域。

## 二、全局变量

有局部变量，那么就有全局变量。所谓的全局变量：既能在一个函数中使用，也能在其他的函数中使用，这样的变量就是全局变量.也就是定义在函数外部的变量就是全局变量。全局变量在任何的地方都可以使用。案例如下：

```
var a int   //变量定义在函数外面
func Test() {
    a := 5
    a += 1
}

func main() {
    a := 9
    Test()
    fmt.Println(a)
}
```

注意：在上面的案例中，我们在函数外面定义了变量`a`,那么该变量就是全局变量，并且`Test()`函数和`main()`函数都可以使用该变量。该程序的执行流程是：先执行`main()`函数，给变量`a`赋值为9,紧接着调用`Test()`函数，在改函数中完成对变量`a`的修改。

由于`main()`函数与`Test()`函数所使用的变量`a`是同一个，所以当`Test( )函数执行完成后，变量的a已经变成了6.回到main( )函数执行后面的代码，也就是 fmt.Println(a),输出的值就是6.

可能有同学已经发现该程序和我们前面写的程序还有一点不同的地方是：第一个程序我们是a:=9，但是第二个程序执行修改成了a=9，现在修改一下第二个程序如下：

```
var a int   //变量定义在函数外面
func Test() {
    a := 5
    a += 1
}

func main() {
    a := 9    // ====》注意这里
    Test()
    fmt.Println(a)
}
```

该程序与上面的程序不同之处在于，该程序是`a:=9`，上面的程序是`a=9`。

现在大家思考一下该程序的结果是多少？最终结果是9。原因是`a:=9`等价于：

```
var a int
a=9
```

也就是定义一个整型变量`a`,并且赋值为`9`。

那么现在的问题是，我们定义了一个全局变量a,同时在`main()`中又定义了一个变量也叫`a`，但是该变量是一个局部变量。

当全局变量与局部变量名称一致时，局部变量的优先级要高于全局变量。所以在`main()`函数中执行`fmt.Println(a)`时输出的是局部变量`a`的值。但是`Test()`函数中的变量`a`还是全局变量。

注意：大家以后在开发中，尽量不要让全局变量的名字与局部变量的名字一样。

所以大家，思考以下程序执行的结果：

```
var a int   //变量定义在函数外面
func Test() {
    a := 5
    a += 1
    fmt.Println("Test",a)

}

func main() {
    a := 9   
    Test()
    fmt.Println("main",a)
}
```

## 三、总结

- 在函数外边定义的变量叫做全局变量。
- 全局变量能够在所有的函数中进行访问。
- 如果全局变量的名字和局部变量的名字相同，那么使用的是局部变量的，小技巧强龙不压地头蛇。

# 工作区介绍

通过前面函数的学习，我们能够体会到函数的优势，就是可以将不同的功能放在不同的函数中实现，主函数（`main()`）可以直接调用。这样结构非常的清晰，也非常方面代码的管理。如果我们把所有的代码都写在`main()`函数中，会出现什么样的情况呢？

代码混乱，非常不容易管理。但是现在我们面临了另外一个问题就是：我们所有自己定义的函数都写在了一个文件中。

如果我们做的项目代码量越来越多，那么该文件会变的非常臃肿，代码也会变得非常难管理。所以，我们在开发中，除了要定义函数，同时还要将代码放在不同的文件中。例如：我们定义了一个`UserInfo.go`文件，里面包含了用户的添加函数，修改函数，删除函数等操作。

这就涉及到项目的工程管理也就是怎样对项目中的文件进行管理。

为了更好的管理项目中的文件，要求将文件都要放在相应的文件夹中。GO语言规定如下的文件夹如下：

- `src`目录：用于以代码包的形式组织并保存Go源码文件。（比如：`.go`、`.c`、`.h`、`.s`等）。
- `pkg`目录：用于存放经由`go install`命令构建安装后的代码包（包含Go库源码文件）的“.a”归档文件。
- `bin`目录：与pkg目录类似，在通过`go install`命令完成安装后，保存由Go命令源码文件生成的可执行文件。

以上目录称为工作区，工作区其实就是一个对应于特定工程的目录。

目录`src`用于包含所有的源代码，是Go命令行工具一个强制的规则，而`pkg`和`bin`则无需手动创建，如果必要Go命令行工具在构建过程中会自动创建这些目录。

## 标准命令概述

Go语言中包含了大量用于处理Go语言代码的命令和工具。其中，go命令就是最常用的一个，它有许多子命令。这些子命令都拥有不同的功能，如下所示。

- build：用于编译给定的代码包或Go语言源码文件及其依赖包。
- clean：用于清除执行其他go命令后遗留的目录和文件。
- doc：用于执行`godoc`命令以打印指定代码包。
- env：用于打印Go语言环境信息。
- fix：用于执行`go tool fix`命令以修正给定代码包的源码文件中包含的过时语法和代码调用。
- fmt：用于执行`gofmt`命令以格式化给定代码包中的源码文件。
- get：用于下载和安装给定代码包及其依赖包(提前安装git或hg)。
- list：用于显示给定代码包的信息。
- run：用于编译并运行给定的命令源码文件。
- install：编译包文件并编译整个程序。
- test：用于测试给定的代码包。
- tool：用于运行Go语言的特殊工具。
- version：用于显示当前安装的Go语言的版本信息。

# 创建同级目录

## 一、 创建src目录，在该目录下创建go源码文件

### 1. 在项目文件夹下新建src目录，如下图所示：

![import01.png](https://www.zlkt.net/media/course/wxM9eiQMBq9cokuCD2ursU.png)

我这里是在D盘的Workspace目录下创建的src目录。

### 2. 在`src`目录下创建不同的go源码文件，如下图所示：

![import02.png](https://www.zlkt.net/media/course/etHYnpUr37KBLnZd8ZjJ36.png)

然后在`src`目录下创建`main.go`文件和`test.go`文件（注意：这个两个文件是在同一个目录下面，都是在src目录下面）。

`main.go`文件下的代码如下所示：

```
package mian
import "fmt"

func main () {
  fmt.Println("main")
}
```

`test.go`文件下的代码如下所示：

```
package main //必须与main.go必须是一个包
import "fmt"

func Test () {
  fmt.Println("Test")
}
```

这也是一个简单的打印语句。

我们现在已经完成两个文件代码的编写，接下来的问题是，我们怎样在`main.go`文件中的入口函数`main()`中调用`test.go`文件中的`Test()`函数呢？这就需要设置环境变量`GOPATH`属性。如果要实现不同文件中函数的调用，必须设置`GOPATH`，否则，即使文件处于同一工作目录（工作区）下，也是无法完成调用的。

## 二、`GOPATH`设置

`GOPATH`设置的具体步骤如下：
![import50.png](https://www.zlkt.net/media/course/YYU7vdkkotMSu8qnp3THzZ.png)

![import51.png](https://www.zlkt.net/media/course/FakArahaKXsQWCZAx4F43Y.png)

最后再配置完成后，可以测试一下是否配置成功。

![import52.png](https://www.zlkt.net/media/course/xCKYe6p9YYnLgPvrGcwYzK.png)

## 三、 在`main.go`文件中完成对`test.go`文件中函数的调用

![import05.png](https://www.zlkt.net/media/course/J39PBLi9dxE4YEav3XsHu8.png)

最后编译执行。
注意：同一个目录下不能定义不同的`package`。

# 创建不同级目录

在上一小节中，将不同的go源代码文件都放在了同一个目录下面，如果将go源代码文件放在不同的目录下面应该怎样进行处理呢？

## 一、步骤：

### 1. 新建项目目录:

如下图所示：
![import54.png](https://www.zlkt.net/media/course/c9HMdKNV6vPjoXKdPMaXnW.png)

在`CmsProject`目录下面，创建`src`目录，在`src`目录下面创建如下目录与文件。
![import55.png](https://www.zlkt.net/media/course/VMDwAWhYCQN8ndega8nx8V.png)

`main.go`定义的是入口函数`main()`。

`userinfo`文件夹下定义的是`user.go`文件。
![import56.png](https://www.zlkt.net/media/course/kUPmueyqnAispUcVc4m8GV.png)

`user.go`文件中的代码如下：

```
//不同目录 包名不一样
package userinfo 
import "fmt"

func Add(){
  fmt.Println("添加用户信息")
}
```

`main.go`文件中的代码如下：

```
package main

import "fmt"
import "userinfo" //注意导包

func main(){
  fmt.Println("main")
  userinfo.Add() //通过包名.函数进行调用
}
```

导入用户的包，然后通过用户调用添加的方法，通过以上两个文件中的代码，可以总结出如下几点:

1. 不同目录，包名不一致（自定义包）。
2. `main.go`中调用`user.go`中的方法时，一定要导包，并且调用的方式是：包名.函数名 的方式。

### 2. `GoLand`中设置`GOPATH`：

![import59.png](https://www.zlkt.net/media/course/vNxVH7FJaXBKxZoP6W6Tdj.png)

会打开如下的窗口，然后进行设置。
![img](https://www.zlkt.net/assets/import57.png)
![import57.png](https://www.zlkt.net/media/course/rH6PWUDDRhQ5QNzDgA4zY9.png)
注意：user.go文件中的函数名首字母必须大写，如果改成小写在main.go中无法进行调用

### 3. 再添加其他文件和函数：

这种不同级目录应用，在以后的项目开发中使用频率非常高。例如：上面我们的案例中，可以将用户管理的操作放在`userinfo`目录下，商品管理模块可以再定义一个目录，例如：`product.go`

如下图所示：
![import60.png](https://www.zlkt.net/media/course/wY82jQN5FGm9XpdWQxMmNR.png)

`product.go`中的代码如下：
![import61.png](https://www.zlkt.net/media/course/EvH2Zd6jSpCLbEToiL75hN.png)

`main.go`中的代码如下：
![import62.png](https://www.zlkt.net/media/course/bJ4CTpsGs3rK76ioxofEBD.png)

## 二、关于包的问题

包就是一个标识，标志着代码是来自哪儿，对代码进行管理。

所以，在`main()`函数中要使用相应的函数，必须进行导包，然后根据包名去调用相应的函数。

通过上面的代码，我们也能够体会出“包”的优势，就是可以在`userinfo`包中定义名叫`Add()`方法，在`product`包中也可以定义`Add()`方法，但是在`main()`函数中进行调用时，通过包名进行调用，就可以很清楚`Add()`方法来自哪个包，不会造成混乱，和名称的冲突。并且相关的功能代码，放在一个包中，可以很好的被复用。例如：可以在`userinfo`包中使用`product`，如下图所示：

![import63.png](https://www.zlkt.net/media/course/TnQBgnyRtHBTFsPKx6mgQR.png)

但是我们创建的的自定义包最好放在`GOPATH`的`src`目录下，在Go语言中，代码包中的源码文件名可以是任意的。但是，这些任意名称的源码文件都必须以包声明语句作为文件中的第一行，每个包都对应一个独立的名字空间。

包中成员以名称⾸字母⼤⼩写决定访问权限：

- public: 首字母大写，可以被包外访问。
- private: 首字母小写，仅可以被包内访问。

注意：同一个目录下不能定义不同的`package`。

## 三、导包的问题

在上面的案例中，要使用包，必须要进行导入，可以通过关键字进行`import`进行导入，它会告诉编译器你想引用该包内的代码。

如果导入的是标准库中的包（GO语言自带，例如：”fmt”包）会在安装Go的位置找到。Go开发者创建的包会在`GOPATH`环境变量指定的目录里查找。所以，`import`关键字的作用就是查找包所在的位置。

如果编译器查遍`GOPATH`也没有找到要导入的包，那么在试图对程序执行`run`或者`build`的时候就会出错。

**注意：如果导入包之后，未调用其中的函数或者类型将会报出编译错误。**
![import64.png](https://www.zlkt.net/media/course/q5QLD2AGPMfEtJMeUHo7uD.png)

我们常规的导包方式是用`import`关键字一个个导入。(Goland会自动帮我们导入包)

例如：
![import65.png](https://www.zlkt.net/media/course/ajMXpnN9N7ng43mmxaGDfg.png)

表示导入三个包，有GO语言自带的包，也有我们自定义的包。但是，这种写法可能比较麻烦，所以为了简化也可以采用如下的方式进行导包：
![import67.png](https://www.zlkt.net/media/course/Nme5RyuvxDZycGQh6eciTW.png)

这种方式，使用的频率是非常高的。

# 数组

## 一、复合类型：

| **类型** | **名称** | **长度** | **默认值** | **说明** |
| :------- | :------- | :------- | :--------- | :------- |
| pointer  | 指针     |          | nil        |          |
| array    | 数组     |          | 0          |          |
| slice    | 切片     |          | nil        | 引⽤类型 |
| map      | 字典     |          | nil        | 引⽤类型 |
| struct   | 结构体   |          |            |          |

## 二、数组

如果要存储班级里所有学生的数学成绩，应该怎样存储呢？可能有同学说，通过定义变量来存储。但是，问题是班级有80个学生，那么要定义80个变量吗？

像以上情况，最好是通过数组的方式来存储。

所谓的数组：是指一系列同一类型数据的集合。

### 1. 数组定义

```
var a [10]int
```

数组定义也是通过`var`关键字，后面是数组的名字`a`，长度是10,类型是整型。表示：数组`a`能够存储10个整型数字。也就是说，数组a的长度是10。

我们可以通过`len()`函数测试数组的长度，如下所示：

```
var a [10]int
fmt.Println(len(a))
```

输出结果为10。

当定义完成数组a后，就在内存中开辟了10个连续的存储空间，每个数据都存储在相应的空间内，数组中包含的每个数据被称为数组元素（element），一个数组包含的元素个数被称为数组的长度。

注意：数组的长度只能是常量。以下定义是错误的：

```
var n int = 10
var a [n]int
```

会报错。

### 2. 数组赋值

数组定义完成后，可以对数组进行赋值操作。数组是通过下标来进行操作的，下标的范围是从0开始到数组长度减1的位置。
![import.png](https://www.zlkt.net/media/course/u6PAhYjQABPUH7SNMgis5A.png)

`var a[10] int`表示的范围是`a[0],a[1],a[2].......,a[9]`。

#### 2.1 第一种赋值方法：

![img](https://www.zlkt.net/assets/import1.png)
![import1.png](https://www.zlkt.net/media/course/DHzxxfQCqjkHncpSb6Uq5i.png)

#### 2.2 第二种赋值方法

第一种赋值方式比较麻烦，可以使用如下所示：

```
var a [10]int
for i := 0; i < 10; i++ {
  a[i] = i + 1
}
for i := 0; i < 10; i++ {
  fmt.Println(a[i])
}
```

第一次循环i等于1，然后赋值给`a[i]`也就是`a[1]`，然后`a[1]`也刚好等于1，然后循环十次。

#### 2.3 使用`len()`函数方式赋值：

对上面的程序，进行如下的修改：

```
var a [10]int
for i := 0; i < len(a); i++ {
  a[i] = i + 1
}
for i := 0; i < len(a); i++ {
  fmt.Println(a[i])
}
```

同上一样的道理，虽然改成了`len(a)`但是其实他的值就是10，而且这样更加方便,你也不用刻意去知道他的长度是多少。

### 3. 数组数据输出：

可以使用`range`。如下所示：

```
for i,data := range a {
  //fmt.Println("下标：",i)
  fmt.Println("元素值：",data)
}
```

`i`变量存储的是数组的下标，`data`变量存储的是数组中的值。

#### 3.1. 输出值，不输出下标：

如果只想输出数组中的元素值，不希望输出下标，可以使用匿名变量：

```
for _,data := range a {
  //fmt.Println("下标：",i)
  fmt.Println("元素值：",data)
}
```

上面的案例中，首先完成了数组的赋值，然后再输出数组中的值。但是，如果定义完成数组后，没有赋值，直接输出会出现什么样的问题呢？

```
var a [10]int
for i := 0; i < len(a); i++ {
  fmt.Println(a[i])
}
```

`a`数组中的元素类型是整型，定义完成后，直接输出，结果全部是0。

#### 3.2 不同类型数组的默认值：

```
var a [10]float64 // 如果不赋值，直接输出，结果默认全部是0
var a [10]string // 如果不赋值，直接输出，结果默认全部是空字符
var a [10]bool // 如果不赋值，直接输出，结果默认全部是false.
```

### 4. 数组初始化

之前我们是定义数组，然后再完成数组的赋值。其实，在定义数组时，也可以完成赋值，这种情况叫做数组的初始化。
具体案例如下：

```go
// 数组初始化
// 1.全部初始化
var a [5]int = [5]int{1,2,3,4,5}
fmt.Println(a)
// 在定义好一个数组的同时就给他赋值

// 自动推导
b := [5]int{1,2,3,4,5}
fmt.Println(b)

// 部分初始化
// 没有初始化的部分 默认为0
c := [5]int{1,2,3}
fmt.Println(c)
// 12300
// 指定某个元素初始化

d := [5]int{2:10,4:20}
fmt.Println(d)
// 其中2的值是10，4的值是20，其他的都为0
// ... 通过初始化确定长度
f:=[...]int{1,2,3}
fmt.Println(len(f))
// 3
```

# 切片

## 一、切片（slice）概念

在讲解切片（slice）之前，大家思考一下数组有什么问题？

1. 数组定义完，长度是固定的。例如：

   ```
   var num [5]int = [5]int{1,2,3,4,5}
   ```

   定义的`num`数组长度是5，表示只能存储5个整型数字，现在向数组`num`中追加一个数字，这时会出错。因为你已经定义死了。

2. 使用数组作为函数参数进行传递时，如果实参为5个元素的整型数组，那么形参也必须5个元素的整型数组，否则出错。

针对以上两个问题，可以使用切片来进行解决。

**切片与数组的区别：** 切片与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大，所以可以将切片理解成“动态数组”，但是，它不是数组。

## 二、 切片与数组区别

通过定义，来比较一下切片与数组的区别。

1. 先回顾数组的基本定义初始化：`a := [5]int{}`，数组中`[]`是一个固定的数字，表示长度。定义完后，长度是固定，最多存储5个数字。

2. 切片的基本定义初始化如下：`s:=[]int{}//定义空切片`看定义的方式，发现与数组很相似.但是注意：切片中的`[]`是空的，或者是“…”。切片的长度和容量可以不固定。现在通过程序演示，动态向切片中追加数据

   ```
   // 初始化切片
   s := []int{1,2,3}
   // 通过append函数向切片中追加数据
   s = append(s,5,6,7)
   fmt.Println(s)
   ```

   `append()`函数，第一个参数表示向哪个切片追加数据，后面表示具体追加的数据。最终输出结果为：

   ```
   [1 2 3 5 6 7]
   ```

## 三、 切片其它定义方式

### 1. 定义空切片：

```
//声明切片和声明数组一样，只是少了长度，此为空(nil)切片
var s1 []int 
```

### 2. 通过`make()`函数实现

```
//借助make函数, 格式 make\(切片类型, 长度, 容量\)
s := make([]int, 5, 10)
```

## 四、切片的长度与容量

长度是已经初始化的空间（以上切片s初始空间默认值都是0）。容量是已经开辟的空间，包括已经初始化的空间和空闲的空间。我们可以通过如下图来理解切片的长度与容量：
![import14.png](https://www.zlkt.net/media/course/KcRbSvCnSiVhz9QraCEeM5.png)

该切片的长度是5（存有数据，注意如果没有赋值，默认值都是0），容量是10，只不过有5个空闲区域。即使没有给切片s赋值，初始化的空间（长度）默认存储的数据都是0。

演示如下：

```
s := make([]int,5,8)
fmt.Println(s)
```

输出的结果是：

```
[0 0 0 0 0]
```

在使用`make()`函数定义切片时，一定要注意，切片长度要小于容量，例如：

```
// 以下是错误的
s := make([]int, 10, 5)
```

`make()`函数中的容量参数是可以省略掉的，如：

```
s := make([]int,10)
```

这时长度与容量是相等的，都是10.

GO语言提供了相应的函数来计算切片的长度与容量，示例如下：

```
s := make([]int,5,10)
fmt.Println("长度是",len(s))
fmt.Println("容量是",cap(s))
```

接下来给切片s赋值，可以通过下标的方式直接来进行赋值。如下所示：

```
s := make([]int,5,10)
s[0] = 1
s[1] = 2
```

也可以通过循环的方式来进行赋值。

```
s := make([]int,5,10)
for i:=0;i<len(s) ;i++ {
  s[i] = i
}
```

在这里一定要注意，循环结束条件是小于切片的长度，而不是容量。因为，切片的长度是指的是初始化的空间。以下方式会出现异常错误。

```
for i:=0;i<cap(s) ;i++ {
  s[i] = i
}
```

给切片赋完值后，怎样将切片中的数据打印出来呢？

1. 第一种方式：直接通过下标的方式输出，例如：`s[0],s[1].....`。

2. 第二种方式: 通过循环的方式，注意循环结束的条件，也是小于切片的长度，如下所示：

   ```
   for i:=0;i<len(s) ;i++ {
     fmt.Println(s[i])
   }
   ```

   或者使用range方式输出：

   ```
   for _,v := range s {
     fmt.Println(v)
   }
   ```

## 四、 切片截取

上一小节中，已经完成了切片的定义，赋值等操作，接下来看一下关于切片的其它操作。首先说一下切片的截取操作，所谓截取就是从切片中获取指定的数据。

我们通过如下程序给大家解释一下：

```
//定义切片 并且完成初始化
s := []int{10,20,30,0,0}

//从切片s中截取数据
slice := s[0:3:5]
fmt.Println(slice)
```

以上程序输出结果：

```
[10 20 30]
```

其中`s[0:3:5]`是什么意思呢？我们来解释一下。每个位置的数字为`s[low:high:max]`：

1. 第一个数`low`表示下标的起点（从该位置开始截取），如果`low`取值为0表示从第一个元素开始截取，也就是对应的切片`s`中的10。
2. 第二个数`high`表示取到哪结束，也就是下标的终点（不包含该位置），3表示取出下标是`0,1,2`的数据`（10,20,30）`，不包括下标为3的数据，那么也就是说取出的数据长度是3。可以根据公式：3-0计算`(len=high-low)`，也就是第二个数减去第一个数，差就是数据长度。在这里可以将长度理解成取出的数据的个数。
3. 第三个数用来计算容量，所谓容量：是指切片目前可容纳的最多元素个数。通过公式5-0计算(`cap=max-low`)，也就是第三个数据减去第一个数。该案例中容量为5。

现在将以上程序进行修改：

```
//定义切片 并且完成初始化
s := []int{10,20,30,40,50}

//从切片s中截取数据
slice := s[0:3:5]
fmt.Println(slice)
```

结果是：

```
[10 20 30]
```

因为起点还是0，也就是10开始，终点还是3也就是到30结束.长度是3，容量是5。

继续修改该程序：

```
//定义切片 并且完成初始化
s := []int{10,20,30,40,50}

//从切片s中截取数据
slice := s[0:4:5]
fmt.Println(slice)
```

结果是：

```
[10 20 30 40]
```

因为起点还是0，也就是10开始，终点还是4也就是到40结束。长度是4，容量是5。
继续修改该程序

```
//定义切片 并且完成初始化
s := []int{10,20,30,40,50}

//从切片s中截取数据
slice := s[1:4:5]
fmt.Println(slice)
```

slice切片结果是：

```
[20 30 40]
```

那么容量是多少呢？容量为4，通过第三个数减去第一个数（5-1）计算。

通过画图的方式来表示slice切片中的容量。

![import15.png](https://www.zlkt.net/media/course/gS5JPLCuz3iuCeGD5P6YJQ.png)

通过上面的图，可以发现切片s经过截取操作以后，将结果赋值给切片`slice`后，长度是3，容量是4，只不过有一块区域是空闲的。

## 切片其他操作。

如下表所示：

| **操作**            | **含义**                                                     |
| :------------------ | :----------------------------------------------------------- |
| **s[n]**            | 切片s中索引位置为n的项                                       |
| **s[:]**            | 从切片s的索引位置0到len(s)-1处所获得的切片                   |
| **s[low:]**         | 从切片s的索引位置low到len(s)-1处所获得的切片                 |
| **s[:high]**        | 从切片s的索引位置0到high处所获得的切片，len=high             |
| **s[low:high]**     | 从切片s的索引位置low到high处所获得的切片，len=high-low       |
| **s[low:high:max]** | 从切片s的索引位置low到high处所获得的切片，len=high-low，cap=max-low |
| **len(s)**          | 切片s的长度，总是<=cap(s)                                    |
| **cap(s)**          | 切片s的容量，总是>=len(s)                                    |

下面通过一个案例，演示一下。

### 1. s[:]：

![import16.png](https://www.zlkt.net/media/course/n39cBwkcBvqF4393fCRRGK.png)
结果是所有的值。

### 2. s[low:]

![import17.png](https://www.zlkt.net/media/course/QF8yFu4AfeBXigoZEhLUca.png)
结果是下标3后面的所有值。

### 3. s[:high]

![import18.png](https://www.zlkt.net/media/course/r8SejaBRcjALs27mYYXkkL.png)
结果是6前面的值。

### 4. s[low:high]

![import19.png](https://www.zlkt.net/media/course/sXRfQ9t8Aa3pkCUarHkMwS.png)
结果是2-5的值。`array[2:5]`表示从下标为2的元素（包含该元素）开始取，到下标为5的元素（不包含该元素）结束。所以切片`s5`的长度是3。切片`s5`的容量是多少呢？是8，根据`array`切片的容量是10,减去`array[2:5]`中的2。

以上就是关于切片的基本操作，这些操作在以后的开发过程中会经常用到，希望大家记住基本的规律。

## 五、 思考题

接下来说，思考如下题，定义一个切片`array`，然后对该切片`array`进行截取操作（范围自定义），得到新的切片`s6`,并修改切片`s6`某个元素的值。代码如下：
![import20.png](https://www.zlkt.net/media/course/ErXi2Jq3gsZnW7BzSyzNhf.png)

`s6`切片的结果是：`[2,3,4]`因为是从下标为2的元素（包含）开始取，到下标为5的元素（不包含）结束,取出3个元素，也就是长度为3。

现在将程序进行如下修改：
![import21.png](https://www.zlkt.net/media/course/pKiVkHRphef9Pb6zgeiJ76.png)

现在程序的输出结果是：

```
s6 = [2 3 888]
```

因为切除了234，然后现在0是2，1是3，2是4，然后把s6[2]也就是s6[4]赋值为888
接下来输出切片`array`的值：

![import23.png](https://www.zlkt.net/media/course/aAzfpJGQHr5VxmTktd6VNZ.png)

输出的结果如下：

```
s6 = [2 3 888]
array = [0 1 2 3 888 5 6 7 8 9]
```

发现切片`array`中的值也发生了变化，也就是修改切片s6的值会影响到原切片`array`的值，下面通过画图的形式来说明其原因。

![import25.png](https://www.zlkt.net/media/course/dXw2NgUzFKRArV8J99Bjhd.png)

在这里重点要理解的是：`s6 := array[2:5]`，将`array`切片中的`array[2]`，`array[3]`，`array[4]`截取作为新切片`s6`，实际上是切片`s6`指向了原切片`array`（在这里并不是为切片s6新建一块区域）。所以修改`s6`，也会影响到array。

下面继续修改上面的程序：
![import26.png](https://www.zlkt.net/media/course/yiy8RtTYVT2wQFVoW32sdT.png)

以上程序中，切片s7的值是多少？

```
结果是：s7 = [888 5 6 7 8]
```

下面也是通过画图的形式，来解释该程序的结果：
![import27.png](https://www.zlkt.net/media/course/qEWch7g99NbWSf2J9SNM95.png)

继续思考,现在在原有的程序中又加了一行，如下图所示：

![import28.png](https://www.zlkt.net/media/course/wCWLXxMPzswpp7HgXHWZoD.png)

最终，切片s7与原来切片array的值分别是多少？

结果所示：

```
s6 = [2 3 888]
s7 = [888 5 999 7 8]
array = [0 1 2 3 888 5 999 7 8 9]
```

## 六、 `append`函数的使用

在第一节中，已经给大家讲解过切片与数组很大的一个区别就是：切片的长度是不固定的，可以向已经定义的切片中追加数据。并且也给大家简单的演示过通过`append`的函数，在原切片的末尾添加元素。

```
arr := []int{1,2,3}
arr = append(arr,4) //追加一个数
arr = append(arr,5,6,7) //追加多个数
fmt.Println(arr)
```

如果容量不够用了，该怎么办呢？

例如有以下切片：

```
s:= make([]int, 5, 8)
```

定义了切片`s`，长度是5，容量是8，`k`。

```
s := make([]int,5,8)
fmt.Printf("len = %d,cap=%d\n",len(s),cap(s))
```

结果是：`len = 5 cap = 8`

并且前面我们讲解过，长度是指已经初始化的空间，现在切片`s`没有赋值，但是默认值为0
验证如下所示：

```
s := make([]int,5,8)
fmt.Printf("len = %d,cap=%d\n",len(s),cap(s))
fmt.Println(s)
```

结果是：

```
len = 5 cap = 8
[0 0 0 0 0]
```

现在开始通过append函数追加数据，如下所示:

```
s := make([]int,5,8)
s = append(s,1)
fmt.Println(s)
fmt.Printf("len = %d,cap=%d\n",len(s),cap(s))
```

输出结果是：

```
[0 0 0 0 0 1]
len = 6 cap = 8
```

从输出的结果上，我们完全能够体会到，append函数的作用是在末尾追加（直接在默认值后面追加数据），由于追加了一个元素，所以长度为6.

但是如果我们把程序修改成如下所示：

```
s := make([]int,5,8)
//s = append(s,1)
s[0] = 1
fmt.Println(s)
fmt.Printf("len = %d,cap=%d\n",len(s),cap(s))
```

输出结果是：

```
[1 0 0 0 0]
len = 5 cap = 8
```

由于s[0]=1是直接给下标为0的元素赋值，并不是追加，所以结果的长度不变。

下面我们继续通过append( )继续追加数据：

```
s := make([]int,5,8)
s = append(s,1)
s = append(s,2)
s = append(s,3)

fmt.Println(s)
fmt.Printf("len = %d,cap=%d\n",len(s),cap(s))
```

结果是：

```
[0 0 0 0 0 1 2 3]
len = 8 cap = 8
```

追加完成3个数据后，长度变为了8，与容量相同。

那么如果现在通过append( )函数，继续向切片s中继续追加一个数据，那么容量会变为多少呢？

代码如下：

```
s := make([]int,5,8)
s = append(s,1)
s = append(s,2)
s = append(s,3)
s = append(s,4)
fmt.Println(s)
fmt.Printf("len = %d,cap=%d\n",len(s),cap(s))
```

输出的结果是：

```
[0 0 0 0 0 1 2 3 4]
len = 9 cap = 16
```

追加完成一个数据后，长度变为9，大于创建切片s时的容量，所以切片s扩容，变为16.

那么切片的容量是否是以2倍容量来进行扩容的呢？

我们可以来验证一下：

![import29.png](https://www.zlkt.net/media/course/BcwkMT5QwwYYEDZKh4c47X.png)

输出结果是：

![img](https://www.zlkt.net/assets/import30.png)
![import30.png](https://www.zlkt.net/media/course/LqTwGSEbb4BELJqeN8zhRh.png)

通过以上结果分析，发现是2倍的容量进行扩容。

但是我们修改一下循环条件看一下结果，将循环结束的条件修改的大一些，如下所示：

![import31.png](https://www.zlkt.net/media/course/6CyGheSniRaon7DwG6rvGD.png)

对应的结果：

![import32.png](https://www.zlkt.net/media/course/gpdPWdAEiXxuVXZyCxXTXF.png)

通过以上的运行结果分析：当容量小于1024时是按照2倍容量扩容，当大于等于1024就不是按照2倍容量扩容。

## 七、 `copy`函数使用

针对切片操作常用的方法除了`append\()`方法以外，还有`copy`方法。

基本语法：`copy(切片1，切片2)`

将第二个切片里面的元素，拷贝到第一个切片中。

下面通过一个案例，看一下该方法的使用：

![import34.png](https://www.zlkt.net/media/course/96fwiMWquPSDKZP739fSMb.png)

上面案例中，将`srcSlice`中的元素拷贝到`destSlice`切片中。结果如下：

```
dst = [1 2 6 6 6]
```

通过以上结果可以分析出，直接将`srcSlice`切片中两个元素拷贝到`dstSlice`元素中相同的位置。而`dstSlice`原有的元素备替换掉。

下面将以上程序修改一下，如下所示：

![img](https://www.zlkt.net/assets/import35.png)
![import35.png](https://www.zlkt.net/media/course/hmpTpkU4K2kuJwPtqkf6mJ.png)

以上程序的结果是：

```
src = [6 6]
```

通过以上两个程序得出如下结论：在进行拷贝时,拷贝的长度为两个`slice`中长度较小的长度值。

思考以下程序输出的结果：
![import36.png](https://www.zlkt.net/media/course/H9myeAePY2QifhYRVobxyK.png)

结果是：

```
slice2 = [1 2 3]
```

现在将程序进行如下修改：
![import37.png](https://www.zlkt.net/media/course/eNNtBScj5jKEqcuRgimXYC.png)

结果是：

```
slice1 = [5 4 3 4 5]
```

## 八、切片作为函数参数

切片也可以作为函数参数，那么与数组作为函数参数有什么区别呢？

接下来通过一个案例，演示一下切片作为函数参数。

![import39.png](https://www.zlkt.net/media/course/dgHGeFzAPsBFVJRjVzP49f.png)

通过以上案例，发现在主函数`main()`中，定义了一个切片`s`，然后调用`InitData()`函数，将切片s作为实参传递到该函数中，并在`InitData()`函数中完成初始化，该函数并没有返回值，但是在主函数中直接打印切片`s`，发现能够输出对应的值。也就是在`InitData()`函数中对形参切片`num`赋值，影响到了`main()`函数中的切片`s`。

但是，大家仔细想一下，如果我们这里传递参数不是切片，而是数组，那么能否完成该操作呢？

那么我们将上面的程序，修改成以数组作为参数进行传递的形式：

![import40.png](https://www.zlkt.net/media/course/wJFyMKG5hFTjCjob9VyS5o.png)

发现以数组的形式作为参数，并不能完成我们的要求，所以切片作为函数实参与数组作为函数实参，进行传递时，传递的方式是不一样的。

在GO语言中，数组作为参数进行传递是值传递，而切片作为参数进行传递是引用传递。

## 九、值传递和引用传递：

- 值传递：方法调用时，实参数把它的值传递给对应的形式参数，方法执行中形式参数值的改变不影响实际参数的值
- 引用传递：也称为传地址。函数调用时，实际参数的引用(地址，而不是参数的值)被传递给函数中相对应的形式参数（实参与形参指向了同一块存储区域），在函数执行中，对形式参数的操作实际上就是对实际参数的操作，方法执行中形式参数值的改变将会影响实际参数的值。

**建议：以后开发中使用切片来代替数组。**

# Map

前面我们学习了GO语言中数组，切片类型，但是我们发现使用数组或者是切片存储的数据量如果比较大，那么通过下标来取出某个具体的数据的时候相对来说，比较麻烦。例如：

```
names := []string{"张三","李四","王五"}
fmt.Println(names[2])
```

现在要取出切片中存储的“王五”，那么需要数一下对应的下标值是多少，这样相对来说就比较麻烦。有没有一种结构能够帮我们快速的取出数据呢？就是字典结构。

说道字典大家想到的就是：

![import666.png](https://www.zlkt.net/media/course/VtHBs2u8szkpFjvSgm7UGj.png)

在使用新华字典查询某个字，我们一般都是根据前面的部首或者是拼音来确定出要查询的该字在什么位置，然后打开对应的页码，查看该字的解释。

GO语言中的字典结构是有键和值构成的。

所谓的键，就类似于新华字典的部首或拼音，可以快速查询出对应的数据。

如下图所示：

![import22.png](https://www.zlkt.net/media/course/3zPtMvcbUTKWsSdK8378gC.png)

通过该图，发现某个键（key）都对应的一个值(value)，如果现在要查询某个值，直接根据键就可以查询出某个值。

在这里需要注意的就是字典中的键是不允许重复的，就像身份证号一样。

## 一、 字典结构定义

```
map[keyType]valueType
```

定义字典结构使用map关键字，[ ]中指定的是键(key)的类型,后面紧跟着的是值的类型。

键的类型，必须是支持`==`和`!=`操作符的类型，切片、函数以及包含切片的结构类型不能作为字典的键，使用这些类型会造成编译错误：

```
//err invalid map key type []string
dict := map[[]string]int{} 
```

下面定义一个字典`m`，键的类型是整型，值的类型是字符串。

```
var m map[int]string
fmt.Println(m)
```

定义完后，直接打印，结果为空`nil`。

注意：字典中不能使用`cap`函数，只能使用`len()`函数。`len()`函数返回`map`拥有的键值对的数量

```
var m map[int]string
fmt.Println(len(m))
```

以上代码值为0，也就是没有值。

当然也可以使用`make()`函数来定义，如下所示：

```
m2 := make(map[int]string)
fmt.Println(m2)
fmt.Println(len(m2))
```

以上代码值为0，也就是没有值。

当然也可以指定容量。

```
m2 := make(map[int]string,3)
fmt.Println(m2)
fmt.Println(len(m2))
```

输出的`len`值还是0，因为这里并没有赋值。

接下来可以给字典`m2`进行赋值，并且指定容量，如果容量不够自动扩容。

```
m2 := make(map[int]string,3)

m2[1] = "张三"
m2[2] = "李四"
m2[3] = "王五"

fmt.Println(m2)
fmt.Println(len(m2))
```

可以直接使用键完成赋值，再次强调键是唯一的，同时发现字典`m2`的输出结果，不一定是按照赋值的顺序输出的，每次运行输出的顺序可能都不一样，所以这里一定要注意：`map`是无序的，我们无法决定它的返回顺序，所以，每次打印结果的顺利有可能不同。

`map`也可以定义完成后直接进行初始化

```
m4 := map[int]string{1:"make",2:"Go"}
fmt.Println(m4[1])
fmt.Println(m4[2])
```

也就是在定义的同时给他直接赋值,然后打印出来

## 二、 打印字典中的值

### 1. 可以直接通过键输出，如下所示：

```
m4 := map[int]string{1:"make",2:"Go"}
fmt.Println(m4[1])//make
fmt.Println(m4[2])//go
```

通过打印键的方式就能得到值

### 2. 通过循环遍历的方式输出

```
m4 := map[int]string{1:"make",2:"Go"}
for key,value := range m4 {
  fmt.Println(key)
  fmt.Println(value)
}
//1 make
//2 go
```

其中key代表的是键，value代表的是值

输出的顺序是无序的。

### 3. 在输出的时候，还可以进行判断。

```
m4 := map[int]string{1:"make",2:"Go"}
value,ok := m4[1]
if ok == true{
  fmt.Println(value)
}else{
  fmt.Println("key不存在")
}
```

第一个返回值为key所对应的value, 第二个返回值为key是否存在的条件，存在ok为true。

删除map中的某个元素。

根据map中的键，删除对应的元素，也是非常的方便。

如下所示：

```
m4 := map[int]string{1:"make",2:"Go"}
delete(m4,1) //删除key为1的内容
fmt.Println(m4)//2 go
```

map作为函数参数是引用传递。

```
func test(m map[int]string){
  delete(m,1)
}

func main(){
  m4 := map[int]string{1:"make",2:"Go"}
  test(m4)
  fmt.Println(m4)// 2 go
}
```

第一个`Test`定义了一个删除键为1的方法，然后在`main`里面又重新初始化了一个map然后调用`Test`的方法,最后输出的结果就是`go`，因为调用`Test`方法的时候把键为1的键值删除了。

# 结构体

现在有一个需求，要求存储学生的详细信息，例如，学生的学号，学生的姓名，年龄，家庭住址等。按照以前学习的存储方式，可以以如下的方式进行存储：

![import41.png](https://www.zlkt.net/media/course/GxioU5jE92moEAcCNatPsi.png)

通过定义变量的信息，进行存储。但是这种方式，比较麻烦，并且不利于数据的管理。

在GO语言中，我们可以通过结构体来存储以上类型的数据，结构体的定义如下：

![import42.png](https://www.zlkt.net/media/course/gzxjBeLWbhAAxo9kzrpP5P.png)

`type`后面跟着的是结构体的名字`Student`, `struct`表示定义的是一个结构体。
大括号中是结构体的成员，注意在定义结构体成员时，不要加`var`。通过以上的定义，大家能够感觉出，通过结构体来定义复杂的数据结构，非常清晰。

结构体定义完成后，可以进行初始化。

## 一、结构体初始

![import43.png](https://www.zlkt.net/media/course/TTVot2x7LeNKGP2REb9b9T.png)

注意：顺序初始化，每个成员必须初始化，在初始化时，值的顺序与结构体成员的顺序保持一致。

![import44.png](https://www.zlkt.net/media/course/bjXddmvfeQwx7E5Jd46ybK.png)

结构体定义完成后，结构体成员的使用。
![import45.png](https://www.zlkt.net/media/course/JHy2pBXFK7WKotrjQMNG64.png)

## 二、结构体比较与赋值

两个结构体可以使用 == 或 != 运算符进行比较，但不支持 > 或 <。

![import46.png](https://www.zlkt.net/media/course/UtR8vNPNHwBazt58WnUygk.png)

同类型的两个结构体变量可以相互赋值。

![import47.png](https://www.zlkt.net/media/course/RisecVTwg3oKzVtMUgWaKW.png)

## 三、结构体数组

上一小节，我们已经对结构体的定义，与基本使用有一定的了解了，下面有一个需求：用结构体存储多个学生的信息。

可以使用上一小节讲解的，通过结构体定义多个结构体变量，也可以定义结构体数组来存储。

结构体数组定义如下所示：
![import48.png](https://www.zlkt.net/media/course/JAeSBPLMBAbZJYk28Bxnv.png)

上面的代码首先是定义了一个`Student`的结构体,然后在`main`方法里面用一个变量`Students`接收了这个新建的`[]Student`结构体数组，结构体是放一个`Student`那么，结构体数组就可想而知了，是放多个结构体,然后循环遍历`Students`的下标也就是有几个结构体就会遍历几次，最后再打印里面的`sutdents`里面的`name`，最后输出的结果"张三",“李四”,“王五”。

## 四、结构体作为函数参数

结构体也可以作为函数参数，进行传递，如下所示：

![import49.png](https://www.zlkt.net/media/course/VYEz7Zaa3SAN2vx3LEAujC.png)

把结构体作为参数的话，参数类型就只能放相对应的结构体，不然会报错。
上面代码首先在`Test`里面修改一下`student`里面的id为`666`，然后在`main`里面新建一个为s的结构体，这个结构体和`student`一样，所以`Test`里面能放，放进去之后就会修改这个`id`，最后打印的结果为：`id:666,name:mike,sex:m,age:18,addr:bj`。

结构体作为函数参数进行传递，是**值传递**。

# 指针

## 一、变量内存与地址

前面我们讲过存储数据的方式，可以通过变量，或者复合类型中的**数组**、**切片**、**Map**、**结构体**。我们不管使用变量存储数据，还是使用符合类型存储数据，都有两层的含义：存储的数据（内存），对应的地址。

接下来，通过变量来说明以上两个含义。例如，定义如下变量：

![import70.png](https://www.zlkt.net/media/course/5yemX4tjtrQKAjD7B6yuc7.png)

第一个`Printf()`函数的输出，大家都很熟悉，输出变量`i`的值，这个实际上就是输出内存中存储的数据。在前面的章节中，已经讲解过，定义一个变量，就是在内存中开辟一个空间，用来存储数据，当给变量`i`赋值为100，其实就是将100存储在改空间内。

第二个`Printf()`函数的输出，输出的是变量`i`在内存中的地址。通过如下图来给大家解释：

![import71.png](https://www.zlkt.net/media/course/ZvdrH378okzpvVp2wLmnc7.png)

这张图，大家也应该非常熟悉，是在讲解变量时，画的一张图，`0x100010`假设是变量`i`的内存地址（通过第二个输出可以获取实际的地址），内存地址的作用：在输出变量中存储的数据时，是通过地址来找到该变量内存空间的。

这个内存地址和实际生活中的地址也很相似，例如：大家可以将内存空间想象成，我们上课的教室，教室中存放有学生，那么现在要找一个学生，必须要知道具体的地址以及教室门牌号。

以上程序输出的结果是：

![import72.png](https://www.zlkt.net/media/course/9AMWU8iXT2DgXL6r2cLPTd.png)

## 二、 指针变量

现在已经知道怎样获取变量在内存中的地址，但是如果想将获取的地址进行保存，应该怎样做呢？

可以通过指针变量来存储，所谓的指针变量：就是用来存储任何一个值的内存地址。

指针变量的定义如下：

![import74.png](https://www.zlkt.net/media/course/WYBUfg6g7Pn3Gfj2wr3sTM.png)

指针变量p的定义是通过`*`这个符号来定义，指针变量p的类型为`*int`,表示存储的是一个整型变量的地址。

如果指针变量`p`存储的是一个字符串类型变量的地址，那么指针变量`p`的类型为`*string p=&i `该行代码的意思是，将变量`i`的地址取出来，并且赋值给指针变量`p`。也就是指针变量p指向了变量`i`的存储单元。

可以通过如下图来表示：

![import75.png](https://www.zlkt.net/media/course/a6X5yUw2mfJsVVsjZGpNNW.png)

在以上图中，一定要注意：指针变量`p`存储的是变量`i`的地址。

大家可以思考一个问题：

既然指针变量`p`指向了变量`i`的存储单元，那么是否可以通过指针变量`p`，来操作变量`i`中存储的数据？

答案是可以的，具体操作方式如下：

![import76.png](https://www.zlkt.net/media/course/zbpDac348rCdGB6so2ifnH.png)

注意：在使用指针变量`p`来修改变量i的值的时候，前面一定要加上`*`.（通过指针访问目标对象）

现在打印变量`i`的值已经有100变为80.

当然，也可以通过指针变量`p`来输出，变量`i`中的值，输出的方式如下所示：

![import77.png](https://www.zlkt.net/media/course/XuPpmLV4KU4TjPshJGsff8.png)

所以，`*p`的作用就是根据存储的变量的地址，来操作变量的存储单元（包括输出变量存储单元中的值，和对值进行修改）

## 三、 注意事项

在使用指针变量时，要注意以下两点。

### 1. 默认值为`nil`

```
var p *int
fmt.Println(p)
```

直接执行上面的程序，结果是：`nil`

### 2. 不要操作没有合法指向的内存。

例如，在上面的案例中，我们定义了指针变量`p`，但是没有让指针变量指向任何一个变量，那么直接运行如下程序，会出现异常。

```
var p *int
*p = 99 //没有指向 直接操作
fmt.Println(p)
```

出现的错误信息如下：

![import78.png](https://www.zlkt.net/media/course/JEJrertc2grVnesDj5JaoD.png)

所以，在使用指针变量时，一定要让指针变量有正确的指向。以下的操作是合法的：

```
var a int
var p *int
p = &a //指向变量a
*p = 99
fmt.Println(p)
```

在该案例中，定义了一个变量`a`，同时定义了一个指针变量`p`,将变量a的地址赋值给指针变量`p`，也就是指针变量`p`指向了变量`a`的存储单元。给指针变量`p`赋值，影响到了变量`a`。最终输出变量a中的值也是56。

## 四、 `new()`函数

指针变量，除了以上介绍的指向以外`(p=&a)`，还可以通过`new()`函数来指向。

具体的应用方式如下:

```
var p *int
p = new(int)
*p = 59
fmt.Println(*p)
```

`new(int)`作用就是创建一个整型大小(4字节)的空间

然后让指针变量`p`指向了该空间，所以通过指针变量`p`进行赋值后，该空间中的值就是57。

`new()`函数的作用就是`C`语言中的动态分配空间。但是在这里与`C`语言不同的地方，就是最后不需要关系该空间的释放。GO语言会自动释放。这也是比C语言使用方便的地方。

也可以使用自动推导类型的方式：

```
q := new(int)
*q = 77
fmt.Println(*q)
```

## 五、指针做函数参数

指针也可以作为函数参数，那么指针作为函数参数在进行传递的时候，是值传递还是引用传递呢？

大家都知道，普通变量作为函数参数进行传递是值传递，如下案例所示：

定义一个函数，实现两个变量值的交换。

![import79.png](https://www.zlkt.net/media/course/5PnLt5kUyY2ByAr2un3xk3.png)

通过以上案例，证实普通类型变量在传递时，为值传递。

那么使用指针作为函数参数呢？现在将以上案例修改成，用指针作为参数，如下所示：

![import80.png](https://www.zlkt.net/media/course/ZUiKNTZEH8Z5SZ75RAGRGc.png)

通过以上案例证实，指针作为参数进行传递时，为引用传递，也就是传递的地址。

在调用`Swap()`函数时，将变量`a`与变量`b`的地址传分别传递给指针变量`num1`，`num2`,这时`num1`和`num2`，分别指向了变量`a`，与变量`b`的内存存储单元，那么操作`num1`，`num2`实际上操作的就是变量`a`与变量`b`，所以变量`a`与变量`b`的值被交换。

## 六、 数组指针

前面在讲解数组的时候，我们用数组作为函数参数，但是数组作为参数进行传递是值传递，如果想引用传递，可以使用数组指针。具体使用方式如下：

![import81.png](https://www.zlkt.net/media/course/N3FMtKomNy7c69nmLdhA8d.png)

定义一个数组，作为函数Swap的实参进行传递，但是这里传递的是数组的地址，所以Swap的形参是数组指针。

这时指针`p`，指向了数组`a`,对指针`p`的操作实际上是对数组`a`的操作，所以如果直接执行如下语句：`fmt.Println(*p)`，会输出数组`a`中的值。也可以通过`*p`结合下标将对应的值取出来进行修改。最终在`main`函数中输出数组`a`，发现其元素也已经修改。

当然，我们也可以通过循环的方式来将数组指针中的数据打印出来：

![import82.png](https://www.zlkt.net/media/course/5vTJEHibutUa4H3Y2omQ9g.png)

## 7、 指针数组

上一小节，讲解到的是数组指针，也就是让一个指针指向数组，然后可以通过该指针来操作数组。还有一个概念叫指针数组，这两个概念很容混淆，指针数组指的是一个数组中存储的都是指针（也就是地址）。也就是一个存储了地址的数组。

下面通过一个案例，看一下指针数组的应用

![import83.png](https://www.zlkt.net/media/course/q8LCFYBYga7uW5HSTefVkB.png)

指针数组的定义方式，与数组指针定义方式是不一样的，注意指针数组是将“*”放在了下标的后面。

由于指针数组存储的都是地址，所以将变量i,与变量j的地址赋值给了指针数组p。

最后输出指针数组p中存储的地址。

思考:既然指针数组p存储了变量i和变量j的的地址，那么怎样通过指针数组p操作变量i与变量j的值呢？

具体实现如下：

![import84.png](https://www.zlkt.net/media/course/Gdepk3ospxK9Z8t5jE3SKL.png)

注意这里输出要注意的问题是，没有加小括号。（注意运算顺序）

当然，我们也可以通过for循环的方式来输出指针数组中对应的值。

![img](https://www.zlkt.net/assets/import85.png)
![import85.png](https://www.zlkt.net/media/course/QVh5eWdHiQMeHs7aEhooFo.png)

## 八、 结构体指针变量

我们前面定义了指针指向了数组，解决了数组引用传递的问题。那么指针是否可以指向结构体，也能够解决结构体引用传递的问题呢？完全可以。

下面我们先来看一下，结构体指针变量的定义：

![import86.png](https://www.zlkt.net/media/course/Z4cuAJPEQNMvbtYEh5BWyK.png)

也可以使用自动推导类型

![import87.png](https://www.zlkt.net/media/course/iLKAuU2vRduHgyYyHDRHK.png)

现在定义了一个结构体指针变量，那么可以通过该指针变量来操作结构体中的成员项。

![import88.png](https://www.zlkt.net/media/course/LjnYoaHJ8JooQG79bmxB2m.png)

前面在讲解结构时，用结构体作为函数的参数，默认的是值传递，那么通过结构体指针，可以实现结构体的引用传递。具体实现的方式如下：

![import89.png](https://www.zlkt.net/media/course/xJxQRbopGyTfgZXqwHZR4G.png)

# 面向对象

前面我们已经将GO语言中各种类型，给大家讲解完毕了，那么接下来要给大家讲解的是面向对象编程思想。

在讲解具体面向对象编程之前，先说一下面向过程编程。我们前面学习都是面向过程的一种编程思想，接下来可以从生活中理解面向过程:

![import07.png](https://www.zlkt.net/media/course/T8EwmNsQMUwa5aWQj9SzWd.png)

如果我们自己来修电脑，应该有哪些步骤呢？

- 第一步：判断问题的原因
- 第二步：找工具
- 第三步：暴力拆卸

这个修理的步骤就是面向过程，所谓的面向过程就是：强调的是步骤、过程、每一步都是自己亲自去实现的。

如果采用面向对象的思想，那么应该怎样修电脑呢？

![import08.png](https://www.zlkt.net/media/course/WWk6rPVNbNN59kwhENCKRj.png)

找维修店的工作人员来帮我们修电脑，但是到底怎么修，我们是不用考虑的，也就是说我们不关心步骤与过程。

大家可以想一下，在生活中还有哪些事情是面向过程，面向对象的。

比如说，做饭，面向过程就是自己做，自己买菜，自己洗，自己炒，整个过程都有自己来完成，但是如果是面向对象，可以叫外卖，不用关心饭是怎么做的。

所以通过以上案例，大家能够体会出，面向过程就是强调的步骤，过程，而面向对象强调的是对象，找个人来做。

在面向对象中，还有两个概念是比较重要的，一是对象，二是类。

## 什么是对象

万物皆对象，例如小明同学是一个对象，小亮同学也是一个对象。那么我们在生活中怎样描述一个对象呢？
比如，描述一下小明同学：
姓名：小明
性别：男
身高：180cm
体重：70kg
年龄：22岁

吃喝拉撒睡一切正常健康，吃喝嫖赌抽。

通过以上的描述，可以总结出在生活中描述对象，可以通过特征（身高，体重，年龄等）和行为（爱好等）来进行描述。

那么在程序中，可以通过属性和方法（函数）来描述对象。属性就是特征，方法（函数）就是行为。所以说，对象必须具有属性和方法。虽然说，万物皆对象，但是在描述一个对象的时候，一定要具体不能泛指，例如，不能说“电灯”是一个对象，而是说具体的哪一台“电灯”。

大家可以思考一下，如果我们现在描述一下教室中某一台电灯，应该有哪些属性（特征）和方法（行为）呢？

下面我们在思考一下，下面这道题：

小明(一个学生)\杨老师\邻居王叔叔\小亮的爸爸\小亮的妈妈

找出这道题中所有对象的共性（所谓共性，指的是相同的属性和方法）。

所以说，我们可以将这些具有相同属性和相同方法的对象进行进一步的封装，抽象出来类这个概念。

类就是个模子，确定了对象应该具有的属性和方法。

对象是根据类创建出来的

例如:上面的案例中，我们可以抽出一个“人”类（都有年龄，性别，姓名等属性，都有吃饭，走路等行为），“小明”这个对象就是根据“人”类创建出来的，也就是说先有类后有对象。

## GO语言中的面向对象

前面我们了解了一下，什么是面向对象，以及类和对象的概念。但是，GO语言中的面向对象在某些概念上和其它的编程语言还是有差别的。

严格意义上说，GO语言中没有类(class)的概念,但是我们可以将结构体比作为类，因为在结构体中可以添加属性（成员），方法（函数）。

面向对象编程的好处比较多，我们先来说一下“继承”，

所谓继承指的是，我们可能会在一些类（结构体）中，写一些重复的成员，我们可以将这些重复的成员，单独的封装到一个类(结构体)中，作为这些类的父类(结构体)，我们可以通过如下图来理解：

![import09.png](https://www.zlkt.net/media/course/i8s27NAR8nBzijzAfLchfh.png)

当然严格意义上，GO语言中是没有继承的，但是我们可以通过”匿名组合”来实现继承的效果。

# 匿名组合

## 一、 匿名字段

一般情况下，定义结构体的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，也就是匿名字段，也称为嵌入字段。

当匿名字段也是一个结构体的时候，那么这个结构体所拥有的全部字段都被隐式地引入了当前定义的这个结构体。

```
//人
type Person struct {
  name string
  sex byte
  age int
}

//学生
type Student struct {
  Person //匿名字段，那么默认Student就包含了Person的所有字段
  id int
  addr string
}
```

`Person`也就是上面定义的这个`Person`结构体。

## 二、 初始化

```
//人
type Person struct {
  name string
  sex byte
  age int
}

//学生
type Student struct {
  Person//匿名字段，那么默认Student就包含了Person的所有字段
  id int
  addr string
}

func main() {
  //顺序初始化
  s1 := Student{Person{"mike",'m',18},1,"sz"}
  //s1 = {Person:{name:mike sex:109 age:18}id:1 addr:sz}
  fmt.Printf("s1=%+v\n",s1)
  //s2 := Student{"mike",'m',18,1,"sz"}//err
  //部分成员初始化1
  s3 := Student{Person:Person{"lily",'f',19},id:2}
  //s3 = {Person:{name:lily sex:102 age:19}id:2 addr:}
  fmt.Printf("s3=%+v\n",s3)
  //部分成员初始化2
  s4 := Student{Person:Person{name:"tom"},id:3}
  //s4 = {Person:{name:tomsex:0age:0}id:3addr:}
  fmt.Printf("s4=%+v\n",s4)
}
```

然后我们在`main`里面调用`Student`就能直接对`Person`里面的属性赋值。

## 三、 成员的操作

```
var s1 Student//变量声明
//给成员赋值
s1.name = "mike"//等价于s1.Person.name="mike"
s1.sex = 'm'
s1.age = 18
s1.id = 1
s1.addr = "sz"
fmt.Println(s1) //{{mike 109 18}1 sz}
var s2 Student//变量声明
s2.Person = Person{"lily",'f',19}
s2.id = 2
s2.addr = "bj"
fmt.Println(s2) //{{lily 102 19}2 bj}
```

或者我们声明一个`Student`的变量也能调用它里面的属性。

## 四、 同名字段

```
//人
type Person struct{
  name string
  sex byte
  age int
}

//学生
type Student struct{
  Person //匿名字段，那么默认Student就包含了Person的所有字段
  id int
  addr string
  name string //和Person中的name同名
}

func main(){
  var s Student//变量声明
  //给Student的name，还是给Person赋值？
  s.name = "mike"
  //{Person:{name:sex:0age:0}id:0addr:name:mike}
  fmt.Printf("%+v\n",s)
  //默认只会给最外层的成员赋值
  //给匿名同名成员赋值，需要显示调用
  s.Person.name = "yoyo"
  //Person:{name:yoyosex:0age:0}id:0addr:name:mike}
  fmt.Printf("%+v\n",s)
}
```

如果命名重名的话我们调用只会给最外层的使用，也就是`Student`，如果说你要给`Person`赋值的话得明确表示。`s.Person.name="张三"`。

## 五、 其它匿名字段

### 1. 非结构体类型

所有的内置类型和自定义类型都是可以作为匿名字段的：

```
type mystr string//自定义类型
type Person struct {
  name string
  sex byte
  age int
}
type Student struct {
  Person //匿名字段，结构体类型
  int //匿名字段，内置类型
  mystr //匿名字段，自定义类型
}
func main() {
  //初始化
  s1 := Student{Person{"mike",'m',18},1,"bj"}
  //{Person:{name:mikesex:109age:18}int:1mystr:bj}
  fmt.Printf("%+v\n",s1)
  //成员的操作，打印结果：mike,m,18,1,bj
  fmt.Printf("%s,%c,%d,%d,%s\n",s1.name,s1.sex,s1.age,s1.int,s1.mystr)
}
```

不一样要结构体才能作为匿名字段，其实定义一个类型也是一样的。

### 2. 结构体指针类型

```
type Person struct { //人
  name string
  sex byte
  age int
}
type Student struct {//学生
  *Person //匿名字段，结构体指针类型
  id int
  addr string
}
func main() {
  //初始化
  s1 := Student{&Person{"mike",'m',18},1,"bj"}
  //{Person:0xc0420023e0id:1addr:bj}
  fmt.Printf("%+v\n",s1)
  //mike,m,18
  fmt.Printf("%s,%c,%d\n",s1.name,s1.sex,s1.age)
  //声明变量
  var s2 Student
  s2.Person = new(Person)//分配空间
  s2.name = "yoyo"
  s2.sex = 'f'
  s2.age = 20
  s2.id = 2
  s2.addr = "sz"
  //yoyo10220220
  fmt.Println(s2.name,s2.sex,s2.age,s2.id,s2.age)
}
```

在匿名方法里面也是能使用指针的，只要在前面加上`&`就行。

# 方法

## 一、 概述

在面向对象编程中，一个对象其实也就是一个简单的值或者一个变量，在这个对象中会包含一些函数，**这种带有接收者的函数，我们称为方法(method)**。本质上，一个方法则是一个和特殊类型关联的函数。

一个面向对象的程序会用方法来表达其属性和对应的操作，这样使用这个对象的用户就不需要直接去操作对象，而是借助方法来做这些事情。

在Go语言中，可以给任意自定义类型（包括内置类型，但不包括指针类型）添加相应的方法。

⽅法总是绑定对象实例，并隐式将实例作为第⼀实参 (receiver)，方法的语法如下：

```
func (receiver ReceiverType) funcName (parameters) (results)
```

- 参数 receiver 可任意命名。如⽅法中未曾使⽤，可省略参数名。
- 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接⼝或指针。
- 不支持重载方法，也就是说，不能定义名字相同但是不同参数的方法。

## 二、 为类型添加方法

### 1. 基础类型作为接收者

```
// 自定义类型，给int改名为MyInt
type MyInt int

// 在函数定义时，在其名字之前放上一个变量，即是一个方法
func (a MyInt) Add(b MyInt) MyInt {
  return a + b
}

//传统方式的定义
func Add(a, b MyInt) MyInt {//面向过程
  return a + b
}

func main() {
  var a MyInt=1
  var b MyInt=1

  //调用func (aMyInt) Add(bMyInt)
  fmt.Println("a.Add(b)=",a.Add(b))//a.Add(b)=2

  //调用func Add(a,bMyInt)
  fmt.Println("Add(a,b)=",Add(a,b))//Add(a,b)=2
}
```

通过上面的例子可以看出，面向对象只是换了一种语法形式来表达。方法是函数的语法糖，因为`receiver`其实就是方法所接收的第1个参数。

注意：虽然方法的名字一模一样，但是如果接收者不一样，那么方法就不一样。

### 2. 结构体作为接收者

方法里面可以访问接收者的字段，调用方法通过点`.`访问，就像`struct``里面访问字段一样：

```
type Person struct {
  name string
  sex byte
  age int
}

func (p Person) PrintInfo(){//给Person添加方法
  fmt.Println(p.name,p.sex,p.age)
}

func main() {
  p:=Person{"mike",'m',18}//初始化
  p.PrintInfo()//调用func(pPerson)PrintInfo()
}
```

打印结果为`mike，m,18`，你方法写的是Person那么这个方法只能传`Person`，不能传别的类型。

## 三、 值语义和引用语义

```
type Person struct {
  name string
  sex byte
  age int
}

// 指针作为接收者，引用语义
func (p *Person) SetInfoPointer(){
  // 给成员赋值
  (*p).name = "yoyo"
  p.sex = 'f'
  p.age = 22
}

// 值作为接收者，值语义
func (p Person) SetInfoValue(){
  // 给成员赋值
  p.name = "yoyo"
  p.sex = 'f'
  p.age = 22
}

func main() {
  // 指针作为接收者，引用语义
  p1 := Person{"mike",'m',18} // 初始化
  fmt.Println("函数调用前=",p1) // 函数调用前={mike10918}
  (&p1).SetInfoPointer()
  fmt.Println("函数调用后=",p1) // 函数调用后={yoyo10222}

  fmt.Println("==========================")

  p2 := Person{"mike",'m',18} // 初始化
  // 值作为接收者，值语义
  fmt.Println("函数调用前=",p2) // 函数调用前={mike10918}
  p2.SetInfoValue()
  fmt.Println("函数调用后=",p2) // 函数调用后={mike10918}
}
```

## 四、方法集

类型的方法集是指可以被该类型的值调用的所有方法的集合。

用实例实例`value`和`pointer`调用方法（含匿名字段）不受⽅法集约束，编译器编总是查找全部方法，并自动转换`receiver`实参。

### 1.类型`*T`方法集

一个指向自定义类型的值的指针，它的方法集由该类型定义的所有方法组成，无论这些方法接受的是一个值还是一个指针。

如果在指针上调用一个接受值的方法，Go语言会聪明地将该指针解引用，并将指针所指的底层值作为方法的接收者。

类型`*T`⽅法集包含全部`receiver T + *T`⽅法：

```
type Person struct{
  name string
  sex byte
  age int
}

// 指针作为接收者，引用语义
func (p *Person) SetInfoPointer(){
  (*p).name="yoyo"
  p.sex='f'
  p.age=22
}

// 值作为接收者，值语义
func (p Person) SetInfoValue(){
  p.name="xxx"
  p.sex='m'
  p.age=33
}

func main() {
  // p为指针类型
  var p*Person = &Person{"mike",'m',18}
  p.SetInfoPointer() // func (p)SetInfoPointer()

  p.SetInfoValue() // func (*p)SetInfoValue()
  (*p).SetInfoValue() // func (*p)SetInfoValue()
}
```

### 2. 类型`T`方法集

一个自定义类型值的方法集则由为该类型定义的接收者类型为值类型的方法组成，但是不包含那些接收者类型为指针的方法。

但这种限制通常并不像这里所说的那样，因为如果我们只有一个值，仍然可以调用一个接收者为指针类型的方法，这可以借助于Go语言传值的地址能力实现。

```
package main

import "fmt"

type Student struct {
  name string
  age int
}

// 指针作为接收者 引用语义
func (s *Student) SetStuPointer() {
  s.name = "Bob"
  s.age = 18
}

// 值作为接收者 值语义
func (s Student) SetStuValue() {
  s.name = "Peter"
  s.age = 18
}

func main() {
  // 指针作为接收者，引用语义
  s1 := Student{"Miller", 18} // 初始化
  fmt.Println("函数调用前 = ", s1) // 函数调用前 = {Miller 18}
  (&s1).SetStuPointer()
  fmt.Println("函数调用后 = ", s1) // 函数调用后 = {Bob 18}

  fmt.Println("==========================")

  s2 := Student{"mike", 18} // 初始化
  //值 作为接收者，值语义
  fmt.Println("函数调用前 = ", s2) // 函数调用前 = {mike 18}
  s2.SetStuValue() 
  fmt.Println("函数调用后 = ", s2) // 函数调用后 = {mike 18}
}
// 总结 : (引用语义:会改变结构体内容) (值语义:不会改变结构体内容)
```

## 五、 匿名字段

### 1. 方法的继承

如果匿名字段实现了一个方法，那么包含这个匿名字段的struct也能调用该方法。

```
type Person struct {
  name string
  sex byte
  age int
}

//Person定义了方法
func (p *Person) PrintInfo() {
  fmt.Printf("%s,%c,%d\n",p.name,p.sex,p.age)
}

type Student struct {
  Person//匿名字段，那么Student包含了Person的所有字段
  id int
  addr string
}

func main() {
  p := Person{"mike",'m',18}
  p.PrintInfo()

  s := Student{Person{"yoyo",'f',20},2,"sz"}
  s.PrintInfo()
}
```

也就是说我用`student`继承了`person`那么我就拥有了`person`的一切不管是字段，还是方法，我都能调用。

### 2. 方法的重写

```
type Person struct {
  name string
  sex byte
  age int
}
//Person定义了方法
func (p *Person) PrintInfo() {
  fmt.Printf("Person:%s,%c,%d\n",p.name,p.sex,p.age)
}
type Student struct {
  Person//匿名字段，那么Student包含了Person的所有字段
  id int
  addr string
}
//Student定义了方法
func (s *Student) PrintInfo() {
  fmt.Printf("Student：%s,%c,%d\n",s.name,s.sex,s.age)
}

func main() {
  p:=Person{"mike",'m',18}
  p.PrintInfo() //Person:mike,m,18
  s:=Student{Person{"yoyo",'f',20},2,"sz"}
  s.PrintInfo() //Student：yoyo,f,20
  s.Person.PrintInfo() //Person:yoyo,f,20
}
```

也就是说我调用了Person的方法，但是我觉得这个方法不行，然后我自己又重新写了个方法，最后调用student方法的时候就只会调用我这个方法，而不会调用person的方法了

## 六、 方法值和方法表达式

类似于我们可以对函数进行赋值和传递一样，方法也可以进行赋值和传递。

根据调用者不同，方法分为两种表现形式：方法值和方法表达式。两者都可像普通函数那样赋值和传参，区别在于方法值绑定实例，⽽方法表达式则须显式传参。

### 1. 方法值

```
type Person struct{
    name string
    sex byte
    age int
}
func (p *Person) PrintInfoPointer() {
    fmt.Printf("%p,%v\n",p,p)
}
func (p Person) PrintInfoValue(){
    fmt.Printf("%p,%v\n",&p,p)
}
//上面是定义的方法
func main() {
    p:=Person{"mike",'m',18}
    p.PrintInfoPointer()    //0xc0420023e0,&{mike 109 18}
    pFunc1:=p.PrintInfoPointer    //方法值，隐式传递 receiver
    pFunc1()    //0xc0420023e0,&{mike 109 18}
    pFunc2:=p.PrintInfoValue
    pFunc2()    //0xc042048420,{mike 109 18}
}
```

### 2. 方法表达式

```
type Person struct {
    name string
    sex byte
    age int
}
func (p *Person) PrintInfoPointer() {
    fmt.Printf("%p,%v\n",p,p)
}
func (p Person) PrintInfoValue() {
    fmt.Printf("%p,%v\n",&p,p)
}
//上面是定义的方法
func main() {
    p:=Person{"mike",'m',18}
    p.PrintInfoPointer()//0xc0420023e0,&{mike 109 18}
    //方法表达式，须显式传参
    //func pFunc1 (p *Person))
    pFunc1:=(*Person).PrintInfoPointer
    pFunc1(&p)    //0xc0420023e0,&{mike 109 18}
    pFunc2:=Person.PrintInfoValue
    pFunc2(p)    //0xc042002460,{mike 109 18}
}
```

# 接口

在讲解具体的接口之前，先看如下问题。

使用面向对象的方式，设计一个加减的计算器

代码如下：

```
package main

import "fmt"

//父类，这是结构体
type Operate struct {
    num1 int
    num2 int
}

//加法子类，这是结构体

type Add struct {
    Operate
}

//减法子类，这是结构体

type Sub struct {
    Operate
}

//加法子类的方法
func (a *Add) Result() int {
    return a.num1 + a.num2
}
可以看到ADD里面是用父类结构体的，然后直接返回num1+num2就行了
//减法子类的方法
func (s *Sub) Result() int {
    return s.num1 - s.num2
}
可以看到Sub里面是用父类结构体的，然后直接返回num1-num2就行了
//方法调用
func main0201() {
    //创建加法对象
    //var a Add
    //a.num1 = 10
    //a.num2 = 20
    //v := a.Result()
    //fmt.Println(v)
//可以看到调用起来还是很简单的，直接给父类结构体的属性赋值，然后调用加法的方法就行。
    //创建减法对象
    var s Sub
    s.num1 = 10
    s.num2 = 20
    v := s.Result()
    fmt.Println(v)
}
//可以看到调用起来还是很简单的，直接给父类结构体的属性赋值，然后调用减法的方法就行
```

以上实现非常简单，但是有个问题，在`main()`函数中，当我们想使用减法操作时，创建减法类的对象，调用其对应的减法的方法。但是，有一天，系统需求发生了变化，要求使用加法，不再使用减法，那么需要对`main()`函数中的代码，做大量的修改。将原有的代码注释掉，创建加法的类对象，调用其对应的加法的方法。有没有一种方法，让`main()`函数，只修改很少的代码就可以解决该问题呢？有，要用到接下来给大家讲解的接口的知识点。

## 一、 什么是接口

接口就是一种规范与标准，在生活中经常见接口，例如:笔记本电脑的USB接口，可以将任何厂商生产的鼠标与键盘，与电脑进行链接。为什么呢？原因就是，USB接口将规范和标准制定好后，各个生产厂商可以按照该标准生产鼠标和键盘就可以了。

在程序开发中，接口只是规定了要做哪些事情，干什么。具体怎么做，接口是不管的。这和生活中接口的案例也很相似，例如：USB接口，只是规定了标准，但是不关心具体鼠标与键盘是怎样按照标准生产的.

在企业开发中，如果一个项目比较庞大，那么就需要一个能理清所有业务的架构师来定义一些主要的接口，这些接口告诉开发人员你需要实现那些功能。

## 二、 接口定义

接口定义的语法如下：

```
//先定义接口 一般以er结尾 根据接口实现功能
type Humaner interface {
  //方法 方法的声明
  sayhi()
}
```

怎样具体实现接口中定义的方法呢？

```
//Student的结构体
type student11 struct {
    name  string
    age   int
    score int
}
//Student的打印方法
func (s *student11)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}
//teacher11的结构体
type teacher11 struct {
    name    string
    age     int
    subject string
}
//teacher11的方法
func (t *teacher11)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的学科是%s\n",t.name,t.age,t.subject)
}
```

具体的调用如下：

```
func main() {
    //接口是一种数据类型 可以接收满足对象的信息
    //接口是虚的  方法是实的
    //接口定义规则  方法实现规则
    //接口定义的规则  在方法中必须有定义的实现
    var h Humaner

    stu := student11{"小明",18,98}
    //stu.sayhi()
    //将对象信息赋值给接口类型变量
    h = &stu
    h.sayhi()
//直接将Student的对象赋值给了h接口，然后就能实现方法的调用
    tea := teacher11{"老王",28,"物理"}
    //tea.sayhi()
    //将对象赋值给接口 必须满足接口中的方法的声明格式
    h = &tea
    h.sayhi()
}
```

只要类(结构体)实现对应的接口，那么根据该类创建的对象，可以赋值给对应的接口类型。

接口的命名习惯以er结尾。

## 三、 多态

接口有什么好处呢？实现多态。

多态就是同一个接口，使用不同的实例而执行不同操作

所谓多态指的是多种表现形式，如下图所示：

![import13.png](https://www.zlkt.net/media/course/Lkawp4MMHy6DDvUnJoZ9XY.png)

使用接口实现多态的方式如下：

```
package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner1 interface {
    //方法  方法的声明
    sayhi()

}

//student12的结构体
type student12 struct {
    name  string
    age   int
    score int
}
//student12的方法
func (s *student12)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}
//teacher12的结构体
type teacher12 struct {
    name    string
    age     int
    subject string
}
//teacher12的方法
func (t *teacher12)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的学科是%s\n",t.name,t.age,t.subject)
}

//多态的实现
//将接口作为函数参数  实现多态
func sayhello(h Humaner1)  {
    h.sayhi()
}

func main() {

    stu := student12{"小明",18,98}
    //调用多态函数
    sayhello(&stu)

    tea := teacher12{"老王",28,"Go"}
    sayhello(&tea)
}
```

关于接口的定义，以及使用接口实现多态，大家都比较熟悉了，但是多态有什么好处呢？现在还是以开始提出的计算器案例给大家讲解一下。

## 四、多态案例

使用多态的功能，实现一个加减计算器。完整代码如下：

```
package main

import "fmt"

//定义接口
type Opter interface {
    //方法声明
    Result() int
}

//父类
type Operate struct {
    num1 int
    num2 int
}
//加法子类
type Add struct {
    Operate
}

//加法子类的方法
func (a *Add) Result() int {
    return a.num1 + a.num2
}

//减法子类
type Sub struct {
    Operate
}

//减法子类的方法
func (s *Sub) Result() int {
    return s.num1 - s.num2
}

//创建一个类负责对象创建
//工厂类
type Factory struct {
}

func (f *Factory) Result(num1 int, num2 int, ch string) {
    switch ch {
    case "+":
        var a Add
        a.num1 = num1
        a.num2 = num2
        Result(&a)
    case "-":
        var s Sub
        s.num1 = num1
        s.num2 = num2
        Result(&s)

    }
}

//通过设计模式调用
func main() {
    //创建工厂对象
    var f Factory
    f.Result(10, 20, "+")
}
```

## 四、 接口继承与转换（了解）

接口也可以实现继承:

```
package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner2 interface {   //子集
    //方法  方法的声明
    sayhi()

}

type Personer interface {  //超集
    Humaner2   //继承sayhi()

    sing(string)
}

type student13 struct {
    name  string
    age   int
    score int
}

func (s *student13)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

func (s *student13)sing(name string)  {
    fmt.Println("我为大家唱首歌",name)
}

func main() {
    //接口类型变量定义
    var h Humaner2
    var stu student13 = student13{"小吴",18,59}
    h = &stu
    h.sayhi()

    //接口类型变量定义
    var p Personer
    p = &stu
    p.sayhi()
    p.sing("大碗面")
}
```

接口继承后，可以实现“超集”接口转换“子集”接口，代码如下：

```
package main

import "fmt"

//先定义接口  一般以er结尾  根据接口实现功能
type Humaner2 interface {   //子集
    //方法  方法的声明
    sayhi()

}

type Personer interface {  //超集
    Humaner2   //继承sayhi()

    sing(string)
}

type student13 struct {
    name  string
    age   int
    score int
}

func (s *student13)sayhi()  {
    fmt.Printf("大家好，我是%s,今年%d岁，我的成绩%d分\n",s.name,s.age,s.score)
}

func (s *student13)sing(name string)  {
    fmt.Println("我为大家唱首歌",name)
}

func main()  {
    //接口类型变量定义
    var h Humaner2  //子集
    var p Personer    //超集
    var stu student13 = student13{"小吴",18,59}

    p = &stu
    //将一个接口赋值给另一个接口
    //超集中包含所有子集的方法
    h = p  //ok

    h.sayhi()

    //子集不包含超集
    //不能将子集赋值给超集
    //p = h  //err
    //p.sayhi()
    //p.sing("大碗面")
}
```

## 五、 空接口

空接口(interface{})不包含任何的方法，正因为如此，所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。例如：

```
var i interface{}
//接口类型可以接收任意类型的数据
//fmt.Println(i)
fmt.Printf("%T\n",i)
i = 10
fmt.Println(i)
fmt.Printf("%T\n",i)
```

当函数可以接受任意的对象实例时，我们会将其声明为`interface{}`，最典型的例子是标准库fmt中PrintXXX系列的函数，例如：

```
func Printf(fmt string, args ...interface{})
func Println(args ...interface{})
```

如果自己定义函数，可以如下：

```
func Test(arg ...interface{}) {

}
```

`Test()`函数可以接收任意个数，任意类型的参数。

# 类型查询

我们知道`interface`的变量里面可以存储任意类型的数值（该类型实现了interface）。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：

- comma-ok断言
- switch测试

## 一、comma-ok断言

Go语言里面有一个语法，可以直接判断是否是该类型的变量：`value, ok = element.(T)`，这里`value`就是变量的值，`ok`是一个`bool`类型，`element`是`interface`变量，`T`是断言的类型。

如果`element`里面确实存储了`T`类型的数值，那么`ok`返回`true`，否则返回`false`。

```
var i []interface{}

i = append(i, 10, 3.14, "aaa", demo15)

for _, v := range i {
    if data, ok := v.(int); ok {
        fmt.Println("整型数据：", data)
    } else if data, ok := v.(float64); ok {
        fmt.Println("浮点型数据：", data)
    } else if data, ok := v.(string); ok {
        fmt.Println("字符串数据：", data)
    } else if data, ok := v.(func()); ok {
        //函数调用
        data()
    }
}
```

如果这个i中有`v.(int)`也就是int类型的数值就返回打印出来。

## 二、switch测试

```
var i []interface{}

i = append(i, 10, 3.14, "aaa", demo15)

for _,data := range i{
    switch value:=data.(type) {
    case int:
        fmt.Println("整型",value)
    case float64:
        fmt.Println("浮点型",value)
    case string:
        fmt.Println("字符串",value)
    case func():
        fmt.Println("函数",value)
    }
}
```

这个也是一样的道理只不过是用了另外一种方法，`data`也就是里面的值，如果里面的值类型是`int`的话就打印出来这个值。

# error接口

Go语言引入了一个关于错误处理的标准模式，即`error`接口，它是Go语言内建的接口类型，该接口的定义如下：

```
type error interface {
	Error() string
}
```

Go语言的标准库代码包`errors`为用户提供如下方法：
![import92.png](https://www.zlkt.net/media/course/2ZoQMZ5GFXQVtrQQg5k4EY.png)

通过以上代码，可以发现`error`接口的使用是非常简单的（`error`是一个接口，该接口只声明了一个方法`Error()`，返回值是`string`类型，用以描述错误）。下面看一下基本使用：

1. 首先导包：

   ```
   import "errors"
   ```

2. 然后调用其对应的方法：
   ![import93.png](https://www.zlkt.net/media/course/EDRsmtEdUmwrWYJPxjRt3R.png)

   当然fmt包中也封装了一个专门输出错误信息的方法，如下所示：

   ![img](https://www.zlkt.net/assets/import94..png)
   ![import94..png](https://www.zlkt.net/media/course/jEpv7JYDtUe78paHNMTRfX.png)

了解完基本的语法以后，接下来使用`error`接口解决`Test()`函数被0整除的问题。如下所示：
![impor95.png](https://www.zlkt.net/media/course/rRZdCp5aGJFE4nUmtfLPRi.png)

在`Test()`函数中，判断变量`b`的取值，如果有误，返回错误信息。并且在`main()`中接收返回的错误信息，并打印出来。

这种用法是非常常见的，例如，后面讲解到文件操作时，涉及到文件的打开，如下：

![import96.png](https://www.zlkt.net/media/course/3noee9o5TtossYugUXZbxH.png)

在打开文件时，如果文件不存在，或者文件在磁盘上存储的路径写错了，都会出现异常，这时可以使用`error`记录相应的错误信息。

# panic函数

`error`返回的是一般性的错误，但是`panic函数`返回的是让程序崩溃的错误。

也就是当遇到不可恢复的错误状态的时候，如数组访问越界、空指针引用等，这些运行时错误会引起`panic异常`，在一般情况下，我们不应通过调用`panic函数`来报告普通的错误，而应该只把它作为报告致命错误的一种方式。当某些不应该发生的场景发生时，我们就应该调用`panic`。

一般而言，当`panic异常`发生时，程序会中断运行。随后，程序崩溃并输出日志信息。日志信息包括`panic value`和函数调用的堆栈跟踪信息。

当然，如果直接调用内置的`panic函数`也会引发`panic异常`，`panic函数`接受任何值作为参数。

下面给大家演示一下，直接调用`panic函数`，是否会导致程序的崩溃。

![import97.png](https://www.zlkt.net/media/course/8XY6sW66Cx9zwCuZbqtvmR.png)

错误信息如下：
![import98.png](https://www.zlkt.net/media/course/CmCfnUFdvBmDATvLRGJEmH.png)

所以，我们在实际的开发过程中并不会直接调用`panic()函数`，但是当我们编程的程序遇到致命错误时，系统会自动调用该函数来终止整个程序的运行，也就是系统内置了`panic函数`。

下面给大家演示一个数组下标越界的问题：
![import99.png](https://www.zlkt.net/media/course/mL8QWQJG8d8X2d98Y4UAHA.png)

错误信息如下：
![import100.png](https://www.zlkt.net/media/course/7ty84urkMsbJTfBhMuZNLi.png)

通过观察错误信息，发现确实是`panic异常`，导致了整个程序崩溃。

# 延迟调用defer

## 一、defer基本使用

函数定义完成后，只有调用函数才能够执行，并且一经调用立即执行。例如：

```
fmt.Println("hello")
fmt.Println("老王")
```

先输出“hello”,然后再输出“老王”。但是关键字`defer`⽤于延迟一个函数（或者当前所创建的匿名函数）的执行。注意，`defer`语句只能出现在函数的内部。

基本用法如下：

```
defer fmt.Println("hello")
fmt.Println("老王")
```

以上两行代码，输出的结果为，先输出“老王”，然后输出“hello”。

`defer`的应用场景：文件操作，先打开文件，执行读写操作，最后关闭文件。为了保证文件的关闭能够正确执行，可以使用`defer`。

### 2、 **defer执行顺序**

先看如下程序执行结果是：

```
defer fmt.Println("hello")
defer fmt.Println("老王")
defer fmt.Println("你好")
```

执行的结果是：

```
你好
老王
hello
```

总结：如果一个函数中有多个defer语句，它们会以**后进先出**的顺序执行。

如下程序执行的结果：

```
func test03(x int) {
  v := 100 / x
  fmt.Println(v)
}
defer fmt.Println("hello")
defer fmt.Println("老王")
defer test03(0)
defer fmt.Println("你好")
```

执行结果：

```
你好
老王
hello
panic: runtime error: integer divide by zero
```

即使函数或某个延迟调用发生错误，这些调用依旧会被执⾏。

## 三、`defer`与匿名函数结合使用

我们先看以下程序的执行结果：

```
a := 10
b := 20
defer func() {
  fmt.Println("匿名函数a", a)
  fmt.Println("匿名函数b", b)
}()

a = 100
b = 200
fmt.Println("main函数a", a)
fmt.Println("main函数b", b)
```

执行的结果如下：

```
main函数a 100
main函数b 200
匿名函数a 100
匿名函数b 200
```

前面讲解过，`defer`会延迟函数的执行，虽然立即调用了匿名函数，但是该匿名函数不会执行，等整个`main()`函数结束之前在去调用执行匿名函数，所以输出结果如上所示。

现在将程序做如下修改：

```
a := 10
b := 20
defer func(a,b int) {	//添加参数
  fmt.Println("匿名函数a", a)
  fmt.Println("匿名函数b", b)
}(a,b) //传参

a = 100
b = 200
fmt.Println("main函数a", a)
fmt.Println("main函数b", b)
```

该程序的执行结果如下：

```
main函数a 100
main函数b 200
匿名函数a 10
匿名函数b 20
```

从执行结果上分析，由于匿名函数前面加上了defer所以，匿名函数没有立即执行。但是问题是，程序从上开始执行当执行到匿名函数时，虽然没有立即调用执行匿名函数，但是已经完成了参数的传递。

# recover函数

运行时`panic异常`一旦被引发就会导致程序崩溃。这当然不是我们愿意看到的，因为谁也不能保证程序不会发生任何运行时错误。

Go语言为我们提供了专用于“拦截”运行时`panic`的内建函数——`recover`。它可以是当前的程序从运行时`panic`的状态中恢复并重新获得流程控制权。

**注意：recover只有在defer调用的函数中有效。**

示例如下：

```
package main

import "fmt"

func testA()  {
    fmt.Println("testA")

}

func testB(x int)  {
    //设置recover()

    //在defer调用的函数中使用recover()
    defer func() {
        //防止程序崩溃
        recover()
    }()  //匿名函数

    var a [3]int
    a[x] = 999
}

func testC()  {
    fmt.Println("testC")
}
func main() {
    testA()
    testB(3)  //发生异常 中断程序
    testC()
}
```

以上程序的运行结果如下：

```
testA
testC
```

通过以上程序，我们发现虽然`TestB()`函数会导致整个应用程序崩溃，但是由于在改函数中调用了`recover()`函数，所以整个函数并没有崩溃。虽然程序没有崩溃，但是我们也没有看到任何的提示信息，那么怎样才能够看到相应的提示信息呢？

可以直接打印`recover()`函数的返回结果，如下所示：

```
func testB(x int)  {
    //设置recover()

    //在defer调用的函数中使用recover()
    defer func() {
        //防止程序崩溃
        //recover()
        fmt.Println(recover())    //直接打印
    }()  //匿名函数

    var a [3]int
    a[x] = 999
}
```

输出结果如下：

```
testA
runtime error: index out of range
testC
```

从输出结果发现，确实打印出了相应的错误信息。

但是，如果程序没有出错，也就是数组下标没有越界，会出现什么情况呢？

```
func testA()  {
    fmt.Println("testA")

}
func testB(x int)  {
    //设置recover()

    //在defer调用的函数中使用recover()
    defer func() {
        //防止程序崩溃
        //recover()
        fmt.Println(recover())
    }()  //匿名函数

    var a [3]int
    a[x] = 999
}

func testC()  {
    fmt.Println("testC")
}
func main() {
    testA()
    testB(0)  //发生异常 中断程序
    testC()
}
```

输入的结果如下：

```
testA
<nil>
testC
```

这时输出的是空，但是我们希望程序没有错误的时候，不输出任何内容。

所以，程序修改如下：

```
func testA()  {
    fmt.Println("testA")
}

func testB(x int)  {
    //设置recover()

    //在defer调用的函数中使用recover()
    defer func() {
        //防止程序崩溃
        //recover()
        //fmt.Println(recover())

        if err := recover();err != nil {
            fmt.Println(err)
        }
    }()  //匿名函数

    var a [3]int
    a[x] = 999
}

func testC()  {
    fmt.Println("testC")
}
func main() {
    testA()
    testB(0)  //发生异常 中断程序
    testC()
}
```

通过以上代码，发现其实就是加了一层判断。这样就不会使得程序崩溃。

# 字符串处理

## 一、 字符串处理函数

我们从文件中将数据读取出来以后，很多情况下并不是直接将数据打印出来，而是要做相应的处理。例如：去掉空格等一些特殊的符号，对一些内容进行替换等。

这里就涉及到对一些字符串的处理。在对字符串进行处理时，需要借助于包“strings”

下面讲解一下常用的字符串处理函数:

### 1. Contains

```
func Contains(s, substr string) bool
```

功能：字符串s中是否包含`substr`，返回`bool`值。演示如下：

```
//查找一个字符串在另一个字符串中是否出现
str1 := "hello world"
str2 := "g"

//Contains(被查找的字符串，查找的字符串)  返回值 bool
//一般用于模糊查找
b := strings.Contains(str1,str2)
//fmt.Println(b)
if b {
    fmt.Println("找到了")
}else {
    fmt.Println("没有找到")
}
```

在使用`Contains`关键字的时候，判断`b`的结果，如果在`str1`中有`str2`的字那么就返回`true`，在判断的时候不写`true`默认就是等于`true`。

### 2. Join

```
func Join(a []string, sep string) string
```

功能：字符串链接，把`slice`通过`sep`链接起来

演示如下：

```
//字符串切片
slice := []string{"123","456","789"}
//fmt.Println(slice)
//Join
//字符串的连接
str := strings.Join(slice,"")
fmt.Println(str)
//fmt.Printf("%T\n",str)
```

结果如下：

```
123456789
```

通过`join`关键字把，`slice`里面的值通过`strings.Join(slice,"")`也就是去除""给从新赋值给了`str`最后打印出来的值就变成了`123456789`。

### 3. Index

```
func Index(s, substr string) int
```

功能：在字符串s中查找sep所在的位置，返回位置值，找不到返回-1

```
str1 := "hello world"
str2 := "e"
//查找一个字符串在另一个字符串中第一次出现的位置 返回值 int 下标 -1 找不到

i := strings.Index(str1,str2)
fmt.Println(i)
```

结果为1。
`i := strings.Index(str1,str2)`通过`index`关键字，在`str1`中查找`str2`的值，然后赋值给`i`，`e`这个值在`hello woeld`中能找到所以就会返回它的下标值，下标值是从`0`开始的，`h`是`0`，`e`就是`1`，所以结果为`1`。如果查找的是一个`g`的话找不到就会返回一个`-1`。

### 4. Repeat

```
func Repeat(s string, count int) string
```

功能：重复`s`字符串`count`次，最后返回重复的字符串。

演示如下：

```
str := "性感网友，在线取名。"
//将一个字符串重复n次
str1 := strings.Repeat(str,100)
fmt.Println(str1)
```

`str1 := strings.Repeat(str,100)`通过repeat关键字重复了str100遍，就和循环遍历str100次是一样的。

### 5. Replace

```
func Replace(s, old, new string, n int) string
```

功能：在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换

```
str := "性感网友在线取名性感性感性感性感性感"
//字符串替换 屏蔽敏感词汇
//如果替换次数小于0 表示全部替换
str1 := strings.Replace(str,"性感","**",-1)
fmt.Println(str1)
```

结果如下：

```
**网友在线取名**********
```

`str1 := strings.Replace(str,"性感","**",-1)`通过关键字`replace`把`str`中的性感替换为了`******`然后给了个`-1`也就是全部替换，当然你给其他的负数也是一样的，只要是小于`0`就全部替换，如果说是`1`的话就是替换一次，输出结果就会是：`**网友在线取名性感性感性感性感性感`。

### 6. Split

```
func Split(s, sep string) []string
```

功能：把s字符串按照`sep`分割，返回`slice`。

```
//将一个字符串按照标志位进行切割变成切片
str1 := "123456789@qq.com"
slice := strings.Split(str1,"@")
fmt.Println(slice[0])
```

结果如下：

```
123456789
```

`slice := strings.Split(str1,"@")`通过`split关`键字对`str1`进行了分割，把`@`后面的给丢弃了，留下了`@`前面的。

### 7. Trim

```
func Trim(s string, cutset string) string
```

功能：在`s`字符串的头部和尾部去除`cutset`指定的字符串。

```
str := "====a===u=ok===="
//去掉字符串头尾的内容
str1:= strings.Trim(str,"=")
fmt.Println(str1)
```

结果如下：

```
a===u=ok
```

`str1:= strings.Trim(str,"=")`通过`Trim`关键字对`str`中的`=`号做了去除头尾的处理，只要是`str`前面有`=`，或者后面有`=`都会去除。

### 8. Fields

```
func Fields(s string) []string
```

功能：去除`s`字符串的空格符，并且按照空格分割返回`slice`。

```
str := " are you ok "
//去除字符串中空格 转成切片 一般用于统计单词个数
slice := strings.Fields(str)
fmt.Println(slice)
```

`slice := strings.Fields(str)`通过Fields关键字对str中的值进行了空格去除。

## 二、字符串转换

通过以上的讲解，发现字符串的处理是非常重要的，GO语言也提供了字符串与其它类型之间相互转换的函数。相应的字符串转换函数都在“strconv”包。

### 1. Format系列函数：

`Format`系列函数把其他类型的转换为字符串。

```
//将其他类型转成字符串 Format
b := false
str := strconv.FormatBool(true)
fmt.Println(str)
fmt.Printf("%T\n",str)

str := strconv.FormatInt(120,10) //计算机中进制 可以表示2-36 2 8 10 16
fmt.Println(str)
// 'f'打印方式 以小数方式 4 指小数位数 64 以float64处理
str:= strconv.FormatFloat(3.14159,'f',4,64)
fmt.Println(str)

str := strconv.Itoa(123)
fmt.Println(str)
```

以上代码只要是通过`Format`关键字就能全部转换为字符串类型输出。

### 2. Parse

`Parse`系列函数把字符串转换为其他类型：

```
//字符串转成其他类型 Parse
b,err := strconv.ParseBool("true")
if err!=nil {
  fmt.Println("类型转换出错")
}else {
  fmt.Println(b)
  fmt.Printf("%T\n",b)
}

v,err := strconv.ParseInt("abc",16,64)
fmt.Println(v,err)

v,_ := strconv.ParseFloat("3.14159",64)
fmt.Println(v)

v,_:=strconv.Atoi("123")
fmt.Println(v)
```

### 3. Append

`Append`系列函数将整数等转换为字符串后，添加到现有的字节数组中。

```
slice := make([]byte,0,1024)
//将其他类型转成字符串 添加到字符切片里

slice = strconv.AppendBool(slice,false)
slice = strconv.AppendInt(slice,123,2)
slice = strconv.AppendFloat(slice,3.14159,'f',4,64)
slice = strconv.AppendQuote(slice,"hello")
fmt.Println(string(slice))
```

对应的结果是：

```
false11110113.1416"hello"
```

通过`Append`关键字把其他类型的值转换成字符串后在拼接到一起，赋值给`slice`，当然赋值给谁你可以自己起名字，所以最后的打印结果是上面这样的。

# 创建文件

将数据存储到文件之前，先要创建文件。GO语言中提供了一个`Create()`函数专门创建文件。

该函数在创建文件时，首先会判断要创建的文件是否存在，如果不存在，则创建，如果存在，会先将文件中已有的数据清空。

同时，当文件创建成功后，该文件会默认的打开，所以不用在执行打开操作，可以直接向该文件中写入数据。

创建文件的步骤：

1. 导入“os”包，创建文件，读写文件的函数都在该包。
2. 指定创建的文件存放路径以及文件名。
3. 执行`Create()`函数，进行文件创建。
4. 关闭文件。

具体代码如下：

```
package main

import (
    "fmt"
    "os"
)

func main() {
    //os.Create(文件名) 文件名  可以写绝对路径和相对路径
    //返回值  文件指针 错误信息
    fp,err := os.Create("./a.txt")
    if err!=nil{
        //文件创建失败
        /*
        1.路径不存在
        2.文件权限
        3.程序打开文件上限
         */
        fmt.Println("文件创建失败")
        return
    }

    //读写文件

    defer fp.Close()

    //关闭文件
    //如果打开文件不关闭 造成内存的浪费  程序打开文件的上限
    //fp.Close()
}
```

执行以上代码后，可以在程序文件存放的目录中，看到有一个`a.txt`的文件。

注意：在创建的文件时，注意需要判断是否出现异常，同时要注意`defer`的应用。

# 写入数据

文件打开以后，可以向文件中写数据，可以使用WriteString( )方法。

```
//\反斜杠 转义字符
//在写路径时可以使用/正斜杠代替\反斜杠
fp,err := os.Create("D:/a.txt")
if err!=nil{
    //文件创建失败
    /*
    1.路径不存在
    2.文件权限
    3.程序打开文件上限
     */
    fmt.Println("文件创建失败")
    return
}

//写文件

//\n不会换行  原因 在windows文本文件中换行\r\n  回车  在linux中换行\n
fp.WriteString("hello world\r\n")
fp.WriteString("性感荷官在线发牌")


defer fp.Close()

//关闭文件
//如果打开文件不关闭 造成内存的浪费  程序打开文件的上限
//fp.Close()
```

`WriteString()`方法默认返回两个参数：

```
count,err1 := fp.WriteString("性感老王在线授课")

if err1!=nil {
    fmt.Println("写入文件失败")
    return
}else {
    fmt.Println(count)
}
```

第一个参数，指的是写入文件的数据长度，第二个参数记录的是错误信息。

`WriteString()`方法默认写到文件中的数据是不换行的。如果想换行，可以采用如下的方式：

```
//\n不会换行 原因 在windows文本文件中换行\r\n 回车 在linux中换行\n
fp.WriteString("hello world\r\n")
fp.WriteString("性感荷官在线发牌")
```

除了使用`WriteString()`函数向文件中写入数据以外，还可以使用`Write()`函数，如下所示：

```
fp,err := os.Create("D:/a.txt")
if err!=nil{
    //文件创建失败
    /*
    1.路径不存在
    2.文件权限
    3.程序打开文件上限
     */
    fmt.Println("文件创建失败")
    return
}

//写操作
//slice := []byte{'h','e','l','l','o'}
//count,err1 := fp.Write(slice)
count,err1 := fp.Write([]byte("性感老王在线授课"))

if err1!=nil {
    fmt.Println("写入文件失败")
    return
}else {
    fmt.Println(count)
}

defer fp.Close()
```

在这里要注意的是，使用`Write()`函数写数据时，参数为字节切片，所以需要将字符串转换成字节切片。该方法返回的也是写入文件数据的长度。

第三种写入的方式使用`WriteAt()`函数，在指定的位置写入数据：

```
fp,err := os.Create("D:/a.txt")
if err!=nil{
    //文件创建失败
    /*
    1.路径不存在
    2.文件权限
    3.程序打开文件上限
     */
    fmt.Println("文件创建失败")
    return
}

//写操作
//获取光标流位置'
//获取文件起始到结尾有多少个字符
//count,_:=fp.Seek(0,os.SEEK_END)
count,_:=fp.Seek(0,io.SeekEnd)
fmt.Println(count)
//指定位置写入
fp.WriteAt([]byte("hello world"),count)
fp.WriteAt([]byte("hahaha"),0)
fp.WriteAt([]byte("秀儿"),19)

defer fp.Close()
```

以上程序中`Seek()`函数返回值存储到变量`n`中，值为文件末尾的位置。`WriteAt()`也返回的是写入的数据长度。

以上就是我们常用的关于向文件中写入数据的方式，但是有同学可能有疑问，每次向文件中写入数据之前，都是先执行了，`Create()`这个函数，而这个函数的作用前面我们也已经说过。有两个作用：

- 第一：创建新文件。
- 第二：如果所创建的文件已经存在，会删除掉文件中存储的数据。

那么，现在怎样向已有的文件中追加数据呢？如果要解决这个问题，那么大家一定要注意的就是，对已经存在的文件不能再执行`Create()`，而是要执行`OpenFile()`。如下所示：

```
//os.Open  只读方式打开
//fp,err := os.Open("D:/a.txt")

//os.OpenFile(文件名，打开方式，打开权限)
fp,err := os.OpenFile("D:/a.txt",os.O_RDWR,6)
if err!=nil {
    fmt.Println("打开文件失败")
}

fp.WriteString("hello")
fp.WriteAt([]byte("hello"),25)


defer fp.Close()
```

`OpenFile()`这个函数有三个参数，第一个参数表示打开文件的路径，第二个参数表示模式，常见的模式有

`O_RDONLY(只读模式)`，`O_WRONLY(只写模式)`，`O_RDWR(可读可写模式)`，`O_APPEND(追加模式)`。

第三个参数，表示权限，取值范围（0-7）

表示如下：

- 0：没有任何权限
- 1：执行权限(如果是可执行文件，是可以运行的)
- 2：写权限
- 3: 写权限与执行权限
- 4：读权限
- 5: 读权限与执行权限
- 6: 读权限与写权限
- 7: 读权限，写权限，执行权限

# 读取文件

## 一、 Read 读取文件

如果文件已经存在，并且也已经有数据了，那么可以直接读取该文件中的内容。

读取文件的基本流程如下：

1. 打开要读取的文件
2. 对文件进行读取
3. 关闭文件

在向文件中写数据的时候，使用的是`Write`，那么读取文件中的数据，使用的是`Read`。

关于`Read()`函数的使用如下：

```
package main

import (
    "fmt"
    "io"
    "os"
)
func main() {
    //打开文件
    fp, err := os.Open("D:/a.txt")
    if err != nil {
        fmt.Println("err=", err)
        return
    }

    buf := make([]byte, 1024*2) //2k大小
    //n代表从文件读取内容的长度
    n, err1 := fp.Read(buf)
    if err1 != nil && err1 != io.EOF {
        fmt.Println("err1=", err1)
        return
    }
    fmt.Println("buf=", string(buf[:n]))

    //关闭文件
    defer fp.Close()
}
```

`Open()`是打开文件，与`OpenFile()`的区别是，`Open()`只有读的权限。

在使用`Read()`函数读取文件中的内容时，需要一个切片类型，而定义切片时类型为字符数组，将文件中的内容保存在切片中，同时除了对其判断是否出错时以外，还要判断是否到文件末尾（这里需要导入`io`包）。

`Read()`函数返回的是从文件中读取的数据的长度。最后，输出切片中存储的文件数据，注意，读取的是从最开始到整个数据长度，因为有可能存储到切片中的数据达不到切片的总长度（也是切片时2k,但是从文件中读取的数据有可能只有1k）。

## 二、 按行读取

上面我们是将文件的内容全部读取出来，然后存放在切片中，我们也可以每次只读取一行数据。

这需要用到`bufio`包中的`ReadBytes`函数。具体如下：

### 1. 打开文件

```
fp, err := os.Open("D:/a.txt")
if err != nil {
    fmt.Println("打开文件失败", err)
    return
}
```

### 2. 创建缓冲区

在使用`ReadBytes()`函数读取数据时，需要用到缓冲区，所谓缓冲区就是存储数据的区域，也就是先将从文件中读取的数据存储在该区域内，然后在将区域中的数据取出来，写到磁盘上。提供缓冲区的原因是：

为了缓和`CPU`与磁盘设备之间速度不匹配矛盾。文件缓冲区是用以暂时存放读写期间的文件数据而在内存区预留的一定空间。

```
//创建文件缓冲区
r := bufio.NewReader(fp)
```

### 3. 循环读取文件中的内容，直到文件末尾位置。

```
for {
    //遇到'\n'结束读取，但是'\n'也读取进入
    buf,err := r.ReadBytes('\n')
    fmt.Println("buf = ",string(buf))
    if err != nil {
        if err == io.EOF {
            break
        }
        fmt.Println("err=",err)
    }
}
```

在使用`ReadBytes()`函数时，传递的参数是`\n`，表示遇到`\n`就结束，所以使用了死循环（每循环一次，读取一行数据），只有到文件末尾了，才退出整个循环。最后，将读取的数据打印出来，注意`ReadBytes()`返回的是字节切片，所以在打印时要转换成字符串。

### 4. 最后关闭文件

```
//关闭文件
defer fp.Close()
```

现在我们已经完成了文件的创建，读取，以及将数据保存到文件的操作，在对文件操作时，我们需要指定文件的路径。

**关于路径，有两种情况：**

- 第一：相对路径，所谓相对路径指的是文件相对于应用程序的路径。例如：上面我们一只使用的a.txt,这个文件，该文件存放的位置与可执行文件存储的路径是一样的。
- 第二：绝对路径：指的是通过给定的这个路径直接能在我的电脑中找到这个文件。例如：D:\Info.txt,
- **建议我们以后在开发中使用相对路径**

# 文件操作案例

文件拷贝，将已有的文件复制一份，同时重新命名。

基本的思路：

1. 让用户输入要拷贝的文件的名称(源文件)以及目的文件的名称
2. 创建目的文件
3. 打开源文件，并且读取该文件中的内容
4. 将从源文件中读取的内容写到目的文件中。

```
var srcFileName string
var dstFileName string
fmt.Printf("请输入源文件名称:")
fmt.Scan(&srcFileName)
fmt.Println("请输入目的文件名称:")
fmt.Scan(&dstFileName)
if srcFileName == dstFileName {
    fmt.Println("源文件和目的文件名字不能相同")
    return
}
//只读方式打开源文件
sF,err1 := os.Open(srcFileName)
if err1 != nil {
    fmt.Println("err1=",err1)
    return
}
//新建目的文件
dF,err2 := os.Create(dstFileName)
if err2 != nil{
    fmt.Println("err2=",err2)
    return
}
//操作完毕，需要关闭文件
defer sF.Close()
defer dF.Close()
//核心处理，从源文件读取内容，往目的文件写，读多少写多少
buf := make([]byte,4*1024)//4k大小临时缓冲区
for{
    n,err := sF.Read(buf)//从源文件读取内容,每次读取一部分
    if err != nil{
        fmt.Println("err=",err)
        if err == io.EOF{//文件读取完毕
            break
        }
    }
    //往目的文件写，读多少写多少
    dF.Write(buf[:n])
}
```
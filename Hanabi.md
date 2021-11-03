<iframe frameborder="no" border="0" marginwidth="0" marginheight="0" width="420" height="86" src="//music.163.com/outchain/player?type=2&id=28578763&auto=1&height=66"></iframe>

[toc]



# Learning in Computer Science

*Ctrl+点击链接*



# Math



# DataBase

## MySQL

## PostgreSQL

## Redis

## Network

## HTTP

### Get

#### URL传参

https://blog.csdn.net/pbhLOVEpp/article/details/80205026



# Data Structure



# Back-end

## Python

​					——内容来源于Python Crash Course(Python编程 ——从入门到实践)第二版。



## Java

[How2J.com -- Java Learning](https://how2j.cn/)

[Alibaba - easyExcel -- Excel相关操作插件-- for Java](https://alibaba-easyexcel.github.io/)

[HuTools -- docs](https://www.hutool.cn/docs/#/)

### 泛型

[Java 泛型详解](https://www.cnblogs.com/lbsjs/p/7367908.html)

### WebService

#### SOAP，rpc/encoding

###### 1.获取WebService接口服务提供商描述文件

WSDL文件：通过URL或者文件方式取得WSDL服务描述文件

###### 2.通过Axis 1.4版本自动生成Client代码

java -Djava.ext.dirs=${lib的目录，绝对路径} org.apache.axis.wsdl.WSDL2Java -o${代码输出路径} -p${package名} ${wsdl的路径(可以是绝对路径或者url)}

示例1.

java -Djava.ext.dirs=D:\Softwares\axis-1_4\axis1.4\lib org.apache.axis.wsdl.WSDL2Java -p wasu http://125.210.115.52:18083/ccps_hz/services/orderService?wsdl

示例2.

java -Djava.ext.dirs=D:\Softwares\axis-1_4\axis1.4\lib org.apache.axis.wsdl.WSDL2Java -oD:\Softwares\axis-1_4 -p packageName D:\Softwares\axis-1_4\orderService.wsdl

###### 3.对WebService服务进行封装

创建Locator

创建stub，new传参数1.URL 2.axis.client.service()

调用stub方法，发送数据

示例代码如下：

```Java
//创建Locator
WoOrderSOAPServiceServiceLocator locator = new WoOrderSOAPServiceServiceLocator();

        try{
            //创建stub
            OrderServiceSoapBindingStub stub = new OrderServiceSoapBindingStub(new URL(locator.getorderServiceAddress()), new org.apache.axis.client.Service());

            //发送WebService请求，获取请求结果
            JSONObject result = JSON.parseObject(stub.ccpsRequest("10019",JSON.toJSONString(param)));
            logger.info(result.toJSONString());
            //返回WebService调用结果
            return Result.success(result);

        } catch (MalformedURLException | RemoteException e){
            e.printStackTrace();
        }
```



### Easy Excel

[Easy Excel 文档 -- 语雀](https://www.yuque.com/easyexcel)

#### 注解判断行、列的宽高

@ContentRowHeight(10) ---- 内容行高

@HeadRowHeight(20) ---- 头高

@ColumnWidth(25) ---- 列宽

#### 自动列宽

```java
EasyExcel.write(response.getOutputStream(), DownloadData.class)
                .registerWriteHandler(new LongestMatchColumnWidthStyleStrategy())
                .sheet("xx")
                .doWrite(data());
```

需要注册写处理器，调用LongestMatchColumnWidthSttyleStrategy()方法



## Flutter

[Flutter-go  -- Flutter developer demo & docs](https://github.com/alibaba/flutter-go)



## Golang

### Go语言入门

[GolangBot](https://golangbot.com)

[Golang from zero to hero](https://milapneupane.com.np/2019/07/06/learning-golang-from-zero-to-hero/)

[一篇文章上手Go语言](https://mp.weixin.qq.com/s/BZZYyPMcVayyPyt5ysVSbw)

[写给新手的 Go 开发指南](https://liujiacai.net/blog/2019/07/17/hello-golang/)

[Golang学习路线](https://www.topgoer.com/%E5%BC%80%E6%BA%90/go%E5%AD%A6%E4%B9%A0%E7%BA%BF%E8%B7%AF%E5%9B%BE.html)

[Go语言爱好者周刊 Vol.1](https://studygolang.com/topics/9700)

### Routine & Channel

[Go中的goroutine&channel](https://blog.csdn.net/qq_37463791/article/details/106919860)



[Five things makes go fast -- dave.cheney.net](https://dave.cheney.net/2014/06/07/five-things-that-make-go-fast)

[The Zen of Go](https://the-zen-of-go.netlify.app/)

[Go中goroutine的理解 -- Go语言中文网](https://studygolang.com/articles/19654)

[Go语言中文网](https://studygolang.com/)

[The Ultimate Go Study Guide](https://github.com/hoanhan101/ultimate-go)

### Go 编写命令行程序(cli)



---

## Interview & Learning Path

[Awesome Interview](https://github.com/Awesome-Interview/Awesome-interview#%E5%90%8E%E7%AB%AF)

[Coding Interview University](https://github.com/jwasham/coding-interview-university/blob/main/translations/README-cn.md)



# Front-end

## Vue

### About Vue

## HTML

## CSS



# Data blank

[Learn Git Branching](https://learngitbranching.js.org/?locale=zh_CN)

[Free E-Books, in web develop/ system management/ coding langs/ SQL/ software developing/ AI / math / others](https://github.com/ruanyf/free-books)




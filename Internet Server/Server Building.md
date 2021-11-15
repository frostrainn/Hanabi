[toc]

### 1.防火墙开放端口

防火墙开放端口，常用端口：
+ SSH：22
+ HTTP：80
+ HTTPS：443

### 2.JDK的安装

```
tar -zxvf jdk-8u74-linux-x64.tar.gz //tar解压

cd /etc
vi profile

export JAVA_HOME=/opt/jdk1.8.0_74   //dir of JAVA_HOME
export PATH=$JAVA_HOME/bin:$PATH
export CLASSPATH=.

source /etc/profile //刷新环境变量配置
```

### 3.安装Tomcat

下载tar.gz压缩包解压后，进入bin目录，通过命令`sh startup.sh`启动tomcat服务器.

### 4.Nginx的安装

####4.1 Nginx 编译安装

####4.2 Nginx 解压安装
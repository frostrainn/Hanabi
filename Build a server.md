[toc]

### 1.防火墙开放端口

### 2.安装 JDK

tar压缩包解压`tar -zxvf jdk-8u74-linux-x64.tar.gz`

`cd /etc`

`vi profile`

```JAVA
export JAVA_HOME=/opt/jdk1.8.0_74
export PATH=$JAVA_HOME/bin:$PATH
export CLASSPATH=.
```

`source /etc/profile`

### 3.安装Tomcat
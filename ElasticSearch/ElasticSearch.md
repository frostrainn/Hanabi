[toc]

# ElasticSearch

[阮一峰的ES入门](https://www.ruanyifeng.com/blog/2017/08/elasticsearch.html)



## 前言

​		ElasticSearch是目前市面上较为流行的搜索引擎。截止笔者开始学习时，其官网最新的版本为7.16.3。

​		本文中所使用的ElasticSearch(下文简称为ES)、其官方附带的ES图形化管理工具Kibana的版本均为7.16.3。

## ElasticSearch简介

​		Elasticsearch是一个基于Lucene的搜索服务器。它提供了一个分布式多用户能力的全文搜索引擎，基于**RESTful web接口**。Elasticsearch是用**Java**语言开发的，并作为Apache许可条款下的开放源码发布，是一种流行的企业级搜索引擎。Elasticsearch用于云计算中，能够达到实时搜索，稳定，可靠，快速，安装使用方便。官方客户端在Java、.NET（C#）、PHP、Python、Apache Groovy、Ruby和许多其他语言中都是可用的。根据DB-Engines的排名显示，Elasticsearch是最受欢迎的企业搜索引擎，其次是Apache Solr，也是基于Lucene。

​		Elasticsearch 是一个分布式、高扩展、高实时的搜索与数据分析引擎。它能很方便的使大量数据具有搜索、分析和探索的能力。充分利用Elasticsearch的水平伸缩性，能使数据在生产环境变得更有价值。Elasticsearch 的实现原理主要分为以下几个步骤，首先用户将数据提交到Elasticsearch 数据库中，再通过分词控制器去将对应的语句分词，将其权重和分词结果一并存入数据，当用户搜索数据时候，再根据权重将结果排名，打分，再将返回结果呈现给用户。
Elasticsearch是与名为Logstash的数据收集和日志解析引擎以及名为Kibana的分析和可视化平台一起开发的。这三个产品被设计成一个集成解决方案，称为“Elastic Stack”（以前称为“ELK stack”）。
​		Elasticsearch可以用于搜索各种文档。它提供可扩展的搜索，具有接近实时的搜索，并支持多租户。Elasticsearch是分布式的，这意味着索引可以被分成分片，每个分片可以有0个或多个副本。每个节点托管一个或多个分片，并充当协调器将操作委托给正确的分片。再平衡和路由是自动完成的。相关数据通常存储在同一个索引中，该索引由一个或多个主分片和零个或多个复制分片组成。一旦创建了索引，就不能更改主分片的数量。
Elasticsearch使用Lucene，并试图通过JSON和Java API提供其所有特性。它支持facetting和percolating，如果新文档与注册查询匹配，这对于通知非常有用。另一个特性称为“网关”，处理索引的长期持久性；例如，在服务器崩溃的情况下，可以从网关恢复索引。Elasticsearch支持实时GET请求，适合作为NoSQL数据存储，但缺少分布式事务。

​		若想使用原生搜索引擎Lucene，需要对信息科学有许多了解，甚至需要考取信息科学学位。而ES则是对Lucene的封装后的结果，其出现大大降低了搜索引擎的门槛。ES的使用非常简单。其基于RESTful的Web接口，你可以使用任何方式与ES服务器交互——包括且不限于各种编程语言、甚至curl。

## ElasticSearch、Kibana的安装及部署

### ElasticSearch的安装、部署

​		访问[ElasticSearch官方网站](https://www.elastic.co/cn/elasticsearch/)获取最新版本的ElasticSearch搜索引擎，该软件为开箱即用。

​		**Windows11**环境下示例，下载`elasticsearch-7.16.3-windows-x86_64.zip`后解压至***无中文路径的目录***下，笔者为`D:\elasticsearch-7.16.3`。

​		进入其中的bin目录`D:\elasticsearch-7.16.3\bin`，运行`elasticsearch.bat`即可以Windows cmd方式本地启动ES服务器。

​		访问ES服务器的默认路径为`http://localhost:9200/`，在浏览器或Postman等工具访问该路径，获得返回结果则为启动成功。

​		获得如下返回结果，则为成功安装、本地部署ES服务器：

```json
{
  "name" : "MAZHE",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "klUuS-QdRFO7rcDEp1ZXpQ",
  "version" : {
    "number" : "7.16.3",
    "build_flavor" : "default",
    "build_type" : "zip",
    "build_hash" : "4e6e4eab2297e949ec994e688dad46290d018022",
    "build_date" : "2022-01-06T23:43:02.825887787Z",
    "build_snapshot" : false,
    "lucene_version" : "8.10.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

​		该返回涵盖了ES服务器的版本、集群名称等各种必要信息，***"You know, for Search"***。

### ElasticSearch部分配置

#### 可用IP配置

默认情况下，Elastic 只允许本机访问，如果需要远程访问，可以修改 Elastic 安装目录的config/elasticsearch.yml文件，去掉network.host的注释，将它的值改成0.0.0.0，然后重新启动 Elastic。

```
network.host: 0.0.0.0
```

上面代码中，设成0.0.0.0让任何人都可以访问。线上服务应当设置为具体的 IP。

### Kibana的安装、部署

​		访问[Kibana官方网站](https://www.elastic.co/cn/kibana/)获取最新版本的Kibana，该软件需要简单配置才可使用。

## ElasticSearch与Kibana的简单使用

https://www.jianshu.com/p/d48c32423789



## ElasticSearch基本概念

#### Node&Cluster

ElasticSearch本质为一个分布式的数据库，允许多台服务器协同工作，每台服务器可以运行多个 Elastic 实例。

单个 Elastic 实例称为一个节点（Node）。一组节点构成一个集群（Cluster）。

#### Index

Elastic 会索引所有字段，经过处理后写入一个反向索引（Inverted Index）。查找数据的时候，直接查找该索引。

所以，Elastic 数据管理的顶层单位就叫做 Index（索引）。它是单个数据库的同义词。每个 Index （即数据库）的名字必须是小写。

下面的命令可以查看当前节点的所有 Index。

# 附录

## ElasticSearch官方网站

https://www.elastic.co/cn/elasticsearch/

## Kibana官方网站

https://www.elastic.co/cn/kibana/


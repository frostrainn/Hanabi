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

​		默认情况下，Elastic 只允许本机访问，如果需要远程访问，可以修改 Elastic 安装目录的config/elasticsearch.yml文件，去掉network.host的注释，将它的值改成0.0.0.0，然后重新启动 Elastic。

```
network.host: 0.0.0.0
```

上面代码中，设成0.0.0.0让任何人都可以访问。线上服务应当设置为具体的 IP。

### Kibana的安装、部署

​		访问[Kibana官方网站](https://www.elastic.co/cn/kibana/)获取最新版本的Kibana，该软件需要简单配置才可使用。

## ElasticSearch与Kibana的使用

https://www.jianshu.com/p/d48c32423789

https://www.jianshu.com/p/4d65ed957e62

https://www.cnblogs.com/JimShi/p/11242657.html

https://www.tizi365.com/archives/845.html



## ElasticSearch核心概念

#### Node&Cluster

​		ElasticSearch本质为一个分布式的数据库，允许多台服务器协同工作，每台服务器可以运行多个 Elastic 实例。单个 Elastic 实例称为一个节点（Node）。一组节点构成一个集群（Cluster）。

#### Index

​		Elastic 会索引所有字段，经过处理后写入一个反向索引（Inverted Index）。查找数据的时候，直接查找该索引。

​		所以，Elastic 数据管理的顶层单位就叫做 Index（索引）。它是单个数据库的同义词。每个 Index （即数据库）的名字必须是小写。

​		下面的命令可以查看当前节点的所有 Index。

```
$ curl -X GET 'http://localhost:9200/_cat/indices?v'
```

#### Document

​		Index 里面单条的记录称为 Document（文档）。许多条 Document 构成了一个 Index。

Document 使用 JSON 格式表示，下面是一个例子。

```JSON
{
  "user": "张三",
  "title": "工程师",
  "desc": "数据库管理"
}
```

#### Type

​		Document 可以分组，比如`weather`这个 Index 里面，可以按城市分组（北京和上海），也可以按气候分组（晴天和雨天）。这种分组就叫做 Type，它是虚拟的逻辑分组，用来过滤 Document。

​		不同的 Type 应该有相似的结构（schema），举例来说，`id`字段不能在这个组是字符串，在另一个组是数值。这是与关系型数据库的表的[一个区别](https://www.elastic.co/guide/en/elasticsearch/guide/current/mapping.html)。性质完全不同的数据（比如`products`和`logs`）应该存成两个 Index，而不是一个 Index 里面的两个 Type（虽然可以做到）。

​		下面的命令可以列出每个 Index 所包含的 Type。

```
$ curl 'localhost:9200/_mapping?pretty=true'
```

​		根据[规划](https://www.elastic.co/blog/index-type-parent-child-join-now-future-in-elasticsearch)，Elastic 6.x 版只允许每个 Index 包含一个 Type，7.x 版将会彻底移除 Type。

## ElasticSearch的使用

[ElasticSearch使用教程、设计到实战](https://blog.csdn.net/ganquanzhong/article/details/108633025)

http://blog.itpub.net/29715045/viewspace-2653369/

https://www.cnblogs.com/dydxw/p/12667874.html

https://www.cnblogs.com/UUUz/p/11170833.html

https://blog.csdn.net/u014646662/article/details/89010759

### 新建和删除Index

​		新建 Index，可以直接向 Elastic 服务器发出 PUT 请求。下面的例子是新建一个名叫`weather`的 Index。

```
$ curl -X PUT 'localhost:9200/weather'
```

​		服务器返回一个 JSON 对象，里面的`acknowledged`字段表示操作成功。

```json
{
  "acknowledged":true,
  "shards_acknowledged":true
}
```

### 布尔查询(BooleanSearch)

​		布尔查询是指将多个查询条件以一定的逻辑组合在一起，进行组合查询获取同时满足结果的查询。

https://www.cnblogs.com/dangdanghepingping/p/14425402.html

### ElasticSearch Java API

```java
package com.wasu.cmp.search.biz.es;

import com.alibaba.fastjson.JSON;
//import com.wasu.cmp.common.exception.BizException;
import com.wasu.cmp.common.exception.BizException;
import com.wasu.cmp.search.SearchResultCode;
import com.wasu.cmp.search.dao.bo.AssetBo;
import com.wasu.cmp.search.vo.BaseSearchVo;
import org.elasticsearch.action.search.SearchRequest;
import org.elasticsearch.action.search.SearchResponse;
import org.elasticsearch.client.RequestOptions;
import org.elasticsearch.client.RestHighLevelClient;
import org.elasticsearch.index.query.BoolQueryBuilder;
import org.elasticsearch.index.query.QueryBuilders;
import org.elasticsearch.rest.RestStatus;
import org.elasticsearch.search.SearchHits;
import org.elasticsearch.search.builder.SearchSourceBuilder;
import org.elasticsearch.search.sort.SortOrder;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@Service
public class EsAssetQuerryService {
    static Logger logger = LoggerFactory.getLogger(EsAssetQuerryService.class);

    @Autowired
    private RestHighLevelClient client;

    private String indexName="_epg_asset";
    private float default_score=1.01f;
    /**
     * 高亮搜索
     * @return
     */
    public BaseSearchVo searcherHighlight(int site_id, String value, int pageNo, int pageSize) {

        SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();//构造搜索对象
        /**
         * 1. 搜索 builder
         */
        BoolQueryBuilder boolQuery = QueryBuilders.boolQuery();
        //boolQuery.must(QueryBuilders.wildcardQuery("name", value+"*")); //模糊匹配
        /**
         * matchQuery：会将搜索词分词，再与目标查询字段进行匹配，若分词中的任意一个词与目标字段匹配上，则可查询到。
         * termQuery：不会对搜索词进行分词处理，而是作为一个整体与目标字段进行匹配，若完全匹配，则可查询到。
         * matchPhraseQuery 短语匹配
         */
        boolQuery.should(QueryBuilders.matchQuery("title.pinyin", value)); //分词匹配
        boolQuery.should(QueryBuilders.matchQuery("singer_name", value)); //分词匹配
        boolQuery.must(QueryBuilders.termQuery("status",1)); //字段匹配

        //boolQuery.filter(QueryBuilders.rangeQuery("@timestamp").gte(start).lte(end));  //范围判断

        searchSourceBuilder.query(boolQuery)
                .sort("_score", SortOrder.DESC);
        searchSourceBuilder.from(pageNo - 1);
        searchSourceBuilder.size(pageSize);
        searchSourceBuilder.minScore(default_score);
        /**
         * 3.筛选字段
         */
        String[] includeFields = new String[] {"asset_id", "mv_id","title","mv_name","status"};
        String[] excludeFields = new String[] {};
        searchSourceBuilder.fetchSource(includeFields, excludeFields);

        /**
         * 4. 搜索调用
         */
        SearchRequest searchRequest = new SearchRequest(site_id+indexName);//创建查询请求对象
        searchRequest.source(searchSourceBuilder);//设置searchSourceBuilder
        /**
         * 5. 结果处理
         */
        SearchResponse searchResponse = null;//执行查询
        try {
            searchResponse = client.search(searchRequest, RequestOptions.DEFAULT);
        } catch (IOException e) {
            e.printStackTrace();
        }

        /**
         * 6.处理搜索结果
         */
        RestStatus restStatus = searchResponse.status();
        if (restStatus != RestStatus.OK){
            logger.error("[EsQueryHighCtrl]搜索错误");
            throw new BizException(SearchResultCode.ES_RESPON_ERROR);
        }

        List<AssetBo> list = new ArrayList<>();
        SearchHits hits = searchResponse.getHits();
        long totalHits = hits.getTotalHits().value;

        hits.forEach(item -> {
                AssetBo bo = JSON.parseObject(item.getSourceAsString(), AssetBo.class);
                list.add(bo);
                }
        );
        BaseSearchVo result=new BaseSearchVo();
        result.setCount(totalHits);
        result.setList(list);
        return result;
    }
}
```

## ElasticSearch评分机制

### 自定义评分

https://blog.csdn.net/laoyang360/article/details/104809787

## ElasticSearch插件

### IK分词器

https://github.com/medcl/elasticsearch-analysis-ik/

https://www.cnblogs.com/xing901022/p/5910139.html

https://blog.csdn.net/opera95/article/details/78594949

https://blog.csdn.net/wwd0501/article/details/80622669

### 拼音分词器(pinyin)

https://www.cnblogs.com/xing901022/p/5910139.html

https://blog.csdn.net/opera95/article/details/78594949

https://blog.csdn.net/wwd0501/article/details/80622669

### 部分推荐插件简介及安装

https://www.cnblogs.com/zlslch/p/6423631.html

https://www.cnblogs.com/xing901022/p/5962722.html

### ElasticSearch性能调优

https://www.cnblogs.com/technologykai/articles/10899582.html

# ElasticSearch底层原理

## Trie树

​		**Trie树**是一种**树形结构**，是一种**哈希树**的变种。典型应用是用于**统计**，**排序**和**保存大量的字符串**（但不仅限于字符串），所以经常被**搜索引擎**系统用于文本词频统计。它的优点是：利用字符串的公共前缀来**减少查询时间**，最大限度地**减少无谓的字符串比较**，查询效率比哈希树高。

https://blog.csdn.net/forever_dreams/article/details/81009580

## 联合索引查询

https://www.infoq.cn/article/database-timestamp-02/?utm_source=infoq&utm_medium=related_content_link&utm_campaign=relatedContent_articles_clk

## 倒排索引

​		Elasticsearch存储文档时，将在1秒钟内几乎实时地对其进行索引和完全搜索。Elasticsearch使用称为倒排索引的数据结构，该结构支持非常快速的全文本搜索。反向索引列出了出现在任何文档中的每个唯一单词，并标识了每个单词出现的所有文档。



# 附录

## ElasticSearch官方网站

https://www.elastic.co/cn/elasticsearch/

## Kibana官方网站

https://www.elastic.co/cn/kibana/


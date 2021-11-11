[toc]

# Vue

#### Mustache语法

{{%变量名%}}

### 创建Vue实例

```html
<body>
    <div id="xxx">//id = xxx
        {{data}}//响应式显示data数据
    </div>
    
    <script src = "../vue.js"></script>
    <script>
    	const app = new Vue({
			el:'xxx',//将id=xxx的元素交给vue进行管理
			data:'xxx'//data数据，可从服务器获取
		})
    </script>
</body>
```


# HTTP 协议详解

## HTTP 请求消息 （Request Message）

### 组成部分

- 请求行（request line）

```http
请求方法    URL        协议版本

GET /images/logo.png HTTP/1.1
```

- 请求头（request header）

```
头部字段名：值

Accept:text/html,image/*
```
- 空行（empty line）

- 请求体（message body）

请求行、每行头部字段必须以<CR><LF>(回车和换行)结尾，空行必须只有<CR><LF>，并且没有其他空白符。

#### 请求方法：
|方法|特性|作用|
|:---:|:---:|:---|
|GET|安全、幂等|请求获得某个特定资源，只能使用GET获取资源，对资源本身不应该有影响|
|HEAD|安全、幂等|和GET一样，请求获得某个资源，但是返回的响应中没有实体内容，用于获取报头|
|POST||向指定资源提交数据进行处理请求（表单、文件），数据包含在请求体中。POST请求有可能会更改已有资源或者建立新资源|
|PUT|幂等|申请修改某些数据，是安全的、幂等的|
|DELETE|幂等|删除指定的资源|
|TRACE|安全、幂等|TRACE方法回显收到的请求，以便客户端可以看到中间服务器进行了哪些更改（或有任何更改）。|
|OPTIONS|安全、幂等|返回该服务器支持的用于该URI的HTTP方法，可以通过请求“*” 而不是特定资源来检查Web服务器的功能|
|CONNECT||HTTP/1.1 协议中预留给能够将连接改为管道方式的代理服务器。|
|PATCH||对已知资源的局部更新|

*安全*： 表示这些方法仅用来获取资源，不应该改变资源的状态
*幂等*： 表示对同样的资源进行同一个操作，无论多次操作还是一次操作的结果都不会产生变化

### 一个典型的请求

```
GET / HTTP/1.1
Host: www.example.com
```

## HTTP 响应消息 （Response Message）

### 组成部分

- 包含状态码和原因信息的状态行（status line）

```
HTTP/1.1 200 OK
```

- 响应头部（response）

- 空行（empty line）

- 实体内容（message body）

### 一个典型的响应

```http
HTTP/1.1 200 OK
Date: Mon, 23 Sept 2019 22:38:34 GMT
Content-Type: text/html; charset=UTF-8
Content-Length: 138
Last-Modified: Tus, 08 Jan 2019 23:11:55 GMT
Server: Apache/1.3.3.7 (Unix) (Red-Hat/Linux)
ETag: "3f80f-1b6-3e1cb03b"
Accept-Ranges: bytes
Connection: close

<html>
  <head>
    <title>An Example Page</title>
  </head>
  <body>
    <p>Hello World, this is a very simple HTML document.</p>
  </body>
</html>
```

## 头部常用字段

|请求头字段（Requset）|说明|例子|
|:----|:---|:---|
|Accept|可接受的响应媒体类型|Accpet: text/html|
|Accept-Charset|可接受的字符编码|Accept-Charset: utf-8|
|Accept-Encoding|一系列可接受的编码方式|Accept-Encoding: gzip, deflate|
|Connection|控制连接方式，但是不能在HTTP/2上使用|Connection: keep-alive 或者 Connection: Upgrade|
|Content-Length|请求正文包含多少字节|Content-Length: 348|
|Content-Type|请求正文的类型（PUT/POST使用）|Content-Type: application/x-www-form-urlencoded|
|User-Agent|请求时使用的浏览器类型和操作系统信息|User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:12.0) Gecko/20100101 Firefox/12.0|

|响应字段（Resposne）|说明|例子|
|:----|:---|:---|
|Allow|特定资源的有效操作方法|Allow: GET, HEAD|

## 状态码

1XX：Informational
2XX：Successfull
3XX：Redirection
4XX：Client Error
5XX：Server Error

|状态码|Statuscode|描述|
|::|||
|100|Continue|客户端继续发送请求|
|101|Switching Protocols|切换HTTP协议至高版本|
|200|OK|请求成功|
|201|Created|成功请求并创建了新资源|
|202|Accepted|已接受请求，但未处理完成|
|206|Partial Content|服务器会根据客户端头部的字段，只传递部分资源|
|300|Multiple Choice|请求的资源包括多种位置|
|400|Bad Request|客户端发送的请求语法错误，服务端无法理解|
|401|Unuthorized|请求要求用户的身份认证|
|403|Forbidden|服务器理解客户端的请求，但是拒绝处理|
|404|Not Found|服务器无法根据客户端的请求找到资源|
|500|Internal Server Error|服务器内部错误，无法完成请求|
|501|Not Inplimented|服务器不支持请求的功能，无法完成请求|
|502|Bad Gateway|作为网关或者代理工作的服务器尝试执行请求时，从远程服务器接收到了一个无效的响应|

## HTTP/1.0 vs. HTTP/1.1 vs. HTTP/2

### HTTP1.0 与 HTTP1.1 的区别

1. **缓存处理**：在 HTTP1.0 中主要使用 header 里的 If-Modified-Since, Expires 来做为缓存判断的标准，HTTP1.1 则引入了更多的缓存控制策略例如 Entity tag，If-Unmodified-Since, If-Match, If-None-Match 等更多可供选择的缓存头来控制缓存策略。
2. **带宽优化及网络连接的使用**：HTTP1.0 中存在一些浪费带宽的现象，例如客户端只是需要某个对象的一部分，而服务器却将整个对象送过来了，并且不支持断点续传功能。HTTP1.1 则在请求头引入了range头域，它允许只请求资源的某个部分，即返回码是206（Partial Content），这样就方便了开发者自由的选择以便于充分利用带宽和连接。
3. **状态码增多**：新增24个状态码。
4. **Host 头处理**：HTTP1.0 中认为每台服务器都绑定一个唯一的 IP 地址，因此，请求消息中的URL并没有传递主机名（hostname）。但随着虚拟主机技术的发展，在一台物理服务器上可以存在多个虚拟主机（Multi-homed Web Servers），并且它们共享一个IP地址。HTTP1.1 的请求消息和响应消息都应支持Host头域，且请求消息中如果没有Host头域会报告一个错误（400 Bad Request）。
5. **长连接**：HTTP 1.1支持长连接（Persistent Connection）和请求的流水线（Pipeline）处理，在一个TCP连接上可以传送多个HTTP请求和响应，减少了建立和关闭连接的消耗和延迟，在HTTP1.1中默认开启 Connection： keep-alive。

#### HTTP1.x 与 HTTP2 的区别

HTTP/2 发展自之前google开发的 SPDY 协议

**SPDY：** 优化了 HTTP1.x 的请求延迟，解决了HTTP1.x的安全性问题，gogle 于 2012 年提出。

1. 降低延迟：SPDY 采用了多路复用，多路复用通过多个请求stream共享一个tcp连接的方式，解决了HOL blocking的问题，降低了延迟同时提高了带宽的利用率
2. 请求优先级（request prioritization）。多路复用带来一个新的问题是，在连接共享的基础之上有可能会导致关键请求被阻塞。SPDY 允许给每个request设置优先级，这样重要的请求就会优先得到响应。比如浏览器加载首页，首页的html内容应该优先展示，之后才是各种静态资源文件，脚本文件等加载，这样可以保证用户能第一时间看到网页内容。

3. header压缩。前面提到HTTP1.x的header很多时候都是重复多余的。选择合适的压缩算法可以减小包的大小和数量。

4. 基于HTTPS的加密协议传输，大大提高了传输数据的可靠性。

5. 服务端推送（server push），采用了 SPDY 的网页，例如我的网页有一个sytle.css的请求，在客户端收到sytle.css数据的同时，服务端会将sytle.js的文件推送给客户端，当客户端再次尝试获取sytle.js时就可以直接从缓存中获取到，不用再发请求了。SPDY构成图：

|SPDY构成图|
|----|
|HTTP|
|SPDY|
|TLS/SSL|
|TCP|

** HTTP2 性能惊人**：

[https://link.zhihu.com/?target=https://http2.akamai.com/demo](https://link.zhihu.com/?target=https://http2.akamai.com/demo) Akamai 公司建立的一个官方的演示，用以说明 HTTP/2 相比于之前的 HTTP/1.1 在性能上的大幅度提升。 同时请求 379 张图片，从Load time 的对比可以看出 HTTP/2 在速度上的优势。

**HTTP2， SPDY 的升级版**，HTTP2 与 SPDY 的区别：

1. HTTP2.0 支持明文 HTTP 传输，而 SPDY 强制使用 HTTPS

2. HTTP2.0 消息头的压缩算法采用 HPACK http://http2.github.io/http2-spec/compression.html，而非 SPDY 采用的 DEFLATE http://zh.wikipedia.org/wiki/DEFLATE



**HTTP2 和 HTTP1.X 相比的新特性**

1. 新的二进制格式（Binary Format），HTTP1.x的解析是基于文本。基于文本协议的格式解析存在天然缺陷，文本的表现形式有多样性，要做到健壮性考虑的场景必然很多，二进制则不同，只认0和1的组合。基于这种考虑HTTP2.0的协议解析决定采用二进制格式，实现方便且健壮。

2. 多路复用（MultiPlexing），即连接共享，即每一个request都是是用作连接共享机制的。一个request对应一个id，这样一个连接上可以有多个request，每个连接的request可以随机的混杂在一起，接收方可以根据request的 id将request再归属到各自不同的服务端请求里面。

3. header压缩，如上文中所言，对前面提到过HTTP1.x的header带有大量信息，而且每次都要重复发送，HTTP2.0使用encoder来减少需要传输的header大小，通讯双方各自cache一份header fields表，既避免了重复header的传输，又减小了需要传输的大小。

4. 服务端推送（server push），同SPDY一样，HTTP2.0也具有server push功能。

---
**References:**

[1]. [runoob, http 教程](https://www.runoob.com/http/http-header-fields.html)
[2]. [wikipedia, *Hypertext_Transfer_Protocol*](https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)
[3]. [wikipedia, *List_of_HTTP_header_fields*](https://en.wikipedia.org/wiki/List_of_HTTP_header_fields#Request_fields)
[4]. [wikipedia, *List_of_HTTP_status_codes*](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
[5]. [_浪潮之巅，*HTTP1.0, HTTP1.1, HTTP2.0的区别*](https://www.cnblogs.com/heluan/p/8620312.html)

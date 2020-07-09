# HTTP 与 HTTPS 的区别

![](../pictures/http-vs-https.jpg)

HTTP 协议传输数据都是明文的，未经加密的，因此使用 HTTP 来传输隐私数据非常不安全。

安全超文本传输协议， HTTPS （Hypertext Transfer Protocol Secure，HTTPS）是 HTTP 的一个扩展，用于安全的传输数据。

在 HTTPS 中，使用 TLS （Transport Layer Security， TLS） 或 SSL （Secure Sockets Layer， SSL）进行文本加密。因此，该协议也经常被称为，基于 TLS 的 HTTP 协议， 或基于 SSL 的 HTTP 协议。

HTTPS 原则就是对访问网站的身份验证，并在传输过程中保证交换数据的隐私和完整性。它可以防止中间人攻击。

简而言之：**HTTPS = HTTP + TLS/SSL**

## 区别

1. HTTP URL 的 Scheme 是 http://，但是 HTTPS 的 URL Schema 是 https://.
2. HTTP 是不安全的， HTTPS 是安全的。
3. HTTP 使用 80 端口， 而 HTTPS 使用 443 端口。
4. HTTP 作用于应用层，而 HTTPS 的 TLS/SSL（协议处于应用层） 作用于传输层。
5. HTTP 不需要 TLS/SSL 许可证，而 HTTPS 需要 CA 签发的许可证
6. HTTP 不需要域验证， 而 HTTPS 需要域验证甚至需要合法的文本验证。
7. HTTP 的数据是不加密的，而 HTTPS 是经过加密的。


-------

**references:**

[1]. [wikipedia， *HTTPS*]（https://en.wikipedia.org/wiki/HTTPS）

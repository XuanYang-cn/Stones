# OSI model

OSI, **Open system interconnection model**

## OSI 参考模型

|层数||协议数据单元（Protocol data unit， PDU ）|功能|协议|
|:---:|:---:|:---:|:---|:---|
|第七层|应用层（Application）|数据|为应用程序提供服务|HTTP，SMTP，FTP，IMAP4， TLS/SSL|
|第六层|表示层（Presentation）|数据|数据格式转化，数据加密||
|第五层|会话层（Session）|数据|负责在网络中的两节点之间建立和维持会话||
|第四层|传输层（Transport）|段，数据报（Segment，Datagram）|建立端到端之间的连接，数据的分段和重组|TCP，UDP|
|第三层|网络层（Network）|包（Packet）|IP地址和路由选择|IP， ARP，ICMP|
|第二层|数据链路层（Data link）|帧（Frame）|提供介质访问和链路管理||
|第一层|物理层（Physical）|0，1 （Symbol）|使用物理介质传输比特流|光纤|

## OSI vs. TCP/IP

---

**reference**

[1]. [wikipedia, *OSI_model*](https://en.wikipedia.org/wiki/OSI_model)
[2]. [wikipedia, *List_of_network_protocols_(OSI_model)*](https://en.wikipedia.org/wiki/List_of_network_protocols_(OSI_model))
[2]. [wikipedia, *Internet_protocol_suite*](https://en.wikipedia.org/wiki/Internet_protocol_suite)

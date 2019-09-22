# 什么是 REST 和 RESTful Web API？

表示性状态转移（**Re**presentational **s**tate **t**ransfer ，REST）是一种架构风格，Web 服务只有遵守它的 6 个约束才能称为 RESTful Web 服务。RESTful 的 Web 服务允许客户端通过预先设定的无状态的统一方法（URI 和 HTTP协议）来获取、操作 Web 资源

重要的一点：**使用 HTTP 协议并不等于 RESTful**

通过使用无状态协议（比如 HTTP 协议）和标准化操作（GET，POST 等），RESTful 的系统有以下优点： 性能更快，稳定性，组件的可重用性、可扩展性和易于维护性。

## 约束

6 大约束确定一个 RESTful 系统。通过遵循这些约束，系统能够获得一些非功能性的提升，比如性能、可扩展性、易维护性、简洁性、可视化、可移植性以及稳定性。违反了任意约束的系统都不能称之为 RESTful 系统。这六个约束如下：

### 1. CS 架构

CS 架构背后的概念就是前后端分离。分离用户界面和数据存储会提高多平台前端的可移植性，也会因为服务端组件的简化而提高可扩展性。最重要的是，前后端分离让各个组件能各自维护，从而支持多领域的互联网扩展需求。

基于事件集成的架构是相对于 CS 架构而言。这种架构中，每个组件都在持续广播事件，同时在监听其他组件的相关事件。它并不是一对一的通信，仅仅是事件的触发和监听。而 REST 则要求一对一的连接，所以基于事件集成的架构并不是 RESTful 的。


### 2. 无状态（Statelessness）

在每个请求之间，服务器不会保存客户端的上下文。 从客户端发送过来的任何请求都包含了能够处理这次请求的全部信息，并且会话记录是由客户端保存的。

### 3. 可缓存性 （Cacheability）

与在万维网上一样，客户端可以缓存响应。 因此，响应必须隐式或显式地将自身定义为可缓存或不缓存，以防止客户端响应过时或使用了不适当的数据。 管理良好的缓存部分或完全消除了某些客户端-服务器交互，从而进一步提高了可伸缩性和性能。

### 4. 层级系统（Layered system）

客户端不需要知道自己是否是直接连接在终端服务器还是连接着中间件。 如果在客户端和服务器之间放置了代理或负载平衡器，既不会影响它们之间的通信，同时也不需要更新客户端或服务器代码。

中间服务器可以通过启用负载平衡和共享缓存来提高系统可伸缩性。 此外，可以将安全认证作为Web服务之上的一层添加，这样可以将业务逻辑与安全性逻辑明确分开。 增加额外的安全认证可以增强安全性，也可以让一个服务器器控制其他的服务器来响应客户端的请求。

### 5. 统一接口（Uniform interface）                                                                                        

The uniform interface constraint is fundamental to the design of any RESTful system. It simplifies and decouples the architecture, which exables each part to evolve independently. The four constraints for this uniform interface are:

**Resource indenrification in requsets**

Individual resources are idendified in requests, for example using **URIs** in RESTful Web services. The resources themselves are conceptually separate from the representations that are returned to the client. For example, the server sould send data from its database as HTML, XML or as JSON--none of which are the server's internal representation

**Resource representations**

Represent the state of a resource for transfer between components

**Self-descriptive messages**

Each message includes enough information to describe how to process the message.

**Hypermedia as the engine of application state**

Having accessed an initial URI for the REST application--analogous to a human Web user accessing the home page of a website--a REST client should then be able to use server-provided links dynamically to discover all the available actinos and resources it needs. As access proceeds, the server responds with text that includes hyperlinks to other actions that are currently available.

### 6 Code on demand(optional)

Servers can temporarily extend or customize the functionality of a client by transferring executable code

## Architectural properties

- Performance in component interactions, which can be dominant factor in user-perceived performance and network efficiency
- Scalability allowing the support of large numbers of components and interactions among components. Roy Fielding describes REST's effect on scalability as follows:

> REST's client-server separation of concerns simplifies component implementation, reduces the complexity of connector semantics, imporves the effectiveness of performance tuning, and increase the scalability of pure server components. Layered system constraints allow intermediaries--proxies, gateways, and firewalls--to be introduced at various points in the communication without changing the interfaces between components, thus allowing them to assist in communication translation or improve performance via large-scale, shared caching. REST enables intermediate processing by contraining messages to be self-descriptive: interaction is stateless between requests, standard methods and media types are used to indicate semantics and exchange information, and responses explicitly indicate cacheability.

- simplicity of a uniform interface
- modifiability of components to meet changing needs(even while the application is running)
- visibility of communication between components by service agents
- portability of components by moving program code with the data
- reliability in the resistance to failure at the system level in the presence of failures within components, connectors, or data

## Applied to Web services

Web service APIs that adhere to the REST architectural constraints are called **RESTful APIs**. HTTP-based RESTful APIs are defined with the following aspects

- a base URI, such as **http://api.example.com/collection/**
- standard HTTP methods(e.g. GET, POST, PUT, PATCh and DELETE)
- a media type that defines state transition data elements(e.g. Atom, microformats, application/vnx.collection+json, etc.). The current representation tells the client how to compose requests for transitions to all the next available application states.

HTTP methods and how the are typically used in a RESTful API:

|HTTP METHODS|Collection resource, such as https://api.example.com/collection/|Member resource, such as https://api.example.com/collection/item3|
|:----------:|:------------------|:--------------|
|GET|**Retrieve** the URIs of the member resources of the collection resource in the resource in the response body.|**Retrieve** representation of the member resource in the response body.|
|POST|**Create** a member resource in the collection resource using the instructions in the request body. The URI of the created member resource is *automatically assigned* and returned in the response *Location* header field.|**Create** a member resource in the member resource using the instructions in the request body. The URI of the created member resource is automatically assigned and returned in the response *Location* header field.|
|PUT|**Replace** all the representations of the member resources of the collection resource with the representation in the request body, or **create** the collection resource if it does not exist.|**Replace** all the representations of the member resource, or **may create** the member resource if it does not exist, using the instructions in the request body.|
|PATCH|**Update** all the representations of the member resources of the collection resource using the instructions in the request body, or **may create** the collection resource if it does not exist.|**Update** all the representations of the member resource, or **may create** the member resource if it does not exist, using the instructions int the request body.|
|DELETE|**Delete** all the representations of the member resources of the collection resource|**Delete** all the representations of the member resource.|

The GET method is **safe**, meaning that applying it to a resource does not result in a state change of the resource(read-only semantics)

The GET, PUT and DELETE methods are **idempotent**, meaning that applying them multiple times to a resource result in the same state change of the resource as applying them once, though the response might differ.

The GET and POST methods are **cacheable**, meaning that responses to them are allowed to be stored for future reuse.

Unlike SOAP-based Web services, there is no "official" standard for RESTful Web APIs. This is because REST is an architectural style, while SOAP is a protocol. REST is not a standard in itself, but RESTful implementations make use of standards, such as HTTP, URI, JSON, and XML. Many developers also describe their APIs as being RESTful, even though these APIs actually don't fulfill all of the architectural constraints described above (especially the uniform interface constraint).

---

References:

\[1\]: Wikipedia. [https://en.wikipedia.org/wiki/Representational_state_transfer](https://en.wikipedia.org/wiki/Representational_state_transfer)

\[2\]: *Lauren Long*, What RESTful actually means. [https://codewords.recurse.com/issues/five/what-restful-actually-means](https://codewords.recurse.com/issues/five/what-restful-actually-means)

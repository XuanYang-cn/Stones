# What is REST and RESTful Web API

**Re**presentational **s**tate **t**ransfer (REST) is a software architectural style that defines a set of constraints to be used for creating Web services.

Web services that conform to the REST architectural style, called RESTful Web Services. RESTful Web services allow the requesting systems to access and manipulate textual representations of **Web resources** by using a **uniform and predefined set of stateless operations**

**Using HTTP is NOT equal to RESTful**

By using a stateless protocol(usually HTTP) and standard operations, RESTful systems aim for fast performance, reliability, and the ability to grow by reusing components that can be managed and updated without affecting the system as a whole, even while it's running.

## Architectural constraints

Six guiding constraints define a RESTful system. These constraints restrict the ways that the server can process a respond to client requests so that, by operating within these constraints, the system gain desirable **non-functional properties**, such as performance, scalabilty, simplicity, modifiability, visibility, portability and reliability. If a system violates any of the required constraints, it cannot be considererd RESTful.

### Client-server architecture

The principle behind the client-server constraints is the separation of concerns. Separating the user interface concerns from the data storage concerns improves the portability of the user interfaces across multiple platforms. It also improves scalability by simplifying the server components. Perhaps most significant to the Web, however, is that the separation allows the components to evolve independently, thus supporting the internet-scale requirement of multiple organizational domains.

A non-RESTful alternative to CS architecture is event-based integration architecture. In this model, each component continuously broadcasts events while listening for pertinent events from other components. These's no one-to-one communication, only broadcasting and eavesdropping. REST requires one-to-one communication, so event-based integration architecture would not be RESTful

### Statelessness

The client-server communication is constrained by no client context being stored on the server between requests. Each request from any client contains all the information necessary to service the request, and the session state is held in the client.

Each request is treated as standalone, the server does not remember anything about the user who uses the API

### Cacheability

As on World Wide Web, clients and intermediaries can cache responses. Responses must therefore, implicitly or explicitly, define themselves as cacheable or not to prevent clients from getting stale or in appropriate data in response to further requests. Well-managed caching partially or completely eliminates some client-server interactions, further improving scalability and performance.

### Layered system

A client cannot ordinarily tell whether it is connected directly to the end server , or to an intermediary along the way. So if a proxy or load balancer is placed between the client and server, it wouldn't affect their communications and there wouldn't be necessities to update the client or server code.

Intermediary servers can improve system scalability by enabling load balancing and by procingding shared caches. Also, security can be added as a layer on top of the web services, and then clearly separate business logic from security logic. By adding security as a separate layer enforces security policies. Finally, it also means that a server can call multiple other servers to generate a response to the client.

### Code on demand(optional)

Servers can temporarily extend or customize the functionality of a client by transferring executable code

### Uniform interface

The uniform interface constraint is fundamental to the design of any RESTful system. It simplifies and decouples the architecture, which exables each part to evolve independently. The four constraints for this uniform interface are:

**Resource indenrification in requsets**

Individual resources are idendified in requests, for example using **URIs** in RESTful Web services. The resources themselves are conceptually separate from the representations that are returned to the client. For example, the server sould send data from its database as HTML, XML or as JSON--none of which are the server's internal representation

**Resource representations**

Represent the state of a resource for transfer between components

**Self-descriptive messages**

Each message includes enough information to describe how to process the message.

**Hypermedia as the engine of application state**

Having accessed an initial URI for the REST application--analogous to a human Web user accessing the home page of a website--a REST client should then be able to use server-provided links dynamically to discover all the available actions and resources it needs. As access proceeds, the server responds with text that includes hyperlinks to other actions that are currently available.

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
- standard HTTP methods(e.g. GET, POST, PUT, PATCH and DELETE)
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

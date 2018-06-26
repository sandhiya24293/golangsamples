package _02_WebServices
/*
LINK:===https://searchmicroservices.techtarget.com/tip/REST-vs-SOAP-Choosing-the-best-web-service

Web services overview
A Web service, in very broad terms, is a method of communication between two applications or electronic devices over the World Wide Web (WWW).
Web services are of two kinds: Simple Object Access Protocol (SOAP) and Representational State Transfer (REST).

SOAP defines a standard communication protocol (set of rules) specification for XML-based message exchange.
SOAP uses different transport protocols, such as HTTP and SMTP.
The standard protocol HTTP makes it easier for SOAP model to tunnel across firewalls and proxies without any modifications to the SOAP protocol.
SOAP can sometimes be slower than middleware technologies like CORBA or ICE due to its verbose XML format.

REST describes a set of architectural principles by which data can be transmitted over a standardized interface (such as HTTP).
REST does not contain an additional messaging layer and focuses on design rules for creating stateless services.
A client can access the resource using the unique URI and a representation of the resource is returned.
With each new resource representation, the client is said to transfer state.
While accessing RESTful resources with HTTP protocol, the URL of the resource serves as the resource identifier and GET, PUT, DELETE, POST and HEAD are the standard HTTP operations to be performed on that resource.

REST vs. SOAP
There are significant differences between SOAP and RESTful web services.
The bullets below break down the features of each web service based on personal experience.

REST
1) RESTful web services are stateless. You can test this condition by restarting the server and checking if interactions survive.
2) For most servers, RESTful web services provide a good caching infrastructure over an HTTP GET method. This can improve the performance if the information the service returns is not altered frequently and is not dynamic.
3) Service producers and consumers must understand the context and content being passed along as there is no standard set of rules to describe the REST web services interface.
4) REST is useful for restricted-profile devices, such as mobile, for which the overhead of additional parameters are less (e.g., headers).
5) REST services are easy to integrate with existing websites and are exposed with XML so the HTML pages can consume the same with ease. There is little need to refactor the existing site architecture. As such, developers are more productive because they don't need to rewrite everything from scratch; instead, they just need to add on the existing functionality.
6) A REST-based implementation is simple compared to SOAP.


SOAP
1) The Web Services Description Language (WSDL) describes a common set of rules to define the messages, bindings, operations and location of the service. WSDL is akin to a contract to define the interface that the service offers.
2) SOAP requires less plumbing code than REST services design (e.g., transactions, security, coordination, addressing and trust). Most real-world applications are not simple and support complex operations, which require conversational state and contextual information to be maintained. With the SOAP approach, developers don't need to write plumbing code into the application layer.
3) SOAP web services, such as JAX-WS, are useful for asynchronous processing and invocation.
4) SOAP supports several protocols and technologies, including WSDL, XSDs and WS-Addressing.
5) Consuming a web service via a database stored procedure allows users to straight away update a database with information from different sources. Users can also schedule a job at regular intervals to get data updated periodically in the database.

REST or SOAP: Which is best for me?
Proponents on both sides of the REST vs. SOAP discussion can be fervent in their advocacy for their web service architecture of choice. Both SOAP and RESTful architectures have proven themselves to be reliable, successful and capable of scaling infinitely, so the decision to use REST or SOAP has less to do with their efficacy and more to do with how either approach fits in with an organization’s software development culture and project needs.

Both SOAP web services and RESTful web services have proven their ability to meet the demands of the largest enterprise organizations in the world, while at the same time being able to service the smallest internet of things devices or embedded applications in production.

When choosing between REST and SOAP, two of the key topics to factor into the decision are:

1) The types of clients that will be supported. For example, REST services offer an effective way for interacting with lightweight clients, such as smartphones.
2) How varying degrees of flexibility and standardization will be either rebuffed or accepted by the organization’s corporate culture.

Keeping these factors in mind will go a long way in helping organizations to choose between SOAP and RESTful web services.





*/

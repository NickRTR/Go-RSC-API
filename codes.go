package main

import (
	"errors"
	"fmt"
)

type code struct {
    Code int16 `json:"code"`
    Title string `json:"Title"`
    Description string `json:"Description"`
}

var codes = [] code {
    {
        Code: 100,
        Title: "Continue",
        Description: "This interim response indicates that the client should continue the request or ignore the response if the request is already finished.",
    }, {
        Code: 101,
        Title: "Switching Protocols",
        Description: "This code is sent in response to an Upgrade request header from the client and indicates the protocol the server is switching to.",
    }, {
        Code: 102,
        Title: "Processing (WebDav)",
        Description: "This code is sent in response to an Upgrade request header from the client and indicates the protocol the server is switching to.",
    }, {
        Code: 103,
        Title: "Early Hints (experimental)",
        Description: "This status code is primarily intended to be used with the Link header, letting the user agent start preloading resources while the server prepares a response.",
    }, {
        Code: 200,
        Title: "OK",
        Description: `The request succeeded. The result meaning of "success" depends on the HTTP method. GET: The resource has been fetched and transmitted in the message body. HEAD: The representation headers are included in the response without any message body. PUT or POST: The resource describing the result of the action is transmitted in the message body. TRACE: The message body contains the request message as received by the server.`,
    }, {
        Code: 201,
        Title: "Created",
        Description: "The request succeeded, and a new resource was created as a result. This is typically the response sent after POST requests, or some PUT requests.",
    }, {
        Code: 202,
        Title: "Accepted",
        Description: "The request has been received but not yet acted upon. It is noncommittal, since there is no way in HTTP to later send an asynchronous response indicating the outcome of the request. It is intended for cases where another process or server handles the request, or for batch processing.",
    }, {
        Code: 203,
        Title: "Non-Authoritative Information",
        Description: "This response code means the returned metadata is not exactly the same as is available from the origin server, but is collected from a local or a third-party copy. This is mostly used for mirrors or backups of another resource. Except for that specific case, the 200 OK response is preferred to this status.",
    }, {
        Code: 204,
        Title: "No Content",
        Description: "There is no content to send for this request, but the headers may be useful. The user agent may update its cached headers for this resource with the new ones.",
    }, {
        Code: 205,
        Title: "Reset Content",
        Description: "Tells the user agent to reset the document which sent this request.",
    }, {
        Code: 206,
        Title: "Partial Content",
        Description: "This response code is used when the Range header is sent from the client to request only part of a resource.",
    }, {
        Code: 207,
        Title: "Multi-Status (WebDAV)",
        Description: "Conveys information about multiple resources, for situations where multiple status codes might be appropriate.",
    }, {
        Code: 208,
        Title: "Already Reported (WebDAV)",
        Description: "Used inside a <dav:propstat> response element to avoid repeatedly enumerating the internal members of multiple bindings to the same collection.",

    }, {
        Code: 226,
        Title: "IM Used",
        Description: "The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.",

    }, {
        Code: 300,
        Title: "Multiple Choices",
        Description: "The request has more than one possible response. The user agent or user should choose one of them. (There is no standardized way of choosing one of the responses, but HTML links to the possibilities are recommended so the user can pick.)",
    }, {
        Code: 301,
        Title: "Moved Permanently",
        Description: "The URL of the requested resource has been changed permanently. The new URL is given in the response.",
    }, {
        Code: 302,
        Title: "Found",
        Description: "This response code means that the URI of requested resource has been changed temporarily. Further changes in the URI might be made in the future. Therefore, this same URI should be used by the client in future requests.",
    }, {
        Code: 303,
        Title: "See Other",
        Description: "The server sent this response to direct the client to get the requested resource at another URI with a GET request.",
    }, {
        Code: 304,
        Title: "Not Modified",
        Description: "This is used for caching purposes. It tells the client that the response has not been modified, so the client can continue to use the same cached version of the response.",
    }, {
        Code: 305,
        Title: "Use Proxy (deprecated)",
        Description: "Defined in a previous version of the HTTP specification to indicate that a requested response must be accessed by a proxy. It has been deprecated due to security concerns regarding in-band configuration of a proxy.",

    }, {
        Code: 306,
        Title: "unused",
        Description: "This response code is no longer used; it is just reserved. It was used in a previous version of the HTTP/1.1 specification.",
    }, {
        Code: 307,
        Title: "Temporary Redirect",
        Description: "The server sends this response to direct the client to get the requested resource at another URI with same method that was used in the prior request. This has the same semantics as the 302 Found HTTP response code, with the exception that the user agent must not change the HTTP method used: if a POST was used in the first request, a POST must be used in the second request.",
    }, {
        Code: 308,
        Title: "Permanent Redirect",
        Description: "This means that the resource is now permanently located at another URI, specified by the Location: HTTP Response header. This has the same semantics as the 301 Moved Permanently HTTP response code, with the exception that the user agent must not change the HTTP method used: if a POST was used in the first request, a POST must be used in the second request.",
    }, {
        Code: 400,
        Title: "Bad Request",
        Description: "The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).",
    }, {
        Code: 401,
        Title: "Unauthorized",
        Description: `Although the HTTP standard specifies "unauthorized", semantically this response means "unauthenticated". That is, the client must authenticate itself to get the requested response.`,
    }, {
        Code: 402,
        Title: "Payment Required (experimental)",
        Description: "This response code is reserved for future use. The initial aim for creating this code was using it for digital payment systems, however this status code is used very rarely and no standard convention exists.",
    }, {
        Code: 403,
        Title: "Forbidden",
        Description: "The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource. Unlike 401 Unauthorized, the client's identity is known to the server.",
    }, {
        Code: 404,
        Title: "Not Found",
        Description: "The server cannot find the requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist. Servers may also send this response instead of 403 Forbidden to hide the existence of a resource from an unauthorized client. This response code is probably the most well known due to its frequent occurrence on the web.",
    }, {
        Code: 405,
        Title: "Method Not Allowed",
        Description: "The request method is known by the server but is not supported by the target resource. For example, an API may not allow calling DELETE to remove a resource.",
    }, {
        Code: 406,
        Title: "Not Acceptable",
        Description: "This response is sent when the web server, after performing server-driven content negotiation, doesn't find any content that conforms to the criteria given by the user agent.",
    }, {
        Code: 407,
        Title: "Proxy Authentication Required",
        Description: "This is similar to 401 Unauthorized but authentication is needed to be done by a proxy.",
    }, {
        Code: 408,
        Title: "Request Timeout",
        Description: "This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection. This response is used much more since some browsers, like Chrome, Firefox 27+, or IE9, use HTTP pre-connection mechanisms to speed up surfing. Also note that some servers merely shut down the connection without sending this message.",
    }, {
        Code: 409,
        Title: "Conflict",
        Description: "This response is sent when a request conflicts with the current state of the server.",
    }, {
        Code: 410,
        Title: "Gone",
        Description: `This response is sent when the requested content has been permanently deleted from server, with no forwarding address. Clients are expected to remove their caches and links to the resource. The HTTP specification intends this status code to be used for "limited-time, promotional services". APIs should not feel compelled to indicate resources that have been deleted with this status code.`,
    }, {
        Code: 411,
        Title: "Length Required",
        Description: "Server rejected the request because the Content-Length header field is not defined and the server requires it.",
    }, {
        Code: 412,
        Title: "Precondition Failed",
        Description: "The client has indicated preconditions in its headers which the server does not meet.",
    }, {
        Code: 413,
        Title: "Payload Too Large",
        Description: "Request entity is larger than limits defined by server. The server might close the connection or return an Retry-After header field.",
    }, {
        Code: 414,
        Title: "URI Too Long",
        Description: "The URI requested by the client is longer than the server is willing to interpret.",
    }, {
        Code: 415,
        Title: "Unsupported Media Type",
        Description: "The media format of the requested data is not supported by the server, so the server is rejecting the request.",
    }, {
        Code: 416,
        Title: "Range Not Satisfiable",
        Description: "The range specified by the Range header field in the request cannot be fulfilled. It's possible that the range is outside the size of the target URI's data.",
    }, {
        Code: 417,
        Title: "Expectation Failed",
        Description: "This response code means the expectation indicated by the Expect request header field cannot be met by the server.",
    }, {
        Code: 418,
        Title: `I'm a teapot`,
        Description: "The server refuses the attempt to brew coffee with a teapot. ☕",
    }, {
        Code: 421,
        Title: "Misdirected Request",
        Description: "The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI. ",
    }, {
        Code: 422,
        Title: "Unprocessable Entity (WebDAV)",
        Description: "The request was well-formed but was unable to be followed due to semantic errors.",
    }, {
        Code: 423,
        Title: "Locked (WebDAV)",
        Description: "The resource that is being accessed is locked.",
    }, {
        Code: 424,
        Title: "Failed Dependency (WebDAV)",
        Description: "The request failed due to failure of a previous request.",
    }, {
        Code: 425,
        Title: "Too Early (experimental)",
        Description: "Indicates that the server is unwilling to risk processing a request that might be replayed.",
    }, {
        Code: 426,
        Title: "Upgrade Required",
        Description: "The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol. The server sends an Upgrade header in a 426 response to indicate the required protocol(s).",
    }, {
        Code: 428,
        Title: "Precondition Required",
        Description: "The origin server requires the request to be conditional. This response is intended to prevent the 'lost update' problem, where a client GETs a resource's state, modifies it and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.",
    }, {
        Code: 429,
        Title: "Too Many Requests",
        Description: `The user has sent too many requests in a given amount of time ("rate limiting").`,
    }, {
        Code: 431,
        Title: "Request Header Fields Too Large",
        Description: "The server is unwilling to process the request because its header fields are too large. The request may be resubmitted after reducing the size of the request header fields.",
    }, {
        Code: 451,
        Title: "Unavailable For Legal Reasons",
        Description: "The user agent requested a resource that cannot legally be provided, such as a web page censored by a government.",
    }, {
        Code: 500,
        Title: "Internal Server Error",
        Description: "The server has encountered a situation it does not know how to handle.",
    }, {
        Code: 501,
        Title: "Not Implemented",
        Description: "The request method is not supported by the server and cannot be handled. The only methods that servers are required to support (and therefore that must not return this code) are GET and HEAD.",
    }, {
        Code: 502,
        Title: "Bad Gateway",
        Description: "This error response means that the server, while working as a gateway to get a response needed to handle the request, got an invalid response.",
    }, {
        Code: 503,
        Title: "Service Unavailable",
        Description: "The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded. Note that together with this response, a user-friendly page explaining the problem should be sent. This response should be used for temporary conditions and the Retry-After HTTP header should, if possible, contain the estimated time before the recovery of the service. The webmaster must also take care about the caching-related headers that are sent along with this response, as these temporary condition responses should usually not be cached.",
    }, {
        Code: 504,
        Title: "Gateway Timeout",
        Description: "This error response is given when the server is acting as a gateway and cannot get a response in time.",
    }, {
        Code: 505,
        Title: "HTTP Version Not Supported",
        Description: "The HTTP version used in the request is not supported by the server.",
    }, {
        Code: 506,
        Title: "Variant Also Negotiates",
        Description: "The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.",
    }, {
        Code: 507,
        Title: "Insufficient Storage (WebDAV)",
        Description: "The method could not be performed on the resource because the server is unable to store the representation needed to successfully complete the request.",
    }, {
        Code: 508,
        Title: "Loop Detected (WebDAV)",
        Description: "The server detected an infinite loop while processing the request.",
    }, {
        Code: 510,
        Title: "Not Extended",
        Description: "Further extensions to the request are required for the server to fulfill it.",
    }, {
        Code: 511,
        Title: "Network Authentication Required",
        Description: "Indicates that the client needs to authenticate to gain network access.",
    },
}

func getCodes()[] code {
    return codes
}

func getJustCodes()[] int16 {
    var cs[] int16
    for _, cod := range codes {
        cs = append(cs, cod.Code)
    }
    return cs
}

func getCode(c int16)(code, error) {
    for _, code := range codes {
        if code.Code == c {
            return code,
                nil
        }
    }
    message := fmt.Sprintf(
        "No matching Response Status Code found for %d.",
        c)
    return code {
        Code: 0,
        Title: "",
        Description: "",
    }, errors.New(message)
}

func findCodeByTitle(input string)(code,
    error) {
    for _, code := range codes {
        if code.Title == input {
            return code,
                nil
        }
    }
    message := fmt.Sprintf(
        "No Response Status Code found matching the Title '%s'.",
        input)
    return code {
        Code: 0,
        Title: "",
        Description: "",
    }, errors.New(message)
}
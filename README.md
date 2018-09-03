# HTTP2 demo server

HTTP2 demo server by GoLang.

## Getting Started

### Prerequisites

* [Go 1.10.3+](https://golang.org/dl/)

### Installing

First, clone git project.

```
git clone https://github.com/usolkim/http2-demo.git
```

And then move to src directory in project.

```
cd http2-demo/src
```

Next build it.

```
go build -o http2-demo main.go
```

Run server now!

```
./http2-demo
```

Now you can access http://localhost via browser.

### Notice

#### Self-signed certificate

HTTP2 demo server has self-signed certificate for localhost.
So your browser blocked to access server via SSL.
Furthermore you ignore security message yourself.

#### Access Port

HTTP2 demo server is using HTTP(80) and HTTPS(443).
Almost Browser is support HTTP/2 Protocol by HTTPS.
So you access to demo server by HTTP, your browser is using HTTP/1.1 protocol.

#### Chrome Browser

Currently Chrome latest version(68.0.3440.106) is not support a Server Push feature of HTTP/2.
So if you can test Server Push, access to demo server via Chrome beta version(69.0.3497.72).
You can downaload Chrome beta vesion [here](https://www.google.com/intl/en/chrome/beta).

## Authors

* **USol Kim** - [usolkim](https://github.com/usolkim)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details


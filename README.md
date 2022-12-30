# Bricks

![Go](https://github.com/go-bricks/bricks/workflows/Go/badge.svg)
[![codecov](https://codecov.io/gh/go-bricks/bricks/branch/master/graph/badge.svg)](https://codecov.io/gh/go-bricks/bricks)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/go-bricks/bricks)](https://pkg.go.dev/mod/github.com/go-bricks/bricks)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-bricks/bricks)](https://goreportcard.com/report/github.com/go-bricks/bricks)

<table>
        <tr>
            <th><p align="left"><img src=wiki/logo.svg align="center" height=256></p></th>
            <th>
                <h1>forked from the amazing work of https://github.com/go-masonry/mortar thanks I've learned a lot from this project</h1>
                <p align="left">Bricks is a GO framework/library for building gRPC (and REST) web services. Bricks has out-of-the-box support for configuration, application metrics, logging, tracing, profiling, dependency injection and more. While it comes with predefined defaults, Bricks gives you total control to fully customize it.            
             </p>
            </th>
        </tr>
</table>

## Demo

Clone this [demo](http://github.com/go-bricks/bricks-demo) repository to better understand some of Bricks capabilities.

When you done, read the [documentation](https://go-bricks.github.io).

## Service Template

To help you bootstrap your services with Bricks [here](https://github.com/go-bricks/bricks-template) you can find a template. Read its README first.

## Features

- Bundled [Grpc-Gateway](https://github.com/grpc-ecosystem/grpc-gateway) (REST Reverse-Proxy).
- Dependency Injection using [Uber-FX](https://github.com/uber-go/fx).
- Pimped `*http.Client` with interceptors support.
- Abstract support for Logging, Configuration, Tracing and Monitoring libraries. Use provided wrappers or your own.
  - [Jaeger wrapper](https://github.com/go-bricks/bjaeger) client for tracing.
  - [Prometheus wrapper](https://github.com/go-bricks/bprometheus) client for monitoring/metrics.
  - [Zerolog wrapper](https://github.com/go-bricks/bzerolog) for logging.
  - [Viper wrapper](https://github.com/go-bricks/bviper) for configuration.
- Internal HTTP [Handlers](providers/handlers.go)
  - _Profiling_ `http://.../debug/pprof`
  - _Debug_ `http://.../debug/*`
  - _Configuration_ `http://.../self/config`
  - _Build Information_ `http://.../self/build`
  - _Health_ `http://.../health`
- [Server/Client](providers) Interceptors both for gRPC and HTTP, you can choose which to use and/or add your own. 
    
    Some examples
    - HTTP Headers can be forwarded to next hop, defined by list.
    - HTTP Headers can be included in logs, defined by list.
    - Made available in `ctx.Context` via gRPC incoming Metadata.
    - Automatic monitoring and tracing (if enabled) for every RPC defined by the API.

...and more.

### Telemetry (Everything connected)

* Logs have Tracing Information `traceId=6ff7e7e38d1e86f` **across services**
    ![logs](wiki/logs.png)

* Also visible in Jaeger `traceId=6ff7e7e38d1e86f` if it's sampled.
    ![jaeger](wiki/jaeger.png)

### Support for `*http.Client` Interceptors, so you can

* Add request and response info to Trace

    <!-- ![jaeger_http](wiki/jaeger_http.png) -->

* Log/Dump requests and/or responses when http request fails.

    ```golang
    return func(req *http.Request, handler client.HTTPHandler) (resp *http.Response, err error) {
        var reqBytes, respBytes []byte
        // If the response is Bad Request, log both Request and Response
        reqBytes, _ = httputil.DumpRequestOut(req, true) // it can be nil and it's ok
        if resp, err = handler(req); err == nil && resp.StatusCode >= http.StatusBadRequest {
            respBytes, _ = httputil.DumpResponse(resp, true) // it can be nil
            logger.WithError(fmt.Errorf("http request failed")).
            WithField("status",resp.StatusCode).
            Warn(req.Context(), "\nRequest:\n%s\n\nResponse:\n%s\n", reqBytes, respBytes)
        }
        return
    }
    ```

    ![http_client](wiki/http_client_dump.png)

* Alter requests and/or responses (useful in [Tests](https://github.com/go-bricks/bricks-demo/blob/master/workshop/app/controllers/workshop_test.go#L162))

    ```golang
    func(*http.Request, clientInt.HTTPHandler) (*http.Response, error) {
        // special case, don't go anywhere just return the response
        return &http.Response{
            Status:        "200 OK",
            StatusCode:    200,
            Proto:         "HTTP/1.1",
            ProtoMajor:    1,
            ProtoMinor:    1,
            ContentLength: 11,
            Body:          ioutil.NopCloser(strings.NewReader("car painted")),
        }, nil
    }
    ```

### Monitoring/Metrics support

Export to either Prometheus/Datadog/statsd/etc, it's your choice. Bricks only provides the Interface and also **caches** the metrics so you don't have to.

```golang
counter := w.deps.Metrics.WithTags(monitor.Tags{
 "color":   request.GetDesiredColor(),
 "success": fmt.Sprintf("%t", err == nil),
}).Counter("paint_desired_color", "New paint color for car")

counter.Inc()
```

> `counter` is actually a *singleton*, uniqueness calculated [here](monitoring/registry.go#L87)

![grafana](wiki/grafana.png)

For more information about Bricks Monitoring read [here](https://go-bricks.github.io/middleware/telemetry/monitoring/).

### Additional Features

* `/debug/pprof` and other useful [handlers](handlers)
* Use `config_test.yml` during [tests](https://github.com/go-bricks/bricks-demo/blob/master/workshop/app/controllers/workshop_test.go#L151) to **override** values in `config.yml`, it saves time.

There are some features not listed here, please check the [Documentation](#documentation) for more.

## Documentation

Bricks is not a drop-in replacement.

It's important to read its [Documentation](https://go-bricks.github.io) first.

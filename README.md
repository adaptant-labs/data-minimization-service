# data-minimization-service

[![Build Status](https://travis-ci.com/adaptant-labs/data-minimization-service.svg?branch=master)](https://travis-ci.com/adaptant-labs/data-minimization-service)
[![Go Report Card](https://goreportcard.com/badge/github.com/adaptant-labs/data-minimization-service)](https://goreportcard.com/report/github.com/adaptant-labs/data-minimization-service)

A simple data minimization microservice wrapped around [go-minimizer].

[go-minimizer]: https://github.com/adaptant-labs/go-minimizer

## JSON interface

The data minimization service receives a JSON payload indicating the data, the type of data, and the level of
minimization to apply. These are written to the `/data` endpoint. The treated data is then returned directly as
result output.

The `type` here refers to the matching data type in [go-minimizer], such as name, email, etc. and is always a named
string. `level` similarly matches the `go-minimizer` minimization levels as a string representation. If `anonymize` is
specified, the `input` data may be omitted.

Input schema:

```json
{ 
  "input": <input data>,
  "type": "string",
  "level": "fine" | "coarse" | "anonymize"
}
```

Output schema:

```json
{
  "result": <output data>
}
```

Consumers of the JSON response must treat the result as a dynamic data type if they wish to use it generically,
as the data type backing the result will vary depending on the type of input data and the minimizer applied.
As there is no way to extract data from the service without entering something first, the caller should always
know which data type it is working with.

## Web Interface

A simplistic web interface is also included, accessible from the top-level index page.

## Features and bugs

Please file feature requests and bugs at the [issue tracker][tracker].

[tracker]: https://github.com/adaptant-labs/data-minimization-service/issues

## License

Licensed under the terms of the Apache 2.0 license, the full version of which can be found in the
[LICENSE](https://raw.githubusercontent.com/adaptant-labs/data-minimization-service/master/LICENSE) file included in the
distribution.

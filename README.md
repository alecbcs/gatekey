# GateKey

[![Go Report Card](https://goreportcard.com/badge/github.com/alecbcs/gatekey)](https://goreportcard.com/report/github.com/alecbcs/gatekey)

GateKey is a simple Go web server which generates and authenticates one time passwords for external connections.

## Installation

#### Dependencies

- `GCC`

- `Golang`

#### Build

1. WIth go installed simply run `go get github.com/alecbcs/gatekey`

or

1. Clone this repository and run

2. `go build` (This will build `gatekey` into a binary you can add to your `bin`.)

3. If you've added your go bin to your system PATH you can also run `go install`

## Configuration

Gatekey automatically builds a configuration directory in the `/home/USER/.config/gatekey` folder.

**Update your configuration with a real username the password before using GateKey**

#### Default Config Example:

```toml
[General]
  Version = "0.0.1"
  Port = 8080

[Authentication]
  User = ""
  Password = ""

[Database]
  Location = "/home/USER/.config/gatekey/tokens.db"

[Relay]
  Location = ""
  User = ""
  Password = ""
  TempFileLocation = "/home/USER/.config/gatekey/temp"

[Tokens]
  Length = 32
```

## Usage

#### Create

`curl -u "USERNAME":"PASSWORD" localhost:PORT/create/`

#### Report

`curl -F myFile=@/path/to/file localhost:PORT/report/TOKEN/`

## License

Copyright 2019-2020 Alec Scott <alecbcs@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

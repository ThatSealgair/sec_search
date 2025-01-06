# sec_search
A Go CLI tool for converting search queries between different security platform syntaxes (Shodan, Censys, Hunter).

Easy searching!

![Searching](https://c.tenor.com/1mOSY9SvIQYAAAAd/tenor.gif)

## Features
- Convert between Shodan, Censys, and Hunter search syntax
- TUI to help generate conversions

## Installation

```bash
go install github.com/ThatSealgair/sec_search@latest
```

## Usage
**Note**: Not all query parameters may have direct equivalents between platforms. The converter attempts to map the most common and relevant search parameters.

```bash
sec_search
```

Follow the TUI instructions.

### Go API:

_coming soon..._

## Building from source

```bash
git clone https://github.com/yourusername/sec_search.git
cd sec_search
go build cmd/sec_search/main.go
```

## Project Structure

```
sec_search/
├── cmd
│   └── sec_search
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── converter
│   │   ├── converter.go
│   │   ├── converter_test.go
│   │   └── types.go
│   ├── platform
│   │   ├── censys
│   │   │   ├── converter.go
│   │   │   ├── converter_test.go
│   │   │   └── filters.go
│   │   ├── hunter
│   │   │   ├── converter.go
│   │   │   ├── converter_test.go
│   │   │   └── filters.go
│   │   ├── mappings
│   │   │   └── mappings.go
│   │   ├── shodan
│   │   │   ├── converter.go
│   │   │   ├── converter_test.go
│   │   │   └── filters.go
│   │   ├── types.go
│   │   └── utils
│   │       ├── convert.go
│   │       └── convert_test.go
│   └── tui
│       ├── model.go
│       ├── state.go
│       └── styles.go
├── LICENSE
└── README.md
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)

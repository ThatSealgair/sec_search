# sec_search
Convert Shodan, Hunter, and Censys search queries.

Easy searching!

![Searching](https://c.tenor.com/1mOSY9SvIQYAAAAd/tenor.gif)

## Features

- Convert between Shodan and Censys search syntax

## Installation

Using UV:

```bash
uv venv
source .venv/bin/activate
uv pip install .
```

## Usage

Command line:
```bash
sec-search convert --from shodan --to censys "port:443 country:US"
```

Python API:
```python
from sec_search import QueryConverter

converter = QueryConverter()
censys_query = converter.convert_query("port:443 country:US", "shodan", "censys")
```

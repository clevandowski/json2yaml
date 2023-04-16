# json2yaml

Transform json input stream into yaml.

## Build

Requires go >= v1.20

```bash
go build
```

## Install

```bash
sudo ln -s ${PWD}/json2yaml /usr/local/bin/json2yaml
```

## Usage

```bash
cat myfile.json | json2yaml
```

```bash
echo '{ "test": 3 }' | json2yaml
```

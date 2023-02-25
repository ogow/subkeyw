# TBD

## Usage
```text
Usage of subkeyw:
  -dashes
    	include words with dashes, otherwise these are removed.
  -file string
    	takes a file with a list of subdomains
```

## Example

subs.txt
```text
api-test-dev1ish.target.com
a-test.target.com
api-blah.target.com
meh-blah.dev.target.com
```

```shell
$ cat subs.txt | subkeyw

api
test
dev1ish
a
blah
meh
dev
```

## Installing

```bash
$ go install -v github.com/dvvedz/subkeyw@latest
```

## Usage

```bash
$ cat subs.txt | subkeyw 
```

```bash
$ subkeyw -file subs.txt
```

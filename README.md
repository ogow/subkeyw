
 ```
            _     _                       
           | |   | |                      
  ___ _   _| |__ | | _____ _   ___      __
 / __| | | | '_ \| |/ / _ \ | | \ \ /\ / /
 \__ \ |_| | |_) |   <  __/ |_| |\ V  V / 
 |___/\__,_|_.__/|_|\_\___|\__, | \_/\_/  
                            __/ |         
                           |___/          
 ```


## Installing

```bash
$ go install -v github.com/dvvedz/subkeyw@latest
```

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

## Usage

```bash
$ cat subs.txt | subkeyw 
```

```bash
$ subkeyw -file subs.txt
```

## Known problems

For some reason stdin does not work on my droplet, but works fine on macos as intended...
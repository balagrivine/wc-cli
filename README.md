# wc-cli
wc command line tool

## Supported Flags
* **-b** **--byte**
* **-w** **--word**
* **-l** **--line**

## Usage
```shell
wc-cli <flag> <filename...>
```

## Basic Commands
* 1. Counting number of lines in a file
```shell
wc-cli -l <filename>
```

You can pass input from stdin
```shell
cat <filename> | wc-cli
```

* 2. Counting number of bytes in a file
```shell
wc-cli -m <filename>
```

* 3. Counting number of words in a file
```shell
wc-cli -w <filename>

## Improvements
Have more robust testing strategies, to test reading input from the stdin

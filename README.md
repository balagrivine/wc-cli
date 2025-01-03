# wc-cli
wc command line tool

## Supported Flags
* **-b** **--byte** - Print byte count
* **-w** **--word** - Print word count
* **-l** **--line** - Print newline count

## Usage
To read from a file use
```shell
wc-cli <flag> <filename...>
```

To read from the standard input use
```shell
cat <filename> | wc-cli <flag>

## Basic Commands
** 1. Printing number of lines in a file **
```shell
wc-cli -l <filename(s)>
```

You can pass input from stdin
```shell
cat <filename> | wc-cli
```

** 2. Printting number of bytes in a file **
```shell
wc-cli -m <filename(s)>
```

** 3. Printing number of words in a file **
```shell
wc-cli -w <filename(s)>
```

## Improvements
1. Have more robust testing strategies, to test reading input from the standard input

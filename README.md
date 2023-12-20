# go-wc
A Go CLI Tool for Word, Line, Byte, and Character Counting (like wc)
go-wc is a simple yet powerful command-line tool written in Go that mimics the functionality of the popular Unix utility wc. It counts the number of words, lines, bytes, and characters in a file or standard input.
![Screen Recording 2023-12-20 at 9 48 45 PM](https://github.com/manthanguptaa/go-wc/assets/42516515/72d42e3a-fbd2-4ea1-b06b-97458ab7fb77)

### Features
- Counts words, lines, bytes, and characters
- Supports reading from files or standard input
- Detailed output with options to show individual counts and total sums

### Getting Started

Follow the steps below to get started with the DNS server and resolver:

1. Clone the repository using Git:

   ```bash
   git clone https://github.com/manthanguptaa/go-wc.git
   ```

2. Change to the project directory:

   ```bash
   cd go-wc
   ```

3. Build the binary:

   ```bash
   go build
   ```

### Usage
```bash
go-wc [flags] [filename]

Flags:
  -c, --c      Outputs the numbers of bytes in the file
  -h, --help   help for go-wc
  -l, --l      Outputs the numbers of lines in the file
  -m, --m      Outputs the numbers of characters in the file
  -w, --w      Outputs the numbers of words in the file
```

You can find this by typing the following command in the terminal
```bash
go-wc -h
```
By default, go-wc outputs the number of words, lines, bytes, and characters for the file. You can use the `-m` and `-c` flags to display only characters and bytes, and `-w` and `-l` to focus on words or lines respectively.

Example: Get the number of words of a file `test.txt`
```bash
./go-wc -w test.txt
```

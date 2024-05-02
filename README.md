# GoShred

GoShred is a secure file deletion tool written in Go. It securely overwrites files with random data multiple times before deleting them, ensuring that the files are irrecoverable. This tool is inspired by the Unix/Linux `shred` command but built using Go's powerful standard libraries.

## Installation

To install GoShred, you need to have Go installed on your system. If you don't have Go installed, you can download and install it from [the official Go website](https://golang.org/dl/).

Once Go is installed, follow these steps to install GoShred:

1. Clone this repository or download the source code:

    ```bash
    git clone https://yourrepository.com/goshred.git
    cd goshred
    ```

2. Build the executable:

    ```bash
    go build -o shred
    ```

3. Optionally, move the executable to a directory in your PATH to make it globally accessible:

    ```bash
    sudo mv shred /usr/local/bin/
    ```

## Usage

To use GoShred, simply run the executable from the command line with the filename you wish to securely delete:

```bash
shred -passes=3 filename.txt
```

### Options

- `-passes`: Specifies the number of times the file should be overwritten with random data. The default is 3 passes.

### Example

Securely delete a file named `example.txt` with 5 passes:

```bash
shred -passes=5 example.txt
```

## How It Works

GoShred opens the target file and determines its size. It then generates a buffer of random bytes, the size of the file, and writes this buffer to the file, directly overwriting the existing data. This process is repeated for the number of passes specified. After the final pass, the file is deleted from the filesystem.

## Limitations

- GoShred is designed for use with conventional magnetic hard drives. It may be less effective on SSDs or other types of storage devices due to their different methods of handling data deletion and storage.
- The tool does not currently handle errors related to file system permissions or locked files.

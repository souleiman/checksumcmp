checksum
=========
Check file fingerprint based on provided hash value using shell commands

Requirements
------------
* Go 1.4.x or greater

## Installation
    # Use go get to install checksum
    go get github.com/souleiman/checksum

Usage
------------
    # To check a single file, filename.txt, via sha1
    checksum sha1sum filename.txt 535c5af68b10df7b04f8728bae3beccd748ee354

    # To check a file that contains multiple hashes for various files
    checksum -c multi_check.txt

    # Suppose the structure of multi_check.txt follows as below
    sha1sum filename.txt 535c5af68b10df7b04f8728bae3beccd748ee354
    md5sum file1.txt 91d7f7bf36fbfd038a4bd1dae736cd1e
    sha256sum file2.txt 8b6ed36d55752ca35871c588c28a6853a998c8e75c64b726d3d5d6afab2402ca

``` 
checksum [--use-exec] <function> <file> <hash>
checksum [--use-exec] -c <file>
checksum -h | --help
checksum -v | --version

--use-exec      Uses your exec to execute the hashsum, as opposed to using internals [Not Supported for Windows]
-c              Check file containing multiple functions followed by file and fingerprint.
                    i.e sha1sum file.txt deadbeef
-h --help       Shows informative information of checksum command.
-v --version    Show version.
```
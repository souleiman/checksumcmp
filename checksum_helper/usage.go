package checksum_helper

const Usage string= `Checksum

Usage:
    checksum [--use-exec] <function> <file> <hash>
    checksum [--use-exec] -c <file>
    checksum -h | --help
    checksum -v | --version

Options:
    --use-exec      Uses your exec to execute the hashsum, as opposed to using internals [Not Supported for Windows]
    -c              Check file containing multiple functions followed by file and fingerprint.
                        i.e sha1sum file.txt deadbeef
    -h --help       Shows informative information of checksum command.
    -v --version    Show version.`
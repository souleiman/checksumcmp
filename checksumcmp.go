package main

/**
 * Soul
 */

import (
    "fmt"

    "github.com/docopt/docopt-go"
    check "github.com/souleiman/checksumcmp/checksum_helper"
)

func main() {
    args, _ := docopt.Parse(check.Usage, nil, true, check.Version, false)

    if args["-c"].(bool) {
        file := args["<file>"].(string)
        check.CompareFiles(file, args["--use-exec"].(bool))
    } else {
        function, file, hash := args["<function>"], args["<file>"], args["<hash>"]
        fmt.Println(check.CompareFile(function.(string), file.(string), hash.(string), args["--use-exec"].(bool)))
    }
}
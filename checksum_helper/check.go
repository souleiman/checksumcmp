package checksum_helper
import (
    "os/exec"
    "strings"
    "io/ioutil"
    "fmt"
    "crypto/sha512"
    "crypto/sha256"
    "crypto/sha1"
    "crypto/md5"
    "hash"
)

func CompareFile(function, file, fingerprint string, u_exec bool) string {
    var hash string
    var raw []byte
    var err error

    if u_exec {
        hash, err = onExec(function, file)
    } else {
        raw, err = computeSum(hashmap[function], file)
        hash = fmt.Sprintf("%x", raw)
    }

    if err != nil {
        return fmt.Sprintf("Error on %s - %s", file, err.Error())
    }


    if hash == fingerprint {
        return fmt.Sprintf("Fingerprint of %s matches.", file)
    }
    return fmt.Sprintf("Fingerprint of %s does not match.", file)
}

func CompareFiles(file string, u_exec bool) {
    content, err := ioutil.ReadFile(file)

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    split := strings.Split(strings.TrimSpace(string(content)), "\n")

    for index, line := range split {
        line_index := index + 1
        line = strings.TrimSpace(line)
        if len(line) == 0 {
            continue
        }

        fields := strings.Fields(line)

        if len(fields) != 3 {
            fmt.Printf("Unable to parse line %d.\n", line_index)
            continue
        }

        if !strings.HasSuffix(fields[0], "sum") {
            fmt.Printf("unknown hashing command at line %d.\n", line_index)
            continue
        }

        fmt.Println(CompareFile(fields[0], fields[1], fields[2], u_exec))
    }
}

func onExec(function, file string) (string, error) {
    cmd := exec.Command(function, file)
    output, err := cmd.Output()
    if err != nil {
        return "", err
    }

    values := strings.Fields(string(output))
    return values[0], nil
}

var hashmap map[string]hash.Hash = map[string]hash.Hash { // No pun intended :P
    "md5sum": md5.New(),
    "sha1sum": sha1.New(),
    "sha224sum": sha256.New224(),
    "sha256sum": sha256.New(),
    "sha384sum": sha512.New384(),
    "sha512sum": sha512.New(),
}
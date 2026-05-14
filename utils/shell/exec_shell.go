// https://github.com/Zeronetsec/TwrapKit

package shell

import (
    "bytes"
    "os/exec"
)

func ExecShell(command string, args ...string) (*Result, error) {
    cmd := exec.Command(command, args...)

    var stdout bytes.Buffer
    var stderr bytes.Buffer

    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    err := cmd.Run()

    result := &Result{
        Stdout: stdout.String(),
        Stderr: stderr.String(),
        Code: 0,
    }

    if err != nil {
        result.Code = 1
        return result, err
    }

    return result, nil
}

// Copyright (c) 2026 Zeronetsec
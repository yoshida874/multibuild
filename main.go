package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
    if err := build(context.Background()); err != nil {
        fmt.Println(err)
    }
}

func build(ctx context.Context) error {
    fmt.Println("Building with Dagger")

    client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
    if err != nil {
        return err
    }
    defer client.Close()

    src := client.Host().Directory(".")

    golang := client.Container().From("golang:latest")

    golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")

    path := "build/"
    golang = golang.WithExec([]string{"go", "build", "-o", path})

    output := golang.Directory(path)

    _, err = output.Export(ctx, path)
    if err != nil {
        return err
    }

    return nil
}

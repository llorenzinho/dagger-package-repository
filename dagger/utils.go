package main

import (
	"context"

	"golang.org/x/exp/slices"
)

func (m *Ci) getDockerDep() *Container {
	return dag.Container().From("python:3.11.7-slim").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "--no-install-recommends", "netcat-traditional"}).
		WithExec([]string{"apt-get", "install", "-y", "gcc", "libpq-dev", "python3-dev", "tini"}).
		WithExec([]string{"rm", "-rf", "/var/lib/apt/lists/*", "/tmp/*", "/var/tmp/*"})
}

func (m *Ci) getDocker(dir *Directory) *Container {
	return m.DockerDep.
		WithWorkdir("/app").
		WithDirectory("/app", dir).
		WithExec([]string{"pip", "install", "poetry==1.6.1"}).
		WithExec([]string{"poetry", "config", "virtualenvs.create", "false"}).
		WithExec([]string{"poetry", "install", "--no-dev", "--no-interaction", "--no-ansi"}).
		WithEntrypoint([]string{"poetry run uvicorn painting_recognition:app --host 0.0.0.0"})
}

func (m *Ci) getSimplePythonDocker(dir *Directory) *Container {
	files, err := dir.Entries(context.Background(), DirectoryEntriesOpts{
		Path: ".",
	})
	if err != nil {
		panic(err)
	}
	if !slices.Contains(files, "requirements.txt") {
		panic("missing requirements.txt")
	}
	return dag.Container().From("python:3.11-alpine3.19").
		WithWorkdir("/app").
		WithDirectory("/app", dir).
		WithExec([]string{"pip", "install", "-r", "requirements.txt"})
}

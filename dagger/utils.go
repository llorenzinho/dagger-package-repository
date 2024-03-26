package main

func (m *Ci) getDocker(dir *Directory) *Container {
	return dag.Container().From("python:3.11-alpine3.19").
		WithWorkdir("/app").
		WithDirectory("/app", dir).
		WithExec([]string{"pip", "install", "poetry==1.8.1"}).
		WithExec([]string{"poetry", "config", "virtualenvs.create", "false"}).
		WithExec([]string{"poetry", "install", "--no-dev", "--no-interaction", "--no-ansi"}).
		WithEntrypoint([]string{"poetry run uvicorn painting_recognition:app --host 0.0.0.0"})
}

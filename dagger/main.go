// A generated module for Ci functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"fmt"
)

type Ci struct {
	DockerDep *Container
}

func (m *Ci) WithDockerDeps() *Ci {
	m.DockerDep = m.getDockerDep()
	return m
}

func (m *Ci) BuildDocker(ctx context.Context, dir *Directory) *Container {
	return m.getDocker(dir)
}

func (m *Ci) Ciao(ctx context.Context) string {
	return "ciao"
}

// Echo returns a string
func (m *Ci) Echo(ctx context.Context, msg string) string {
	return fmt.Sprintf("echo %s", msg)
}

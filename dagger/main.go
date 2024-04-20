// A generated module for Toledo functions
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
	"math"
	"math/rand"
)

type Toledo struct{}

// Builds and publishes a container image from a given source directory.
func (m *Toledo) BuildAndPublish(ctx context.Context, buildSrc *Directory, appName string) (string, error) {
	ctr := dag.Wolfi().Container()

	return dag.
		Golang().
		BuildContainer(GolangBuildContainerOpts{Source: buildSrc, Base: ctr}).
		Publish(ctx, fmt.Sprintf("ttl.sh/%s-%.0f", appName, math.Floor(rand.Float64()*10000000))) //#nosec
}

// Builds and publishes a container image from a given source directory.
func (m *Toledo) Run(ctx context.Context, buildSrc *Directory, appName string) *Service {
	ctr := dag.Wolfi().Container()
	appPath := fmt.Sprintf("/usr/local/bin/%s", appName)

	return dag.
		Golang().
		BuildContainer(GolangBuildContainerOpts{Source: buildSrc, Base: ctr}).
		WithEntrypoint([]string{appPath}).
		WithExposedPort(8080).
		AsService()
}

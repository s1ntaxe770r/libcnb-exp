package main

import (
	"context"
	"os"

	"github.com/buildpacks/pack/pkg/client"
	"github.com/buildpacks/pack/pkg/image"
	"github.com/sirupsen/logrus"
)

func main() {

	c, err := client.NewClient()
	ctx := context.Background()

	if err != nil {
		logrus.Error(err)
	}
		
	envs := make(map[string]string)
	
	envs["CGO_ENABLED"] = "0"
	envs["BP_GO_BUILD_FLAGS"] = "-buildmode=default"
	
	wd, err := os.Getwd()

	if err != nil {
		logrus.Error(err)
		return
	}


    // should be the equivalent of  pack build my-app  --buildpack paketo-buildpacks/go  --builder paketobuildpacks/builder-jammy-buildpackless-static
    buildOpts := client.BuildOptions{
        Image:           "pong",
        Builder:         "paketobuildpacks/builder-jammy-buildpackless-static",
        Buildpacks:      []string{"paketo-buildpacks/go"},
        AppPath:         "./sample",
        PullPolicy:      image.PullIfNotPresent,
        RelativeBaseDir: wd + "/sample",
        DockerHost:      "unix:///var/run/docker.sock",
        TrustBuilder:    client.IsSuggestedBuilderFunc,
        Env:             envs,
        LifecycleImage:  "buildpacksio/lifecycle",
    }
    err = c.Build(ctx, buildOpts)
    if err != nil {
        logrus.Error(err)
    }

}

package plugins

import (
	"fmt"

	"github.com/evanw/esbuild/pkg/api"
)

func buildPluginScript(script string, name string) (string, error) {
	res := api.Transform(script, api.TransformOptions{
		Format:     api.FormatIIFE,
		GlobalName: "smmPlugins['" + name + "']",
	})
	if len(res.Errors) > 0 {
		fmt.Println(res.Errors)
		return "", fmt.Errorf("Error transforming plugin script.")
	}

	return string(res.Code), nil
}
package appdefinition

import (
	"testing"

	"github.com/acorn-io/baaah/pkg/router/tester"
	v1 "github.com/acorn-io/runtime/pkg/apis/internal.acorn.io/v1"
	"github.com/acorn-io/runtime/pkg/appdefinition"
	"github.com/acorn-io/runtime/pkg/scheme"
)

func TestParseAppImage(t *testing.T) {
	tester.DefaultTest(t, scheme.Scheme, "testdata/parseappimage", ParseAppImage)
}

func TestParseAppImageDevMode(t *testing.T) {
	tester.DefaultTest(t, scheme.Scheme, "testdata/parsedevmode", ParseAppImage)
}

func TestParseAppImageBug(t *testing.T) {
	appImage := &v1.AppImage{
		ImageData: v1.ImagesData{
			Containers: map[string]v1.ContainerData{},
		},
		Acornfile: "",
		ID:        "",
	}
	app, err := appdefinition.FromAppImage(appImage)
	if err != nil {
		t.Fatal(err)
	}

	_, err = app.AppSpec()
	if err != nil {
		t.Fatal(err)
	}
}

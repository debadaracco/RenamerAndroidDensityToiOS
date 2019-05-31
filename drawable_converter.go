package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Resource struct {
	folder string
	iosDensity string
}

func NewResource(folder string, iosDensity string) *Resource {
	return &Resource{folder: folder, iosDensity: iosDensity}
}


func main() {

	var resources = [3]*Resource{ NewResource("drawable-mdpi", ""),
		NewResource("drawable-xhdpi", "@2x"),
		NewResource("drawable-xxhdpi", "@3x")}

	currentPath, error := os.Getwd()

	if error != nil {
		panic(error)
	}

	fmt.Println(currentPath)

	_, error = os.Stat(currentPath + "/drawables-ios")

	if  error == nil {
		error = os.RemoveAll(currentPath + "/drawables-ios")
		if error != nil {
			panic(error)
		}else{
			fmt.Println("Remove folder")
		}
	}

	iosFolder := currentPath + "/drawables-ios"

	error = os.MkdirAll(iosFolder, os.ModePerm)

	if error != nil {
		panic(error)
	}

	for e := range resources {
		fmt.Println("==================" + resources[e].folder + "==================")

		resourcePath :=  currentPath + "/" + resources[e].folder

		resourceDensity := resources[e].iosDensity

		_, error := os.Stat(resourcePath)

		if error == nil {

			files, err := ioutil.ReadDir(resourcePath)

			if err == nil {
				for _, f := range files {

					name := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))

					cmd := exec.Command("cp",
						resourcePath + "/" + f.Name(),
						iosFolder + "/" + name + resourceDensity + filepath.Ext(f.Name()))

					err = cmd.Run()
					if err != nil {
						panic(err)
					}
				}
			}else{
				panic(err)
			}

		}else{
			fmt.Println(error.Error())
		}
	}




}

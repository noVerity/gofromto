//gobuild:ignore

package main

import (
	"github.com/noVerity/gofromto"
	"log"
	"os"
	"text/template"
)

func main() {
	f, err := os.Create("conversion_map.go")
	die(err)
	defer f.Close()

	conversionMap := make(map[gofromto.Unit]map[gofromto.Unit][]int)

	for index, c := range gofromto.ConversionList {
		if _, prs := conversionMap[c.From]; !prs {
			conversionMap[c.From] = make(map[gofromto.Unit][]int)
		}
		conversionMap[c.From][c.To] = []int{index}
	}

	madeAdditions := true

	for madeAdditions {
		madeAdditions = false

		for from, toMap := range conversionMap {
			for to, conversionPath := range toMap {
				if extensionsViaTo, prs := conversionMap[to]; prs {
					for newTo, newConversionPath := range extensionsViaTo {
						if from == newTo {
							continue
						}
						if _, prs := toMap[newTo]; !prs {
							conversionMap[from][newTo] = append(newConversionPath, conversionPath...)
							madeAdditions = true
						}
					}
				}
			}
		}
	}

	packageTemplate, err := template.New("conversion_map.tmpl").ParseFiles("conversion_map.tmpl")
	die(err)

	err = packageTemplate.Execute(f, struct {
		ConversionMap map[gofromto.Unit]map[gofromto.Unit][]int
	}{
		ConversionMap: conversionMap,
	})
	die(err)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

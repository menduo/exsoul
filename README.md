# ExSoul

Golang excel utils waraps 360EntSecGroup-Skylar/excelize.


## Usage


```go

// Created by @menduo @ 2020/7/29
package main

import (
	"fmt"
	"github.com/menduo/exsoul"
)

/*
Date 	Year	Name 	Caption
2020/1/1	2012	Galaxy Leo IV	Leo IV is one of more than a dozen ultra-faint dwarf galaxies near the Milky Way...
2020/1/2	2005	Galaxy Cluster SDSS J1004+4112	This picture captures a galaxy ..
2020/1/3	2017	NGC 4302 and NGC 4298	This image captures two spiral galaxies...
*/

func main() {
	esfile, err := exsoul.NewFromFile("./example/hubble-birthdays-example-menduo.xlsx")

	if err != nil {
		fmt.Println("failed to load exsoul", err)
		return
	}
	rows := esfile.LoadRowsFromFirstSheet()
	for _, row := range rows {
		isValidData, err := row.IsColValidDateStr(0, "2006/1/2")
		if err != nil {
			fmt.Println("err == ", err)
			continue
		}

		fmt.Println("isValidData", isValidData)

		year, err := row.GetColAsInt(1)
		if err != nil {
			fmt.Println("failed to get year")
			continue
		}
		fmt.Println("year===", year)
	}

	e2, err := esfile.Clone()
	if err != nil {
		fmt.Println("failed to Clone exsoul", err)
		return
	}

	e2.SetRow(5, &[]interface{}{"2020/1/4", 2020, "menduo/exsoul", "Golang excel utils github.com/menduo/exsoul"})
	e2.SaveAs("./example/hubble-birthdays-example-menduo-new.xlsx")
}

```

## License

Mit


## Contact

shimenduo@gmail.com

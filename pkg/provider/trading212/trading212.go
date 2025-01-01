/*
Copyright © 2019 Ce Gao

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trading212

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/deb-sig/double-entry-generator/pkg/io/reader"
	"github.com/deb-sig/double-entry-generator/pkg/ir"
)

// Trading212 is the provider for Trading212.
type Trading212 struct {
	Statistics Statistics `json:"statistics,omitempty"`
	LineNum    int        `json:"line_num,omitempty"`
	Orders     []Order    `json:"orders,omitempty"`

	// TitleParsed is a workaround to ignore the title row.
	TitleParsed bool `json:"title_parsed,omitempty"`
}

// New creates a new Trading212 provider.
func New() *Trading212 {
	return &Trading212{
		Statistics:  Statistics{},
		LineNum:     0,
		Orders:      make([]Order, 0),
		TitleParsed: false,
	}
}

// Translate translates the trading212 bill records to IR.
func (a *Trading212) Translate(filename string) (*ir.IR, error) {
	log.SetPrefix("[Provider-Trading212] ")

	billReader, err := reader.GetReader(filename)
	if err != nil {
		return nil, fmt.Errorf("can't get bill reader, err: %v", err)
	}

	csvReader := csv.NewReader(billReader)
	csvReader.LazyQuotes = true
	// If FieldsPerRecord is negative, no check is made and records
	// may have a variable number of fields.
	csvReader.FieldsPerRecord = -1
	// Read into a map by the header
	header, err := csvReader.Read()
	a.LineNum++
	if err != nil {
		return nil, fmt.Errorf("failed to read the header: %v", err)
	}
	log.Printf("Header: %v", header)

	// Read the records into a list of map
	for {
		line, err := csvReader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		a.LineNum++
		m := make(map[string]string)
		for idx, a := range line {
			a = strings.Trim(a, " ")
			a = strings.Trim(a, "\t")
			m[header[idx]] = a
		}
		err = a.translateToOrdersFromMap(m)
		if err != nil {
			return nil, fmt.Errorf("Failed to translate bill: line %d: %v",
				a.LineNum, err)
		}
	}

	log.Printf("Finished to parse the file %s", filename)

	ir := a.convertToIR()
	return ir, nil
}

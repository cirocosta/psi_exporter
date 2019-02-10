package parser

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Row struct {
	Type   string
	Avg10  float64
	Avg60  float64
	Avg300 float64
	Total  float64
}

func Parse(r io.Reader) (rows []*Row, err error) {
	if r == nil {
		err = errors.Errorf("a reader must be provided")
		return
	}

	var (
		scanner = bufio.NewScanner(r)
		row     *Row
	)

	rows = make([]*Row, 0)
	for scanner.Scan() {
		row, err = parseRow(scanner.Text())
		if err != nil {
			err = errors.Wrapf(err,
				"failed parsing psi row")
			return
		}

		rows = append(rows, row)
	}

	return
}

func parseRow(input string) (row *Row, err error) {
	if input == "" {
		err = errors.Errorf("input must not be empty")
		return
	}

	fields := strings.Fields(input)
	if len(fields) != 5 {
		err = errors.Errorf("expected exactly 5 fields from '%s', got '%d'",
			input, len(fields))
		return
	}

	row = &Row{Type: fields[0]}
	structVals := []*float64{
		&row.Avg10,
		&row.Avg60,
		&row.Avg300,
		&row.Total,
	}

	for idx, field := range fields[1:5] {
		parts := strings.Split(field, "=")
		if len(parts) != 2 {
			err = errors.Errorf("expected two parts for '%s', got %d",
				field, len(parts))
			return
		}

		*structVals[idx], err = strconv.ParseFloat(parts[1], 64)
		if err != nil {
			err = errors.Wrapf(err,
				"failed to parse counter in field %s", field)
			return
		}
	}

	return
}

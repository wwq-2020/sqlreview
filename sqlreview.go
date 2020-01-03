package main

import (
	"fmt"
	"github.com/pingcap/parser"
	"github.com/pingcap/parser/ast"
	_ "github.com/pingcap/tidb/types/parser_driver"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("fail to read from stdin")
		os.Exit(0)
	}

	p := parser.New()
	stmt, warns, err := p.Parse(string(data), "utf8mb4", "utf8mb4")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	if len(warns) != 0 {
		for _, warn := range warns {
			fmt.Println(warn.Error())
		}
	}
	for _, each := range stmt {
		createTableStmt, ok := each.(*ast.CreateTableStmt)
		if !ok {
			continue
		}
		columns := map[string]struct{}{}
		for _, col := range createTableStmt.Cols {
			columns[col.Name.Name.String()] = struct{}{}
		}

		for _, constraint := range createTableStmt.Constraints {
			parts := strings.Split(constraint.Name, "-")

			if constraint.Tp == ast.ConstraintIndex {
				if len(parts) < 2 || parts[0] != "idx" {
					fmt.Println("index should been idx-<col>-<col>")
					os.Exit(0)
				}
				for _, part := range parts[1:] {
					if _, exist := columns[part]; !exist {
						fmt.Println("idx should been idx-<col>-<col>")
						os.Exit(0)
					}
				}
			}
			if constraint.Tp == ast.ConstraintUniq {
				if len(parts) < 2 || parts[0] != "idx" {
					fmt.Println("idx should been idx-<col>-<col>")
					os.Exit(0)
				}
				for _, part := range parts[1:] {
					if _, exist := columns[part]; !exist {
						fmt.Println("uniq should been uniq-<col>-<col>")
						os.Exit(0)
					}
				}
			}
		}
	}
}

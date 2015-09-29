package handlers

import (
	"bitbucket.org/harshadp/ontotestpad/src/common"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func Loincsearchterm(w http.ResponseWriter, r *http.Request) {
	searchstring := r.URL.Query().Get("q")
	searchtype := r.URL.Query().Get("st")
	// TODO: Check for injection
	if searchstring != "" {
		resp, err := getloincJSON(searchstring, searchtype)
		check(err, false)
		fmt.Fprintf(w, resp)
	}
}

/*
wordsanyorder
phrasematch
identicalterm
startswith
endswith
*/

func getloincJSON(searchstring string, searchtype string) (string, error) {
	var queryString string
	var NumRows int64
	switch searchtype {
	}
	queryString = "select long_common_name as title, loinc_num as description from loinc where match('@long_common_name " + searchstring + "*') LIMIT 0, 1000;"
	rows, err := common.Dbloinc.Query(queryString)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
		NumRows = NumRows + 1
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		return "", err
	}
	return "{ \x22success\x22: true,\x22results\x22: { \x22terms\x22: { \x22name\x22: \x22Found " + strconv.FormatInt(NumRows, 10) + " Terms\x22, \x22results\x22: " + string(jsonData) + "}}}", nil
}

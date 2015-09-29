package handlers

import (
	"bitbucket.org/harshadp/ontotestpad/src/common"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Snomedsearchterm(w http.ResponseWriter, r *http.Request) {
	searchstring := r.URL.Query().Get("q")
	searchtype := r.URL.Query().Get("st")
	// TODO: Check for injection
	if searchstring != "" {
		resp, err := getJSON(searchstring, searchtype)
		check(err, false)
		fmt.Fprintf(w, resp)
	}
}

func check(e error, exit bool) {
	if e != nil {
		fmt.Println(e)
		if exit == true {
			panic(e)
			os.Exit(1)
		}
	}
}

/*
wordsanyorder
phrasematch
identicalterm
startswith
endswith
*/

func getJSON(searchstring string, searchtype string) (string, error) {
	var queryString string
	var NumRows int64
	switch searchtype {
	case "wordsanyorder":
		queryString = "select title, description from sct_pd where match('@title " + searchstring + "*') LIMIT 0, 1000;"
	case "phrasematch":
		queryString = "select title, description from sct_pd where match('@title \"" + searchstring + "\"') LIMIT 0, 1000;"
	case "identicalterm":
		queryString = "select title, description from sct_pd where match('@title ^" + searchstring + "$') LIMIT 0, 1000;"
	case "startswith": //TODO: a space at the end of the search string matches with everything
		queryString = "select title, description from sct_pd where match('@title ^" + searchstring + "') LIMIT 0, 1000;"
	case "endswith":
		queryString = "select title, description from sct_pd where match('@title " + searchstring + "$') LIMIT 0, 1000;"
	default:
		queryString = "select title, description from sct_pd where match('@title ^" + searchstring + "*') LIMIT 0, 1000;"
	}
	rows, err := common.Dbsctpd.Query(queryString)
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

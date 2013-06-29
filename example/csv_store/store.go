package main

import (
    "encoding/csv"
    "github.com/davecgh/go-spew/spew"
    "io"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

type DataBase struct {
    filePath string
    columns  []string
    records  []map[string]string
}

func Open (filePath string) (*DataBase, error) {

    // opening and reading a file
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // creating new database object
    db := new (DataBase)
    db.filePath = filePath

    // starting reading a file
    reader := csv.NewReader(file)
    count := 0

    for {

        // reading a row of a file
        record, err := reader.Read()

        // breaking if end of file reached
        if err == io.EOF {
            break
        } else {

            // filling columns
            if count == 0 {
                for _, record_field := range record {
                    db.columns = append(db.columns, strings.TrimSpace(record_field))
                }
            } else {

                // filling the database records
                databaseRecord := map[string]string{}

                for index, record_field := range record {
                    databaseRecord[db.columns[index]] = strings.TrimSpace(record_field)
                }

                db.records = append(db.records, databaseRecord)
            }
            count++
        }
    }
    return db, nil
}


func (db *DataBase) Where (needleKey, needleValue string) (result []map[string]string) {
    for _, record := range db.records {
        for key, value := range record  {
            if needleKey == key && needleValue == value {
              result = append (result, record)
            }
        }
    }
    return
}


func (db *DataBase) WhereLike (needleKey, needleValue string) (result []map[string]string) {
    for _, record := range db.records {
        for key, value := range record  {
            if needleKey == key &&
               strings.Index(
                 strings.ToLower(value),
                 strings.ToLower(needleValue)) >= 0 {
              result = append (result, record)
            }
        }
    }
    return
}


func (db *DataBase) Add (record map[string]string) {
    db.records = append (db.records, record)
    db.update ()
}


func (db *DataBase) Remove (needleKey, needleValue string) {
    tmpRecords := []map[string]string{}

    // filtering records
    R: for _, record := range db.records {
        for key, value := range record {
            if key == needleKey && value == needleValue {
                break R
            }
        }
        tmpRecords = append (tmpRecords, record)
    }

    // updating db redord
    db.records = tmpRecords
    db.update ()
}


func (db *DataBase) update () {
    newData := ""

    // writing column names
    newData += strings.Join(db.columns, ",") + "\n"

    // writing records
    for _, record := range db.records {

        // collecting values in correct order
        values := []string{}
        for _, column := range db.columns {
            values = append (values, record[column])
        }
        newData += strings.Join(values, ",") + "\n"
    }

    // updating database file and object
    err := ioutil.WriteFile(db.filePath, []byte(newData), 0644)
    if err != nil {
        log.Fatal (err)
    }
}


func main () {
    db, _ := Open ("names.txt")
    /* db.Remove ("last_name", "Zibert") */

    db.Add (map[string]string {
      "first_letter": "z",
      "first_name": "Zahria",
      "last_name": "Johnes"
    })

    results := db.WhereLike("last_name", "joh")

    spew.Dump(results)
}

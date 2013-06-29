package store

import (
    "encoding/csv"
    "github.com/davecgh/go-spew/spew"
    "io"
    "io/ioutil"
    "log"
    "os"
    "strings"
)

type Store struct {
    filePath string
    columns  []string
    records  []map[string]string
}

/**
 *  Opening new store file and creating new Store object
 */
func Open (filePath string) (*Store, error) {

    // opening and reading a file
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    store := new (Store)
    store.filePath = filePath

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
                    store.columns = append(store.columns, strings.TrimSpace(record_field))
                }
            } else {

                // filling the store records
                storeRecord := map[string]string{}

                for index, record_field := range record {
                    storeRecord[store.columns[index]] = strings.TrimSpace(record_field)
                }

                store.records = append(store.records, storeRecord)
            }
            count++
        }
    }
    return store, nil
}


/**
 *  Searching for an exact matching
 */
func (store *Store) Where (needleKey, needleValue string) (result []map[string]string) {
    for _, record := range store.records {
        for key, value := range record  {
            if needleKey == key && needleValue == value {
              result = append (result, record)
            }
        }
    }
    return
}


/**
 *  Searching for an partial entry ignoring case
 */
func (store *Store) WhereLike (needleKey, needleValue string) (result []map[string]string) {
    for _, record := range store.records {
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


/**
 *  Adding record to store
 */
func (store *Store) Add (record map[string]string) {
    store.records = append (store.records, record)
    store.update ()
}


/**
 *  Updating a record in a store
 */
func (store *Store) Update (needleKey, needleValue string, record map[string]string) {}


/**
 *  Deleting record from store
 */
func (store *Store) Remove (needleKey, needleValue string) {
    tmpRecords := []map[string]string{}

    // filtering records
    R: for _, record := range store.records {
        for key, value := range record {
            if key == needleKey && value == needleValue {
                break R
            }
        }
        tmpRecords = append (tmpRecords, record)
    }

    // updating store redord
    store.records = tmpRecords
    store.update ()
}


/**
 *  Updating a store file
 */
func (store *Store) update () {
    newData := ""

    // writing column names
    newData += strings.Join(store.columns, ",") + "\n"

    // writing records
    for _, record := range store.records {

        // collecting values in correct order
        values := []string{}
        for _, column := range store.columns {
            values = append (values, record[column])
        }
        newData += strings.Join(values, ",") + "\n"
    }

    // updating store file and object
    err := ioutil.WriteFile(store.filePath, []byte(newData), 0644)
    if err != nil {
        log.Fatal (err)
    }
}

// Usage example
func main () {
    store, _ := Open ("names.csv")
    /* store.Remove ("last_name", "Zibert") */

    store.Add (map[string]string {
      "first_letter": "z",
      "first_name": "Zahria",
      "last_name": "Johnes" })

    results := store.WhereLike("last_name", "joh")

    spew.Dump(results)
}

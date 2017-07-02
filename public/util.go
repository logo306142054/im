/*
* created by Orz on 2017/6/11
*/
package public

import (
    "time"
    "encoding/json"
)

func GetCurrTime() string {
    return time.Now().Format("2006-01-02 15:04:06")
}

func ParseToJsonBytes(v interface{}) []byte {
    if b, err := json.Marshal(v); err != nil {
        return make([]byte, 0)
    } else {
        return b
    }
}

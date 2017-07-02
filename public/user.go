/*
* created by Orz on 2017/6/11
*/
package public

type Gender int

const (
    MALE = iota
    FEMALE
)

func (g *Gender) String() string {
    if *g == MALE {
        return "male"
    } else if *g == FEMALE {
        return "female"
    } else {
        return "unknowm"
    }
}

type User struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Gender Gender `json:"gender"`
}

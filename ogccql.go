package tegola

import (
    "fmt"
    "log"
    "regexp"
)

type CQLFilter struct {
    Lang string
    FilterText string
}

func NewCQLFilter(filter_lang string, filter_text string) (*CQLFilter, error) {
    cf := CQLFilter{}
    cf.Lang = filter_lang
    cf.FilterText = filter_text

    return &cf, nil
}

func (c *CQLFilter) Parse() error {
    return nil
}

func (c *CQLFilter) ToSQL(sql_statement string, srid uint64) (string, error) {
    var filter_text string
    var value string

    log.Println("Adjusting spatial predicates")
    re := regexp.MustCompile("(?i)(BEYOND|CONTAINS|CROSSES|DISJOINT|WITHIN|EQUALS|INTERSECTS|OVERLAPS|TOUCHES)")
	filter_text = re.ReplaceAllString(c.FilterText, "ST_$1")

    log.Println("Adjusting geometry representation")
	re2 := regexp.MustCompile(`(?i)(POINT\s?\(.*\d\)|LINESTRING\s?\(.*\d\)|POLYGON\s?\(\(.*\d\)\))`)
 	filter_text = fmt.Sprintf(re2.ReplaceAllString(filter_text, "ST_Transform('SRID=4326;$1'::geometry, %d)"), srid)

    log.Println("Adjusting temporal predicates")
    re3 := regexp.MustCompile("(?i)BEFORE,ENDS")
	filter_text = re3.ReplaceAllString(filter_text, "<")
    re4 := regexp.MustCompile("(?i)AFTER,BEGINS")
	filter_text = re4.ReplaceAllString(filter_text, ">")
    re5 := regexp.MustCompile("(?i)TEQUALS")
	filter_text = re5.ReplaceAllString(filter_text, "=")

    log.Println("Assembling final SQL statement")

    matched, err := regexp.MatchString("(?i) WHERE ", sql_statement)

    if err != nil {
        fmt.Errorf("error matching SQL statement")
    }

    if matched {
        log.Println("appending to existing where clause")
        value = sql_statement + " AND " + filter_text
    } else {
        log.Println("defining where clause")
        value = sql_statement + " WHERE " + filter_text
    }
    return value, nil
}

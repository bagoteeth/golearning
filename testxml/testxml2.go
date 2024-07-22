/*
*

	@author rfjiang
	@date 2024/7/15
	@desc

*
*/
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"regexp"
	"strings"
)

type Xml21 struct {
	XMLName xml.Name `xml:"xml21"`
	Xml22   string   `xml:"xml22"`
	Xml23   string   `xml:"xml23"`
}

//type Xml22 struct {
//	Name string `xml:"name"`
//}
//
//type Xml23 struct {
//	Age int `xml:"age"`
//}

func main() {
	xml0 := `
<?xml version='1.1' encoding='GBK'?>
<xml21 xmlns="http://xmlns.oracle.com/weblogic/domain" xmlns:sec="http://xmlns.oracle.com/weblogic/security" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:wls="http://xmlns.oracle.com/weblogic/security/wls" xsi:schemaLocation="http://xmlns.oracle.com/weblogic/security/wls http://xmlns.oracle.com/weblogic/security/wls/1.0/wls.xsd http://xmlns.oracle.com/weblogic/domain http://xmlns.oracle.com/weblogic/1.0/domain.xsd http://xmlns.oracle.com/weblogic/security/xacml http://xmlns.oracle.com/weblogic/security/xacml/1.0/xacml.xsd http://xmlns.oracle.com/weblogic/security/providers/passwordvalidator http://xmlns.oracle.com/weblogic/security/providers/passwordvalidator/1.0/passwordvalidator.xsd http://xmlns.oracle.com/weblogic/security http://xmlns.oracle.com/weblogic/1.0/security.xsd">
  <xml22>wl_server</xml22>
  <xml23>14.1.1.0.0</xml23>
</xml21>
`
	re := regexp.MustCompile("(?i)<\\?xml[^>]*\\?>")
	match := re.FindString(xml0)
	fmt.Printf("match: %s\n", match)
	reversion := regexp.MustCompile("version=[^ ]+")
	match2 := reversion.ReplaceAllString(match, "version=\"1.0\"")
	xml1 := strings.Replace(xml0, match, match2, 1)
	fmt.Printf("xml1: %s\n", xml1)

	var a Xml21
	//b := []byte(xml0)
	//b = []byte(strings.Replace(string(b), "version='1.1'", "version='1.0'", 1))
	//fmt.Println(string(a))
	reader := bytes.NewBuffer([]byte(xml1))
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err := decoder.Decode(&a)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v\n", a)
}

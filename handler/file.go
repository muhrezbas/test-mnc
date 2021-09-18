package handler

import (
	"mnc-api/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"mnc-api/model"
	"strconv"
	// "fmt"
)

// Context godoc
type Context struct {
	// UC usecase.Context
}
var languages []model.LanguageHardCoded

func (h *Context) Palindrome(c *gin.Context) {
	cfg := config.IsPalindrome(c.Query("text"), false)
	if cfg == true {
		SendSuccess(c, "Palindrome")
		return
	} else {
		SendBadRequets(c, "Not Palindrome")
		return
	}
	return
}
func (h *Context) Hello(c *gin.Context) {
	c.JSONP(200, "Hello Go Developers")
	return
}
func (h *Context) Language(c *gin.Context) {
	var language model.LanguageHardCoded
	var relation model.Relation


	language.Language = "C"
	language.Appeared = 1972
	language.Created  =[]string{"Dennis Ritchie"}
	language.Functional = true
	language.ObjectOriented = false
	relation.InfluencedBy =[]string{"B","ALGOL 68","Assembly","FORTRAN"}
	relation.Influences = []string{"C++","Objective-C","C#","Java","JavaScript","PHP","Go"}
	language.Relation = relation
	c.JSONP(200, language)

	return
}
func (h *Context) PostLanguage(c *gin.Context) {
	var f model.LanguageHardCoded
        // Read ones
        if err := c.ShouldBindBodyWith(&f, binding.JSON); err != nil {
            // fmt.Printf("%+v", err)
        }
        // fmt.Printf("%+v", f)
	languages = append(languages, f)

	c.JSONP(200, f)

	return
}
func (h *Context) GetAllLanguages(c *gin.Context) {

	c.JSONP(200, languages)

	return
}
func (h *Context) GetOneLanguage(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	// fmt.Println(intID)
	// fmt.Println(languages[intID])
	c.JSONP(200, languages[intID])

	return
}
func (h *Context) UpdateOneLanguage(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	// fmt.Println(intID)
	
	// fmt.Println(languages[intID])
	languages[intID] = languages[len(languages)-1]
	languages[len(languages)-1] = model.LanguageHardCoded{}
	languages= languages[:len(languages)-1]
	// fmt.Println(languages)
	var f model.LanguageHardCoded
        // Read ones
        if err := c.ShouldBindBodyWith(&f, binding.JSON); err != nil {
            // fmt.Printf("%+v", err)
        }
        // fmt.Printf("%+v", f, "anjay")

	languages = append(languages, model.LanguageHardCoded{})
	copy(languages[intID+1:], languages[intID:])
	languages[intID] = f
	// fmt.Println(languages[intID],"anjay")
	c.JSONP(200, languages)

	return
}
func (h *Context) DeleteOneLanguage(c *gin.Context) {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	if err != nil {
		SendError(c, 200, err.Error())
		return
	}
	// fmt.Println(intID)
	
	// fmt.Println(languages[intID])
	languages[intID] = languages[len(languages)-1]
	languages[len(languages)-1] = model.LanguageHardCoded{}
	languages= languages[:len(languages)-1]
	c.JSONP(200, languages)

	return
}
// Upload godoc

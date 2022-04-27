package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/AmanSrivastava2000/bookings/cmd/pkg/config"
	"github.com/AmanSrivastava2000/bookings/cmd/pkg/models"
)

var functions = template.FuncMap{}

//get the Template cache in "app" var
var app *config.AppConfig
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData{
	//add default data here.

	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//getting the template cache
	var tc map[string]*template.Template
	//if useCache is true then we are not in development mode so we can use pre built template
	if app.UseCache{
		tc = app.TemplateCache
	} else{ // useCache == false ----> means we are in development mode
		tc, _ = CreateTemplateCache()
	}

	//getting the required template of the tmpl string or page passed from template cache
	t, ok := tc[tmpl] // ok will have wether we were able to find that string or not.
	if !ok {
		log.Fatal("Could not get template from template cache ")
	}

	buf := new(bytes.Buffer)

	//adding default data to td using AddDefaultData function.
	td = AddDefaultData(td)

	//executing the template set of the current page
	_ = t.Execute(buf, td) //td is the data passed from handlers to display
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser")
	}
}

//Creates a Template Cache as a map.
func CreateTemplateCache() (map[string]*template.Template, error) {
	//will store the name(key):pointer to templateSet(value) for each file(not the layout one.)
	myCache := map[string]*template.Template{}

	//getting all the pages in an array(slice) ending with .page.tmpl(not the layout one)
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//iterating in this slice or array and working on each page
	for _, page := range pages {
		//getting the base name from the whole path name of the current page
		name := filepath.Base(page)

		//making template set of this page
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//finding the layout page
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		//if layout page found then matches length > 0 so again parsing this layout page to get final template set.
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		//putting the page and corresponding pointer to its template set into the myCache map.
		myCache[name] = ts
	}
	return myCache, err

}

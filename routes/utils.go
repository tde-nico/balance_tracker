package routes

import (
	"balance/db"
	"balance/middleware"
	"fmt"
	"html/template"
)

type Data struct {
	User    *db.User
	Flashes []middleware.Flash
}

type DataInterface interface {
	UpdateFlashes([]middleware.Flash)
}

func (d *Data) UpdateFlashes(flashes []middleware.Flash) {
	d.Flashes = flashes
}

func getTemplates(ctx *middleware.Ctx, pages ...string) *template.Template {
	files := make([]string, 0, len(pages)+1)
	files = append(files, "templates/base.html")
	for _, page := range pages {
		files = append(files, fmt.Sprintf("templates/%s.html", page))
	}
	tmpl, err := template.New("").ParseFiles(files...)
	if err != nil {
		ctx.InternalError(fmt.Errorf("error parsing template %v", err))
		return nil
	}
	return tmpl
}

func getTemplate(ctx *middleware.Ctx, page string) *template.Template {
	tmpl, err := template.New("").ParseFiles("templates/base.html", fmt.Sprintf("templates/%s.html", page))
	if err != nil {
		ctx.InternalError(fmt.Errorf("error parsing template %v", err))
		return nil
	}
	return tmpl
}

func getSingleTemplate(ctx *middleware.Ctx, page string) *template.Template {
	tmpl, err := template.New("").ParseFiles(fmt.Sprintf("templates/%s.html", page))
	if err != nil {
		ctx.InternalError(fmt.Errorf("error parsing template %v", err))
		return nil
	}
	return tmpl
}

func executeTemplate(ctx *middleware.Ctx, tmpl *template.Template, data DataInterface) {
	flashes, err := ctx.GetFlashes()
	if err != nil {
		ctx.InternalError(fmt.Errorf("error getting flashes: %v", err))
		return
	}
	data.UpdateFlashes(flashes)

	err = tmpl.ExecuteTemplate(ctx.Writer, "base", data)
	if err != nil {
		ctx.InternalError(fmt.Errorf("error executing template %v", err))
		return
	}
}

func executeSingleTemplate(ctx *middleware.Ctx, tmpl *template.Template, name string, data any) {
	err := tmpl.ExecuteTemplate(ctx.Writer, name, data)
	if err != nil {
		ctx.InternalError(fmt.Errorf("error executing template %v", err))
		return
	}
}

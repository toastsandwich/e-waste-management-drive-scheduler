// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.771
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/toastsandwich/ewaste-collection-caller/views/base"

func scheduleDrive() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"container mt-4\"><h2>Schedule a New E-Waste Drive</h2><form action=\"/submit-drive\" method=\"post\"><div class=\"mb-3\"><label for=\"drive-name\" class=\"form-label\">Drive Name</label> <input type=\"text\" class=\"form-control\" id=\"drive-name\" name=\"drive-name\" required></div><div class=\"mb-3\"><label for=\"address\" class=\"form-label\">Address</label> <input type=\"text\" class=\"form-control\" id=\"address\" name=\"address\" required></div><div class=\"mb-3\"><label for=\"name\" class=\"form-label\">Your Name</label> <input type=\"text\" class=\"form-control\" id=\"name\" name=\"name\" required></div><div class=\"mb-3\"><label for=\"email\" class=\"form-label\">Email</label> <input type=\"email\" class=\"form-control\" id=\"email\" name=\"email\" required></div><div class=\"mb-3\"><label for=\"phone\" class=\"form-label\">Phone Number</label> <input type=\"tel\" class=\"form-control\" id=\"phone\" name=\"phone\" required pattern=\"[0-9]{10}\"></div><div class=\"mb-3\"><label for=\"scheduled-date\" class=\"form-label\">Scheduled Date & Time</label> <input type=\"datetime-local\" class=\"form-control\" id=\"scheduled-date\" name=\"scheduled_date\" required></div><button type=\"submit\" class=\"btn btn-success\">Submit</button></form></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func ScheduleDrive() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = base.Base("Schedule Drive", scheduleDrive()).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate

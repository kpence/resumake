package templates

import (
	"regexp"
	"strings"
	"text/template"
)

// Plaintext returns the go template of the plaintext resume template
func Plaintext(censor *bool) *template.Template {
	fns := template.FuncMap{"censor": func(text string) string {
		if *censor {
			re := regexp.MustCompile(`\|\|.*?\|\|`)
			text = re.ReplaceAllStringFunc(text, func(s string) string {
				return strings.Repeat("#", len(s)-4)
			})
		} else {
			text = strings.ReplaceAll(text, "||", "")
		}

		return text
	}}
	tmpl, err := template.New("plaintext").Funcs(fns).Parse(plaintext)
	if err != nil {
		panic(err)
	}

	return tmpl
}

var plaintext = `
{{- .Header.Name | censor }}
{{ .Header.Email | censor }}
==============================

EDUCATION
==============================
{{- range $eduEntry := .EducationEntries }}
{{ $eduEntry.Degree | censor }}, {{ $eduEntry.School | censor }}, {{ $eduEntry.TimeSpan.Display | censor }}
{{ if (and $eduEntry.GPA (not $eduEntry.MajorGPA)) }}GPA: {{ $eduEntry.GPA | censor }}{{ "\n" }}{{ end }}
{{- if (and $eduEntry.GPA $eduEntry.MajorGPA) }}Cumulative GPA: {{ $eduEntry.GPA | censor }}{{ "\n" }}{{ end }}
{{- if $eduEntry.MajorGPA }}Major GPA: {{ $eduEntry.MajorGPA | censor }}{{ "\n" }}{{ end }}
{{- end }}
PROFESSIONAL EXPERIENCE
===============================
{{- range $jobEntry := .JobEntries }}
{{ $jobEntry.Employer | censor }}, {{ $jobEntry.Location | censor }}, {{ $jobEntry.TimeSpan.Display | censor }}
{{ $jobEntry.Title | censor }}
{{- range $bullet := $jobEntry.Bullets }}
* {{ $bullet | censor }} 
{{- end }}
{{ if $jobEntry.Skills }}Technologies used: {{ $jobEntry.Skills.Display | censor }}{{ "\n" }}{{ end }}
{{- if $jobEntry.Description }}{{ "\n" }}{{ $jobEntry.Description | censor }}{{- end }}
{{- end }}
SKILLS
==============================
Languages: {{ .Languages.Display | censor }}
Technologies: {{ .Technologies.Display | censor }}

PROJECTS
==============================
{{- range $project := .Projects }}
{{ $project.Name | censor }}
{{- if $project.Url }}{{ "\n" }}{{ $project.Url | censor }}{{- end }}
{{- if $project.Description }}{{ "\n" }}{{ $project.Description | censor }}{{- end }}
{{- range $bullet := $project.Bullets }}
* {{ $bullet | censor }} 
{{- end }}
Technologies used: {{ $project.Skills.Display | censor }}{{ "\n" }}
{{- end }}`

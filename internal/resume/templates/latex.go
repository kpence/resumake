package templates

import (
	"regexp"
	"strings"
	"text/template"
)

// Latex returns the go template for the Latex resume template
func Latex(censor *bool) *template.Template {
	fns := template.FuncMap{"escape": latexEscape, "toUpper": strings.ToUpper, "censor": func(text string) string {
		if *censor {
			re := regexp.MustCompile(`\|\|(.*?)\|\|`)
			text = re.ReplaceAllString(text, "\\censor{$1}")
		} else {
			text = strings.ReplaceAll(text, "||", "")
		}

		return text
	}}

	tmpl, err := template.New("latex").Funcs(fns).Delims("[[", "]]").Parse(latexDocument)
	if err != nil {
		panic(err)
	}

	return tmpl
}

func latexEscape(text string) string {
	text = strings.ReplaceAll(text, "&", "\\&")
	text = strings.ReplaceAll(text, "%", "\\%")
	text = strings.ReplaceAll(text, "$", "\\$")
	text = strings.ReplaceAll(text, "#", "\\#")
	text = strings.ReplaceAll(text, "_", "\\_")
	return text
}

func latexCensor(text string) string {
	re := regexp.MustCompile(`\|\|(.*?)\|\|`)
	text = re.ReplaceAllString(text, "\\censor{$1}")
	return text
}

var latexDocument = `
\documentclass[letterpaper]{article}
    \usepackage{fullpage}
    \usepackage{amsmath}
    \usepackage{amssymb}
    \usepackage{textcomp}
    \usepackage{enumitem}
    \usepackage[utf8]{inputenc}
    \usepackage[T1]{fontenc}
    \usepackage[margin=0.75in]{geometry}
    \textheight=10in
    \pagestyle{empty}
    \raggedright
    \usepackage{censor}
    \usepackage{fontawesome}
    \usepackage{helvet}


%%%%%%%%%%%%%%%%%%%%%%% DEFINITIONS FOR RESUME %%%%%%%%%%%%%%%%%%%%%%%

\newcommand{\lineunder} {
    \vspace*{-8pt} \\
    \hspace*{-18pt} \hrulefill \\
}

\newcommand{\header} [1] {
    {\hspace*{-18pt}\vspace*{6pt} {#1}}
    \vspace*{-6pt} \lineunder
}

%%%%%%%%%%%%%%%%%%%%%%% END RESUME DEFINITIONS %%%%%%%%%%%%%%%%%%%%%%%

\begin{document}
\vspace*{-40pt}

\sffamily

%==== Profile ====%
\vspace*{-10pt}
\begin{center}
    {\Huge [[ .Header.Name | toUpper | censor ]]}\\
    \vspace{2.5pt}
    \faEnvelope \ [[ .Header.Email | censor ]]
    [[ if .Header.Phone ]] $|$ \faPhone \ [[ .Header.Phone | censor ]][[ end -]]
    [[ if .Header.Linkedin ]] $|$ \faLinkedinSquare \ [[ .Header.Linkedin | censor ]][[ end -]]
    [[ if .Header.Github ]] $|$ \faGithub \ [[ .Header.Github | censor ]][[ end -]]\\
\end{center}




%==== Education ====%
\vspace{-5pt}
\header{Education}
[[ range $eduEntry := .EducationEntries ]]
\textbf{[[ $eduEntry.School | escape | censor]]}
\hfill [[ if $eduEntry.Location ]][[ $eduEntry.Location | censor]][[ end -]]\\
[[ $eduEntry.Degree | censor]][[- if (and $eduEntry.GPA (not $eduEntry.MajorGPA)) ]] \textit{GPA: [[ $eduEntry.GPA | censor]]}[[ end ]]
\hfill [[ $eduEntry.TimeSpan.Display | censor]]\\
[[- if (and $eduEntry.MajorGPA (not $eduEntry.GPA)) ]] \textit{Major GPA: [[ $eduEntry.MajorGPA | censor]]}\\[[ end ]]
[[- if (and $eduEntry.MajorGPA $eduEntry.GPA) ]] \textit{Major GPA: [[ $eduEntry.MajorGPA | censor]]} | \textit{Cumulative GPA: [[ $eduEntry.GPA | censor]]}\\[[ end ]]
\vspace{2mm}
[[ end ]]



%==== Skills ====%
\header{Skills}
\vspace{1mm}
\begin{tabular}{ l l }
    Languages:    & [[ .Languages.Display | escape | censor]] \\
    Technologies: & [[ .Technologies.Display | escape | censor]] \\
\end{tabular}
\vspace{2mm}




%==== Experience ====%
\header{Experience}
\vspace{1mm}
[[ range $jobEntry := .JobEntries ]]
\textbf{[[ $jobEntry.Employer | escape | censor]] \textbar{} [[ $jobEntry.Title | censor]]}
\hfill [[ $jobEntry.Location | censor]]\\
\vspace{0.75mm}
[[ if $jobEntry.Skills ]]\textit{[[ $jobEntry.Skills.Display | escape | censor]]}[[ "\n" ]][[ end -]]
\hfill [[ $jobEntry.TimeSpan.Display | censor]]\\
[[ if $jobEntry.Description ]]\textit{[[ $jobEntry.Description | escape | censor]]} \\[[ "\n" ]][[ end -]]
[[ if $jobEntry.Skills ]]\vspace{-2.5mm}[[ else ]]\vspace{-7mm}[[ end ]]
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
[[- range $bullet := $jobEntry.Bullets ]]
    \item [[ $bullet | escape | censor]] 
[[- end ]]
\end{itemize}
[[ end ]]



%==== Projects ====%
\header{Projects}
\vspace{1mm}
[[- range $project := .Projects ]]
{\textbf{[[ $project.Name | escape | censor]]}}\\
\vspace{0.75mm}
[[ if $project.Skills ]]\textit{[[ $project.Skills.Display | escape | censor]]} \\[[ "\n" ]][[ end -]]
[[ if $project.Url ]]\textit{[[ $project.Url | escape | censor]]} \\[[ "\n" ]][[ end -]]
[[ if $project.Description ]]\vspace{0.75mm}[[ "\n" ]][[ $project.Description | censor]] \\[[ "\n" ]][[ end -]]
[[- if $project.Bullets ]]
\vspace{-2.25mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
[[- range $bullet := $project.Bullets ]]
    \item [[ $bullet | escape | censor]] 
[[- end ]]
\vspace{-2.25mm}
\end{itemize}
[[ end -]]
\vspace*{2mm}
[[ end ]]
\end{document}
`

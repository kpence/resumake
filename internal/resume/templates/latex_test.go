package templates_test

import (
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	"resumake/internal/resume/templates"
)

func TestWholeLatexTemplate(t *testing.T) {
	g := NewGomegaWithT(t)

	g.Expect(func() { templates.Latex() }).ToNot(Panic())
	tmpl := templates.Latex()

	b := &strings.Builder{}
	err := tmpl.Execute(b, testResume)
	g.Expect(err).ToNot(HaveOccurred())

	g.Expect(b.String()).To(Equal(latexResume))
}

var latexResume = `
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
\StopCensoring

%==== Profile ====%
\vspace*{-10pt}
\begin{center}
    {\Huge JOHN SMITH}\\
    \vspace{2.5pt}
    \faEnvelope \ john.smith@gmail.com
    \\
\end{center}




%==== Education ====%
\vspace{-5pt}
\header{Education}

\textbf{Georgia Institute of Technology}
\hfill\\
M.S. in Computer Science \textit{GPA: 3.9}
\hfill Jan. 2004 - Current\\
\vspace{2mm}

\textbf{University of Philadelphia}
\hfill\\
B.S. in Computer Science
\hfill Jan. 2004 - Dec. 2006\\
\vspace{2mm}





%==== Experience ====%
\header{Experience}
\vspace{1mm}

\textbf{Microsoft \textbar{} Senior Software Engineer}
\hfill Seattle, WA\\
\vspace{0.75mm}
\textit{C\#, C++}
\hfill Jan. 2004 - Current\\
\vspace{-2.5mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}

\textbf{IBM \textbar{} Software Engineer}
\hfill Seattle, WA\\
\vspace{0.75mm}
\textit{Java}
\hfill Jan. 2004 - Dec. 2006\\
\vspace{-2.5mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}

\textbf{SAP \textbar{} Software Engineer Intern}
\hfill Seattle, WA\\
\vspace{0.75mm}
\textit{ABAP}
\hfill Winter 2004\\
\vspace{-2.5mm}
\begin{itemize}[leftmargin=10pt] \itemsep -1pt
    \item did a thing
    \item did another thing
\end{itemize}





%==== Projects ====%
\header{Projects}
\vspace{1mm}
{\textbf{Compiler}}\\
\vspace{0.75mm}
\textit{C\#, ANTLR, LLVM} \\
\vspace{0.75mm}
Compiles stuff \\
\vspace*{2mm}

{\textbf{Linker}}\\
\vspace{0.75mm}
\textit{Java, Bison, GCC} \\
\textit{https://example.com} \\
\vspace{0.75mm}
Links stuff \\
\vspace*{2mm}

{\textbf{Gameboy Emulator}}\\
\vspace{0.75mm}
\textit{C++} \\
\vspace{0.75mm}
Emulates stuff \\
\vspace*{2mm}





%==== Skills ====%
\header{Skills}
\vspace{1mm}
\begin{tabular}{ l l }
    Languages:    & C++, Java, C\# \\
    Technologies: & git, Docker \\
\end{tabular}
\vspace{2mm}
\end{document}
`

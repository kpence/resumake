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
    \usepackage[margin=1in]{geometry}
    \textheight=10in
    \pagestyle{empty}
    \raggedright

%%%%%%%%%%%%%%%%%%%%%%% DEFINITIONS FOR RESUME %%%%%%%%%%%%%%%%%%%%%%%

\newcommand{\lineunder} {
    \vspace*{-8pt} \\
    \hspace*{-18pt} \hrulefill \\
}

\newcommand{\header} [1] {
    {\hspace*{-18pt}\vspace*{6pt} \textsc{#1}}
    \vspace*{-6pt} \lineunder
}

%%%%%%%%%%%%%%%%%%%%%%% END RESUME DEFINITIONS %%%%%%%%%%%%%%%%%%%%%%%

\begin{document}
\vspace*{-40pt}



%==== Profile ====%
\vspace*{-10pt}
\begin{center}
    {\Huge \scshape {John Smith}}\\
    john.smith@gmail.com\\
\end{center}




%==== Education ====%
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




%==== Skills ====%
\header{Skills}
\begin{tabular}{ l l }
    Languages:    & C++, Java, C\# \\
    Technologies: & git, Docker \\
\end{tabular}
\vspace{2mm}




%==== Projects ====%
\header{Projects}
{\textbf{Compiler}} {\sl C\#, ANTLR, LLVM} \\
Compiles stuff \\
\vspace*{2mm}

{\textbf{Gameboy Emulator}} {\sl C++} \\
Emulates stuff \\
\vspace*{2mm}


\end{document}
`
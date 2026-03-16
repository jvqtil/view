package view

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Show(text string, paintFunc ...func(a ...any) string) {
	if len(paintFunc) == 0 {
		paintFunc = append(paintFunc, func(a ...any) string {
			return fmt.Sprint(a...)
		})
	}

	lines := strings.Split(text, "\n")
	for i, line := range lines {
		lines[i] = paintFunc[0](line)
	}
	toProcess := strings.Join(lines, "\n")

	pager := os.Getenv("PAGER")
	if pager == "" {
		if _, err := exec.LookPath("bat"); err == nil {
			pager = "bat"
		} else if _, err := exec.LookPath("glow"); err == nil {
			pager = "glow"
		} else {
			pager = "less"
		}
	}

	cmd := exec.Command(pager)
	cmd.Stdin = strings.NewReader(toProcess)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running pager:", err)
	}
}

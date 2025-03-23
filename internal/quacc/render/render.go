package render

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

func RenderNoteContent(content string) error {
	rc, err := glamour.Render(content, "dark")

	if err != nil {
		return err
	}

	fmt.Println(rc)

	return nil
}

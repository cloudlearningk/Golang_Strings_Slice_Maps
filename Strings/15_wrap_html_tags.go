package strings_problems

import "fmt"

// AddHTMLTags uses fmt.Sprintf for template-style string construction — clear and safe.
func AddHTMLTags(tag, content string) string {
	return fmt.Sprintf("<%s>%s</%s>", tag, content, tag)
}

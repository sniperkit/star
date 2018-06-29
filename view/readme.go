package view

import (
	"fmt"
	"strings"

	"github.com/voidint/star/store"
)

// GenReadme 生成markdown格式的README内容。
func GenReadme(tags []store.RepoedTag) (content []byte, err error) {
	var sb strings.Builder
	sb.WriteString("# Awesome\n\nThis file is autogenerated by [star](https://github.com/voidint/star), do not edit.\n\n")
	for i := range tags {
		sb.WriteString(fmt.Sprintf("## %s\n\n%s", tags[i].Tag.Name, tags[i].Tag.Desc))

		for j := range tags[i].Repos {
			sb.WriteString(fmt.Sprintf("- [%s](%s) %s\n",
				tags[i].Repos[j].FullName,
				tags[i].Repos[j].HTMLURL,
				tags[i].Repos[j].Description,
			))
		}
	}
	return []byte(sb.String()), nil
}
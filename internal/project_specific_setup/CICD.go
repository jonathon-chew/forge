package projectspecificsetup

import (
	"fmt"
)

func cicd() {

	var content string = `
# Commit message must use Conventional Commits style
if ! printf '%s\n' "$subject" |
    grep -qEi '^(feat|fix|build|chore|ci|delete|docs|perf|new|refactor|revert|style|test|update)(\([^)]+\))?!?:[[:space:]]+'; then
    error "commit message must follow Conventional Commits style"
fi
	`

	fmt.Printf(content)
}

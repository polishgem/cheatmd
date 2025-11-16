package sheet

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

// Parse parses cheatsheet frontmatter
func parse(markdown string) (frontmatter, string, error) {

	// determine the appropriate line-break for the platform
	linebreak := "\n"
	if runtime.GOOS == "windows" {
		linebreak = "\r\n"
	}

	// specify the frontmatter delimiter
	delim := fmt.Sprintf("---%s", linebreak)

	// initialize a frontmatter struct
	var fm frontmatter

	// if the markdown does not contain frontmatter, pass it through unmodified
	if !strings.HasPrefix(markdown, delim) {
		return fm, markdown, nil
	}

	// otherwise, split the frontmatter and cheatsheet text
	parts := strings.SplitN(markdown, delim, 3)

	// return an error if the frontmatter parses into the wrong number of parts
	if len(parts) != 3 {
		return fm, markdown, fmt.Errorf("failed to delimit frontmatter")
	}

	// return an error if the YAML cannot be unmarshalled
	if err := yaml.Unmarshal([]byte(parts[1]), &fm); err != nil {
		return fm, markdown, fmt.Errorf("failed to unmarshal frontmatter: %v", err)
	}

	return fm, parts[2], nil
}

// extensionToLanguage maps file extensions to syntax highlighting languages
var extensionToLanguage = map[string]string{
	".md":   "markdown",
	".sh":   "bash",
	".bash": "bash",
	".py":   "python",
	".js":   "javascript",
	".ts":   "typescript",
	".go":   "go",
	".rb":   "ruby",
	".yml":  "yaml",
	".yaml": "yaml",
	".json": "json",
	".php":  "php",
	".java": "java",
	".c":    "c",
	".cpp":  "cpp",
	".rs":   "rust",
	".sql":  "sql",
}

// guessLanguageFromExtension attempts to determine the syntax highlighting
// language from a file's extension
func guessLanguageFromExtension(path string) string {
	ext := filepath.Ext(path)
	if lang, ok := extensionToLanguage[ext]; ok {
		return lang
	}
	return ""
}

package dos

import (
	"fmt"
	"strings"
)

var catalog = map[string]string{
	"HELLO": "PRINT \"HELLO WORLD\"",
}

func Execute(cmd string) string {
	upper := strings.ToUpper(strings.TrimSpace(cmd))
	args := strings.Fields(upper)

	if len(args) == 0 {
		return ""
	}

	switch args[0] {
	case "CATALOG":
		return handleCatalog()
	case "LOAD":
		if len(args) < 2 {
			return "MISSING FILE NAME"
		}
		return loadFile(args[1])
	case "SAVE":
		if len(args) < 2 {
			return "MISSING FILE NAME"
		}
		return saveFile(args[1])
	case "DELETE":
		if len(args) < 2 {
			return "MISSING FILE NAME"
		}
		return deleteFile(args[1])
	default:
		return "SYNTAX ERROR"
	}
}

func handleCatalog() string {
	var b strings.Builder
	b.WriteString("DISK VOLUME 254\nNAME       TYPE\n")
	for name := range catalog {
		b.WriteString(fmt.Sprintf("%-10s  TXT\n", name))
	}
	return b.String()
}

func loadFile(name string) string {
	name = strings.ToUpper(name)
	if val, ok := catalog[name]; ok {
		return fmt.Sprintf("LOADED \"%s\"\n%s", name, val)
	}
	return "FILE NOT FOUND"
}

func saveFile(name string) string {
	name = strings.ToUpper(name)
	if _, exists := catalog[name]; exists {
		return "FILE ALREADY EXISTS"
	}
	catalog[name] = "NEW FILE"
	return fmt.Sprintf("SAVED \"%s\"", name)
}

func deleteFile(name string) string {
	name = strings.ToUpper(name)
	if _, exists := catalog[name]; !exists {
		return "FILE NOT FOUND"
	}
	delete(catalog, name)
	return fmt.Sprintf("DELETED \"%s\"", name)
}

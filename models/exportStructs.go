package models

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
)

func goTypeToTypeScript(goType string) string {
	switch goType {
	case "int", "int32", "int64", "float32", "float64":
		return "number"
	case "string":
		return "string"
	case "bool":
		return "boolean"
	default:
		return "any"
	}
}

func GenerateTypescriptFiles(filePath string, fileOutputPath string) {
	// Abrir el archivo structs.go (cambia el archivo si es necesario)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Crear un nuevo token.FileSet
	fset := token.NewFileSet()

	// Parsear el archivo
	node, err := parser.ParseFile(fset, filePath, file, parser.ParseComments)
	if err != nil {
		fmt.Println("Error al parsear el archivo:", err)
		return
	}

	var tsOutput string
	// Iterar por las estructuras encontradas
	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.TYPE {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					structName := typeSpec.Name.Name
					if structType, ok := typeSpec.Type.(*ast.StructType); ok {

						// Crear el equivalente en TypeScript
						tsFields := []string{}

						// Iterar sobre los campos del struct y convertirlos a TypeScript
						for _, field := range structType.Fields.List {
							for _, fieldName := range field.Names {
								// Convertir el tipo del campo a string
								goType := fmt.Sprintf("%s", field.Type)
								tsType := goTypeToTypeScript(goType)

								// Crear el campo en formato TypeScript
								tsFields = append(tsFields, fmt.Sprintf("%s: %s;", toSnakeCase(fieldName.Name), tsType))
							}
						}

						// Unir los campos y generar el tipo TypeScript final
						ts := fmt.Sprintf("export type %s = {\n\t%s\n};", structName, strings.Join(tsFields, "\n\t"))
						tsOutput = tsOutput + "\n" + ts
					}
				}
			}
		}
	}
	err = os.WriteFile(fileOutputPath, []byte(tsOutput), 0644)
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Archivo " + fileOutputPath + " creado exitosamente.")
}

func toSnakeCase(str string) string {
	// Expresión regular para encontrar lugares donde hay mayúsculas precedidas por letras minúsculas o números
	re := regexp.MustCompile("([a-z0-9])([A-Z])")

	// Insertar un guion bajo entre las letras que coinciden con el patrón
	snake := re.ReplaceAllString(str, "${1}_${2}")

	// Convertir todo el resultado a minúsculas
	return strings.ToLower(snake)
}

// Función para convertir tipos SQL a tipos Go
func sqlTypeToGoType(sqlType string) string {
	sqlType = strings.ToUpper(sqlType)
	switch {
	case strings.Contains(sqlType, "INT"):
		return "int"
	case strings.Contains(sqlType, "VARCHAR"), strings.Contains(sqlType, "TEXT"), strings.Contains(sqlType, "CHAR"):
		return "string"
	case strings.Contains(sqlType, "FLOAT"), strings.Contains(sqlType, "DOUBLE"), strings.Contains(sqlType, "DECIMAL"):
		return "float64"
	case strings.Contains(sqlType, "BOOL"):
		return "bool"
	case strings.Contains(sqlType, "DATE"), strings.Contains(sqlType, "TIME"):
		return "time.Time"
	default:
		return "interface{}" // Tipo genérico en caso de que no se reconozca el tipo SQL
	}
}

// Función para procesar un archivo SQL y generar structs en Go
func processSQLFile(filePath string) {
	// Abrir el archivo SQL
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Expresión regular para capturar la definición de tablas y columnas
	tableRegex := regexp.MustCompile(`(?i)CREATE\s+TABLE\s+(\w+)\s*\(([^;]+)\)`)
	columnRegex := regexp.MustCompile(`(\w+)\s+(\w+(\(\d+\))?)`)

	// Leer el archivo línea por línea
	scanner := bufio.NewScanner(file)
	var sqlContent string
	for scanner.Scan() {
		sqlContent += scanner.Text() + "\n"
	}

	// Buscar definiciones de tablas
	tables := tableRegex.FindAllStringSubmatch(sqlContent, -1)

	// Para cada tabla encontrada, generar un struct Go
	for _, table := range tables {
		tableName := table[1]
		columns := table[2]

		fmt.Printf("type %s struct {\n", tableName)

		// Buscar definiciones de columnas
		columnMatches := columnRegex.FindAllStringSubmatch(columns, -1)
		for _, column := range columnMatches {
			columnName := column[1]
			sqlType := column[2]
			goType := sqlTypeToGoType(sqlType)

			// Imprimir el campo del struct en Go
			fmt.Printf("\t%s %s `json:\"%s\"`\n", strings.Title(columnName), goType, columnName)
		}

		fmt.Printf("}\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}

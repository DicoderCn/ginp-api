package desc

import (
	"fmt"
	"ginp-api/internal/app/gapi/start"
	"ginp-api/internal/gen"
	"io/ioutil"
	"reflect"
)

func GenFields() {
	for _, entity_ := range start.EntityAutoMigrateList {
		t := reflect.TypeOf(entity_).Elem()
		// fileName := strings.ToLower(t.Name()) + ".go"
		packageName := "m" + gen.NameToAllSmall(t.Name())
		content := "package " + packageName + " \n\n"
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			constName := t.Name() + field.Name
			fieldName := gen.NameToLine(field.Name)
			if fieldName == "model" { //gorm.model
				content += fmt.Sprintf("const %s = \"%s\"\n\n", t.Name()+"ID", "id")
				content += fmt.Sprintf("const %s = \"%s\"\n\n", t.Name()+"Created", "created_at")
				content += fmt.Sprintf("const %s = \"%s\"\n\n", t.Name()+"Updated", "updated_at")
			} else {
				content += fmt.Sprintf("const %s = \"%s\"\n\n", constName, fieldName)
			}
		}

		filePath := PathFields(t.Name())
		err := ioutil.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}
		fmt.Println("Data written to file.")
	}
}

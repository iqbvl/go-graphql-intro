package schema

import (
	"bytes"

	"iqbvl/go-graphql-intro/services"
)

func GetSchema() string {
	registerDepSchema := []string{
		"schema/schema.graphql",
		"schema/type/admin-user.graphql",
	}
	buf := bytes.Buffer{}
	for _, name := range AssetNames() {
		if service.ArrayContains(registerDepSchema, name) {
			b := MustAsset(name)
			buf.Write(b)

			// Add a newline if the file does not end in a newline.
			if len(b) > 0 && b[len(b)-1] != '\n' {
				buf.WriteByte('\n')
			}
		}
	}

	return buf.String()
}

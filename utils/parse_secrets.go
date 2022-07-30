package utils

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/jessevdk/go-flags"
	"gocloud.dev/runtimevar"
)

func Parse(config interface{}, baseProjectURL string) error {

	_, err := flags.Parse(config)
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Now process "parse_to" fields.
	configValue := reflect.ValueOf(config).Elem()
	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {

		fieldValue := configValue.Field(i)
		sourceField := configType.Field(i).Tag.Get("secret-key")
		if sourceField == "" {
			continue
		}
		if _, has := configType.FieldByName(sourceField); !has {
			log.Fatalf("No such field: %s.", sourceField)
		}
		sourceFieldValue := configValue.FieldByName(sourceField)
		sourceFieldStringValue := sourceFieldValue.Interface().(string)

		pathToUrl := fmt.Sprintf("%s/%s?decoder=%s", baseProjectURL, sourceFieldStringValue, "string")
		if strings.Contains("://", sourceFieldStringValue) {
			pathToUrl = fmt.Sprintf("%s?decoder=%s", sourceFieldStringValue, "string")
		}

		v, err := runtimevar.OpenVariable(ctx, pathToUrl)
		if err != nil {
			return fmt.Errorf("failed to open variable %s: %w", pathToUrl, err)
		}
		defer v.Close()
		latestVal, err := v.Latest(ctx)
		if err != nil {
			return fmt.Errorf("failed to get latest value for %s: %w", pathToUrl, err)
		}
		fieldValue.SetString(latestVal.Value.(string))

	}
	return nil
}

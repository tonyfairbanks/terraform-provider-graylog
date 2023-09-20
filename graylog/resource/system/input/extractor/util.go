package extractor

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/tonyfairbanks/terraform-provider-graylog/graylog/convert"
	"github.com/tonyfairbanks/terraform-provider-graylog/graylog/util"
)

const (
	keyInputID         = "input_id"
	keyExtractorID     = "extractor_id"
	keyExtractorConfig = "extractor_config"
	keyID              = "id"
	keyConfig          = "config"
	keyConverters      = "converters"
	keyType            = "type"
)

func getDataFromResourceData(d *schema.ResourceData) (map[string]interface{}, error) {
	data, err := convert.GetFromResourceData(d, Resource())
	if err != nil {
		return nil, err
	}
	util.SetDefaultValue(data, "target_field", "")
	util.SetDefaultValue(data, "condition_value", "")

	if err := convert.JSONToData(data, keyExtractorConfig); err != nil {
		return nil, err
	}
	util.RenameKey(data, keyExtractorID, keyID)

	converters := data[keyConverters].([]interface{})
	for i, a := range converters {
		elem := a.(map[string]interface{})
		err = convert.JSONToData(elem, keyConfig)
		if err != nil {
			return nil, err
		}
		converters[i] = elem
	}

	data[keyConverters] = converters

	return data, nil
}

func setDataToResourceData(d *schema.ResourceData, data map[string]interface{}) error {
	if err := convert.DataToJSON(data, keyExtractorConfig); err != nil {
		return err
	}
	util.RenameKey(data, keyID, keyExtractorID)

	converters := data[keyConverters].([]interface{})
	for i, a := range converters {
		elem := a.(map[string]interface{})
		b, err := json.Marshal(elem[keyConfig])
		if err != nil {
			return err
		}
		elem[keyConfig] = string(b)
		converters[i] = elem
	}

	data[keyConverters] = converters

	if err := convert.SetResourceData(d, Resource(), data); err != nil {
		return err
	}

	d.SetId(d.Get(keyInputID).(string) + "/" + d.Get(keyExtractorID).(string))
	return nil
}

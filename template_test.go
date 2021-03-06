package elit

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMappingsMarshalJSON(t *testing.T) {
	table := []struct {
		mappings map[string]Type
		out      string
	}{
		{
			mappings: map[string]Type{
				"_defualt_": Type{
					All: &All{
						Enabled: true,
					},
				},
				"some_type": Type{
					Properties: map[string]Property{
						"@timestamp": Property{
							Type:   PropertyTypeDate,
							Format: "yyyy-MM-dd'T'HH:mm:ssZ",
						},
						"count": Property{
							Type: PropertyTypeInteger,
						},
						"location": Property{
							Type: PropertyTypeGeoPoint,
						},
						"word": Property{
							Type:      PropertyTypeText,
							FieldData: true,
							Fields: map[string]Property{
								"keyword": Property{
									Type:        PropertyTypeKeyword,
									IgnoreAbove: 256,
								},
							},
						},
						"object": Property{
							Type: PropertyTypeNested,
							Properies: map[string]Property{
								"title": Property{
									Type:        PropertyTypeKeyword,
									IgnoreAbove: 256,
								},
								"user": Property{
									Type: PropertyTypeNested,
									Properies: map[string]Property{
										"first_name": Property{
											Type: PropertyTypeKeyword,
										},
										"last_name": Property{
											Type: PropertyTypeKeyword,
										},
										"age": Property{
											Type: PropertyTypeInteger,
										},
									},
								},
							},
						},
					},
				},
			},
			out: `{"_defualt_":{"_all":{"enabled":true}},"some_type":{"properties":{"@timestamp":{"type":"date","format":"yyyy-MM-dd'T'HH:mm:ssZ"},"count":{"type":"integer"},"location":{"type":"geo_point"},"object":{"type":"nested","properties":{"title":{"type":"keyword","ignore_above":256},"user":{"type":"nested","properties":{"age":{"type":"integer"},"first_name":{"type":"keyword"},"last_name":{"type":"keyword"}}}}},"word":{"type":"text","fielddata":true,"fields":{"keyword":{"type":"keyword","ignore_above":256}}}}}}`,
		},
	}

	for _, row := range table {
		_, err := json.Marshal(row.mappings)
		if err != nil {
			t.Fatalf("json.Marshal got error: %s", err)
		}

		m := map[string]Type{}
		if err := json.Unmarshal([]byte(row.out), &m); err != nil {
			t.Fatalf("json.Unmarshal got error: %s", err)
		}

		if !reflect.DeepEqual(row.mappings, m) {
			t.Errorf("deep equal missed. expected(%v) but (%v)", row.mappings, m)
		}
	}
}

func TestTemplateMarshalJSON(t *testing.T) {
	table := []struct {
		template Template
		out      string
	}{
		{
			template: Template{
				Template: "te*",
				Settings: Settings{
					NumberOfShards: 1,
				},
				Mappings: map[string]Type{
					"type1": Type{
						Source: &Source{
							Enabled: false,
						},
						Properties: map[string]Property{
							"host_name": {
								Type: "keyword",
							},
							"created_at": {
								Type:   "date",
								Format: "EEE MMM dd HH:mm:ss Z YYYY",
							},
						},
					},
				},
			},
			out: `{"template":"te*","settings":{"number_of_shards":1},"mappings":{"type1":{"_source":{"enabled":false},"properties":{"created_at":{"type":"date","format":"EEE MMM dd HH:mm:ss Z YYYY"},"host_name":{"type":"keyword"}}}}}`,
		},
	}

	for _, row := range table {
		_, err := json.Marshal(row.template)
		if err != nil {
			t.Fatalf("json.Marshal got error: %s", err)
		}

		tpl := Template{}
		if err := json.Unmarshal([]byte(row.out), &tpl); err != nil {
			t.Fatalf("json.Unmarshal got error: %s", err)
		}

		if !reflect.DeepEqual(row.template, tpl) {
			t.Errorf("deep equal missed. expected(%v) but (%v)", row.template, tpl)
		}
	}
}

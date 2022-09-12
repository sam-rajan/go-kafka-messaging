package converter

import (
	schemaregistry "go-kafka-messaging/internal/pkg/schema-registry"
)

const (
	JSON     string = "JSON"
	AVRO     string = "AVARO"
	PROTOBUF string = "PROTO"
)

type DataConverter interface {
	Convert(message []byte, schemaName string) ([]byte, int)
}

func GetDataConverter(dataFormat string, registry *schemaregistry.SchemaRegistry) DataConverter {

	switch dataFormat {
	case JSON:
		return NewJsonDataFormat(registry)
	case AVRO:
		return NewAvaroDataFormat(registry)
	default:
		return NewJsonDataFormat(registry)
	}
}

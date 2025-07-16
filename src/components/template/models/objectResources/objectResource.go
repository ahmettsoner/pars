package objectResources

import (
	data_resource_payload_structs "parsdevkit.net/modules/resource/data_resource_payload/structs"

	object_resource_payload_structs "parsdevkit.net/modules/resource/object_resource_payload/structs"

	"parsdevkit.net/modules/resource/data_resource"
	"parsdevkit.net/modules/resource/object_resource"
)

type DataLayerComposite struct {
	data_resource.DataLayer
	Original data_resource_payload_structs.Layer
}

type DataSectionComposite struct {
	data_resource.DataSection
	Original data_resource_payload_structs.Section
}

type ObjectLayerComposite struct {
	object_resource.ObjectLayer
	Original object_resource_payload_structs.Layer
}

type ObjectSectionComposite struct {
	object_resource.ObjectSection
	Original object_resource_payload_structs.Section
}

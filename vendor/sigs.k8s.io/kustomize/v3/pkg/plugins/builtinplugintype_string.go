// Code generated by "stringer -type=BuiltinPluginType"; DO NOT EDIT.

package plugins

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[SecretGenerator-1]
	_ = x[ConfigMapGenerator-2]
	_ = x[ReplicaCountTransformer-3]
	_ = x[NamespaceTransformer-4]
	_ = x[PatchJson6902Transformer-5]
	_ = x[PatchStrategicMergeTransformer-6]
	_ = x[PatchTransformer-7]
	_ = x[LabelTransformer-8]
	_ = x[AnnotationsTransformer-9]
	_ = x[PrefixSuffixTransformer-10]
	_ = x[ImageTagTransformer-11]
	_ = x[HashTransformer-12]
	_ = x[InventoryTransformer-13]
	_ = x[LegacyOrderTransformer-14]
}

const _BuiltinPluginType_name = "UnknownSecretGeneratorConfigMapGeneratorReplicaCountTransformerNamespaceTransformerPatchJson6902TransformerPatchStrategicMergeTransformerPatchTransformerLabelTransformerAnnotationsTransformerPrefixSuffixTransformerImageTagTransformerHashTransformerInventoryTransformerLegacyOrderTransformer"

var _BuiltinPluginType_index = [...]uint16{0, 7, 22, 40, 63, 83, 107, 137, 153, 169, 191, 214, 233, 248, 268, 290}

func (i BuiltinPluginType) String() string {
	if i < 0 || i >= BuiltinPluginType(len(_BuiltinPluginType_index)-1) {
		return "BuiltinPluginType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _BuiltinPluginType_name[_BuiltinPluginType_index[i]:_BuiltinPluginType_index[i+1]]
}

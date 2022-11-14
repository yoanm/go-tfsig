package tokens

import (
	"github.com/zclconf/go-cty/cty"
)

// IsCapsuleType returns true if provided `cty.Type` is a special capsule encapsulating `hclwrite.Tokens`.
func IsCapsuleType(t cty.Type) bool {
	return t.IsCapsuleType() && t.FriendlyName() == HclwriteTokensCtyTypeName
}

// ContainsCapsule will deep check if provided value contains a special capsule encapsulating `hclwrite.Tokens`
// (and therefore requires special process to de-encapsulate it).
func ContainsCapsule(valPtr *cty.Value) bool {
	if valPtr == nil {
		return false
	}

	val := *valPtr
	valType := val.Type()

	switch {
	case valType.IsListType() || valType.IsSetType() || valType.IsMapType():
		// Map, List and Set contain only one type of value
		return IsCapsuleType(valType.ElementType())
	case valType.IsTupleType() || valType.IsObjectType():
		// Tuple and Object contain multiple value type => iterate over each of them and check
		for it := val.ElementIterator(); it.Next(); {
			_, eVal := it.Element()
			if ContainsCapsule(&eVal) {
				return true
			}
		}

	case IsCapsuleType(valType):
		return true
	}

	return false
}

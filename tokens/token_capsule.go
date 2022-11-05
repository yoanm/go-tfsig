package tokens

import (
	"github.com/zclconf/go-cty/cty"
)

func IsCapsule(attrType cty.Type) bool {
	return attrType.IsCapsuleType() && attrType.FriendlyName() == HclwriteTokensCtyTypeName
}

func ContainsCapsule(valPtr *cty.Value) bool {
	if valPtr == nil {
		return false
	}

	val := *valPtr
	valType := val.Type()
	switch {
	case valType.IsListType() || valType.IsSetType() || valType.IsTupleType():
		// List and set contain only one type of value
		if valType.IsListType() || valType.IsSetType() {
			return IsCapsule(valType.ElementType())
		}
		for it := val.ElementIterator(); it.Next(); {
			_, eVal := it.Element()
			if ContainsCapsule(&eVal) {
				return true
			}
		}

	case valType.IsMapType() || valType.IsObjectType():
		// Map contains only one type of value
		if valType.IsMapType() {
			return IsCapsule(valType.ElementType())
		}
		// TODO check keys also ????
		for it := val.ElementIterator(); it.Next(); {
			_, eVal := it.Element()
			if ContainsCapsule(&eVal) {
				return true
			}
		}

	case IsCapsule(valType):
		return true
	}

	return false
}

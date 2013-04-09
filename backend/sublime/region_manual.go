package sublime

import (
	"fmt"
	"lime/3rdparty/libs/gopy/lib"
	"lime/backend/primitives"
)

func (o *Region) PyRichCompare(other py.Object, op py.Op) (py.Object, error) {
	if op != py.EQ && op != py.NE {
		return nil, fmt.Errorf("Can only do EQ and NE compares")
	}
	var o2 primitives.Region
	switch t := other.(type) {
	case *Region:
		o2 = t.data
	case *py.Tuple:
		if s := t.Size(); s != 2 {
			return nil, fmt.Errorf("Invalid tuple size: %d != 2", s)
		}
		if a, err := t.GetItem(0); err != nil {
			return nil, err
		} else if b, err := t.GetItem(1); err != nil {
			return nil, err
		} else if a2, ok := a.(*py.Int); !ok {
			return nil, fmt.Errorf("Can only compare with int tuples and other regions")
		} else if b2, ok := b.(*py.Int); !ok {
			return nil, fmt.Errorf("Can only compare with int tuples and other regions")
		} else {
			o2 = primitives.Region{a2.Int(), b2.Int()}
		}
	default:
		return nil, fmt.Errorf("Can only compare with int tuples and other regions")
	}
	if op == py.EQ {
		return toPython(o.data == o2)
	} else {
		return toPython(o.data != o2)
	}
}
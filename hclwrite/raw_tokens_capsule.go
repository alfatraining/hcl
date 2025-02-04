package hclwrite

import (
	"fmt"
	"reflect"

	"github.com/zclconf/go-cty/cty"
)

type rawTokensValue struct{}

var rawTokensKey = rawTokensValue{}

func rawTokensCapsule(tokens Tokens) cty.Value {
	ty := cty.CapsuleWithOps("RawTokens", reflect.TypeOf(Tokens{}), &cty.CapsuleOps{
		GoString: func(v interface{}) string {
			iPtr := v.(*Tokens)
			return fmt.Sprintf("RawTokens(%#v)", *iPtr)
		},
		TypeGoString: func(ty reflect.Type) string {
			return fmt.Sprintf("RawTokens(%s)", ty)
		},
		ExtensionData: func(key interface{}) interface{} {
			if key == rawTokensKey {
				return tokens
			}

			return nil
		},
	})

	return cty.CapsuleVal(ty, &tokens)
}

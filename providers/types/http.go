package types

import "github.com/go-bricks/bricks/v2/constructors/partial"

// HTTPHandlerPatternPair defines pattern -> handler pair
//
// Use this type to provide new Internal HTTP Handler
type HTTPHandlerPatternPair = partial.HTTPHandlerPatternPair

// HTTPHandlerFuncPatternPair defines patter -> handler func pair
//
// Use this type to provide new Internal HTTP Handler
type HTTPHandlerFuncPatternPair = partial.HTTPHandlerFuncPatternPair

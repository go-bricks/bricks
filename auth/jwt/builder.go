package jwt

import (
	"container/list"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/go-bricks/bricks/interfaces/auth/jwt"
)

// JSONDecoder json.Unmarshaler
type JSONDecoder func(data []byte, v interface{}) error

// ExtractorBuilder defines what can be configured when building JWT Token Extractor
type ExtractorBuilder interface {
	SetDecoder(dec JSONDecoder) ExtractorBuilder
	SetContextExtractor(extractor jwt.ContextExtractor) ExtractorBuilder
	SetBase64Decoder(dec *base64.Encoding) ExtractorBuilder
	Build() jwt.TokenExtractor
}

type extractorConfig struct {
	jsonDecoder      JSONDecoder
	base64Enc        *base64.Encoding
	contextExtractor jwt.ContextExtractor
}

type builder struct {
	ll *list.List
}

// Builder creates a fresh instance of Extractor Builder
func Builder() ExtractorBuilder {
	return &builder{
		ll: list.New(),
	}
}

func (b *builder) SetDecoder(dec JSONDecoder) ExtractorBuilder {
	b.ll.PushBack(func(cfg *extractorConfig) {
		cfg.jsonDecoder = dec
	})
	return b
}

func (b *builder) SetContextExtractor(extractor jwt.ContextExtractor) ExtractorBuilder {
	b.ll.PushBack(func(cfg *extractorConfig) {
		cfg.contextExtractor = extractor
	})
	return b
}

func (b *builder) SetBase64Decoder(dec *base64.Encoding) ExtractorBuilder {
	b.ll.PushBack(func(cfg *extractorConfig) {
		cfg.base64Enc = dec
	})
	return b
}

func (b *builder) Build() jwt.TokenExtractor {
	var cfg = new(extractorConfig)
	for e := b.ll.Front(); e != nil; e = e.Next() {
		f := e.Value.(func(config *extractorConfig))
		f(cfg)
	}
	if cfg.base64Enc == nil {
		cfg.base64Enc = base64.RawURLEncoding
	}
	if cfg.jsonDecoder == nil {
		cfg.jsonDecoder = json.Unmarshal
	}
	if cfg.contextExtractor == nil {
		cfg.contextExtractor = func(ctx context.Context) (string, error) {
			return "", fmt.Errorf("no context extractor provided")
		}
	}
	return newTokenExtractor(cfg)
}

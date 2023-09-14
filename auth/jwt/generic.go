package jwt

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-bricks/bricks/v2/interfaces/auth/jwt"
)

type tokenExtractorImpl struct {
	cfg *extractorConfig
}

func newTokenExtractor(cfg *extractorConfig) jwt.TokenExtractor {
	return &tokenExtractorImpl{
		cfg: cfg,
	}
}

func (t *tokenExtractorImpl) FromContext(ctx context.Context) (jwt.Token, error) {
	tokenString, err := t.cfg.contextExtractor(ctx)
	if err == nil {
		return t.FromString(tokenString)
	}
	return nil, err
}

func (t *tokenExtractorImpl) FromString(str string) (token jwt.Token, err error) {
	if parts := strings.Split(str, "."); len(parts) == 3 {
		var payload []byte
		if payload, err = t.cfg.base64Enc.DecodeString(parts[1]); err == nil {
			return newToken(str, payload, t.cfg.jsonDecoder), nil
		}
		return nil, fmt.Errorf("error decoding from base 64 %w", err)
	}
	return nil, fmt.Errorf("%s is not a JWT", str)
}

type tokenInstance struct {
	raw         string
	payload     []byte
	jsonDecoder JSONDecoder
}

func newToken(jwtAsString string, justPayload []byte, decoder JSONDecoder) jwt.Token {
	return &tokenInstance{
		raw:         jwtAsString,
		payload:     justPayload,
		jsonDecoder: decoder,
	}
}

func (t *tokenInstance) Raw() string {
	return t.raw
}

func (t *tokenInstance) Payload() []byte {
	return t.payload
}

func (t *tokenInstance) Map() (output map[string]interface{}, err error) {
	output = make(map[string]interface{})
	err = t.jsonDecoder(t.payload, &output)
	return
}

func (t *tokenInstance) Decode(target interface{}) error {
	return t.jsonDecoder(t.payload, target)
}

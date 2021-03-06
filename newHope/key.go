package newHope

import (
	"io"

	"go.dedis.ch/onet/v3/glyph"
)

//GenerateKey generates a public, private key pair for Glyph
//for large coefficients
func (g *GlyphSuite) GenerateKey(rand io.Reader) (PublicKey, PrivateKey, error) {
	if rand != nil {
		//TODO: Use it
		secretBuffer := make([]byte, NewHopePrivateKeySize)
		_, e := rand.Read(secretBuffer)
		if e != nil {
			return nil, nil, e
		}
		//TODO: make this an actual thing
		return nil, nil, nil
	}
	ctx := glyph.GetCtx()
	private, e := glyph.NewPrivateKey(ctx, glyph.GetA(ctx))
	if e != nil {
		return nil, nil, e
	}
	public := private.PK()
	publicData, e1 := public.Marshall()
	if e1 != nil {
		return nil, nil, e1
	}
	/*t, ep := checkPublicKey(publicData, ctx)
	if ep != nil {
		return nil, nil, ep
	}*/
	privateData, e2 := private.Marshall()
	if e2 != nil {
		return nil, nil, e2
	}
	return publicData, privateData, nil
}

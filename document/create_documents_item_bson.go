package document // nolint:dupl

import (
	bsonenc "github.com/spikeekips/mitum/util/encoder/bson"
	"go.mongodb.org/mongo-driver/bson"
)

func (it CreateDocumentsItemImpl) MarshalBSON() ([]byte, error) {
	return bsonenc.Marshal(
		bsonenc.MergeBSONM(bsonenc.NewHintedDoc(it.Hint()),
			bson.M{
				// "doctype":  it.doctype,
				"doc":      it.doc,
				"currency": it.cid,
			}),
	)
}

type CreateDocumentsItemImplBSONUnpacker struct {
	// DT string   `bson:"doctype"`
	DD bson.Raw `bson:"doc"`
	CI string   `bson:"currency"`
}

func (it *CreateDocumentsItemImpl) UnpackBSON(b []byte, enc *bsonenc.Encoder) error {
	var ucd CreateDocumentsItemImplBSONUnpacker
	if err := bson.Unmarshal(b, &ucd); err != nil {
		return err
	}

	return it.unpack(
		enc,
		// ucd.DT,
		ucd.DD,
		ucd.CI,
	)
}

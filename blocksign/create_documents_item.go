package blocksign

import (
	"github.com/soonkuk/mitum-data/currency"
	"github.com/spikeekips/mitum/base"
	"github.com/spikeekips/mitum/util"
	"github.com/spikeekips/mitum/util/hint"
)

type BaseCreateDocumentsItem struct {
	hint     hint.Hint
	fileHash FileHash
	signers  []base.Address
	cid      currency.CurrencyID
}

func NewBaseCreateDocumentsItem(ht hint.Hint, filehash FileHash, signers []base.Address, cid currency.CurrencyID) BaseCreateDocumentsItem {
	return BaseCreateDocumentsItem{
		hint:     ht,
		fileHash: filehash,
		signers:  signers,
		cid:      cid,
	}
}

func (it BaseCreateDocumentsItem) Hint() hint.Hint {
	return it.hint
}

func (it BaseCreateDocumentsItem) Bytes() []byte {
	bs := make([][]byte, 2)
	bs[0] = it.fileHash.Bytes()
	bs[1] = it.cid.Bytes()

	return util.ConcatBytesSlice(bs...)
}

func (it BaseCreateDocumentsItem) IsValid([]byte) error {

	/*
		signers := map[string]bool{}
		for i := range it.signers {
			_, found := signers[it.signers[i].String()]
			if found {
				return xerrors.Errorf("duplicated signer, %v", it.signers[i])
			}
			signers[it.signers[i].String()] = true
		}
	*/
	if err := it.fileHash.IsValid(nil); err != nil {
		return err
	}

	return nil
}

// FileHash return BaseCreateDocumetsItem's owner address.
func (it BaseCreateDocumentsItem) FileHash() FileHash {
	return it.fileHash
}

func (it BaseCreateDocumentsItem) Signers() []base.Address {
	return it.signers
}

// FileData return BaseCreateDocumentsItem's fileData.
func (it BaseCreateDocumentsItem) Currency() currency.CurrencyID {
	return it.cid
}

func (it BaseCreateDocumentsItem) Rebuild() CreateDocumentsItem {
	return it
}
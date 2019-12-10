package strgpwr

import (
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p-core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

/* This file was generated by github.com/whyrusleeping/cbor-gen */

var _ = xerrors.Errorf

func (t *CreateStorageMinerParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.t.Owner (address.Address) (struct)
	if err := t.Owner.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.Worker (address.Address) (struct)
	if err := t.Worker.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.SectorSize (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.SectorSize))); err != nil {
		return err
	}

	// t.t.PeerID (peer.ID) (string)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.PeerID)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.PeerID)); err != nil {
		return err
	}
	return nil
}

func (t *CreateStorageMinerParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Owner (address.Address) (struct)

	{

		if err := t.Owner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.t.Worker (address.Address) (struct)

	{

		if err := t.Worker.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.t.SectorSize (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.SectorSize = uint64(extra)
	// t.t.PeerID (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.PeerID = peer.ID(sval)
	}
	return nil
}

func (t *UpdateStorageParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.t.Delta (types.BigInt) (struct)
	if err := t.Delta.MarshalCBOR(w); err != nil {
		return err
	}

	// t.t.NextProvingPeriodEnd (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.NextProvingPeriodEnd))); err != nil {
		return err
	}

	// t.t.PreviousProvingPeriodEnd (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.PreviousProvingPeriodEnd))); err != nil {
		return err
	}
	return nil
}

func (t *UpdateStorageParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Delta (types.BigInt) (struct)

	{

		if err := t.Delta.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	// t.t.NextProvingPeriodEnd (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.NextProvingPeriodEnd = uint64(extra)
	// t.t.PreviousProvingPeriodEnd (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.PreviousProvingPeriodEnd = uint64(extra)
	return nil
}

func (t *StoragePowerState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{133}); err != nil {
		return err
	}

	// t.t.Miners (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.Miners); err != nil {
		return xerrors.Errorf("failed to write cid field t.Miners: %w", err)
	}

	// t.t.ProvingBuckets (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.ProvingBuckets); err != nil {
		return xerrors.Errorf("failed to write cid field t.ProvingBuckets: %w", err)
	}

	// t.t.MinerCount (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.MinerCount))); err != nil {
		return err
	}

	// t.t.LastMinerCheck (uint64) (uint64)
	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.LastMinerCheck))); err != nil {
		return err
	}

	// t.t.TotalStorage (types.BigInt) (struct)
	if err := t.TotalStorage.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *StoragePowerState) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Miners (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Miners: %w", err)
		}

		t.Miners = c

	}
	// t.t.ProvingBuckets (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ProvingBuckets: %w", err)
		}

		t.ProvingBuckets = c

	}
	// t.t.MinerCount (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.MinerCount = uint64(extra)
	// t.t.LastMinerCheck (uint64) (uint64)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint64 field")
	}
	t.LastMinerCheck = uint64(extra)
	// t.t.TotalStorage (types.BigInt) (struct)

	{

		if err := t.TotalStorage.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *PledgeCollateralParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.t.Size (types.BigInt) (struct)
	if err := t.Size.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *PledgeCollateralParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Size (types.BigInt) (struct)

	{

		if err := t.Size.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

func (t *PowerLookupParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *PowerLookupParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(br); err != nil {
			return err
		}

	}
	return nil
}

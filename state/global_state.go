package state

import "github.com/SemyonHoyrish/GoPlayEngine/data_structures"

type FlagType uint32

const (
	GF_SublayerRebuild FlagType = iota
)

var GlobalFlags map[FlagType]bool

type EntryType uint32

const (
	ET_Node             EntryType = iota
	ET_OverlapInterface EntryType = iota
)

type StoreEntry struct {
	EntryType EntryType

	// TODO: HACK: We are not allowing GC to remove object because we keep pointer to that object!
	Pointer any
}

var AllNodes data_structures.Set[StoreEntry] = data_structures.CreateSet[StoreEntry]()

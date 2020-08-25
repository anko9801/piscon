package parameter

import "time"

const (
	NumOfSearchChairInScenario     = 5
	NumOfSearchEstateInScenario    = 5
	NumOfCheckChairSearchPaging    = 3
	NumOfCheckEstateSearchPaging   = 3
	NumOfCheckChairDetailPage      = 7
	NumOfCheckEstateDetailPage     = 3
	PerPageOfChairSearch           = 30
	PerPageOfEstateSearch          = 30
	MaxLengthOfNazotteResponse     = 50
	NeighborhoodRadiusOfNazotte    = 1e-6
	SleepTimeOnFailScenario        = 1500 * time.Millisecond
	SleepSwingOnFailScenario       = 500 // * time.Millisecond
	SleepTimeOnUserAway            = 500 * time.Millisecond
	SleepSwingOnUserAway           = 100 // * time.Millisecond
	SleepTimeOnBotInterval         = 500 * time.Millisecond
	SleepSwingOnBotInterval        = 100 // * time.Millisecond
	SleepBeforePostDraft           = 500 * time.Millisecond
	SleepSwingBeforePostDraft      = 100 // * time.Millisecond
	IntervalForCheckWorkers        = 5 * time.Second
	ThresholdTimeOfAbandonmentPage = 1000 * time.Millisecond
	DefaultAPITimeout              = 2000 * time.Millisecond
	InitializeTimeout              = 30 * time.Second
	VerifyTimeout                  = 10 * time.Second
	LoadTimeout                    = 60 * time.Second
)

var BoundaryOfLevel []int64 = []int64{
	400, 800, 1200, 1600, 2000,
	2400, 2800, 3200, 3600, 4000,
	4400, 4800, 5200, 5600, 6000,
	6400, 6800, 7200, 7600, 8000,
	8400, 8800, 9200, 9600, 10000,
}

type incWorkers struct {
	ChairSearchWorker         int
	EstateSearchWorker        int
	EstateNazotteSearchWorker int
	BotWorker                 int
	ChairDraftPostWorker      int
	EstateDraftPostWorker     int
}

// IncListOfWorkers 前のレベルとのWorkerの個数の差分を保持するList
var ListOfIncWorkers = []incWorkers{
	{ // level 00
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      0,
		EstateDraftPostWorker:     0,
	},
	{ // level 01
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      0,
		EstateDraftPostWorker:     0,
	},
	{ // level 02
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      0,
		EstateDraftPostWorker:     0,
	},
	{ // level 03
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      0,
		EstateDraftPostWorker:     0,
	},
	{ // level 04
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      0,
		EstateDraftPostWorker:     0,
	},
	{ // level 05
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 06
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 07
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 08
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 09
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 10
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 11
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 12
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 13
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 14
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 15
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 16
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 17
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 18
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
	{ // level 19
		ChairSearchWorker:         1,
		EstateSearchWorker:        1,
		EstateNazotteSearchWorker: 1,
		BotWorker:                 1,
		ChairDraftPostWorker:      1,
		EstateDraftPostWorker:     1,
	},
}

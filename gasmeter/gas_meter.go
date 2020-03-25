package gasmeter

import (
	"container/list"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/filecoin-project/chain-validation/box"
	"github.com/filecoin-project/chain-validation/chain/types"
)

const ValidationDataEnvVar = "CHAIN_VALIDATION_DATA"

type trackerElement struct {
	receipt types.MessageReceipt
}

func (te *trackerElement) fileKey() string {
	return fmt.Sprintf("%d", te.receipt.GasUsed)
}

type GasMeter struct {
	tracker *list.List
	T       testing.TB

	// index in gasUnits of expected gas
	gasIdx int
	// slice of gas units used by the test
	expectedGasUnits []int64
}

func NewGasMeter(t testing.TB) *GasMeter {
	return &GasMeter{
		tracker:          list.New(),
		T:                t,
		gasIdx:           0,
		expectedGasUnits: LoadGasForTest(t),
	}
}

func (gm *GasMeter) Track(receipt types.MessageReceipt) {
	gm.tracker.PushBack(&trackerElement{
		receipt: receipt,
	})
}

func (gm *GasMeter) NextExpectedGas() (types.GasUnits, bool) {
	defer func() { gm.gasIdx += 1 }()
	if gm.gasIdx > len(gm.expectedGasUnits)-1 {
		// didn't find any gas
		return 0, false
	}
	return types.GasUnits(gm.expectedGasUnits[gm.gasIdx]), true
}

// write the contents of gm.tracker to a file using the format:
// GasUnit
// GasUnit
// ...
func (gm *GasMeter) Record() {
	file := getTestDataFilePath(gm.T)
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		gm.T.Log(err)
		return
	}
	defer f.Close()

	for e := gm.tracker.Front(); e != nil; e = e.Next() {
		_, err := fmt.Fprintf(f, "%s\n", e.Value.(*trackerElement).fileKey())
		if err != nil {
			gm.T.Fatal(err)
		}
	}
}

// Given a testing T, load the gas file associated with it and return a slice of the gas used by the test
// an index in the slice represents the order of apply message calls.
func LoadGasForTest(t testing.TB) []int64 {
	fileName := filenameFromTest(t)
	f, found := box.Get(fileName)
	if !found {
		t.Logf("WARNING (does NOT indicate test failure): can't find file: %s", fileName)
		// return an empty slice here since `NextExpectedGas` performs bounds checking
		return []int64{}
	}
	return f
}

func getTestDataFilePath(t testing.TB) string {
	dataPath := os.Getenv(ValidationDataEnvVar)
	if dataPath == "" {
		t.Fatalf("failed to find validation data path, make sure %s is set", ValidationDataEnvVar)
	}
	return filepath.Join(dataPath, filenameFromTest(t))
}

// return a string containing only letters and number.
func filenameFromTest(t testing.TB) string {
	// only want letters and numbers
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		t.Fatal(err)
	}
	return fmt.Sprintf("/%s", reg.ReplaceAllString(t.Name(), ""))
}

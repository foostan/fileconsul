package fileconsul

import (
	"testing"
)

func TestDiff(t *testing.T) {
	NewRFList := RFList{
		Remotefile{Path: "/path/to/sample1", Url: "http://path/to/sample1", Hash: "12"},
		Remotefile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "12"},
		Remotefile{Path: "/path/to/sample5", Url: "http://path/to/sample4", Hash: "90"},
		Remotefile{Path: "/path/to/sample6", Url: "http://path/to/sample4", Hash: "12"},
	}

	OldRFList := RFList{
		Remotefile{Path: "/path/to/sample1", Url: "http://path/to/sample1", Hash: "12"},
		Remotefile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "34"},
		Remotefile{Path: "/path/to/sample3", Url: "http://path/to/sample3", Hash: "56"},
		Remotefile{Path: "/path/to/sample4", Url: "http://path/to/sample4", Hash: "78"},
	}

	rfDiff := NewRFList.Diff(OldRFList)

	addMfDiff := RFList{
		Remotefile{Path: "/path/to/sample5", Url: "http://path/to/sample4", Hash: "90"},
		Remotefile{Path: "/path/to/sample6", Url: "http://path/to/sample4", Hash: "12"}}
	if len(rfDiff.Add) != len(addMfDiff) {
		t.Fatalf("expected result is %s, but %s", addMfDiff, rfDiff.Add)
	}
	for i := 0; i < len(addMfDiff); i++ {
		if !addMfDiff[i].EqVer(rfDiff.Add[i]) {
			t.Fatalf("expected result is %s, but %s", addMfDiff[i], rfDiff.Add[i])
		}
	}

	delMfDiff := RFList{
		Remotefile{Path: "/path/to/sample3", Url: "http://path/to/sample3", Hash: "56"},
		Remotefile{Path: "/path/to/sample4", Url: "http://path/to/sample4", Hash: "78"}}
	if len(rfDiff.Del) != len(delMfDiff) {
		t.Fatalf("expected result is %s, but %s", delMfDiff, rfDiff.Del)
	}
	for i := 0; i < len(delMfDiff); i++ {
		if !delMfDiff[i].EqVer(rfDiff.Del[i]) {
			t.Fatalf("expected result is %s, but %s", delMfDiff[i], rfDiff.Del[i])
		}
	}

	newMfDiff := RFList{
		Remotefile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "12"}}
	if len(rfDiff.New) != len(newMfDiff) {
		t.Fatalf("expected result is %s, but %s", newMfDiff, rfDiff.New)
	}
	for i := 0; i < len(newMfDiff); i++ {
		if !newMfDiff[i].EqVer(rfDiff.New[i]) {
			t.Fatalf("expected result is %s, but %s", newMfDiff[i], rfDiff.New[i])
		}
	}

	oldMfDiff := RFList{
		Remotefile{Path: "/path/to/sample2", Url: "http://path/to/sample2", Hash: "34"}}
	if len(rfDiff.Old) != len(oldMfDiff) {
		t.Fatalf("expected result is %s, but %s", oldMfDiff, rfDiff.Old)
	}
	for i := 0; i < len(oldMfDiff); i++ {
		if !oldMfDiff[i].EqVer(rfDiff.Old[i]) {
			t.Fatalf("expected result is %s, but %s", oldMfDiff[i], rfDiff.Old[i])
		}
	}

	eqMfDiff := RFList{
		Remotefile{Path: "/path/to/sample1", Url: "http://path/to/sample1", Hash: "12"}}
	if len(rfDiff.Eq) != len(eqMfDiff) {
		t.Fatalf("expected result is %s, but %s", eqMfDiff, rfDiff.Eq)
	}
	for i := 0; i < len(eqMfDiff); i++ {
		if !eqMfDiff[i].EqVer(rfDiff.Eq[i]) {
			t.Fatalf("expected result is %s, but %s", eqMfDiff[i], rfDiff.Eq[i])
		}
	}
}

func TestReadRFList(t *testing.T) {
	client, err := NewClient(&ClientConfig{
		ConsulAddr: "localhost:8500",
		ConsulDC:   "dc1",
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	_, err = client.ReadRFList("fileconsul")
	if err != nil {
		t.Skipf("err: %v", err)
	}
}

func TestStrToRFValue(t *testing.T) {
	str := "http://path/to/sample1,12"
	rfValue := StrToRFValue(str)
	res := rfValue.ToStr()

	if str != res {
		t.Fatalf("expected result is %s, but %s", str, res)
	}
}

package fileconsul

import (
	"testing"
)

func TestDiff(t *testing.T) {
	NewRFList := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample1", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample2", Hash: "656b38f3402a1e8b4211fac826efd433", Data: []byte("sample2")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample3", Hash: "d35f70211135de265bc7c66df4dd3605", Data: []byte("sample3")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample4", Hash: "247f4201f214ff279da3a24570d642f1", Data: []byte("sample4")},
	}

	OldRFList := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample1", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample2", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample5", Hash: "828974c6c954abd2ada226a48c7d6090", Data: []byte("sample5")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample6", Hash: "1d1756986764035547f4a1e1a106d7d1", Data: []byte("sample6")},
	}

	rfDiff := NewRFList.Diff(OldRFList)

	addMfDiff := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample3", Hash: "d35f70211135de265bc7c66df4dd3605", Data: []byte("sample3")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample4", Hash: "247f4201f214ff279da3a24570d642f1", Data: []byte("sample4")}}
	if len(rfDiff.Add) != len(addMfDiff) {
		t.Fatalf("expected result is %s, but %s", addMfDiff, rfDiff.Add)
	}
	for i := 0; i < len(addMfDiff); i++ {
		if !addMfDiff[i].EqVer(rfDiff.Add[i]) {
			t.Fatalf("expected result is %s, but %s", addMfDiff[i], rfDiff.Add[i])
		}
	}

	delMfDiff := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample5", Hash: "828974c6c954abd2ada226a48c7d6090", Data: []byte("sample5")},
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample6", Hash: "1d1756986764035547f4a1e1a106d7d1", Data: []byte("sample6")}}
	if len(rfDiff.Del) != len(delMfDiff) {
		t.Fatalf("expected result is %s, but %s", delMfDiff, rfDiff.Del)
	}
	for i := 0; i < len(delMfDiff); i++ {
		if !delMfDiff[i].EqVer(rfDiff.Del[i]) {
			t.Fatalf("expected result is %s, but %s", delMfDiff[i], rfDiff.Del[i])
		}
	}

	newMfDiff := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample2", Hash: "656b38f3402a1e8b4211fac826efd433", Data: []byte("sample2")}}
	if len(rfDiff.New) != len(newMfDiff) {
		t.Fatalf("expected result is %s, but %s", newMfDiff, rfDiff.New)
	}
	for i := 0; i < len(newMfDiff); i++ {
		if !newMfDiff[i].EqVer(rfDiff.New[i]) {
			t.Fatalf("expected result is %s, but %s", newMfDiff[i], rfDiff.New[i])
		}
	}

	oldMfDiff := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample2", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")}}
	if len(rfDiff.Old) != len(oldMfDiff) {
		t.Fatalf("expected result is %s, but %s", oldMfDiff, rfDiff.Old)
	}
	for i := 0; i < len(oldMfDiff); i++ {
		if !oldMfDiff[i].EqVer(rfDiff.Old[i]) {
			t.Fatalf("expected result is %s, but %s", oldMfDiff[i], rfDiff.Old[i])
		}
	}

	eqMfDiff := RFList{
		Remotefile{Prefix: "fileconsul", Path: "/path/to/sample1", Hash: "ac46374a846d97e22f917b6863f690ad", Data: []byte("sample1")}}
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

package fpm

import (
	"testing"
	"time"

	"github.com/teeratpitakrat/hora/adm"
	"github.com/teeratpitakrat/hora/rbridge"
)

func TestCreate(t *testing.T) {
	m := make(adm.ADM)

	compA := adm.Component{"A", "host1"}
	compB := adm.Component{"B", "host2"}
	compC := adm.Component{"C", "host3"}
	compD := adm.Component{"D", "host4"}

	depA := adm.DependencyInfo{compA, make([]adm.Dependency, 2, 2)}
	depA.Component = compA
	depA.Dependencies[0] = adm.Dependency{compB, 0.5}
	depA.Dependencies[1] = adm.Dependency{compC, 0.5}
	m[compA.UniqName()] = depA

	depB := adm.DependencyInfo{compB, make([]adm.Dependency, 1, 1)}
	depB.Component = compB
	depB.Dependencies[0] = adm.Dependency{compD, 1}
	m[compB.UniqName()] = depB

	depC := adm.DependencyInfo{compC, make([]adm.Dependency, 1, 1)}
	depC.Component = compC
	depC.Dependencies[0] = adm.Dependency{compD, 1}
	m[compC.UniqName()] = depC

	depD := adm.DependencyInfo{}
	depD.Component = compD
	m[compD.UniqName()] = depD

	// Configure R bridge
	rbridge.SetHostname("localhost")
	rbridge.SetPort(6311)

	var f FPMBNR
	f.LoadADM(m)
	err := f.Create()
	if err != nil {
		t.Error("Error creating FPM", err)
	}

	res, err := f.Predict()
	if err != nil {
		t.Error("Error making prediction", err)
	}
	// TODO: more precision checks
	fprobA := res[adm.Component{"A", "host1"}]
	if fprobA != 0 {
		t.Error("Expected: 0 but got", fprobA)
	}
	fprobB := res[adm.Component{"B", "host2"}]
	if fprobB != 0 {
		t.Error("Expected: 0 but got", fprobB)
	}
	fprobC := res[adm.Component{"C", "host3"}]
	if fprobC != 0 {
		t.Error("Expected: 0 but got", fprobC)
	}
	fprobD := res[adm.Component{"D", "host4"}]
	if fprobD != 0 {
		t.Error("Expected: 0 but got", fprobD)
	}

	f.Update(adm.Component{"D", "host4"}, 0.1)
	time.Sleep(100 * time.Millisecond)
	res, err = f.Predict()
	if err != nil {
		t.Error("Error making prediction", err)
	}
	fprobA = res[adm.Component{"A", "host1"}]
	if fprobA > 0.12 {
		t.Error("Expected: 0 but got", fprobA)
	}
	fprobB = res[adm.Component{"B", "host2"}]
	if fprobB > 0.12 {
		t.Error("Expected: 0 but got", fprobB)
	}
	fprobC = res[adm.Component{"C", "host3"}]
	if fprobC > 0.12 {
		t.Error("Expected: 0 but got", fprobC)
	}
	fprobD = res[adm.Component{"D", "host4"}]
	if fprobD > 0.12 {
		t.Error("Expected: 0 but got", fprobD)
	}

	f.Update(adm.Component{"D", "host4"}, 0.9)
	time.Sleep(100 * time.Millisecond)
	res, err = f.Predict()
	if err != nil {
		t.Error("Error making prediction", err)
	}
	fprobA = res[adm.Component{"A", "host1"}]
	if fprobA < 0.89 {
		t.Error("Expected: 0 but got", fprobA)
	}
	fprobB = res[adm.Component{"B", "host2"}]
	if fprobB < 0.89 {
		t.Error("Expected: 0 but got", fprobB)
	}
	fprobC = res[adm.Component{"C", "host3"}]
	if fprobC < 0.89 {
		t.Error("Expected: 0 but got", fprobC)
	}
	fprobD = res[adm.Component{"D", "host4"}]
	if fprobD < 0.89 {
		t.Error("Expected: 0 but got", fprobD)
	}

	f.Update(adm.Component{"D", "host4"}, 0.0)
	f.Update(adm.Component{"B", "host2"}, 0.1)
	time.Sleep(100 * time.Millisecond)
	res, err = f.Predict()
	if err != nil {
		t.Error("Error making prediction", err)
	}
	fprobA = res[adm.Component{"A", "host1"}]
	if fprobA > 0.12 {
		t.Error("Expected: 0 but got", fprobA)
	}
	fprobB = res[adm.Component{"B", "host2"}]
	if fprobB > 0.12 {
		t.Error("Expected: 0 but got", fprobB)
	}
	fprobC = res[adm.Component{"C", "host3"}]
	if fprobC != 0 {
		t.Error("Expected: 0 but got", fprobC)
	}
	fprobD = res[adm.Component{"D", "host4"}]
	if fprobD != 0 {
		t.Error("Expected: 0 but got", fprobD)
	}

	f.Update(adm.Component{"B", "host2"}, 0.0)
	f.Update(adm.Component{"A", "host1"}, 0.1)
	time.Sleep(100 * time.Millisecond)
	res, err = f.Predict()
	if err != nil {
		t.Error("Error making prediction", err)
	}
	fprobA = res[adm.Component{"A", "host1"}]
	if fprobA > 0.12 {
		t.Error("Expected: 0 but got", fprobA)
	}
	fprobB = res[adm.Component{"B", "host2"}]
	if fprobB > 0.1 {
		t.Error("Expected: 0 but got", fprobB)
	}
	fprobC = res[adm.Component{"C", "host3"}]
	if fprobC != 0 {
		t.Error("Expected: 0 but got", fprobC)
	}
	fprobD = res[adm.Component{"D", "host4"}]
	if fprobD != 0 {
		t.Error("Expected: 0 but got", fprobD)
	}
}
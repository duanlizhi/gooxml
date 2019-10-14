// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml_test

import (
	"encoding/xml"
	"testing"

	"baliance.com/gooxml/schema/soo/dml"
)

func TestTblConstructor(t *testing.T) {
	v := dml.NewTbl()
	if v == nil {
		t.Errorf("dml.NewTbl must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.Tbl should validate: %s", err)
	}
}

func TestTblMarshalUnmarshal(t *testing.T) {
	v := dml.NewTbl()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewTbl()
	xml.Unmarshal(buf, v2)
}

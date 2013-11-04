// Copyright 2013 Justin Wilson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"time"
)

type NullTime struct {
	Time  time.Time
	Valid bool
}

func (nt *NullTime) Scan(value interface{}) error {
	if value == nil {
		nt.Time, nt.Valid = time.Time{}, false
		return nil
	}

	t, ok := value.(time.Time)
	if !ok {
		s := fmt.Sprintf("%v", value)
		v := reflect.ValueOf(value)
		return fmt.Errorf("converting %v %q to a time.Time", v.Kind(), s)
	}
	nt.Valid, nt.Time = true, t
	return nil
}

func (nt *NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

// Copyright 2013 Justin Wilson. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package sql implements some sql helpers to be used on top of the core lib database/sql.

NullTime implements a nullable time.Time

	var t time.NullTime
	row := db.QueryRow(`SELECT timestamp FROM TimesTable LIMIT(1)`)
	row.Scan(&t)
	timestamp := t.Value()

*/
package sql

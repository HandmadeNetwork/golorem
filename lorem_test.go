// Copyright 2012 Derek A. Rhodes.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lorem

import (
	"log"
	"testing"
)

func TestAll(t *testing.T) {
	for i := 0; i < 5; i++ {
		log.Print(word(i))
		for j := 0; j < 5; j++ {
			t.Logf("     Word (%d, %d): %s", i, j, Word(i, j))
			t.Logf(" Sentence (%d, %d): %s", i, j, Sentence(i, j))
			t.Logf("Paragraph (%d, %d): %s", i, j, Paragraph(i, j))
		}
	}
	t.Log(Url())
	t.Log(Host())
	t.Log(Email())
}

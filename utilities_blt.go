/*
Copyright (c) 2020, Tomasz "VedVid" Nowakowski
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package main

import (
	blt "bearlibterminal"
	"strconv"
)

func SetGlyph(path, number, filter string, size int) string {
	/* Function SetTile allows to use special tiles (glyphs, bitmaps)
	   as font elements; it returns unicode number.
	   Number variable has to be formatted in that way:
	   U+<unicode-number>, like: U+E001
	   Later, that U+E001 identifier may be used in printing functions, like
	   wall := 0xE001 (note different format!); blt.Print(x, y, wall). */
	blt.Set(number + ": " + path + ", resize=" + strconv.Itoa(size) + ", resize-filter=" + filter)
	return "0x" + number[2:]
}

func SetColor(name, number string) string {
	/* Function SetColor allows to declare specified colors
	   by passing custom name and its code.
	   By default, it uses hex values, but BearLibTerminal
	   supports others formats as well:
	   check blt documentation available on
	   http://foo.wyrd.name/en:bearlibterminal:reference
	   SetColor returns name string. */
	blt.Set("palette: " + name + " = " + number)
	return name
}

func SetBkColor(color string) {
	blt.BkColor(blt.ColorFromName("black"))
}

func RuneCountInBltString(s string) int {
	length := 0
	var r = []rune(s)
	internal := false
	for i, v := range r {
		nextRune := ' '
		prevRune := ' '
		if i < len(r)-1 {
			nextRune = r[i+1]
		}
		if i > 0 {
			prevRune = r[i-1]
		}
		if internal == false {
			if v == '[' && nextRune != '[' && prevRune != '[' {
				internal = true
			}
		} else {
			if v == ']' && nextRune != ']' && prevRune != ']' {
				internal = false
			}
		}
		if internal == false {
			if v != ']' {
				length++
			} else {
				if prevRune == ']' || nextRune == ']' {
					length++
				}
			}
		}
	}
	return length
}
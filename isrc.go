/*
 * Copyright (c) 2014 Michael Wendland
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the "Software"),
 * to deal in the Software without restriction, including without limitation
 * the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 * FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
 * IN THE SOFTWARE.
 *
 * 	Authors:
 * 		Michael Wendland <michael@michiwend.com>
 */

package gomusicbrainz

import "encoding/xml"

// ISRC represents generally a musician, a group of musicians, a collaboration
// of multiple musicians or other music professionals.
type ISRC struct {
	recordingListResult

	ID MBID `xml:"id,attr"`
}

func (mbe *ISRC) lookupResult() interface{} {
	var res struct {
		XMLName xml.Name `xml:"metadata"`
		Ptr     *ISRC    `xml:"isrc"`
	}
	res.Ptr = mbe
	return &res
}

func (mbe *ISRC) apiEndpoint() string {
	return "/isrc"
}

func (mbe *ISRC) Id() MBID {
	return mbe.ID
}

// LookupISRC performs an isrc lookup request for the given MBID.
func (c *WS2Client) LookupISRC(id MBID, inc ...string) (*ISRC, error) {
	a := &ISRC{ID: id}
	err := c.Lookup(a, inc...)

	return a, err
}

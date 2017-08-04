// Copyright 2015 ThoughtWorks, Inc.
// Copyright (C) 2012 vova616 <vova616@gmail.com>

// This file is part of Gauge-Screenshot.

// Gauge-Screenshot is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge-Screenshot is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge-Screenshot.  If not, see <http://www.gnu.org/licenses/>.

// +build !darwin

package capture

import (
	"image"
	"image/png"
	"os"

	"github.com/vova616/screenshot"
)

func CaptureScreen(filename string) error {
	img, err := screenshot.CaptureScreen()
	if err != nil {
		return err
	}
	pngImg := image.Image(img)

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	err = png.Encode(f, pngImg)
	if err != nil {
		return err
	}
	return nil
}

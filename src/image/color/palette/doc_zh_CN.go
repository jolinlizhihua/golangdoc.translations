// Copyright The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

// Package palette provides standard color palettes.

// palette包提供了标准的调色板。
package palette

// Plan9 is a 256-color palette that partitions the 24-bit RGB space into 4×4×4
// subdivision, with 4 shades in each subcube. Compared to the WebSafe, the idea is
// to reduce the color resolution by dicing the color cube into fewer cells, and to
// use the extra space to increase the intensity resolution. This results in 16
// gray shades (4 gray subcubes with 4 samples in each), 13 shades of each primary
// and secondary color (3 subcubes with 4 samples plus black) and a reasonable
// selection of colors covering the rest of the color cube. The advantage is better
// representation of continuous tones.
//
// This palette was used in the Plan 9 Operating System, described at
// http://plan9.bell-labs.com/magic/man2html/6/color

// Plan9是一个256色的调色板，将24位的RGB色彩立方划分为4x4x4的子立方，每个子立方体都有4个色调。比起WebSafe，本方案通过将颜色立方划分成更少的方格减少了颜色的分辨率，以便用额外的空间换取增加强度的分辨率。结果是灰色有16个色调（4个灰色子立方每个有4个采样），初级和次级颜色各有13个色调（3个子立方每个有4个采样，再加上黑色），加上对色彩立方其余部分的合理色彩选择，实现了更好的连续色调的表现。
//
// 本调色板用于Plan9操作系统，参见http://plan9.bell-labs.com/magic/man2html/6/color
//
//	var WebSafe = []color.Color{
//	    color.RGBA{0x00, 0x00, 0x00, 0xff},
//	    color.RGBA{0x00, 0x00, 0x33, 0xff},
//	    color.RGBA{0x00, 0x00, 0x66, 0xff},
//	    color.RGBA{0x00, 0x00, 0x99, 0xff},
//	    color.RGBA{0x00, 0x00, 0xcc, 0xff},
//	    color.RGBA{0x00, 0x00, 0xff, 0xff},
//	    color.RGBA{0x00, 0x33, 0x00, 0xff},
//	    color.RGBA{0x00, 0x33, 0x33, 0xff},
//	    color.RGBA{0x00, 0x33, 0x66, 0xff},
//	    color.RGBA{0x00, 0x33, 0x99, 0xff},
//	    color.RGBA{0x00, 0x33, 0xcc, 0xff},
//	    color.RGBA{0x00, 0x33, 0xff, 0xff},
//	    color.RGBA{0x00, 0x66, 0x00, 0xff},
//	    color.RGBA{0x00, 0x66, 0x33, 0xff},
//	    color.RGBA{0x00, 0x66, 0x66, 0xff},
//	    color.RGBA{0x00, 0x66, 0x99, 0xff},
//	    color.RGBA{0x00, 0x66, 0xcc, 0xff},
//	    color.RGBA{0x00, 0x66, 0xff, 0xff},
//	    color.RGBA{0x00, 0x99, 0x00, 0xff},
//	    color.RGBA{0x00, 0x99, 0x33, 0xff},
//	    color.RGBA{0x00, 0x99, 0x66, 0xff},
//	    color.RGBA{0x00, 0x99, 0x99, 0xff},
//	    color.RGBA{0x00, 0x99, 0xcc, 0xff},
//	    color.RGBA{0x00, 0x99, 0xff, 0xff},
//	    color.RGBA{0x00, 0xcc, 0x00, 0xff},
//	    color.RGBA{0x00, 0xcc, 0x33, 0xff},
//	    color.RGBA{0x00, 0xcc, 0x66, 0xff},
//	    color.RGBA{0x00, 0xcc, 0x99, 0xff},
//	    color.RGBA{0x00, 0xcc, 0xcc, 0xff},
//	    color.RGBA{0x00, 0xcc, 0xff, 0xff},
//	    color.RGBA{0x00, 0xff, 0x00, 0xff},
//	    color.RGBA{0x00, 0xff, 0x33, 0xff},
//	    color.RGBA{0x00, 0xff, 0x66, 0xff},
//	    color.RGBA{0x00, 0xff, 0x99, 0xff},
//	    color.RGBA{0x00, 0xff, 0xcc, 0xff},
//	    color.RGBA{0x00, 0xff, 0xff, 0xff},
//	    color.RGBA{0x33, 0x00, 0x00, 0xff},
//	    color.RGBA{0x33, 0x00, 0x33, 0xff},
//	    color.RGBA{0x33, 0x00, 0x66, 0xff},
//	    color.RGBA{0x33, 0x00, 0x99, 0xff},
//	    color.RGBA{0x33, 0x00, 0xcc, 0xff},
//	    color.RGBA{0x33, 0x00, 0xff, 0xff},
//	    color.RGBA{0x33, 0x33, 0x00, 0xff},
//	    color.RGBA{0x33, 0x33, 0x33, 0xff},
//	    color.RGBA{0x33, 0x33, 0x66, 0xff},
//	    color.RGBA{0x33, 0x33, 0x99, 0xff},
//	    color.RGBA{0x33, 0x33, 0xcc, 0xff},
//	    color.RGBA{0x33, 0x33, 0xff, 0xff},
//	    color.RGBA{0x33, 0x66, 0x00, 0xff},
//	    color.RGBA{0x33, 0x66, 0x33, 0xff},
//	    color.RGBA{0x33, 0x66, 0x66, 0xff},
//	    color.RGBA{0x33, 0x66, 0x99, 0xff},
//	    color.RGBA{0x33, 0x66, 0xcc, 0xff},
//	    color.RGBA{0x33, 0x66, 0xff, 0xff},
//	    color.RGBA{0x33, 0x99, 0x00, 0xff},
//	    color.RGBA{0x33, 0x99, 0x33, 0xff},
//	    color.RGBA{0x33, 0x99, 0x66, 0xff},
//	    color.RGBA{0x33, 0x99, 0x99, 0xff},
//	    color.RGBA{0x33, 0x99, 0xcc, 0xff},
//	    color.RGBA{0x33, 0x99, 0xff, 0xff},
//	    color.RGBA{0x33, 0xcc, 0x00, 0xff},
//	    color.RGBA{0x33, 0xcc, 0x33, 0xff},
//	    color.RGBA{0x33, 0xcc, 0x66, 0xff},
//	    color.RGBA{0x33, 0xcc, 0x99, 0xff},
//	    color.RGBA{0x33, 0xcc, 0xcc, 0xff},
//	    color.RGBA{0x33, 0xcc, 0xff, 0xff},
//	    color.RGBA{0x33, 0xff, 0x00, 0xff},
//	    color.RGBA{0x33, 0xff, 0x33, 0xff},
//	    color.RGBA{0x33, 0xff, 0x66, 0xff},
//	    color.RGBA{0x33, 0xff, 0x99, 0xff},
//	    color.RGBA{0x33, 0xff, 0xcc, 0xff},
//	    color.RGBA{0x33, 0xff, 0xff, 0xff},
//	    color.RGBA{0x66, 0x00, 0x00, 0xff},
//	    color.RGBA{0x66, 0x00, 0x33, 0xff},
//	    color.RGBA{0x66, 0x00, 0x66, 0xff},
//	    color.RGBA{0x66, 0x00, 0x99, 0xff},
//	    color.RGBA{0x66, 0x00, 0xcc, 0xff},
//	    color.RGBA{0x66, 0x00, 0xff, 0xff},
//	    color.RGBA{0x66, 0x33, 0x00, 0xff},
//	    color.RGBA{0x66, 0x33, 0x33, 0xff},
//	    color.RGBA{0x66, 0x33, 0x66, 0xff},
//	    color.RGBA{0x66, 0x33, 0x99, 0xff},
//	    color.RGBA{0x66, 0x33, 0xcc, 0xff},
//	    color.RGBA{0x66, 0x33, 0xff, 0xff},
//	    color.RGBA{0x66, 0x66, 0x00, 0xff},
//	    color.RGBA{0x66, 0x66, 0x33, 0xff},
//	    color.RGBA{0x66, 0x66, 0x66, 0xff},
//	    color.RGBA{0x66, 0x66, 0x99, 0xff},
//	    color.RGBA{0x66, 0x66, 0xcc, 0xff},
//	    color.RGBA{0x66, 0x66, 0xff, 0xff},
//	    color.RGBA{0x66, 0x99, 0x00, 0xff},
//	    color.RGBA{0x66, 0x99, 0x33, 0xff},
//	    color.RGBA{0x66, 0x99, 0x66, 0xff},
//	    color.RGBA{0x66, 0x99, 0x99, 0xff},
//	    color.RGBA{0x66, 0x99, 0xcc, 0xff},
//	    color.RGBA{0x66, 0x99, 0xff, 0xff},
//	    color.RGBA{0x66, 0xcc, 0x00, 0xff},
//	    color.RGBA{0x66, 0xcc, 0x33, 0xff},
//	    color.RGBA{0x66, 0xcc, 0x66, 0xff},
//	    color.RGBA{0x66, 0xcc, 0x99, 0xff},
//	    color.RGBA{0x66, 0xcc, 0xcc, 0xff},
//	    color.RGBA{0x66, 0xcc, 0xff, 0xff},
//	    color.RGBA{0x66, 0xff, 0x00, 0xff},
//	    color.RGBA{0x66, 0xff, 0x33, 0xff},
//	    color.RGBA{0x66, 0xff, 0x66, 0xff},
//	    color.RGBA{0x66, 0xff, 0x99, 0xff},
//	    color.RGBA{0x66, 0xff, 0xcc, 0xff},
//	    color.RGBA{0x66, 0xff, 0xff, 0xff},
//	    color.RGBA{0x99, 0x00, 0x00, 0xff},
//	    color.RGBA{0x99, 0x00, 0x33, 0xff},
//	    color.RGBA{0x99, 0x00, 0x66, 0xff},
//	    color.RGBA{0x99, 0x00, 0x99, 0xff},
//	    color.RGBA{0x99, 0x00, 0xcc, 0xff},
//	    color.RGBA{0x99, 0x00, 0xff, 0xff},
//	    color.RGBA{0x99, 0x33, 0x00, 0xff},
//	    color.RGBA{0x99, 0x33, 0x33, 0xff},
//	    color.RGBA{0x99, 0x33, 0x66, 0xff},
//	    color.RGBA{0x99, 0x33, 0x99, 0xff},
//	    color.RGBA{0x99, 0x33, 0xcc, 0xff},
//	    color.RGBA{0x99, 0x33, 0xff, 0xff},
//	    color.RGBA{0x99, 0x66, 0x00, 0xff},
//	    color.RGBA{0x99, 0x66, 0x33, 0xff},
//	    color.RGBA{0x99, 0x66, 0x66, 0xff},
//	    color.RGBA{0x99, 0x66, 0x99, 0xff},
//	    color.RGBA{0x99, 0x66, 0xcc, 0xff},
//	    color.RGBA{0x99, 0x66, 0xff, 0xff},
//	    color.RGBA{0x99, 0x99, 0x00, 0xff},
//	    color.RGBA{0x99, 0x99, 0x33, 0xff},
//	    color.RGBA{0x99, 0x99, 0x66, 0xff},
//	    color.RGBA{0x99, 0x99, 0x99, 0xff},
//	    color.RGBA{0x99, 0x99, 0xcc, 0xff},
//	    color.RGBA{0x99, 0x99, 0xff, 0xff},
//	    color.RGBA{0x99, 0xcc, 0x00, 0xff},
//	    color.RGBA{0x99, 0xcc, 0x33, 0xff},
//	    color.RGBA{0x99, 0xcc, 0x66, 0xff},
//	    color.RGBA{0x99, 0xcc, 0x99, 0xff},
//	    color.RGBA{0x99, 0xcc, 0xcc, 0xff},
//	    color.RGBA{0x99, 0xcc, 0xff, 0xff},
//	    color.RGBA{0x99, 0xff, 0x00, 0xff},
//	    color.RGBA{0x99, 0xff, 0x33, 0xff},
//	    color.RGBA{0x99, 0xff, 0x66, 0xff},
//	    color.RGBA{0x99, 0xff, 0x99, 0xff},
//	    color.RGBA{0x99, 0xff, 0xcc, 0xff},
//	    color.RGBA{0x99, 0xff, 0xff, 0xff},
//	    color.RGBA{0xcc, 0x00, 0x00, 0xff},
//	    color.RGBA{0xcc, 0x00, 0x33, 0xff},
//	    color.RGBA{0xcc, 0x00, 0x66, 0xff},
//	    color.RGBA{0xcc, 0x00, 0x99, 0xff},
//	    color.RGBA{0xcc, 0x00, 0xcc, 0xff},
//	    color.RGBA{0xcc, 0x00, 0xff, 0xff},
//	    color.RGBA{0xcc, 0x33, 0x00, 0xff},
//	    color.RGBA{0xcc, 0x33, 0x33, 0xff},
//	    color.RGBA{0xcc, 0x33, 0x66, 0xff},
//	    color.RGBA{0xcc, 0x33, 0x99, 0xff},
//	    color.RGBA{0xcc, 0x33, 0xcc, 0xff},
//	    color.RGBA{0xcc, 0x33, 0xff, 0xff},
//	    color.RGBA{0xcc, 0x66, 0x00, 0xff},
//	    color.RGBA{0xcc, 0x66, 0x33, 0xff},
//	    color.RGBA{0xcc, 0x66, 0x66, 0xff},
//	    color.RGBA{0xcc, 0x66, 0x99, 0xff},
//	    color.RGBA{0xcc, 0x66, 0xcc, 0xff},
//	    color.RGBA{0xcc, 0x66, 0xff, 0xff},
//	    color.RGBA{0xcc, 0x99, 0x00, 0xff},
//	    color.RGBA{0xcc, 0x99, 0x33, 0xff},
//	    color.RGBA{0xcc, 0x99, 0x66, 0xff},
//	    color.RGBA{0xcc, 0x99, 0x99, 0xff},
//	    color.RGBA{0xcc, 0x99, 0xcc, 0xff},
//	    color.RGBA{0xcc, 0x99, 0xff, 0xff},
//	    color.RGBA{0xcc, 0xcc, 0x00, 0xff},
//	    color.RGBA{0xcc, 0xcc, 0x33, 0xff},
//	    color.RGBA{0xcc, 0xcc, 0x66, 0xff},
//	    color.RGBA{0xcc, 0xcc, 0x99, 0xff},
//	    color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
//	    color.RGBA{0xcc, 0xcc, 0xff, 0xff},
//	    color.RGBA{0xcc, 0xff, 0x00, 0xff},
//	    color.RGBA{0xcc, 0xff, 0x33, 0xff},
//	    color.RGBA{0xcc, 0xff, 0x66, 0xff},
//	    color.RGBA{0xcc, 0xff, 0x99, 0xff},
//	    color.RGBA{0xcc, 0xff, 0xcc, 0xff},
//	    color.RGBA{0xcc, 0xff, 0xff, 0xff},
//	    color.RGBA{0xff, 0x00, 0x00, 0xff},
//	    color.RGBA{0xff, 0x00, 0x33, 0xff},
//	    color.RGBA{0xff, 0x00, 0x66, 0xff},
//	    color.RGBA{0xff, 0x00, 0x99, 0xff},
//	    color.RGBA{0xff, 0x00, 0xcc, 0xff},
//	    color.RGBA{0xff, 0x00, 0xff, 0xff},
//	    color.RGBA{0xff, 0x33, 0x00, 0xff},
//	    color.RGBA{0xff, 0x33, 0x33, 0xff},
//	    color.RGBA{0xff, 0x33, 0x66, 0xff},
//	    color.RGBA{0xff, 0x33, 0x99, 0xff},
//	    color.RGBA{0xff, 0x33, 0xcc, 0xff},
//	    color.RGBA{0xff, 0x33, 0xff, 0xff},
//	    color.RGBA{0xff, 0x66, 0x00, 0xff},
//	    color.RGBA{0xff, 0x66, 0x33, 0xff},
//	    color.RGBA{0xff, 0x66, 0x66, 0xff},
//	    color.RGBA{0xff, 0x66, 0x99, 0xff},
//	    color.RGBA{0xff, 0x66, 0xcc, 0xff},
//	    color.RGBA{0xff, 0x66, 0xff, 0xff},
//	    color.RGBA{0xff, 0x99, 0x00, 0xff},
//	    color.RGBA{0xff, 0x99, 0x33, 0xff},
//	    color.RGBA{0xff, 0x99, 0x66, 0xff},
//	    color.RGBA{0xff, 0x99, 0x99, 0xff},
//	    color.RGBA{0xff, 0x99, 0xcc, 0xff},
//	    color.RGBA{0xff, 0x99, 0xff, 0xff},
//	    color.RGBA{0xff, 0xcc, 0x00, 0xff},
//	    color.RGBA{0xff, 0xcc, 0x33, 0xff},
//	    color.RGBA{0xff, 0xcc, 0x66, 0xff},
//	    color.RGBA{0xff, 0xcc, 0x99, 0xff},
//	    color.RGBA{0xff, 0xcc, 0xcc, 0xff},
//	    color.RGBA{0xff, 0xcc, 0xff, 0xff},
//	    color.RGBA{0xff, 0xff, 0x00, 0xff},
//	    color.RGBA{0xff, 0xff, 0x33, 0xff},
//	    color.RGBA{0xff, 0xff, 0x66, 0xff},
//	    color.RGBA{0xff, 0xff, 0x99, 0xff},
//	    color.RGBA{0xff, 0xff, 0xcc, 0xff},
//	    color.RGBA{0xff, 0xff, 0xff, 0xff},
//	}
//
// WebSafe是一个216色的调色板，被早期版本的Netscape
// Navigator（一种浏览器）广泛使用，它也被称为Netcape色彩立方。细节参见http://en.wikipedia.org/wiki/Web_colors#Web-safe_colors
var Plan9 = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0x44, 0xff},
	color.RGBA{0x00, 0x00, 0x88, 0xff},
	color.RGBA{0x00, 0x00, 0xcc, 0xff},
	color.RGBA{0x00, 0x44, 0x00, 0xff},
	color.RGBA{0x00, 0x44, 0x44, 0xff},
	color.RGBA{0x00, 0x44, 0x88, 0xff},
	color.RGBA{0x00, 0x44, 0xcc, 0xff},
	color.RGBA{0x00, 0x88, 0x00, 0xff},
	color.RGBA{0x00, 0x88, 0x44, 0xff},
	color.RGBA{0x00, 0x88, 0x88, 0xff},
	color.RGBA{0x00, 0x88, 0xcc, 0xff},
	color.RGBA{0x00, 0xcc, 0x00, 0xff},
	color.RGBA{0x00, 0xcc, 0x44, 0xff},
	color.RGBA{0x00, 0xcc, 0x88, 0xff},
	color.RGBA{0x00, 0xcc, 0xcc, 0xff},
	color.RGBA{0x00, 0xdd, 0xdd, 0xff},
	color.RGBA{0x11, 0x11, 0x11, 0xff},
	color.RGBA{0x00, 0x00, 0x55, 0xff},
	color.RGBA{0x00, 0x00, 0x99, 0xff},
	color.RGBA{0x00, 0x00, 0xdd, 0xff},
	color.RGBA{0x00, 0x55, 0x00, 0xff},
	color.RGBA{0x00, 0x55, 0x55, 0xff},
	color.RGBA{0x00, 0x4c, 0x99, 0xff},
	color.RGBA{0x00, 0x49, 0xdd, 0xff},
	color.RGBA{0x00, 0x99, 0x00, 0xff},
	color.RGBA{0x00, 0x99, 0x4c, 0xff},
	color.RGBA{0x00, 0x99, 0x99, 0xff},
	color.RGBA{0x00, 0x93, 0xdd, 0xff},
	color.RGBA{0x00, 0xdd, 0x00, 0xff},
	color.RGBA{0x00, 0xdd, 0x49, 0xff},
	color.RGBA{0x00, 0xdd, 0x93, 0xff},
	color.RGBA{0x00, 0xee, 0x9e, 0xff},
	color.RGBA{0x00, 0xee, 0xee, 0xff},
	color.RGBA{0x22, 0x22, 0x22, 0xff},
	color.RGBA{0x00, 0x00, 0x66, 0xff},
	color.RGBA{0x00, 0x00, 0xaa, 0xff},
	color.RGBA{0x00, 0x00, 0xee, 0xff},
	color.RGBA{0x00, 0x66, 0x00, 0xff},
	color.RGBA{0x00, 0x66, 0x66, 0xff},
	color.RGBA{0x00, 0x55, 0xaa, 0xff},
	color.RGBA{0x00, 0x4f, 0xee, 0xff},
	color.RGBA{0x00, 0xaa, 0x00, 0xff},
	color.RGBA{0x00, 0xaa, 0x55, 0xff},
	color.RGBA{0x00, 0xaa, 0xaa, 0xff},
	color.RGBA{0x00, 0x9e, 0xee, 0xff},
	color.RGBA{0x00, 0xee, 0x00, 0xff},
	color.RGBA{0x00, 0xee, 0x4f, 0xff},
	color.RGBA{0x00, 0xff, 0x55, 0xff},
	color.RGBA{0x00, 0xff, 0xaa, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0x33, 0x33, 0x33, 0xff},
	color.RGBA{0x00, 0x00, 0x77, 0xff},
	color.RGBA{0x00, 0x00, 0xbb, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0x77, 0x00, 0xff},
	color.RGBA{0x00, 0x77, 0x77, 0xff},
	color.RGBA{0x00, 0x5d, 0xbb, 0xff},
	color.RGBA{0x00, 0x55, 0xff, 0xff},
	color.RGBA{0x00, 0xbb, 0x00, 0xff},
	color.RGBA{0x00, 0xbb, 0x5d, 0xff},
	color.RGBA{0x00, 0xbb, 0xbb, 0xff},
	color.RGBA{0x00, 0xaa, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x44, 0x00, 0x44, 0xff},
	color.RGBA{0x44, 0x00, 0x88, 0xff},
	color.RGBA{0x44, 0x00, 0xcc, 0xff},
	color.RGBA{0x44, 0x44, 0x00, 0xff},
	color.RGBA{0x44, 0x44, 0x44, 0xff},
	color.RGBA{0x44, 0x44, 0x88, 0xff},
	color.RGBA{0x44, 0x44, 0xcc, 0xff},
	color.RGBA{0x44, 0x88, 0x00, 0xff},
	color.RGBA{0x44, 0x88, 0x44, 0xff},
	color.RGBA{0x44, 0x88, 0x88, 0xff},
	color.RGBA{0x44, 0x88, 0xcc, 0xff},
	color.RGBA{0x44, 0xcc, 0x00, 0xff},
	color.RGBA{0x44, 0xcc, 0x44, 0xff},
	color.RGBA{0x44, 0xcc, 0x88, 0xff},
	color.RGBA{0x44, 0xcc, 0xcc, 0xff},
	color.RGBA{0x44, 0x00, 0x00, 0xff},
	color.RGBA{0x55, 0x00, 0x00, 0xff},
	color.RGBA{0x55, 0x00, 0x55, 0xff},
	color.RGBA{0x4c, 0x00, 0x99, 0xff},
	color.RGBA{0x49, 0x00, 0xdd, 0xff},
	color.RGBA{0x55, 0x55, 0x00, 0xff},
	color.RGBA{0x55, 0x55, 0x55, 0xff},
	color.RGBA{0x4c, 0x4c, 0x99, 0xff},
	color.RGBA{0x49, 0x49, 0xdd, 0xff},
	color.RGBA{0x4c, 0x99, 0x00, 0xff},
	color.RGBA{0x4c, 0x99, 0x4c, 0xff},
	color.RGBA{0x4c, 0x99, 0x99, 0xff},
	color.RGBA{0x49, 0x93, 0xdd, 0xff},
	color.RGBA{0x49, 0xdd, 0x00, 0xff},
	color.RGBA{0x49, 0xdd, 0x49, 0xff},
	color.RGBA{0x49, 0xdd, 0x93, 0xff},
	color.RGBA{0x49, 0xdd, 0xdd, 0xff},
	color.RGBA{0x4f, 0xee, 0xee, 0xff},
	color.RGBA{0x66, 0x00, 0x00, 0xff},
	color.RGBA{0x66, 0x00, 0x66, 0xff},
	color.RGBA{0x55, 0x00, 0xaa, 0xff},
	color.RGBA{0x4f, 0x00, 0xee, 0xff},
	color.RGBA{0x66, 0x66, 0x00, 0xff},
	color.RGBA{0x66, 0x66, 0x66, 0xff},
	color.RGBA{0x55, 0x55, 0xaa, 0xff},
	color.RGBA{0x4f, 0x4f, 0xee, 0xff},
	color.RGBA{0x55, 0xaa, 0x00, 0xff},
	color.RGBA{0x55, 0xaa, 0x55, 0xff},
	color.RGBA{0x55, 0xaa, 0xaa, 0xff},
	color.RGBA{0x4f, 0x9e, 0xee, 0xff},
	color.RGBA{0x4f, 0xee, 0x00, 0xff},
	color.RGBA{0x4f, 0xee, 0x4f, 0xff},
	color.RGBA{0x4f, 0xee, 0x9e, 0xff},
	color.RGBA{0x55, 0xff, 0xaa, 0xff},
	color.RGBA{0x55, 0xff, 0xff, 0xff},
	color.RGBA{0x77, 0x00, 0x00, 0xff},
	color.RGBA{0x77, 0x00, 0x77, 0xff},
	color.RGBA{0x5d, 0x00, 0xbb, 0xff},
	color.RGBA{0x55, 0x00, 0xff, 0xff},
	color.RGBA{0x77, 0x77, 0x00, 0xff},
	color.RGBA{0x77, 0x77, 0x77, 0xff},
	color.RGBA{0x5d, 0x5d, 0xbb, 0xff},
	color.RGBA{0x55, 0x55, 0xff, 0xff},
	color.RGBA{0x5d, 0xbb, 0x00, 0xff},
	color.RGBA{0x5d, 0xbb, 0x5d, 0xff},
	color.RGBA{0x5d, 0xbb, 0xbb, 0xff},
	color.RGBA{0x55, 0xaa, 0xff, 0xff},
	color.RGBA{0x55, 0xff, 0x00, 0xff},
	color.RGBA{0x55, 0xff, 0x55, 0xff},
	color.RGBA{0x88, 0x00, 0x88, 0xff},
	color.RGBA{0x88, 0x00, 0xcc, 0xff},
	color.RGBA{0x88, 0x44, 0x00, 0xff},
	color.RGBA{0x88, 0x44, 0x44, 0xff},
	color.RGBA{0x88, 0x44, 0x88, 0xff},
	color.RGBA{0x88, 0x44, 0xcc, 0xff},
	color.RGBA{0x88, 0x88, 0x00, 0xff},
	color.RGBA{0x88, 0x88, 0x44, 0xff},
	color.RGBA{0x88, 0x88, 0x88, 0xff},
	color.RGBA{0x88, 0x88, 0xcc, 0xff},
	color.RGBA{0x88, 0xcc, 0x00, 0xff},
	color.RGBA{0x88, 0xcc, 0x44, 0xff},
	color.RGBA{0x88, 0xcc, 0x88, 0xff},
	color.RGBA{0x88, 0xcc, 0xcc, 0xff},
	color.RGBA{0x88, 0x00, 0x00, 0xff},
	color.RGBA{0x88, 0x00, 0x44, 0xff},
	color.RGBA{0x99, 0x00, 0x4c, 0xff},
	color.RGBA{0x99, 0x00, 0x99, 0xff},
	color.RGBA{0x93, 0x00, 0xdd, 0xff},
	color.RGBA{0x99, 0x4c, 0x00, 0xff},
	color.RGBA{0x99, 0x4c, 0x4c, 0xff},
	color.RGBA{0x99, 0x4c, 0x99, 0xff},
	color.RGBA{0x93, 0x49, 0xdd, 0xff},
	color.RGBA{0x99, 0x99, 0x00, 0xff},
	color.RGBA{0x99, 0x99, 0x4c, 0xff},
	color.RGBA{0x99, 0x99, 0x99, 0xff},
	color.RGBA{0x93, 0x93, 0xdd, 0xff},
	color.RGBA{0x93, 0xdd, 0x00, 0xff},
	color.RGBA{0x93, 0xdd, 0x49, 0xff},
	color.RGBA{0x93, 0xdd, 0x93, 0xff},
	color.RGBA{0x93, 0xdd, 0xdd, 0xff},
	color.RGBA{0x99, 0x00, 0x00, 0xff},
	color.RGBA{0xaa, 0x00, 0x00, 0xff},
	color.RGBA{0xaa, 0x00, 0x55, 0xff},
	color.RGBA{0xaa, 0x00, 0xaa, 0xff},
	color.RGBA{0x9e, 0x00, 0xee, 0xff},
	color.RGBA{0xaa, 0x55, 0x00, 0xff},
	color.RGBA{0xaa, 0x55, 0x55, 0xff},
	color.RGBA{0xaa, 0x55, 0xaa, 0xff},
	color.RGBA{0x9e, 0x4f, 0xee, 0xff},
	color.RGBA{0xaa, 0xaa, 0x00, 0xff},
	color.RGBA{0xaa, 0xaa, 0x55, 0xff},
	color.RGBA{0xaa, 0xaa, 0xaa, 0xff},
	color.RGBA{0x9e, 0x9e, 0xee, 0xff},
	color.RGBA{0x9e, 0xee, 0x00, 0xff},
	color.RGBA{0x9e, 0xee, 0x4f, 0xff},
	color.RGBA{0x9e, 0xee, 0x9e, 0xff},
	color.RGBA{0x9e, 0xee, 0xee, 0xff},
	color.RGBA{0xaa, 0xff, 0xff, 0xff},
	color.RGBA{0xbb, 0x00, 0x00, 0xff},
	color.RGBA{0xbb, 0x00, 0x5d, 0xff},
	color.RGBA{0xbb, 0x00, 0xbb, 0xff},
	color.RGBA{0xaa, 0x00, 0xff, 0xff},
	color.RGBA{0xbb, 0x5d, 0x00, 0xff},
	color.RGBA{0xbb, 0x5d, 0x5d, 0xff},
	color.RGBA{0xbb, 0x5d, 0xbb, 0xff},
	color.RGBA{0xaa, 0x55, 0xff, 0xff},
	color.RGBA{0xbb, 0xbb, 0x00, 0xff},
	color.RGBA{0xbb, 0xbb, 0x5d, 0xff},
	color.RGBA{0xbb, 0xbb, 0xbb, 0xff},
	color.RGBA{0xaa, 0xaa, 0xff, 0xff},
	color.RGBA{0xaa, 0xff, 0x00, 0xff},
	color.RGBA{0xaa, 0xff, 0x55, 0xff},
	color.RGBA{0xaa, 0xff, 0xaa, 0xff},
	color.RGBA{0xcc, 0x00, 0xcc, 0xff},
	color.RGBA{0xcc, 0x44, 0x00, 0xff},
	color.RGBA{0xcc, 0x44, 0x44, 0xff},
	color.RGBA{0xcc, 0x44, 0x88, 0xff},
	color.RGBA{0xcc, 0x44, 0xcc, 0xff},
	color.RGBA{0xcc, 0x88, 0x00, 0xff},
	color.RGBA{0xcc, 0x88, 0x44, 0xff},
	color.RGBA{0xcc, 0x88, 0x88, 0xff},
	color.RGBA{0xcc, 0x88, 0xcc, 0xff},
	color.RGBA{0xcc, 0xcc, 0x00, 0xff},
	color.RGBA{0xcc, 0xcc, 0x44, 0xff},
	color.RGBA{0xcc, 0xcc, 0x88, 0xff},
	color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	color.RGBA{0xcc, 0x00, 0x00, 0xff},
	color.RGBA{0xcc, 0x00, 0x44, 0xff},
	color.RGBA{0xcc, 0x00, 0x88, 0xff},
	color.RGBA{0xdd, 0x00, 0x93, 0xff},
	color.RGBA{0xdd, 0x00, 0xdd, 0xff},
	color.RGBA{0xdd, 0x49, 0x00, 0xff},
	color.RGBA{0xdd, 0x49, 0x49, 0xff},
	color.RGBA{0xdd, 0x49, 0x93, 0xff},
	color.RGBA{0xdd, 0x49, 0xdd, 0xff},
	color.RGBA{0xdd, 0x93, 0x00, 0xff},
	color.RGBA{0xdd, 0x93, 0x49, 0xff},
	color.RGBA{0xdd, 0x93, 0x93, 0xff},
	color.RGBA{0xdd, 0x93, 0xdd, 0xff},
	color.RGBA{0xdd, 0xdd, 0x00, 0xff},
	color.RGBA{0xdd, 0xdd, 0x49, 0xff},
	color.RGBA{0xdd, 0xdd, 0x93, 0xff},
	color.RGBA{0xdd, 0xdd, 0xdd, 0xff},
	color.RGBA{0xdd, 0x00, 0x00, 0xff},
	color.RGBA{0xdd, 0x00, 0x49, 0xff},
	color.RGBA{0xee, 0x00, 0x4f, 0xff},
	color.RGBA{0xee, 0x00, 0x9e, 0xff},
	color.RGBA{0xee, 0x00, 0xee, 0xff},
	color.RGBA{0xee, 0x4f, 0x00, 0xff},
	color.RGBA{0xee, 0x4f, 0x4f, 0xff},
	color.RGBA{0xee, 0x4f, 0x9e, 0xff},
	color.RGBA{0xee, 0x4f, 0xee, 0xff},
	color.RGBA{0xee, 0x9e, 0x00, 0xff},
	color.RGBA{0xee, 0x9e, 0x4f, 0xff},
	color.RGBA{0xee, 0x9e, 0x9e, 0xff},
	color.RGBA{0xee, 0x9e, 0xee, 0xff},
	color.RGBA{0xee, 0xee, 0x00, 0xff},
	color.RGBA{0xee, 0xee, 0x4f, 0xff},
	color.RGBA{0xee, 0xee, 0x9e, 0xff},
	color.RGBA{0xee, 0xee, 0xee, 0xff},
	color.RGBA{0xee, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x55, 0xff},
	color.RGBA{0xff, 0x00, 0xaa, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x55, 0x00, 0xff},
	color.RGBA{0xff, 0x55, 0x55, 0xff},
	color.RGBA{0xff, 0x55, 0xaa, 0xff},
	color.RGBA{0xff, 0x55, 0xff, 0xff},
	color.RGBA{0xff, 0xaa, 0x00, 0xff},
	color.RGBA{0xff, 0xaa, 0x55, 0xff},
	color.RGBA{0xff, 0xaa, 0xaa, 0xff},
	color.RGBA{0xff, 0xaa, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x55, 0xff},
	color.RGBA{0xff, 0xff, 0xaa, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

// WebSafe is a 216-color palette that was popularized by early versions of
// Netscape Navigator. It is also known as the Netscape Color Cube.
//
// See http://en.wikipedia.org/wiki/Web_colors#Web-safe_colors for details.
var WebSafe = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0x00, 0x33, 0xff},
	color.RGBA{0x00, 0x00, 0x66, 0xff},
	color.RGBA{0x00, 0x00, 0x99, 0xff},
	color.RGBA{0x00, 0x00, 0xcc, 0xff},
	color.RGBA{0x00, 0x00, 0xff, 0xff},
	color.RGBA{0x00, 0x33, 0x00, 0xff},
	color.RGBA{0x00, 0x33, 0x33, 0xff},
	color.RGBA{0x00, 0x33, 0x66, 0xff},
	color.RGBA{0x00, 0x33, 0x99, 0xff},
	color.RGBA{0x00, 0x33, 0xcc, 0xff},
	color.RGBA{0x00, 0x33, 0xff, 0xff},
	color.RGBA{0x00, 0x66, 0x00, 0xff},
	color.RGBA{0x00, 0x66, 0x33, 0xff},
	color.RGBA{0x00, 0x66, 0x66, 0xff},
	color.RGBA{0x00, 0x66, 0x99, 0xff},
	color.RGBA{0x00, 0x66, 0xcc, 0xff},
	color.RGBA{0x00, 0x66, 0xff, 0xff},
	color.RGBA{0x00, 0x99, 0x00, 0xff},
	color.RGBA{0x00, 0x99, 0x33, 0xff},
	color.RGBA{0x00, 0x99, 0x66, 0xff},
	color.RGBA{0x00, 0x99, 0x99, 0xff},
	color.RGBA{0x00, 0x99, 0xcc, 0xff},
	color.RGBA{0x00, 0x99, 0xff, 0xff},
	color.RGBA{0x00, 0xcc, 0x00, 0xff},
	color.RGBA{0x00, 0xcc, 0x33, 0xff},
	color.RGBA{0x00, 0xcc, 0x66, 0xff},
	color.RGBA{0x00, 0xcc, 0x99, 0xff},
	color.RGBA{0x00, 0xcc, 0xcc, 0xff},
	color.RGBA{0x00, 0xcc, 0xff, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x33, 0xff},
	color.RGBA{0x00, 0xff, 0x66, 0xff},
	color.RGBA{0x00, 0xff, 0x99, 0xff},
	color.RGBA{0x00, 0xff, 0xcc, 0xff},
	color.RGBA{0x00, 0xff, 0xff, 0xff},
	color.RGBA{0x33, 0x00, 0x00, 0xff},
	color.RGBA{0x33, 0x00, 0x33, 0xff},
	color.RGBA{0x33, 0x00, 0x66, 0xff},
	color.RGBA{0x33, 0x00, 0x99, 0xff},
	color.RGBA{0x33, 0x00, 0xcc, 0xff},
	color.RGBA{0x33, 0x00, 0xff, 0xff},
	color.RGBA{0x33, 0x33, 0x00, 0xff},
	color.RGBA{0x33, 0x33, 0x33, 0xff},
	color.RGBA{0x33, 0x33, 0x66, 0xff},
	color.RGBA{0x33, 0x33, 0x99, 0xff},
	color.RGBA{0x33, 0x33, 0xcc, 0xff},
	color.RGBA{0x33, 0x33, 0xff, 0xff},
	color.RGBA{0x33, 0x66, 0x00, 0xff},
	color.RGBA{0x33, 0x66, 0x33, 0xff},
	color.RGBA{0x33, 0x66, 0x66, 0xff},
	color.RGBA{0x33, 0x66, 0x99, 0xff},
	color.RGBA{0x33, 0x66, 0xcc, 0xff},
	color.RGBA{0x33, 0x66, 0xff, 0xff},
	color.RGBA{0x33, 0x99, 0x00, 0xff},
	color.RGBA{0x33, 0x99, 0x33, 0xff},
	color.RGBA{0x33, 0x99, 0x66, 0xff},
	color.RGBA{0x33, 0x99, 0x99, 0xff},
	color.RGBA{0x33, 0x99, 0xcc, 0xff},
	color.RGBA{0x33, 0x99, 0xff, 0xff},
	color.RGBA{0x33, 0xcc, 0x00, 0xff},
	color.RGBA{0x33, 0xcc, 0x33, 0xff},
	color.RGBA{0x33, 0xcc, 0x66, 0xff},
	color.RGBA{0x33, 0xcc, 0x99, 0xff},
	color.RGBA{0x33, 0xcc, 0xcc, 0xff},
	color.RGBA{0x33, 0xcc, 0xff, 0xff},
	color.RGBA{0x33, 0xff, 0x00, 0xff},
	color.RGBA{0x33, 0xff, 0x33, 0xff},
	color.RGBA{0x33, 0xff, 0x66, 0xff},
	color.RGBA{0x33, 0xff, 0x99, 0xff},
	color.RGBA{0x33, 0xff, 0xcc, 0xff},
	color.RGBA{0x33, 0xff, 0xff, 0xff},
	color.RGBA{0x66, 0x00, 0x00, 0xff},
	color.RGBA{0x66, 0x00, 0x33, 0xff},
	color.RGBA{0x66, 0x00, 0x66, 0xff},
	color.RGBA{0x66, 0x00, 0x99, 0xff},
	color.RGBA{0x66, 0x00, 0xcc, 0xff},
	color.RGBA{0x66, 0x00, 0xff, 0xff},
	color.RGBA{0x66, 0x33, 0x00, 0xff},
	color.RGBA{0x66, 0x33, 0x33, 0xff},
	color.RGBA{0x66, 0x33, 0x66, 0xff},
	color.RGBA{0x66, 0x33, 0x99, 0xff},
	color.RGBA{0x66, 0x33, 0xcc, 0xff},
	color.RGBA{0x66, 0x33, 0xff, 0xff},
	color.RGBA{0x66, 0x66, 0x00, 0xff},
	color.RGBA{0x66, 0x66, 0x33, 0xff},
	color.RGBA{0x66, 0x66, 0x66, 0xff},
	color.RGBA{0x66, 0x66, 0x99, 0xff},
	color.RGBA{0x66, 0x66, 0xcc, 0xff},
	color.RGBA{0x66, 0x66, 0xff, 0xff},
	color.RGBA{0x66, 0x99, 0x00, 0xff},
	color.RGBA{0x66, 0x99, 0x33, 0xff},
	color.RGBA{0x66, 0x99, 0x66, 0xff},
	color.RGBA{0x66, 0x99, 0x99, 0xff},
	color.RGBA{0x66, 0x99, 0xcc, 0xff},
	color.RGBA{0x66, 0x99, 0xff, 0xff},
	color.RGBA{0x66, 0xcc, 0x00, 0xff},
	color.RGBA{0x66, 0xcc, 0x33, 0xff},
	color.RGBA{0x66, 0xcc, 0x66, 0xff},
	color.RGBA{0x66, 0xcc, 0x99, 0xff},
	color.RGBA{0x66, 0xcc, 0xcc, 0xff},
	color.RGBA{0x66, 0xcc, 0xff, 0xff},
	color.RGBA{0x66, 0xff, 0x00, 0xff},
	color.RGBA{0x66, 0xff, 0x33, 0xff},
	color.RGBA{0x66, 0xff, 0x66, 0xff},
	color.RGBA{0x66, 0xff, 0x99, 0xff},
	color.RGBA{0x66, 0xff, 0xcc, 0xff},
	color.RGBA{0x66, 0xff, 0xff, 0xff},
	color.RGBA{0x99, 0x00, 0x00, 0xff},
	color.RGBA{0x99, 0x00, 0x33, 0xff},
	color.RGBA{0x99, 0x00, 0x66, 0xff},
	color.RGBA{0x99, 0x00, 0x99, 0xff},
	color.RGBA{0x99, 0x00, 0xcc, 0xff},
	color.RGBA{0x99, 0x00, 0xff, 0xff},
	color.RGBA{0x99, 0x33, 0x00, 0xff},
	color.RGBA{0x99, 0x33, 0x33, 0xff},
	color.RGBA{0x99, 0x33, 0x66, 0xff},
	color.RGBA{0x99, 0x33, 0x99, 0xff},
	color.RGBA{0x99, 0x33, 0xcc, 0xff},
	color.RGBA{0x99, 0x33, 0xff, 0xff},
	color.RGBA{0x99, 0x66, 0x00, 0xff},
	color.RGBA{0x99, 0x66, 0x33, 0xff},
	color.RGBA{0x99, 0x66, 0x66, 0xff},
	color.RGBA{0x99, 0x66, 0x99, 0xff},
	color.RGBA{0x99, 0x66, 0xcc, 0xff},
	color.RGBA{0x99, 0x66, 0xff, 0xff},
	color.RGBA{0x99, 0x99, 0x00, 0xff},
	color.RGBA{0x99, 0x99, 0x33, 0xff},
	color.RGBA{0x99, 0x99, 0x66, 0xff},
	color.RGBA{0x99, 0x99, 0x99, 0xff},
	color.RGBA{0x99, 0x99, 0xcc, 0xff},
	color.RGBA{0x99, 0x99, 0xff, 0xff},
	color.RGBA{0x99, 0xcc, 0x00, 0xff},
	color.RGBA{0x99, 0xcc, 0x33, 0xff},
	color.RGBA{0x99, 0xcc, 0x66, 0xff},
	color.RGBA{0x99, 0xcc, 0x99, 0xff},
	color.RGBA{0x99, 0xcc, 0xcc, 0xff},
	color.RGBA{0x99, 0xcc, 0xff, 0xff},
	color.RGBA{0x99, 0xff, 0x00, 0xff},
	color.RGBA{0x99, 0xff, 0x33, 0xff},
	color.RGBA{0x99, 0xff, 0x66, 0xff},
	color.RGBA{0x99, 0xff, 0x99, 0xff},
	color.RGBA{0x99, 0xff, 0xcc, 0xff},
	color.RGBA{0x99, 0xff, 0xff, 0xff},
	color.RGBA{0xcc, 0x00, 0x00, 0xff},
	color.RGBA{0xcc, 0x00, 0x33, 0xff},
	color.RGBA{0xcc, 0x00, 0x66, 0xff},
	color.RGBA{0xcc, 0x00, 0x99, 0xff},
	color.RGBA{0xcc, 0x00, 0xcc, 0xff},
	color.RGBA{0xcc, 0x00, 0xff, 0xff},
	color.RGBA{0xcc, 0x33, 0x00, 0xff},
	color.RGBA{0xcc, 0x33, 0x33, 0xff},
	color.RGBA{0xcc, 0x33, 0x66, 0xff},
	color.RGBA{0xcc, 0x33, 0x99, 0xff},
	color.RGBA{0xcc, 0x33, 0xcc, 0xff},
	color.RGBA{0xcc, 0x33, 0xff, 0xff},
	color.RGBA{0xcc, 0x66, 0x00, 0xff},
	color.RGBA{0xcc, 0x66, 0x33, 0xff},
	color.RGBA{0xcc, 0x66, 0x66, 0xff},
	color.RGBA{0xcc, 0x66, 0x99, 0xff},
	color.RGBA{0xcc, 0x66, 0xcc, 0xff},
	color.RGBA{0xcc, 0x66, 0xff, 0xff},
	color.RGBA{0xcc, 0x99, 0x00, 0xff},
	color.RGBA{0xcc, 0x99, 0x33, 0xff},
	color.RGBA{0xcc, 0x99, 0x66, 0xff},
	color.RGBA{0xcc, 0x99, 0x99, 0xff},
	color.RGBA{0xcc, 0x99, 0xcc, 0xff},
	color.RGBA{0xcc, 0x99, 0xff, 0xff},
	color.RGBA{0xcc, 0xcc, 0x00, 0xff},
	color.RGBA{0xcc, 0xcc, 0x33, 0xff},
	color.RGBA{0xcc, 0xcc, 0x66, 0xff},
	color.RGBA{0xcc, 0xcc, 0x99, 0xff},
	color.RGBA{0xcc, 0xcc, 0xcc, 0xff},
	color.RGBA{0xcc, 0xcc, 0xff, 0xff},
	color.RGBA{0xcc, 0xff, 0x00, 0xff},
	color.RGBA{0xcc, 0xff, 0x33, 0xff},
	color.RGBA{0xcc, 0xff, 0x66, 0xff},
	color.RGBA{0xcc, 0xff, 0x99, 0xff},
	color.RGBA{0xcc, 0xff, 0xcc, 0xff},
	color.RGBA{0xcc, 0xff, 0xff, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x33, 0xff},
	color.RGBA{0xff, 0x00, 0x66, 0xff},
	color.RGBA{0xff, 0x00, 0x99, 0xff},
	color.RGBA{0xff, 0x00, 0xcc, 0xff},
	color.RGBA{0xff, 0x00, 0xff, 0xff},
	color.RGBA{0xff, 0x33, 0x00, 0xff},
	color.RGBA{0xff, 0x33, 0x33, 0xff},
	color.RGBA{0xff, 0x33, 0x66, 0xff},
	color.RGBA{0xff, 0x33, 0x99, 0xff},
	color.RGBA{0xff, 0x33, 0xcc, 0xff},
	color.RGBA{0xff, 0x33, 0xff, 0xff},
	color.RGBA{0xff, 0x66, 0x00, 0xff},
	color.RGBA{0xff, 0x66, 0x33, 0xff},
	color.RGBA{0xff, 0x66, 0x66, 0xff},
	color.RGBA{0xff, 0x66, 0x99, 0xff},
	color.RGBA{0xff, 0x66, 0xcc, 0xff},
	color.RGBA{0xff, 0x66, 0xff, 0xff},
	color.RGBA{0xff, 0x99, 0x00, 0xff},
	color.RGBA{0xff, 0x99, 0x33, 0xff},
	color.RGBA{0xff, 0x99, 0x66, 0xff},
	color.RGBA{0xff, 0x99, 0x99, 0xff},
	color.RGBA{0xff, 0x99, 0xcc, 0xff},
	color.RGBA{0xff, 0x99, 0xff, 0xff},
	color.RGBA{0xff, 0xcc, 0x00, 0xff},
	color.RGBA{0xff, 0xcc, 0x33, 0xff},
	color.RGBA{0xff, 0xcc, 0x66, 0xff},
	color.RGBA{0xff, 0xcc, 0x99, 0xff},
	color.RGBA{0xff, 0xcc, 0xcc, 0xff},
	color.RGBA{0xff, 0xcc, 0xff, 0xff},
	color.RGBA{0xff, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0x33, 0xff},
	color.RGBA{0xff, 0xff, 0x66, 0xff},
	color.RGBA{0xff, 0xff, 0x99, 0xff},
	color.RGBA{0xff, 0xff, 0xcc, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

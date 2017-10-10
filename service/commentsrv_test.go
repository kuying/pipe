// Solo.go - A small and beautiful blogging platform written in golang.
// Copyright (C) 2017, b3log.org
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package service

import (
	"testing"

	"github.com/b3log/solo.go/model"
)

func TestConsoleGetComments(t *testing.T) {
	comments, pagination := Comment.ConsoleGetComments(1, 1)

	if 0 != len(comments) {
		t.Errorf("expected is [%d], actual is [%d]", 0, len(comments))

		return
	}
	if 0 != pagination.RecordCount {
		t.Errorf("expected is [%d], actual is [%d]", 0, pagination.RecordCount)
	}
}

func TestRemoveComment(t *testing.T) {
	comment := &model.Comment{
		ArticleID:       1,
		AuthorName:      "Daniel",
		AuthorAvatarURL: "https://img.hacpai.com/avatar/1353745196354_1500432853138.png?imageView2/1/w/80/h/80/interlace/0/q/100",
		Content:         "写博客需要坚持，相信积累后必然会有收获，我们一起努力加油 :smile:",
		BlogID:          1,
	}
	if err := Comment.AddComment(comment); nil != err {
		t.Errorf("add comment failed: " + err.Error())

		return
	}

	if err := Comment.RemoveComment(comment.ID); nil != err {
		t.Error(err)

		return
	}

	comments, _ := Comment.ConsoleGetComments(1, 1)
	if 0 != len(comments) {
		t.Error("remove comment failed")
	}
}
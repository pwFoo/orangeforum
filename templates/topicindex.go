// Copyright (c) 2017 Sagar Gubbi. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package templates

const topicindexSrc = `
{{ define "content" }}

<div class="btn-row">
	{{ if not .IsClosed }}
	<a class="link-btn" href="/comments/new?tid={{ .TopicID }}">Reply</a>
	{{ end }}
	{{ if or .IsAdmin .IsMod .IsSuperAdmin (and .IsOwner (not .IsClosed)) }}
	<a class="link-btn" href="/topics/edit?id={{ .TopicID }}">Edit topic</a>
	{{ end }}
	{{ if and .Common.UserName .Common.IsTopicSubAllowed }}
	{{ if .SubToken }}
	<form action="/topics/unsubscribe?token={{ .SubToken }}" method="POST">
		<input type="hidden" name="csrf" value="{{ .Common.CSRF }}">
		<input class="btn" type="submit" value="Unsubscribe">
	</form>
	{{ else }}
	<form action="/topics/subscribe?id={{ .TopicID }}" method="POST">
		<input type="hidden" name="csrf" value="{{ .Common.CSRF }}">
		<input class="btn" type="submit" value="Subscribe">
	</form>
	{{ end }}
	{{ end }}
</div>

<h2 id="title"><a href="/topics?id={{ .TopicID }}">{{ .TopicName }}{{ if .IsClosed }} [closed]{{ end }}</a></h2>
<div class="comment-title muted"><a href="/users?u={{ .OwnerName }}">{{ .OwnerName }}</a> in <a href="/groups?name={{ .GroupName }}">{{ .GroupName }}</a> {{ .CreatedDate }}</div>
<div class="comment-row">
	<div class="comment">
		<p>{{ .Content }}</p>
	</div>
</div>
<hr class="sep">

{{ if .Comments }}
{{ range .Comments }}
<div class="comment-row" id="comment-{{ .ID }}">
	<div class="comment-title muted"><a href="/users?u={{ .UserName }}">{{ .UserName }}</a> <a href="/comments?id={{ .ID }}">{{ .CreatedDate }}</a>{{ if or .IsOwner $.IsAdmin $.IsMod $.IsSuperAdmin }} | <a href="/comments/edit?id={{ .ID }}">edit</a> {{end}}</div>
	{{ if .IsDeleted }}
		<div class="comment">[DELETED]</div>
	{{ else }}
		<div class="comment">{{ .Content }}</div>
		{{ if .ImgSrc }}<div><img src="/img?name={{ .ImgSrc }}"></div>{{ end }}
	{{ end }}
</div>
<hr class="sep">
{{ end }}
{{ else }}
<div class="row">
	<div class="muted">No comments to show.</div>
</div>
{{ end }}
<div id="comment-last"></div>

{{ if .LastCommentDate }}
<div class="row">
	<div><a href="/topics?id={{ .TopicID }}&lcd={{ .LastCommentDate }}">More</a></div>
</div>
{{ end }}

{{ end }}`

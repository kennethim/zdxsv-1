<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="utf-8">
<meta http-equiv="Refresh" content="60">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
<link href="/assets/css/bootstrap.min.css" rel="stylesheet">
<!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
<!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
<!--[if lt IE 9]>
  <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
  <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
<![endif]-->
<style>
body { padding-top: 70px; }
</style>

<title>Gundam vs Z Gundam 私服狀況</title>
</head>

<body>
<nav class="navbar navbar-inverse navbar-fixed-top">
<div class="container">
<div class="navbar-header">
  <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar" aria-expanded="false" aria-controls="navbar">
	<span class="sr-only">Toggle navigation</span>
	<span class="icon-bar"></span>
	<span class="icon-bar"></span>
	<span class="icon-bar"></span>
  </button>
  <a class="navbar-brand" href="#">zdxsv</a>
</div>
</div>
</nav>

<div class="container">
<div class="starter-template">

{{if .Lives}}
{{range .Lives}}
<div class="panel panel-primary">
  <div class="panel-heading">
    <h3 class="panel-title"><span class="glyphicon glyphicon-facetime-video"></span> 直播中</h3>
  </div>
  <div class="panel-body">
    <div class="media">
      <div class="media-left">
        <img class="media-object" src="{{.ThumbUrl}}" alt="">
      </div>
      <div class="media-body">
        <h4 class="media-heading"><a target="_blank" href="{{.LiveUrl}}">{{.Title}}</a></h4>
        <p>{{.Description}}</p>
		<p>({{.CommunityName}})
		<a target="_blank" href="{{.LiveUrl}}" class="btn btn-primary btn-lg btn-danger pull-right" role="button">觀看</a></p>
      </div>
    </div>
  </div>
</div>
{{end}}
{{end}}

<h1>SERVER STATUS</h1>
<p> {{.NowDate}} 目前連接狀態  <span class="glyphicon glyphicon-info-sign"></span> 每60秒自動更新</p>
<h3>Lobby {{.LobbyUserCount}} 人</h3>
<table class="table table-inverse table-sm">
<thead>
<tr><th>ID</th><th>HN</th><th>部隊名</th><th>UDP</th></tr>
</thead>
<tbody>
{{range .LobbyUsers}}
<tr><td>【{{.UserId}}】</td><td>{{.Name}}</td><td>{{.Team}}</td><td>{{.UDP}}</td></tr>
{{end}}
</tbody>
</table>
<h3>對戰中 {{.BattleUserCount}} 人</h3>
<table class="table table-inverse table-sm">
<thead>
<tr><th>ID</th><th>HN</th><th>部隊名</th><th>UDP</th></tr>
</thead>
<tbody>
{{range .BattleUsers}}
<tr><td>【{{.UserId}}】</td><td>{{.Name}}</td><td>{{.Team}}</td><td>{{.UDP}}</td></tr>
{{end}}
</tbody>
</table>

<h2>聊天室</h2>
{{if .ChatInviteUrl}}
<a target="_blank" href="{{.ChatInviteUrl}}" class="btn btn-primary btn-sm btn-primary" role="button">參加</a>
{{end}}
{{if .ChatUrl}}
<a target="_blank" href="{{.ChatUrl}}" class="btn btn-primary btn-sm btn-success" role="button">打開</a>
{{end}}

<p>在線{{len .OnlineChatUsers}}人 離線{{len .OfflineChatUsers}}人</p>
<table class="table table-sm table-striped table-condensed table-inverse">
<tbody>
{{range .OnlineChatUsers}}
<tr>
	<td>
		<span class="glyphicon glyphicon-signal" style="color:green"></span>
		{{if .VoiceChat}}
			<span class="glyphicon glyphicon-volume-up" style="color:green"></span>
		{{else}}
			<span class="glyphicon glyphicon-volume-off" style="color:gray"></span>
		{{end}}
		{{if .Avatar}}
			<img class="img-rounded" width="30" height="30" src="https://cdn.discordapp.com/avatars/{{.ID}}/{{.Avatar}}.jpg">
		{{else}}
			<img class="img-rounded" width="30" height="30" src="/assets/discord.png">
		{{end}}
		{{.Username}}
	</td>
</tr>
{{end}}

{{range .OfflineChatUsers}}
<tr>
	<td>
		<span class="glyphicon glyphicon-signal" style="color:gray"></span>
		<span class="glyphicon glyphicon-volume-off" style="color:gray"></span>
		{{if .Avatar}}
			<img class="img-rounded" width="30" height="30" src="https://cdn.discordapp.com/avatars/{{.ID}}/{{.Avatar}}.jpg">
		{{else}}
			<img class="img-rounded" width="30" height="30" src="/assets/discord.png">
		{{end}}
		{{.Username}}
	</td>
</tr>
{{end}}
</tbody>
</table>


</div>
</div>

<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.12.4/jquery.min.js"></script>
<script src="/assets/js/bootstrap.min.js"></script>

</body>
</html>

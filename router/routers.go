package router

type routers struct {
	User     User
	Feed     feed
	Favorite favorite
	Relation relation
	Comment  comment
	Message  message
	Publish  publish
}

var Routers = new(routers)

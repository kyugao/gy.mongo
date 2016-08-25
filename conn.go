package mongo

import (
	"gopkg.in/mgo.v2"
)

func GetContext() *DialContext {
	return dialContext
}

func getSession() *Session {
	context := dialContext.Ref()
	return context
}

func returnSession(session *Session) {
	dialContext.UnRef(session)
}

func GetCollection(name string) *mgo.Collection {
	session := getSession()
	defer returnSession(session)
	return session.Session.DB(Config.DBName).C(name)
}

func GetGridFS(category string) *mgo.GridFS {
	session := getSession()
	defer returnSession(session)
	return session.DB(Config.DBName).GridFS(category)
}
package database

import "github.com/gocql/gocql"

var Cluster *gocql.ClusterConfig

func InitCluster() *gocql.ClusterConfig {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "test_axxonsoft"
	cluster.Consistency = gocql.Quorum
	cluster.Port = 9042
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: "test_axxonsoft", Password: "123456"}
	Cluster = cluster
	return cluster
}

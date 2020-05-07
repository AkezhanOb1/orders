package config

import "fmt"

//postgresAddress is the address of the postgres
const postgresAddress = "127.0.0.1"
//const postgresAddress = "46.101.138.224"

//postgresPort is the port of the postgres
const postgresPort = "5432"

//postgresDataBase is the name of the database
const postgresDataBase = "diploma"

//postgresUsername is the name of the user inside DBA
const postgresUsername = "postgres"

//postgresPassword is the password of the user
const postgresPassword = "postgres"

//PostgresConnection is the connection string to the database
var PostgresConnection = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	postgresAddress, postgresPort, postgresUsername, postgresPassword, postgresDataBase)

//RpcServerAddress is an address to the gRPC server
var RpcServerAddress = "server:50051"
//var RpcServerAddress = "localhost:50051"

//TokenServer is an address to the authorization
var TokenServer = "46.101.138.224:50002"
//rabbitUserName is the name of the user in Rabbit MQ
const rabbitUserName = "admin"
//rabbitPassword is the password of the user
const rabbitPassword = "admin"
//RabbitConnection is the connection string to the Rabbit MQ
var RabbitConnection = fmt.Sprintf("amqp://%s:%s@46.101.138.224:5672/", rabbitUserName, rabbitPassword)

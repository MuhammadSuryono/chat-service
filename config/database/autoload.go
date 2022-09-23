package database

func Init() {
	InitConnectionFromEnvironment().CreateNewConnection()
}

package main
import(

	"build_architecture/datasource"
	"build_architecture/web"
)

func main () {
	datasource.СonnectToDatabase()
	defer datasource.DataSource.Close()
	web.RunServer()


}

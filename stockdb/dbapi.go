package dbapi

import(
	"fmt"

	"utl"

	//
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

type DB_SerV struct{
	__session *mgo.Session	
}
 
var(
	dbserv DB_SerV 
	//__session *mgo.Session
)

//const DB_URL = "192.168.0.208:27017"
const DB_URL = "localhost:27017"

func getSession()  (*mgo.Session, error) {

	if( dbserv.__session == nil ){
		var err error
		dbip := loadCfg() 
		dbserv.__session, err = mgo.Dial(  dbip )
		//fmt.Println("sess:", dbserv.__session)
		if(err != nil){
			fmt.Println("connect db error:", err )
			return dbserv.__session, err
		}else{
			dbserv.__session.SetMode(mgo.Eventual, true  )
			//__session.SetMode(mgo.Monotonic   )
		}		
	}

	return dbserv.__session, nil 
}

func Initdb(){
	getSession()
}

func Closedb(){
	dbserv.__session.Close()
}

func GetCollect( dbname, collectname string ) (*mgo.Collection, error) {

	sess, err := getSession()
 
	if( sess == nil ){
		return nil,  err  
	}else{
		return sess.DB( dbname ).C( collectname ), nil
	}
 
}

/////////////////////////////////////////////////////////////////////////////////////
func loadCfg() string {
    ini_parser := utl.IniParser{}

    conf_file_name := "config.ini"
    if err := ini_parser.Load("config.ini"); err != nil {
        fmt.Printf("try load config file[%s] error[%s]\n", conf_file_name, err.Error())
        return DB_URL 
    }

	ip := ini_parser.GetString("server", "dbip")	
	fmt.Println("ip:", ip)
	return ip
}
package main

import (
	"github.com/himanshuo/appdirs/appdirs"
	"fmt"
)

func main(){
	//dataDir,err := appdirs.UserDataDir("myapp","appauthor","1", false)
	//if err != nil{
		//panic(err)
	//}
	//fmt.Println(dataDir)
	
	//siteDataDir, err := appdirs.SiteDataDir("myapp", "appauthor", "1", false)
	//fmt.Println(siteDataDir)
	
	//userConfigDir, err := appdirs.UserConfigDir("myapp", "appauthor", "1", false)
	//fmt.Println(userConfigDir)
	
	siteDataDir, _ := appdirs.SiteDataDir("s", "s", "0", false)
	fmt.Println(siteDataDir)


}

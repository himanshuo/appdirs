package appdirs
import (
	//"log"
	//"runtime"
	//"path"
	"os"
	//"os/user"
	//"strconv"
	"path/filepath"
	"strings"
	//"unicode/utf8"
	"fmt"
	"errors"
)




//todo: iota

var system  string
var pathListSeperatorString string
var badLinuxAppNames []string


func init(){
	system = "linux"
	pathListSeperatorString = fmt.Sprintf("%c", os.PathListSeparator)
	badLinuxAppNames = []string{".", "","..", "/", "/", "./", "/.", "//"}

}

//todo: can make this more versatile
func badAppName(appname string) bool{
	if system == "linux"{
		for _, item := range badLinuxAppNames {
			if item == appname {
				return true
			}
		}
		return false
	}
	return false
}


func UserDataDir(appname string, appauthor string, version string, roaming bool) (string, error){
	if badAppName(appname){
		return "", errors.New("Invalid Application Name")
	}
	
	home, err := homeDir(); 
	if err !=nil {
		return "", err
	}
	
	data_dir := filepath.Join(home, ".local", "share")
	
	//add appname to app_data_dir
	data_dir = filepath.Join(data_dir, appname)
	
	
	//add version to app_data_dir	
	if version != "0" {
		data_dir = filepath.Join(data_dir, version ) 
	} 
	
	return data_dir, nil
}

func SiteDataDir(appname string, appauthor string, version string, multipath bool) (string, error){
	baseDirs := os.Getenv("XDG_DATA_HOME")
	if baseDirs == "" {
		defaultBaseDirs := []string{}
		defaultBaseDirs = append(defaultBaseDirs, "/usr/local/share")
		defaultBaseDirs = append(defaultBaseDirs, "/usr/share")
		
		baseDirs =  strings.Join(defaultBaseDirs, pathListSeperatorString)
		
	}
	
	paths := filepath.SplitList(baseDirs)
	paths = apply(paths, expandTilde)
	
	resultPaths := make([]string, 0)
	
	if !badAppName(appname){
		if version != "0" {
			
			appname = filepath.Join(appname, version)
			
		}
		
		for _, path := range paths {
			
			resultPaths = append(resultPaths, filepath.Join(path,appname))
		}
		
		
	} else {
		return "", errors.New("invalid appname")
	}
	
	if multipath{
		return strings.Join(resultPaths, pathListSeperatorString), nil
	} else {
		return resultPaths[0], nil
	}
	
}

func apply(o []string, f func(string) (string,error)) []string {
    r := make([]string, len(o))
    for i, p := range o {
        //not handling error here
        r[i], _ = f(p)
    }
    return r
}


func UserConfigDir(appname string, appauthor string, version string, roaming bool) (string, error){
	baseDir := os.Getenv("XDG_CONFIG_HOME")
	err := errors.New("")
	if baseDir == "" {
		baseDir,err = expandTilde("~/.config")
		if err !=nil{
			return "", err
		}
	}
	
	path := baseDir
	if appname != "" {
		path = filepath.Join(path, appname)
	}
	if appname != "" && version != "0"{
		path = filepath.Join(path, version)
	}
	return path, nil
}





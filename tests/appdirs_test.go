package tests

import (
	"github.com/himanshuo/appdirs/appdirs"
	"testing"
	"path/filepath"
	"strings"
	"runtime"
)


func determineSystem() appdirs.SystemType {
	
	/* GOOS values
	 * darwin
	 * dragonfly
	 * freebsd
	 * linux
	 * netbsd
	 * openbsd
	 * plan9
	 * solaris
	 * windows
	 */
	if runtime.GOOS == "windows"{
		return appdirs.WINDOWS
	} else if runtime.GOOS == "darwin" {
		return appdirs.MAC
	} else {
		return appdirs.LINUX
	}
	 
}
var platform  appdirs.SystemType
var home string


func init(){
	platform = determineSystem()
	home = "/home/himanshu"
}


var testSiteDataDir = []struct {
	name  string
	author string
	version string
	multipath bool
	ok  bool
}{  //good
	{"a_name","a_author","", false, true},            // simple
	{"a_name","a_author","0", false, true},            // 0 version
	{"a_name","a_author","1", false, true},            // version
	{"a_name","a_author","0.9.1", false, true},        // complex version
	
	{"s", "s", "0", false, true},					   // single char
	{" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	{"a_name","a_author","0", true, true},            // simple multipath
	{"a_name","a_author","1", true, true},            // version multipath
	{"a_name","a_author","0.9.1", true, true},        // complex version multipath
	{"s", "s", "0", false, true},					   // single char multipath
	{"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	//bad
	{".", ".", "0", false, false},					   // dot input
	{"..", "..", "0", false, false},				   // double dot input
	{"//", "//", "0", false, false},				   // double slash
	
}


func TestSiteDataDir(t *testing.T) {
	for i, test := range testSiteDataDir {
		ret, err := appdirs.SiteDataDir(test.name, test.author, test.version, test.multipath)
		
		// ok
		// !ok
		// ret exist
		// err exists
		// ret no exist
		// error no exist
		//t.Errorf("/usr/local/share/a_name == /usr/local/share/a_name is %v", "/usr/local/share/a_name" == "/usr/local/share/a_name")
		
		if platform == appdirs.LINUX {
			// ok/!ok, ret exist, err exist
			if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error exist: ret value and error: %v and %v",i, ret, err)
			// ok/!ok, ret not exist, err not exist
			} else if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error are nil", i)
			// ok, ret exist, err not exist
			} else if  test.ok && ret != "" && err == nil {
				shouldBe := filepath.Join("/usr/local/share/",test.name)
				shouldBe2 := filepath.Join("/usr/share/", test.name)
				if test.version != "" {
					shouldBe = filepath.Join(shouldBe, test.version)
					shouldBe2 = filepath.Join(shouldBe2, test.version)
				}

				if test.multipath {
					shouldBe =  shouldBe + ":" + shouldBe2
				}
				if ret != shouldBe {
					t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, shouldBe, ret)
				}
			// ok, ret not exist, err exist
			} else if test.ok && ret == "" && err != nil {
				t.Errorf("#%d: result was not ok. Expected ok.", i)
			// !ok, ret exist, err not exist
			} else if !test.ok && ret != "" && err == nil {
				t.Errorf("#%d: result was ok. Expected error. Input was: {%v, %v, %v, %v}. Result was:", i, test.name, test.author, test.version, test.multipath, ret)
			// !ok, ret not exist, err exist
			} else if !test.ok && ret == "" && err != nil{
				//expected error, got it. all good.
			} 
	
		}
		
	
	}
}



var testUserDataDir = []struct {
	name  string
	author string
	version string
	roaming bool
	ok  bool
}{  //good
	{"a_name","a_author","", false, true},            // simple
	{"a_name","a_author","0", false, true},            // simple version
	{"a_name","a_author","1", false, true},            // version
	{"a_name","a_author","0.9.1", false, true},        // complex version
	
	{"s", "s", "0", false, true},					   // single char
	{" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	{"a_name","a_author","0", true, true},            // simple multipath
	{"a_name","a_author","1", true, true},            // version multipath
	{"a_name","a_author","0.9.1", true, true},        // complex version multipath
	{"s", "s", "0", false, true},					   // single char multipath
	{"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	//bad
	
	{".", ".", "0", false, false},					   // dot input
	{"..", "..", "0", false, false},				   // double dot input
	{"//", "//", "0", false, false},				   // double slash
	
}


func TestUserDataDir(t *testing.T) {
	for i, test := range testUserDataDir {
		ret, err := appdirs.UserDataDir(test.name, test.author, test.version, test.roaming)
		
	
		if platform == appdirs.LINUX {
			// ok/!ok, ret exist, err exist
			if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error exist: ret value and error: %v and %v",i, ret, err)
			// ok/!ok, ret not exist, err not exist
			} else if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error are nil", i)
			// ok, ret exist, err not exist
			} else if  test.ok && ret != "" && err == nil {
				shouldBe := filepath.Join("/home/himanshu/.local/share",test.name)
				if test.version != "" {
					shouldBe = filepath.Join(shouldBe, test.version)
				}
				if ret != shouldBe {
					t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, shouldBe, ret)
				}
			// ok, ret not exist, err exist
			} else if test.ok && ret == "" && err != nil {
				t.Errorf("#%d: result was not ok. Expected ok.", i)
			// !ok, ret exist, err not exist
			} else if !test.ok && ret != "" && err == nil {
				t.Errorf("#%d: result was ok. Expected error. Input was: {%v, %v, %v, %v}. Result was:", i, test.name, test.author, test.version, test.roaming, ret)
			// !ok, ret not exist, err exist
			} else if !test.ok && ret == "" && err != nil{
				//expected error, got it. all good.
			} 
	
		}
		
	
	}
}


var testSiteConfigDir = []struct {
	name  string
	author string
	version string
	multipath bool
	ok  bool
}{  //good multipath
	{"a_name","a_author","", false, true},            // simple
	{"a_name","a_author","0", false, true},            // simple 0 version
	{"a_name","a_author","1", false, true},            // version
	{"a_name","a_author","0.9.1", false, true},        // complex version
	
	{"s", "s", "0", false, true},					   // single char
	{" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	{"a_name","a_author","0", true, true},            // simple multipath
	{"a_name","a_author","1", true, true},            // version multipath
	{"a_name","a_author","0.9.1", true, true},        // complex version multipath
	{"s", "s", "0", false, true},					   // single char multipath
	{"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	//bad multipath
	{".", ".", "0", true, false},					   // dot input
	{"..", "..", "0", true, false},				   // double dot input
	{"//", "//", "0", true, false},				   // double slash
	
	//good single path
	{"a_name","a_author","0", false, true},            // simple
	{"a_name","a_author","1", false, true},            // version
	{"a_name","a_author","0.9.1", false, true},        // complex version
	
	{"s", "s", "0", false, true},					   // single char
	{" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	{"a_name","a_author","0", true, true},            // simple multipath
	{"a_name","a_author","1", true, true},            // version multipath
	{"a_name","a_author","0.9.1", true, true},        // complex version multipath
	{"s", "s", "0", false, true},					   // single char multipath
	{"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	//bad single path
	{".", ".", "0", false, false},					   // dot input
	{"..", "..", "0", false, false},				   // double dot input
	{"//", "//", "0", false, false},				   // double slash
	
}


func TestSiteConfigDir(t *testing.T) {
	for i, test := range testSiteConfigDir {
		ret, err := appdirs.SiteConfigDir(test.name, test.author, test.version, test.multipath)
		
		if platform == appdirs.LINUX {
			// ok/!ok, ret exist, err exist
			if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error exist: ret value and error: %v and %v",i, ret, err)
			// ok/!ok, ret not exist, err not exist
			} else if ret != "" && err != nil {
				t.Errorf("#%d: Both return value and error are nil", i)
			// ok, ret exist, err not exist
			} else if  test.ok && ret != "" && err == nil {
				if !test.multipath{
					shouldBe := filepath.Join("/etc/xdg/xdg-ubuntu/",test.name)
					if test.version != "" {
						shouldBe = filepath.Join(shouldBe, test.version)
					}
					if ret != shouldBe {
						t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, shouldBe, ret)
					}
				} else {
					shouldBe := []string{"/etc/xdg/xdg-ubuntu","/usr/share/upstart/xdg","/etc/xdg/"}
					for i, path := range shouldBe{
						shouldBe[i] = filepath.Join(path,test.name)
						if test.version != ""{
							shouldBe[i] = filepath.Join(shouldBe[i], test.version)
						}
					}					
					if ret != strings.Join(shouldBe, ":") {
						t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, strings.Join(shouldBe, ":"), ret)
					}
				}
				
			// ok, ret not exist, err exist
			} else if test.ok && ret == "" && err != nil {
				t.Errorf("#%d: result was not ok. Expected ok.", i)
			// !ok, ret exist, err not exist
			} else if !test.ok && ret != "" && err == nil {
				t.Errorf("#%d: result was ok. Expected error. Input was: {%v, %v, %v, %v}. Result was:", i, test.name, test.author, test.version, test.multipath, ret)
			// !ok, ret not exist, err exist
			} else if !test.ok && ret == "" && err != nil{
				//expected error, got it. all good.
			} 
	
		}
		
	
	}
}


var testUserCacheDir = []struct {
	name  string
	author string
	version string
	opinion bool
	ok  bool
}{  //good
	{"a_name","a_author","", false, true},            // simple
	{"a_name","a_author","0", false, true},            // simple 0 version
	{"a_name","a_author","1", false, true},            // version
	{"a_name","a_author","0.9.1", false, true},        // complex version
	
	{"s", "s", "0", false, true},					   // single char
	{" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	{"a_name","a_author","0", true, true},            // simple multipath
	{"a_name","a_author","1", true, true},            // version multipath
	{"a_name","a_author","0.9.1", true, true},        // complex version multipath
	{"s", "s", "0", false, true},					   // single char multipath
	{"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	//bad
	{".", ".", "0", true, false},					   // dot input
	{"..", "..", "0", true, false},				   // double dot input
	{"//", "//", "0", true, false},				   // double slash
	{"a", "a", "", false, false},					   // dot input
	{"a", "a", "0", false, false},				   // double dot input
		
}


//func TestUserCacheDir(t *testing.T) {
	//for i, test := range testUserCacheDir {
		//ret, err := appdirs.UserCacheDir(test.name, test.author, test.version, test.multipath)
		
		//if platform == "linux" {
			//// ok/!ok, ret exist, err exist
			//if ret != "" && err != nil {
				//t.Errorf("#%d: Both return value and error exist: ret value and error: %v and %v",i, ret, err)
			//// ok/!ok, ret not exist, err not exist
			//} else if ret != "" && err != nil {
				//t.Errorf("#%d: Both return value and error are nil", i)
			//// ok, ret exist, err not exist
			//} else if  test.ok && ret != "" && err == nil {
				//shouldBe := filepath.Join("/home/himanshu/.cache",test.name)
				//if test.version != "" {
					//shouldBe = filepath.Join(shouldBe, test.version)
				//}
				//if ret != shouldBe {
					//t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, shouldBe, ret)
				//}	
			//// ok, ret not exist, err exist
			//} else if test.ok && ret == "" && err != nil {
				//t.Errorf("#%d: result was not ok. Expected ok.", i)
			//// !ok, ret exist, err not exist
			//} else if !test.ok && ret != "" && err == nil {
				//t.Errorf("#%d: result was ok. Expected error. Input was: {%v, %v, %v, %v}. Result was:", i, test.name, test.author, test.version, test.multipath, ret)
			//// !ok, ret not exist, err exist
			//} else if !test.ok && ret == "" && err != nil{
				////expected error, got it. all good.
			//} 
	
		//}
		
	
	//}
//}



//var testUserConfigDir = []struct {
	//name  string
	//author string
	//version string
	//roaming bool
	//ok  bool
//}{  //good multipath
	//{"a_name","a_author","", false, true},            // simple
	//{"a_name","a_author","0", false, true},            // simple 0 version
	//{"a_name","a_author","1", false, true},            // version
	//{"a_name","a_author","0.9.1", false, true},        // complex version
	
	//{"s", "s", "0", false, true},					   // single char
	//{" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	//{"a_name","a_author","0", true, true},            // simple multipath
	//{"a_name","a_author","1", true, true},            // version multipath
	//{"a_name","a_author","0.9.1", true, true},        // complex version multipath
	//{"s", "s", "0", false, true},					   // single char multipath
	//{"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	////bad multipath
	//{".", ".", "0", true, false},					   // dot input
	//{"..", "..", "0", true, false},				   // double dot input
	//{"//", "//", "0", true, false},				   // double slash
	
	////good single path
	//{"a_name","a_author","0", false, true},            // simple
	//{"a_name","a_author","1", false, true},            // version
	//{"a_name","a_author","0.9.1", false, true},        // complex version
	
	//{"s", "s", "0", false, true},					   // single char
	//{" 1234567890/.,;'[]\"@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char
		
	//{"a_name","a_author","0", true, true},            // simple multipath
	//{"a_name","a_author","1", true, true},            // version multipath
	//{"a_name","a_author","0.9.1", true, true},        // complex version multipath
	//{"s", "s", "0", false, true},					   // single char multipath
	//{"1234567890/.,;'[]\" ~!@#$%^&*()_+}{|:\"?><", "s", "0", false, true},	// weird char multipath
	
	////bad single path
	//{".", ".", "0", false, false},					   // dot input
	//{"..", "..", "0", false, false},				   // double dot input
	//{"//", "//", "0", false, false},				   // double slash
	
//}


//func TestUserConfigDir(t *testing.T) {
	//for i, test := range testUserConfigDir {
		//ret, err := appdirs.UserConfigDir(test.name, test.author, test.version, test.multipath)
		
		//if platform == appdirs.LINUX {
			//// ok/!ok, ret exist, err exist
			//if ret != "" && err != nil {
				//t.Errorf("#%d: Both return value and error exist: ret value and error: %v and %v",i, ret, err)
			//// ok/!ok, ret not exist, err not exist
			//} else if ret != "" && err != nil {
				//t.Errorf("#%d: Both return value and error are nil", i)
			//// ok, ret exist, err not exist
			//} else if  test.ok && ret != "" && err == nil {
				//if !test.multipath{
					//shouldBe := filepath.Join("/etc/xdg/xdg-ubuntu/",test.name)
					//if test.version != "" {
						//shouldBe = filepath.Join(shouldBe, test.version)
					//}
					//if ret != shouldBe {
						//t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, shouldBe, ret)
					//}
				//} else {
					//shouldBe := []string{"/etc/xdg/xdg-ubuntu","/usr/share/upstart/xdg","/etc/xdg/"}
					//for i, path := range shouldBe{
						//shouldBe[i] = filepath.Join(path,test.name)
						//if test.version != ""{
							//shouldBe[i] = filepath.Join(shouldBe[i], test.version)
						//}
					//}					
					//if ret != strings.Join(shouldBe, ":") {
						//t.Errorf("#%d: Incorrect result. Expected:%v Got:%v", i, strings.Join(shouldBe, ":"), ret)
					//}
				//}
				
			//// ok, ret not exist, err exist
			//} else if test.ok && ret == "" && err != nil {
				//t.Errorf("#%d: result was not ok. Expected ok.", i)
			//// !ok, ret exist, err not exist
			//} else if !test.ok && ret != "" && err == nil {
				//t.Errorf("#%d: result was ok. Expected error. Input was: {%v, %v, %v, %v}. Result was:", i, test.name, test.author, test.version, test.multipath, ret)
			//// !ok, ret not exist, err exist
			//} else if !test.ok && ret == "" && err != nil{
				////expected error, got it. all good.
			//} 
	
		//}
		
	
	//}
//}

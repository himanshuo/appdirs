Appdirs
==



Help! I have a cross-platform application and 
I don't know where to store all my user/site configuration, 
data, cache, and log file/folders!!!!!! 

Never fear, Appdirs is here! :)

Appdirs is based on the python module Appdirs 
(https://github.com/ActiveState/appdirs).

The purpose of this module is to provide simple importable functions
to determine where settings, data, configurations, cache, and logs
directories should be for a given application.

This package supports applications with multiple versions.

Example:

import appdirs

appdirs.UserLogDir("myappname", "myappauthor") returns 
/home/user_name/.local/log/myappname/
on linux.



Functions:
user data dir (UserDataDir)
user config dir (UserConfigDir)
user cache dir (UserCacheDir)
site data dir (SiteDataDir)
site config dir (SiteConfigDir)
user log dir (UserLogDir)

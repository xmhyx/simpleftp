package main

import (
    "./ftp"
)

func main() {

    ftp := new(ftp.FTP)                                                         
    // debug default false
    ftp.Debug = true

    ftp.Connect("127.0.0.1", 2121)  
    ftp.Login("test", "test")
    
    ftp.Pwd()
    ftp.Mkd("test")
    ftp.List()
    
    ftp.Size("HomeFtpServer.exe")
    
    ftp.Request("TYPE I") 
    ftp.Retr("HomeFtpServer.exe","HomeFtpServer.exe") 

    ftp.Request("TYPE I") 
    ftp.Retr("None.exe","None.exe") 

    ftp.Stor("ftp3.exe")
    ftp.Stor("none.exe")
    
    ftp.Quit()
 
}

# LogIt

A Logger interface that prints to the console and append messages to a log file. 

#### By default files are created in the root folder of your project, within the logs/ folder using the current date of the server as its name and .log extension:
*logs/2019_06_11.log* 

#### That can be easily changed by:
```Go
logit.Log.Filepath = "anotherFolder/mylogfile.txt"
```
*mylogfile.txt* would be created in the anotherFolder, in root folder of the project
#### or 
````Go
logit.Log.Filepath = build.Default.GOPATH + "/logs/myapp.log"
````
*myapp.log* would be created in your GOPATH folder
#### or
```Go
logit.Log.Filepath = fmt.Sprintf("%s/%s%s%s", build.Default.GOPATH, "myapp_logs/logfile_", time.Now().Format("2006_01_02"), ".txt")
```
/home/server/go/myapp_logs/logfile_2019_06_11.txt would be created in your GOPATH folder, inside a folder called myapp_logs

#### Automatic folder creation
The app will try to create any folder and subfolders in case of non existance and will throw an error in case of permition error or other exception


#### To write in the log file you must call the function *WriteLog(category string, msg string, errorTrace string)*:
```Go
logit.Log.WriteLog("error", "This is an error message", logit.Log.GetTraceMsg())
```

#### Default categories:
- error
- emergency
- alert
- critical
- warning
- notice
- info
- debug

#### More categories can be added by calling the function *AppendCategories(map[string][]string)*:
```Go
    nc := map[string][]string{                                                                                      
        "custom1": {"Custom1:", "msg..."},                                                                          
        "cutom2":  {"Custom2:", "mgs..."},                                                                          
    }                                                                                                               
    logit.Log.AppendCategories(nc)         
```

#### Basic use 
```Go
package main

import (
    "github.com/thegorgeouslang/logit"
)

func main() {
    logit.Log.WriteLog("error", "This is an error message", logit.Log.GetTraceMsg())
    logit.Log.WriteLog("emergency", "This is an emergency message", logit.Log.GetTraceMsg())
    logit.Log.WriteLog("alert", "This is an alert message", logit.Log.GetTraceMsg())
    logit.Log.WriteLog("critical", "This is a critical message", logit.Log.GetTraceMsg())
    logit.Log.WriteLog("warning", "This is a warning message", logit.Log.GetTraceMsg())
    logit.Log.WriteLog("notice", "This is a notice message", logit.Log.GetTraceMsg())
    logit.Log.WriteLog("info", "This is an info message", logit.Log.GetTraceMsg())
    logit.Log.WriteLog("debug", "This is a debug message", logit.Log.GetTraceMsg())
}
```

2019/06/12 18:21:17 Error: This is an error message on /server/go/src/app/main.go:8 PID: 37777   
2019/06/12 18:21:17 Emergency: This is an emergency message on /server/go/src/app/main.go:9 PID: 37777   
2019/06/12 18:21:17 Alert: This is an alert message on /yserver/go/src/app/main.go:10 PID: 37777   
2019/06/12 18:21:17 Critical: This is a critical message on /server/go/src/app/main.go:11 PID: 37777    
2019/06/12 18:21:17 Warning: This is a warning message on /server/go/src/app/main.go:12 PID: 37777   
2019/06/12 18:21:17 Notice: This is a notice message on /server/go/src/app/main.go:13 PID: 37777    
2019/06/12 18:21:17 Info: This is a info message on /server/go/src/app/main.go:14 PID: 37777    
2019/06/12 18:21:17 Debug: This is a debug message on /server/go/src/app/main.go:15 PID: 37777    


#### Better use
#### Calling from a setup container and using [godotenv] to retrieve .env file values for path and extension:
By calling if from a container you can have fixed customized settings as well as use other dependencies such as [godotenv] for a main configuration file. 
#### .env
```
logfile_path = "logs/"
logfil_ext = "log"
```

##### libs/logger/logit-container.go
```Go
package logger                                                                           
                                                                                                                    
import (                                                                                                            
    "fmt"                                                                                                           
    "github.com/joho/godotenv"                                                                                      
    "github.com/thegorgeouslang/logit"                                                                              
    "go/build"                                                                                                      
    "os"                                                                                                            
    "time"                                                                                                          
)                                                                                                                   
                                                                                                                    
var LogIt = *logit.Log                                                                                           
                                                                                                                    
func init() {                                                                                                       
                                                                                                                    
    // loading godotenv                                                                                             
    e := godotenv.Load()                                                                                            
    if e != nil {                                                                                  
        fmt.Print(e)                                                                                                
    }                                                                                                               
                                                                                                                    
    // changing the default log file path                                                                           
    LogIt.Filepath = fmt.Sprintf("%s/%s%s%s", build.Default.GOPATH,                  
        os.Getenv("logfile_path"), //                                                                           
        time.Now().Format("2006_01_02"),                                                                            
        os.Getenv("logfile_ext"), //                                                                            
    )                                                                                                           
                                                                                                                    
    // appending custom categories                                                                                  
    nc := map[string][]string{                                                                                      
        "custom1": {"Custom1:", "msg..."},                                                                          
        "cutom2":  {"Custom2:", "mgs..."},                                                                      
    }                                                                                                               
    LogIt.AppendCategories(nc)                                                                                  
}                                                                                                             
```
##### main.go
```Go
package main                                                                                                        
                                                                                                                    
import (                                                                                                            
    log "[yourapp]/logger"                                                                                         
)                                                                                                                   
                                                                                                                    
func main() {                      
    log.LogIt.WriteLog("error", "This is an error message", log.LogIt.GetTraceMsg()) 
    log.LogIt.WriteLog("custom1", "This is a custom message", log.LogIt.GetTraceMsg()) 
}
```
2019/06/12 18:21:17 Error: This is an error message on /server/go/src/app/main.go:8 PID: 37777   
2019/06/12 18:21:17 Cutom1: This is an custom message on /server/go/src/app/main.go:8 PID: 37777    

**by [James Mallon]**

[James Mallon]: <https://www.linkedin.com/in/thiago-mallon/>
[godotenv]: <https://github.com/joho/godotenv>

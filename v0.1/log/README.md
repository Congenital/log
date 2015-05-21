#Log Example

###has five :  DEBUG < INFO < WARN < ERROR < FATAL
###  level:	5       4       3       2       1


------------------------------------
###the fatal level can panic, you need recover it. 
####example : 
    func Test(){
        defer func(){
            if err:= recover(); err != nil {
                //handle 
            }
        }()


        log.Fatal("panic")
    }   

---------------------------------------------
###you can set log level:
####example:
    log.SetLevel(level)
    
    /*
    level should less than DEBUG(5) and large than FATAL(1).
     if less than 0, it's TRASH, can't output all log.
     it's be equal to log.Off()
     WARN, ERROR, FATAL has file and fileine output,
     DEBUG and INFO is no;
     */
     
     
--------------------------------------
###You can close and open the log.

####example : log.xxxOff(), log.xxxOn(), log.On(), log.Off()
        log.DebugOff()
        log.DebugOn()
        log.On()
        log.Off()

------------------------------------
###you can set output file

####example:
	log.SetFile(file *os.File)
	
--------------------------------
####a simple project.
####You can change the level or color of the source.





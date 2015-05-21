
#Log Example Version v0.2

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
		log.FatalNoLine("panic")
		log.FatalTo(io io.Writer)
    }   

---------------------------------------------
     
     
--------------------------------------
###You can close and open the log.

####example : log.xxxOff(), log.xxxOn(), log.On(), log.Off()
        log.DebugOff()
        log.DebugOn()
        log.On()
        log.Off()

------------------------------------
####a simple project.
####You can change the level or color of the source.


###you can set output file
####example:
	log.SetFile(file *os.File)


Change Processing structure, Improve execution efficiency,

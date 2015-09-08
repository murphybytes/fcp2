package tufer

type Logger interface {
  LogDebug(  v ...interface{} )
  LogInfo( v ...interface{} )
  LogWarn( v ...interface{} )
  LogFatal( v ...interface{} )
}

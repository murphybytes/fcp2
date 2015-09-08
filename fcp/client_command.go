////////////////////////////////////////////////////////////////////////////////////////////////
//
// This file contains code used to parse file,  host and user information from source
// and destination command line arguments.  
//
////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"regexp"
	"os/user"
	"os"
	"errors"
)

type FileInformation struct {
	userName string 
	hostName string
	fileName string
	isLocalFile    bool 
}

type ClientCommandParser struct {
	ctx *Context
}

func NewFileInformation( )( *FileInformation, error ) {
	return &FileInformation{ "", "", "", false }, nil
}

func NewClientCommandParser( ctx *Context )( *ClientCommandParser, error ) {
	return &ClientCommandParser{ ctx }, nil
}


func (p *ClientCommandParser) Parse() ( source *FileInformation, dest *FileInformation, err error ) {
	if len( p.ctx.arguments.fileArgs ) != 2 {
		err = errors.New( `File source and/or destination arguments are missing` )
		return
	}

	source, err = p.ParseFileInformation( p.ctx.arguments.fileArgs[0] )

	if err != nil {
		return
	}

	source.isLocalFile = true
	dest, err = p.ParseFileInformation( p.ctx.arguments.fileArgs[1] )

	if err != nil {
		return nil, nil, err
	}

	return 
}

func (p *ClientCommandParser) ParseWithUserAndHost( command string )( info *FileInformation, err error ) {
	info, _ = NewFileInformation()
	userExpr := regexp.MustCompile( `^[a-zA-Z0-9._-]+@`)
	loc := userExpr.FindStringIndex( command )

	if loc != nil {
		info.userName = command[0:(loc[1]-1)]
		command = command[loc[1]:]
	} else {
		user, e := user.Current() 
		if e != nil {
			return nil, e 
		}
		info.userName = user.Username
	}

	hostExpr := regexp.MustCompile( `^[a-zA-Z0-9._-]+:` )
	loc = hostExpr.FindStringIndex( command )

	if loc == nil {
		return nil, errors.New( "Invalid file specification" )
	}

	info.hostName = command[0:loc[1]-1]
	info.fileName = command[loc[1]:]
	return 
}

func (p* ClientCommandParser) ParseWithoutUserAndHost( command string)( info *FileInformation, err error ) {

	expr := regexp.MustCompile( `^[\/a-zA-Z0-9_-]+$` )
	loc := expr.FindStringIndex( command )

	if loc == nil {
		return nil, errors.New( "Invalid file specification" )
	}

	info, _ = NewFileInformation() 
	user, e := user.Current()

	if e != nil {
		return nil, e
	}

	info.userName = user.Username
	info.hostName, err = os.Hostname()
	
	if err != nil {
		return nil, err
	}

	info.fileName = command 

	return 	
}

func (p *ClientCommandParser) ParseFileInformation( command string )( info *FileInformation, err error ) {
	info = new(FileInformation)
	expr := regexp.MustCompile( `^[a-z0-9A-Z._-]+@?[a-z0-9A-Z._-]+:[\/a-z0-9A-Z._-]+$`)
	loc := expr.FindStringIndex( command )

	if loc == nil {
		return p.ParseWithoutUserAndHost( command )
	} 

	return p.ParseWithUserAndHost( command ) 


}



package main

import (
	"testing"
)


func TestGetService( t *testing.T ) {
	context := new(Context)
	context.arguments = new(Arguments)
	context.arguments.port = new(int)
	*context.arguments.port = 8089
	
	fileInfo, _ := NewFileInformation()
	fileInfo.hostName = "foo.com"
	client, _ := NewClient( context )
	expected := "foo.com:8089"
	actual := client.GetService( fileInfo )
	if actual != expected {
		t.Error( "Actual ", actual, " does not match expected ", expected )
	}
	
}

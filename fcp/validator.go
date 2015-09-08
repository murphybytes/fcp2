package main

import (
  "errors"
  "fmt"
  "github.com/murphybytes/fcp2/tufer"
  )


  // fatal validator logs fatal error which will
  // cause application to exit
  type FatalValidator struct {
    logger      tufer.Logger
  }

  func NewFatalValidator( logger tufer.Logger )( *FatalValidator ) {
    return &FatalValidator{ logger }
  }


  func (v *FatalValidator) ValidateInt( actual int, expected int ) {
    if actual != expected {
      v.logger.LogFatal( "Validation failed, actual ", actual, " expected ", expected )
    }
  }

  func (v *FatalValidator) ValidateString( actual string, expected string ) {
    if actual != expected {
      v.logger.LogFatal( "Validation failed, actual ", actual, " expected ", expected )
    }
  }

  type Validator struct {

  }

  // these return error
  func NewValidator(  )( *Validator ) {
    return &Validator{}
  }


  // this return error code and don't exit the
  func (v *Validator) ValidateInt( actual int, expected int )( err error ) {
    if actual != expected {
      err = errors.New(fmt.Sprintf( "Validation failed actual %d expected %d", actual, expected) )
    }
    return
  }

  func (v *Validator) ValidateString( actual string, expected string )( err error ) {
    if actual != expected {
      err = errors.New(fmt.Sprintf("Validation failed actual %s expected %s", actual, expected ))
    }
    return
  }

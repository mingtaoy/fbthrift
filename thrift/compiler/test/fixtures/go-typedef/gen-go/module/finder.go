// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package module

import (
	"bytes"
	"context"
	"sync"
	"fmt"
	thrift "github.com/facebook/fbthrift-go"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = sync.Mutex{}
var _ = bytes.Equal
var _ = context.Background

type Finder interface {
  // Parameters:
  //  - Plate
  ByPlate(plate Plate) (r *Automobile, err error)
  // Parameters:
  //  - Plate
  AliasByPlate(plate Plate) (r *Car, err error)
  // Parameters:
  //  - Plate
  PreviousPlate(plate Plate) (r Plate, err error)
}

type FinderClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
}

func (client *FinderClient) Close() error {
  return client.Transport.Close()
}

func NewFinderClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *FinderClient {
  return &FinderClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewFinderClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *FinderClient {
  return &FinderClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

// Parameters:
//  - Plate
func (p *FinderClient) ByPlate(plate Plate) (r *Automobile, err error) {
  if err = p.sendByPlate(plate); err != nil { return }
  return p.recvByPlate()
}

func (p *FinderClient) sendByPlate(plate Plate)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("byPlate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := FinderByPlateArgs{
  Plate : plate,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *FinderClient) recvByPlate() (value *Automobile, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "byPlate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "byPlate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "byPlate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error3 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error4 error
    error4, err = error3.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error4
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "byPlate failed: invalid message type")
    return
  }
  result := FinderByPlateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Plate
func (p *FinderClient) AliasByPlate(plate Plate) (r *Car, err error) {
  if err = p.sendAliasByPlate(plate); err != nil { return }
  return p.recvAliasByPlate()
}

func (p *FinderClient) sendAliasByPlate(plate Plate)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("aliasByPlate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := FinderAliasByPlateArgs{
  Plate : plate,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *FinderClient) recvAliasByPlate() (value *Car, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "aliasByPlate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "aliasByPlate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "aliasByPlate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error5 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error6 error
    error6, err = error5.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error6
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "aliasByPlate failed: invalid message type")
    return
  }
  result := FinderAliasByPlateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Plate
func (p *FinderClient) PreviousPlate(plate Plate) (r Plate, err error) {
  if err = p.sendPreviousPlate(plate); err != nil { return }
  return p.recvPreviousPlate()
}

func (p *FinderClient) sendPreviousPlate(plate Plate)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("previousPlate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := FinderPreviousPlateArgs{
  Plate : plate,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *FinderClient) recvPreviousPlate() (value Plate, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "previousPlate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "previousPlate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "previousPlate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error7 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error8 error
    error8, err = error7.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error8
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "previousPlate failed: invalid message type")
    return
  }
  result := FinderPreviousPlateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type FinderThreadsafeClient struct {
  Transport thrift.Transport
  ProtocolFactory thrift.ProtocolFactory
  InputProtocol thrift.Protocol
  OutputProtocol thrift.Protocol
  SeqId int32
  Mu sync.Mutex
}

func NewFinderThreadsafeClientFactory(t thrift.Transport, f thrift.ProtocolFactory) *FinderThreadsafeClient {
  return &FinderThreadsafeClient{Transport: t,
    ProtocolFactory: f,
    InputProtocol: f.GetProtocol(t),
    OutputProtocol: f.GetProtocol(t),
    SeqId: 0,
  }
}

func NewFinderThreadsafeClient(t thrift.Transport, iprot thrift.Protocol, oprot thrift.Protocol) *FinderThreadsafeClient {
  return &FinderThreadsafeClient{Transport: t,
    ProtocolFactory: nil,
    InputProtocol: iprot,
    OutputProtocol: oprot,
    SeqId: 0,
  }
}

func (p *FinderThreadsafeClient) Threadsafe() {}

// Parameters:
//  - Plate
func (p *FinderThreadsafeClient) ByPlate(plate Plate) (r *Automobile, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendByPlate(plate); err != nil { return }
  return p.recvByPlate()
}

func (p *FinderThreadsafeClient) sendByPlate(plate Plate)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("byPlate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := FinderByPlateArgs{
  Plate : plate,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *FinderThreadsafeClient) recvByPlate() (value *Automobile, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "byPlate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "byPlate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "byPlate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error9 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error10 error
    error10, err = error9.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error10
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "byPlate failed: invalid message type")
    return
  }
  result := FinderByPlateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Plate
func (p *FinderThreadsafeClient) AliasByPlate(plate Plate) (r *Car, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendAliasByPlate(plate); err != nil { return }
  return p.recvAliasByPlate()
}

func (p *FinderThreadsafeClient) sendAliasByPlate(plate Plate)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("aliasByPlate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := FinderAliasByPlateArgs{
  Plate : plate,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *FinderThreadsafeClient) recvAliasByPlate() (value *Car, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "aliasByPlate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "aliasByPlate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "aliasByPlate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error11 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error12 error
    error12, err = error11.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error12
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "aliasByPlate failed: invalid message type")
    return
  }
  result := FinderAliasByPlateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}

// Parameters:
//  - Plate
func (p *FinderThreadsafeClient) PreviousPlate(plate Plate) (r Plate, err error) {
  p.Mu.Lock()
  defer p.Mu.Unlock()
  if err = p.sendPreviousPlate(plate); err != nil { return }
  return p.recvPreviousPlate()
}

func (p *FinderThreadsafeClient) sendPreviousPlate(plate Plate)(err error) {
  oprot := p.OutputProtocol
  if oprot == nil {
    oprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.OutputProtocol = oprot
  }
  p.SeqId++
  if err = oprot.WriteMessageBegin("previousPlate", thrift.CALL, p.SeqId); err != nil {
      return
  }
  args := FinderPreviousPlateArgs{
  Plate : plate,
  }
  if err = args.Write(oprot); err != nil {
      return
  }
  if err = oprot.WriteMessageEnd(); err != nil {
      return
  }
  return oprot.Flush()
}


func (p *FinderThreadsafeClient) recvPreviousPlate() (value Plate, err error) {
  iprot := p.InputProtocol
  if iprot == nil {
    iprot = p.ProtocolFactory.GetProtocol(p.Transport)
    p.InputProtocol = iprot
  }
  method, mTypeId, seqId, err := iprot.ReadMessageBegin()
  if err != nil {
    return
  }
  if method != "previousPlate" {
    err = thrift.NewApplicationException(thrift.WRONG_METHOD_NAME, "previousPlate failed: wrong method name")
    return
  }
  if p.SeqId != seqId {
    err = thrift.NewApplicationException(thrift.BAD_SEQUENCE_ID, "previousPlate failed: out of sequence response")
    return
  }
  if mTypeId == thrift.EXCEPTION {
    error13 := thrift.NewApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
    var error14 error
    error14, err = error13.Read(iprot)
    if err != nil {
      return
    }
    if err = iprot.ReadMessageEnd(); err != nil {
      return
    }
    err = error14
    return
  }
  if mTypeId != thrift.REPLY {
    err = thrift.NewApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "previousPlate failed: invalid message type")
    return
  }
  result := FinderPreviousPlateResult{}
  if err = result.Read(iprot); err != nil {
    return
  }
  if err = iprot.ReadMessageEnd(); err != nil {
    return
  }
  value = result.GetSuccess()
  return
}


type FinderProcessor struct {
  processorMap map[string]thrift.ProcessorFunction
  handler Finder
}

func (p *FinderProcessor) AddToProcessorMap(key string, processor thrift.ProcessorFunction) {
  p.processorMap[key] = processor
}

func (p *FinderProcessor) GetProcessorFunction(key string) (processor thrift.ProcessorFunction, err error) {
  if processor, ok := p.processorMap[key]; ok {
    return processor, nil
  }
  return nil, nil // generic error message will be sent
}

func (p *FinderProcessor) ProcessorMap() map[string]thrift.ProcessorFunction {
  return p.processorMap
}

func NewFinderProcessor(handler Finder) *FinderProcessor {
  self15 := &FinderProcessor{handler:handler, processorMap:make(map[string]thrift.ProcessorFunction)}
  self15.processorMap["byPlate"] = &finderProcessorByPlate{handler:handler}
  self15.processorMap["aliasByPlate"] = &finderProcessorAliasByPlate{handler:handler}
  self15.processorMap["previousPlate"] = &finderProcessorPreviousPlate{handler:handler}
  return self15
}

type finderProcessorByPlate struct {
  handler Finder
}

func (p *finderProcessorByPlate) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := FinderByPlateArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *finderProcessorByPlate) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("byPlate", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *finderProcessorByPlate) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*FinderByPlateArgs)
  var result FinderByPlateResult
  if retval, err := p.handler.ByPlate(args.Plate); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing byPlate: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}

type finderProcessorAliasByPlate struct {
  handler Finder
}

func (p *finderProcessorAliasByPlate) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := FinderAliasByPlateArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *finderProcessorAliasByPlate) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("aliasByPlate", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *finderProcessorAliasByPlate) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*FinderAliasByPlateArgs)
  var result FinderAliasByPlateResult
  if retval, err := p.handler.AliasByPlate(args.Plate); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing aliasByPlate: " + err.Error())
      return x, x
    }
  } else {
    result.Success = retval
  }
  return &result, nil
}

type finderProcessorPreviousPlate struct {
  handler Finder
}

func (p *finderProcessorPreviousPlate) Read(iprot thrift.Protocol) (thrift.Struct, thrift.Exception) {
  args := FinderPreviousPlateArgs{}
  if err := args.Read(iprot); err != nil {
    return nil, err
  }
  iprot.ReadMessageEnd()
  return &args, nil
}

func (p *finderProcessorPreviousPlate) Write(seqId int32, result thrift.WritableStruct, oprot thrift.Protocol) (err thrift.Exception) {
  var err2 error
  messageType := thrift.REPLY
  switch result.(type) {
  case thrift.ApplicationException:
    messageType = thrift.EXCEPTION
  }
  if err2 = oprot.WriteMessageBegin("previousPlate", messageType, seqId); err2 != nil {
    err = err2
  }
  if err2 = result.Write(oprot); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
    err = err2
  }
  if err2 = oprot.Flush(); err == nil && err2 != nil {
    err = err2
  }
  return err
}

func (p *finderProcessorPreviousPlate) Run(argStruct thrift.Struct) (thrift.WritableStruct, thrift.ApplicationException) {
  args := argStruct.(*FinderPreviousPlateArgs)
  var result FinderPreviousPlateResult
  if retval, err := p.handler.PreviousPlate(args.Plate); err != nil {
    switch err.(type) {
    default:
      x := thrift.NewApplicationException(thrift.INTERNAL_ERROR, "Internal error processing previousPlate: " + err.Error())
      return x, x
    }
  } else {
    result.Success = &retval
  }
  return &result, nil
}


// HELPER FUNCTIONS AND STRUCTURES

// Attributes:
//  - Plate
type FinderByPlateArgs struct {
  Plate Plate `thrift:"plate,1" db:"plate" json:"plate"`
}

func NewFinderByPlateArgs() *FinderByPlateArgs {
  return &FinderByPlateArgs{}
}


func (p *FinderByPlateArgs) GetPlate() Plate {
  return p.Plate
}
func (p *FinderByPlateArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FinderByPlateArgs)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Plate(v)
  p.Plate = temp
}
  return nil
}

func (p *FinderByPlateArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("byPlate_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FinderByPlateArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("plate", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:plate: ", p), err) }
  if err := oprot.WriteString(string(p.Plate)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.plate (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:plate: ", p), err) }
  return err
}

func (p *FinderByPlateArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FinderByPlateArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FinderByPlateResult struct {
  Success *Automobile `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewFinderByPlateResult() *FinderByPlateResult {
  return &FinderByPlateResult{}
}

var FinderByPlateResult_Success_DEFAULT *Automobile
func (p *FinderByPlateResult) GetSuccess() *Automobile {
  if !p.IsSetSuccess() {
    return FinderByPlateResult_Success_DEFAULT
  }
return p.Success
}
func (p *FinderByPlateResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *FinderByPlateResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FinderByPlateResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewAutomobile()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *FinderByPlateResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("byPlate_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FinderByPlateResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *FinderByPlateResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FinderByPlateResult(%+v)", *p)
}

// Attributes:
//  - Plate
type FinderAliasByPlateArgs struct {
  Plate Plate `thrift:"plate,1" db:"plate" json:"plate"`
}

func NewFinderAliasByPlateArgs() *FinderAliasByPlateArgs {
  return &FinderAliasByPlateArgs{}
}


func (p *FinderAliasByPlateArgs) GetPlate() Plate {
  return p.Plate
}
func (p *FinderAliasByPlateArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FinderAliasByPlateArgs)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Plate(v)
  p.Plate = temp
}
  return nil
}

func (p *FinderAliasByPlateArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("aliasByPlate_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FinderAliasByPlateArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("plate", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:plate: ", p), err) }
  if err := oprot.WriteString(string(p.Plate)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.plate (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:plate: ", p), err) }
  return err
}

func (p *FinderAliasByPlateArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FinderAliasByPlateArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FinderAliasByPlateResult struct {
  Success *Car `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewFinderAliasByPlateResult() *FinderAliasByPlateResult {
  return &FinderAliasByPlateResult{}
}

var FinderAliasByPlateResult_Success_DEFAULT *Car
func (p *FinderAliasByPlateResult) GetSuccess() *Car {
  if !p.IsSetSuccess() {
    return FinderAliasByPlateResult_Success_DEFAULT
  }
return p.Success
}
func (p *FinderAliasByPlateResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *FinderAliasByPlateResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FinderAliasByPlateResult)  ReadField0(iprot thrift.Protocol) error {
  p.Success = NewAutomobile()
  if err := p.Success.Read(iprot); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
  }
  return nil
}

func (p *FinderAliasByPlateResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("aliasByPlate_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FinderAliasByPlateResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := p.Success.Write(oprot); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
    }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *FinderAliasByPlateResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FinderAliasByPlateResult(%+v)", *p)
}

// Attributes:
//  - Plate
type FinderPreviousPlateArgs struct {
  Plate Plate `thrift:"plate,1" db:"plate" json:"plate"`
}

func NewFinderPreviousPlateArgs() *FinderPreviousPlateArgs {
  return &FinderPreviousPlateArgs{}
}


func (p *FinderPreviousPlateArgs) GetPlate() Plate {
  return p.Plate
}
func (p *FinderPreviousPlateArgs) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 1:
      if err := p.ReadField1(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FinderPreviousPlateArgs)  ReadField1(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 1: ", err)
} else {
  temp := Plate(v)
  p.Plate = temp
}
  return nil
}

func (p *FinderPreviousPlateArgs) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("previousPlate_args"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField1(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FinderPreviousPlateArgs) writeField1(oprot thrift.Protocol) (err error) {
  if err := oprot.WriteFieldBegin("plate", thrift.STRING, 1); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:plate: ", p), err) }
  if err := oprot.WriteString(string(p.Plate)); err != nil {
  return thrift.PrependError(fmt.Sprintf("%T.plate (1) field write error: ", p), err) }
  if err := oprot.WriteFieldEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write field end error 1:plate: ", p), err) }
  return err
}

func (p *FinderPreviousPlateArgs) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FinderPreviousPlateArgs(%+v)", *p)
}

// Attributes:
//  - Success
type FinderPreviousPlateResult struct {
  Success *Plate `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewFinderPreviousPlateResult() *FinderPreviousPlateResult {
  return &FinderPreviousPlateResult{}
}

var FinderPreviousPlateResult_Success_DEFAULT Plate
func (p *FinderPreviousPlateResult) GetSuccess() Plate {
  if !p.IsSetSuccess() {
    return FinderPreviousPlateResult_Success_DEFAULT
  }
return *p.Success
}
func (p *FinderPreviousPlateResult) IsSetSuccess() bool {
  return p.Success != nil
}

func (p *FinderPreviousPlateResult) Read(iprot thrift.Protocol) error {
  if _, err := iprot.ReadStructBegin(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
  }


  for {
    _, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
    if err != nil {
      return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
    }
    if fieldTypeId == thrift.STOP { break; }
    switch fieldId {
    case 0:
      if err := p.ReadField0(iprot); err != nil {
        return err
      }
    default:
      if err := iprot.Skip(fieldTypeId); err != nil {
        return err
      }
    }
    if err := iprot.ReadFieldEnd(); err != nil {
      return err
    }
  }
  if err := iprot.ReadStructEnd(); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
  }
  return nil
}

func (p *FinderPreviousPlateResult)  ReadField0(iprot thrift.Protocol) error {
  if v, err := iprot.ReadString(); err != nil {
  return thrift.PrependError("error reading field 0: ", err)
} else {
  temp := Plate(v)
  p.Success = &temp
}
  return nil
}

func (p *FinderPreviousPlateResult) Write(oprot thrift.Protocol) error {
  if err := oprot.WriteStructBegin("previousPlate_result"); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err) }
  if err := p.writeField0(oprot); err != nil { return err }
  if err := oprot.WriteFieldStop(); err != nil {
    return thrift.PrependError("write field stop error: ", err) }
  if err := oprot.WriteStructEnd(); err != nil {
    return thrift.PrependError("write struct stop error: ", err) }
  return nil
}

func (p *FinderPreviousPlateResult) writeField0(oprot thrift.Protocol) (err error) {
  if p.IsSetSuccess() {
    if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err) }
    if err := oprot.WriteString(string(*p.Success)); err != nil {
    return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err) }
    if err := oprot.WriteFieldEnd(); err != nil {
      return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err) }
  }
  return err
}

func (p *FinderPreviousPlateResult) String() string {
  if p == nil {
    return "<nil>"
  }
  return fmt.Sprintf("FinderPreviousPlateResult(%+v)", *p)
}



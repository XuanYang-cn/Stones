# Thrift basics

### content

  - Tutorial: [thrift](https://thrift.apache.org/tutorial/)
  -	download: [building from source](https://thrift.apache.org/docs/BuildingFromSource)
  or [download](https://thrift.apache.org/download)
  - write .thrift by [IDL(interface definition language)](https://thrift.apache.org/docs/idl)
  - generate code
  	

	thrift --gen py <Thrift filename>

To recursivly generate source code from a Thrift file and all other Thrift files included by it, run

	thrift -r --gen <language> <Thrift filename>


package conf

// We use zookeeper to store the flag value, and update value of the flag
// once there is a changing of the value stored in zookeeper.

// For example, declaring a flag such as:
// var testFlag = flag.Int("test-flag", 10, "test flag usage.")

// We can modify value of testFlag with following command in zookddper.
// create(set) /confPath/prefix/mysvr.test-flag 15
// carete(set) /confPath/prefix/mysvr.nodeName.test-flag 20

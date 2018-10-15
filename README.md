# flashflag
Flashflag can change value of flag defined in go flag without break your daemon.


<pre><code>
	// create(or set) /conf/mysvr.f-int 5  // for all nodes.
	// create(or set) /conf/mysvr.node1.f-int 5  // for node1 only.
	// carete(or set) /conf/mysvr.node1.f-str ok // for node1 only, change f-str to "ok".
	conf.NewConf("/conf", "mysvr", "node1", zk)
</code></pre>

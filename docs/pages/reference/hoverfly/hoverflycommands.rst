hoverfly commands
-----------------

This page contains the equivalent of running:

.. code:: bash
    
    hoverfly --help

The command's help content has been placed here for convenience.    

:: 

    Usage of hoverfly:
      -add
        	add new user '-add -username hfadmin -password hfpass'
      -admin
        	supply '-admin false' to make this non admin user (defaults to 'true')  (default true)
      -ap string
        	admin port - run admin interface on another port (i.e. '-ap 1234' to run admin UI on port 1234)
      -auth
        	enable authentication, currently it is disabled by default
      -capture
        	start Hoverfly in capture mode - transparently intercepts and saves requests/response
      -cert string
        	CA certificate used to sign MITM certificates
      -cert-name string
        	cert name (default "hoverfly.proxy")
      -cert-org string
        	organisation name for new cert (default "Hoverfly Authority")
      -db string
        	Persistance storage to use - 'boltdb' or 'memory' which will not write anything to disk (default "boltdb")
      -db-path string
        	database location - supply it to provide specific database location (will be created there if it doesn't exist)
      -dest value
        	specify which hosts to process (i.e. '-dest fooservice.org -dest barservice.org -dest catservice.org') - other hosts will be ignored will passthrough'
      -destination string
        	destination URI to catch (default ".")
      -dev
        	supply -dev flag to serve directly from ./static/dist instead from statik binary
      -generate-ca-cert
        	generate CA certificate and private key for MITM
      -httptest.serve string
        	if non-empty, httptest.NewServer serves on this address and blocks
      -import value
        	import from file or from URL (i.e. '-import my_service.json' or '-import http://mypage.com/service_x.json'
      -key string
        	private key of the CA used to sign MITM certificates
      -metrics
        	supply -metrics flag to enable metrics logging to stdout
      -middleware string
        	should proxy use middleware
      -modify
        	start Hoverfly in modify mode - applies middleware (required) to both outgoing and incomming HTTP traffic
      -password string
        	password for new user
      -pp string
        	proxy port - run proxy on another port (i.e. '-pp 9999' to run proxy on port 9999)
      -synthesize
        	start Hoverfly in synthesize mode (middleware is required)
      -test.bench string
        	regular expression per path component to select benchmarks to run
      -test.benchmem
        	print memory allocations for benchmarks
      -test.benchtime duration
        	approximate run time for each benchmark (default 1s)
      -test.blockprofile string
        	write a goroutine blocking profile to the named file after execution
      -test.blockprofilerate int
        	if >= 0, calls runtime.SetBlockProfileRate() (default 1)
      -test.count n
        	run tests and benchmarks n times (default 1)
      -test.coverprofile string
        	write a coverage profile to the named file after execution
      -test.cpu string
        	comma-separated list of number of CPUs to use for each test
      -test.cpuprofile string
        	write a cpu profile to the named file during execution
      -test.memprofile string
        	write a memory profile to the named file after execution
      -test.memprofilerate int
        	if >=0, sets runtime.MemProfileRate
      -test.outputdir string
        	directory in which to write profiles
      -test.parallel int
        	maximum test parallelism (default 8)
      -test.run string
        	regular expression to select tests and examples to run
      -test.short
        	run smaller test suite to save time
      -test.timeout duration
        	if positive, sets an aggregate time limit for all tests
      -test.trace string
        	write an execution trace to the named file after execution
      -test.v
        	verbose: print additional output
      -tls-verification
        	turn on/off tls verification for outgoing requests (will not try to verify certificates) - defaults to true (default true)
      -username string
        	username for new user
      -v	should every proxy request be logged to stdout
      -version
        	get the version of hoverfly
      -webserver
        	start Hoverfly in webserver mode (simulate mode)

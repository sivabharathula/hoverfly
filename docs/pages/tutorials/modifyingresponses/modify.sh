hoverctl start
hoverctl mode capture
curl --proxy http://localhost:8500 http://time.jsontest.com
hoverctl mode simulate
hoverctl middleware "python ./middleware.py"
curl --proxy http://localhost:8500 http://time.jsontest.com
hoverctl stop

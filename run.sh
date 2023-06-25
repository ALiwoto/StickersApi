

buildApp() 
{
	# clear the screen (the terminal)
	clear

	cd docs/
	./build.sh
	cd ..

	echo -e "building stickers Api, please wait a bit..."
	
	rm stickersApi -f
	go build -o stickersApi
}

runApp()
{
	# clear the screen (the terminal)
	clear

	echo -e "we are done building it,\n->now running stickers api...\n-------------------"

	./stickersApi
}

testApp()
{
	# clear the screen (the terminal)
	clear

	echo -e "we are running all test files (*_test.go);\nplease wait a bit"

	go test -v ./...
}

if [ "$1" == "test" ];
then
	testApp;
	exit 0
fi;

operations=0

if [ -z "$1" ] || [ "$1" == "true" ] || [ "$1" == "1" ];
then
	buildApp;
	operations=$((i+1))
fi;

if [ -z "$2" ] || [ "$2" == "true" ] || [ "$2" == "1" ];
then
	runApp;
	operations=$((i+1))
fi;

if [ $operations == 0 ]
then
	echo "You have done nothing!"
fi;

# hail woto.
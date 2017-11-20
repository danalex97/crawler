lsof -ti:30000 | xargs kill -9
go install
cd run
go install
cd ..
run $1&
cd app
npm start

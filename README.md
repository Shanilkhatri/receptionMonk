# reakgo
Simple Framework to quickly build webapps in GoLang


curl -X POST http://localhost:4001/loginbyemail -H "Content-Type: application/json" -d '{"authEmailId": "nihal@123yopmail.com","authSignInOTP":""}' 


curl -X POST http://localhost:4001/matchotp -H "Content-Type: application/json" -d '{"authEmailId": "nihal@123yopmail.com","authSignInOTP":"1234567","emailVerToken":"21345672337867367gsgduydygs"}' 


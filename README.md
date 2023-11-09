# reakgo
Simple Framework to quickly build webapps in GoLang


curl -X POST http://localhost:4001/loginbyemail -H "Content-Type: application/json" -d '{"authEmailId": "nihal@123yopmail.com","authSignInOTP":""}' 


curl -X POST http://localhost:4001/matchotp -H "Content-Type: application/json" -d '{"authEmailId": "admin@yopmail.com","authSignInOTP":"973101","emailVerToken":"$2a$10$Tcuu/9hX.ItYMrXVZiku/.0L0bgY/t55dvdyTNEFrIx23ysAurJXq"}' 


curl -X POST http://localhost:4001/loginbyemail -H "Content-Type: application/json" -d '{"authEmailId": "admin@yopmail.com","authSignInOTP":""}' 
{"Status":"200","Message":"New OTP has been sent, Please check your inbox","Payload":"$2a$10$Tcuu/9hX.ItYMrXVZiku/.0L0bgY/t55dvdyTNEFrIx23ysAurJXq","LastId":0,"TotalRecord":0}

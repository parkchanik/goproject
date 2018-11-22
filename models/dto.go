package models

type SendBox struct {
	Boxidx 			int 
	Sender_address 	string 
	Boxmsg 			string 
	Send_wei 		int 
}

type SendBoxRanking struct {
	Rank 				int
	Sender_address 		string
    Last_boxmsg			string
	Total_take_token 	int
			
}

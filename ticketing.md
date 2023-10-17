# Ticketing System

### Two tables :
- > Tickets
- > Responses

### Tickets :-
> This table will contain each new ticket about an issue that is raised and each row will have two columns `query` & `last_response` which can further be used for making quick responses to `FAQs`.

### Responses :-
> This table will have each response from our side on a particular `query`, each row in `responses` table will have columns called `response_sent` & `query_received` against that response.

## Full Schema of table : tickets
    - id                int64           `UNIQUE/PRIMARY`
    - email/userId      string/int64
    - created_time      int64           `EPOCH-TIMESTAMP`
    - last_updated_on   int64           `EPOCH-TIMESTAMP`
    - status            string          `ENUM(open,in_process,closed)`
    - query             string
    - feedback          string          `ENUm (satisfied,not_satisfied)`
    - last_response     string

## Full Schema of table : responses
    - id                int64           `UNIQUE/PRIMARY`
    - response_sent     string          
    - response_received string
    - ticket_id         int64           `FOREIGN KEY`
    - sent_time         int64           `EPOCH-TIMESTAMP`
    - rec_time          int64           `EPOCH-TIMESTAMP`

## Flow
- > customer raises a ticket using interface and marks his entry into the `tickets` table populating all the columns except `feedback` & `last_response`.

- > We send the first response marking our entry into the `responses` table, populating all the fields except `response_received` & `rec_time` and also changing `status` in tickets table to `in_process`.

- > When customer reverts back on our response the otherwise empty columns are populated as well.

- > Once they are satisfied I am hoping they'll have an option to close that ticket and if they don't then based on their responses we can from their side.

- > if the customer chooses to close it they will also be given a `feedback choice` and yup! that will update our `feedback` field in tickets table and will also update the `last_response` by getting the latest response according to timestamp from the `responses` table.


### Why this flow?

***-This flow is made keeping in mind that the customer just can't spam with responses, he/she can only send their response against ours, once after each response from our side, no matter how long they want to make it but they have to write it in one go. In short it's not a chat.***

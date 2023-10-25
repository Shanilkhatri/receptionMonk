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
    - customer_name     string
    - created_time      int64           `EPOCH-TIMESTAMP`
    - last_updated_on   int64           `EPOCH-TIMESTAMP`
    - status            string          `ENUM(open,in_process,closed)`
    - query             string
    - feedback          string          `ENUM (satisfied,not_satisfied,no_feedback)`
    - last_response     string
    - companyId         int64
## Full Schema of table : responses
    - id                int64           `UNIQUE/PRIMARY`
    - response          string          
    - ticket_id         int64           `FOREIGN KEY`
    - response_time     int64           `EPOCH-TIMESTAMP`
    - type              string          `ENUM(customer,employee,super-admin)`
    - respondee_id      int64           

## Flow
- > customer raises a ticket using interface and marks his entry into the `tickets` table populating all the columns except `last_response`, column `feedback` will be initialised with `no_feedback` at first. 

- > We send the first response marking our entry into the `responses` table, populating all the fields and also changing `status` in tickets table to `in_process`.

- > At each response, whether from Employee or from Customer, a new entry will be made under `responses`.

- > Once they are satisfied I am hoping they'll have an option to close that ticket and if they don't then based on their responses we can from their side.

- > if the customer chooses to close it they will also be given a `feedback choice` and yup! that will update our `feedback` field in tickets table.

- > Whatever will be the latest employee response before closure will be updated as `last_response` under tickets table.


### what is `respondee_id`?

***- It's the id of the person who is responding, if it's the customer the `type` field would be set to customer and `respondee_id` would be customers id and vice-versa.***

### Why this flow?

***- Flow has now been updated and follows a more chat-like approach and if in future we want to switch things up as a chat-based ticketing platform the transition would be smooth.***




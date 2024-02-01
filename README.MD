
# SEARCH ENCRYPT

Pencarian customer data berdasarkan Id yang tidak ter-encrypt dengan baik

Error:
>  ENCRYPT [P0001] ERROR: ERROR : SCP Encrypt Fuction. see file damo_scp_error.log.

ENC-DEC Fields :
- full_name
- birth_place
- surgate_mother_name
- mobile_phone
- email
- residence_address
- residence_phone1
- legal_address
- legal_phone1
- company_address
- company_phone1
- emergency_contact_name
- emergency_contact_address
- emergency_contact_home_phone1
- emergency_contact_office_phone1
- emergency_contact_mobile_phone


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

`app.json`

    {  
      "database_customers": {  
        "debug": true,  
        "driver": "postgres",  
        "host": "",  
        "username": "",  
        "password": "",  
        "name": "",  
        "port": 5432,  
        "options": "TimeZone=Asia/Jakarta",  
        "maxOpenConn": 100,  
        "maxIdleConn": 100,  
        "maxLifeTimeConn": 450,  
        "maxIdleTimeConn": 450  
      },  
      "database_customers_history": {  
          "debug": true,  
          "driver": "postgres",  
          "host": "",  
          "username": "",  
          "password": "",  
          "name": "historical",  
          "port": ,  
          "options": "TimeZone=Asia/Jakarta",  
          "maxOpenConn": 100,  
          "maxIdleConn": 100,  
          "maxLifeTimeConn": 450,  
          "maxIdleTimeConn": 450  
        },  
      "database_invalid_encryption_customers": {  
        "debug": true,  
        "driver": "",  
        "host": "",  
        "username": "",  
        "password": "",  
        "name": "invalid_encryption",  
        "port": 5432,  
        "options": "TimeZone=Asia/Jakarta",  
        "maxOpenConn": 100,  
        "maxIdleConn": 100,  
        "maxLifeTimeConn": 450,  
        "maxIdleTimeConn": 450  
      }  
    }


`database_customers`  -> database customer_domain

`database_customers_history`  -> database historical

`database_invalid_encryption_customers`  -> database menampung hasil data pencarian (localhost)

## Demo

![enter image description here](https://photos.app.goo.gl/2gXVvBSThgdcCrTE6)

example:

    go run main.go  --help  

    go run main.go countAll

    go run main.go findInvalidHistory [min] [max]
    go run main.go findInvalidHistory 1000 40000  

## Appendix

- Data structure : Binary Search Tree
- DB: Postgresql
- Language: Go
- CLI-Lib: github.com/spf13/cobra

## Acknowledgements

-   [Binary Search Tree](https://flaviocopes.com/golang-data-structure-binary-search-tree/)

    




WORK IN PROGRESS!
USE AT YOUR OWN RISK!

# MO Pool

Deploy the MineOnlium stack in two commands :)  

# Get started

* `git clone https://github.com/JeffNeff/MO-Pool.git && cd MO-Pool/`
* replace my address with yours in the `.env` file.
* update the PRIVATE_KEY env with a private key and matching ADDRESS
* update the `MO-Stack/Miningcore.UI-master/be/cmd/be/kodata/js/miningcore.js` file (at a minimum just the API variable to reflect your server's public address. )
* `docker compose up -d` 
* `docker compose -f docker-compose.mc.yml up  -d`
* `docker compose -f docker-compose.bs.yml up  -d`
* ...profit?

# Services Rendered:
* [Grafana](https://grafana.com/): 3001
* BlockExplorer: 40001
* Stratum Mining Server: 4073
* Mining Server Frontend: 3000

( dont forget to open port 60606 for p2p networking :) ) 

# MO Pool

A MineOnlium mining pool service stack.

## System Requirements
* Fast (150+mbs) and stable internet connection 
* SSD (recomended)
* memory >= 8gb  
* cpu >= 4 cores


## Services Rendered
* Stratum Mining Server: 4073
* Mining Server Frontend: 3000

## Get Started

* Clone the repository 
```cmd
git clone https://github.com/MineOnliumOfficial/Mo-Pool.git
```

* Generate a new account

```cmd
cd Mo-Pool/ethsigner/keygen/
./keygen
mv wallet/* ../config/keyFile
mv passwordFile ../config/passwordFile
```

* Modify the `.env` file with the public address that was returned by the keygen.

* Modify the `mcconfig.json` file with the same public address.

* From the root of the repository bring the stack up with `docker compose up -d`.

* (If you plan to use the frontend) Modify the `Miningcore.UI-master/be/cmd/be/kodata/js/miningcore.js` file to reflect the servers public IP address

* Open and forward port `60606` for p2p networking 

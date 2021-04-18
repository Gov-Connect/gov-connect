# Gov-Connect  
Toy application to connect people to their government  

## Dependencies (Raspberry Pi / Debian)  
- Python3  
    `sudo apt update`  
    `sudo apt-get install python3`  
    `sudo apt-get install python3-pandas`  
- NPM and Node.js  
    `curl -sL https://deb.nodesource.com/setup_10.x | sudo bash -`  
    `sudo apt-get install -y nodejs`  
- Golang (https://www.digitalocean.com/community/tutorials/how-to-install-go-on-debian-10)
    `sudo apt-get install golang`  
    `nano ~/.profile`  
    paste in the following to the bottom of the file:  
    `export GOROOT=$HOME/go`  
    `export GOPATH=$HOME/Documents`  
    `export PATH=$PATH:$GOROOT/bin:$GOPATH/bin`  
    then refresh your source file:  
    `source ~/.profile`  



## Start Application

### Kafka Option
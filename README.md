#about
rewrite in go of https://github.com/Nabaixin/TFOLED  
tested on RaspberryPi 4b+ .raspios arm64 , ubuntu arm64,openWrt arm64
#install 
upload to RaspberryPi
### in raspios
do raspi-config enable i2c
### in openwrt
opkg install kmod-i2c-core
opkg install kmod-i2c-bcm2835

## cd /path of go-tfoled/  
sudo chmod +x ./go_tfoled_rpi_linux #arm64  
sudo chmod +x ./*.sh  
sudo ./start.sh


#auto start
## in raspios or ubuntu
sudo ./regAndStartGoTfoledService.sh

## in opwnWrt

vi /etc/rc.local

add line:  
/path of go-tfoled/start.sh
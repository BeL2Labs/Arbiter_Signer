#!/bin/bash

echo_error()
{
    echo -e "\033[1;31mERROR:\033[0m $1"
}

echo_info()
{
    echo -e "\033[1;34mINFO:\033[0m $1"
}

echo_info_green() {
    echo -e "\033[1;32mINFO:\033[0m $1"
}

echo_info_red() {
    echo -e "\033[1;31mERROR:\033[0m $1"
}

#
# record update log
#
update_log()
{
  if [ ! -f $SCRIPT_PATH/update.log ]; then
    echo_error "$SCRIPT_PATH/update.log is not exist"
    echo_info "Create update.log"
    touch $SCRIPT_PATH/update.log
  fi

  local time=$(date "+%Y-%m-%d %H:%M:%S")
  echo_info "$time">>$SCRIPT_PATH/update.log
  echo_info "==========">>$SCRIPT_PATH/update.log
  echo_info "deploy arbiter">>$SCRIPT_PATH/update.log
  echo_info "">>$SCRIPT_PATH/update.log
  if [ $1 == "succeeded" ]; then
      echo_info "$time deploy arbiter succeeded!"
  else
      echo_error "$time deploy arbiter failed!"
  fi
  echo_info "Please check update log via command: cat $SCRIPT_PATH/update.log"
}

#
# check status
#
check_status()
{
 PROCESS_NAME="arbiter"
 if pgrep -x "$PROCESS_NAME" > /dev/null
 then
     echo_info "$PROCESS_NAME is running."
     echo_info_green "Succeed!"
 else
     echo_info "$PROCESS_NAME is not running."
     echo_info_red "Failed!"
 fi
}

#
# deploy arbiter
#
deploy_arbiter()
{
	echo_info $SCRIPT_PATH
  if [ $# -ne 4 ]; then
        echo "Need to use deploy_loan_arbiter.sh [your_arbiter_esc_address] [hex_encoded_btc_private_key] [hex_encoded_esc_private_key] [your_keystore_password]"
        exit 1
  fi

  if [ ! -d "$SCRIPT_PATH" ]; then
		mkdir -p $SCRIPT_PATH/data/logs
		mkdir -p $SCRIPT_PATH/keys
	fi
	cd $SCRIPT_PATH

	#prepare config.yaml
	wget -O conf.tgz https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/conf.tgz
	tar xf conf.tgz
	#mv conf/config.yaml .
  sed -i "s/0x0262aB0ED65373cC855C34529fDdeAa0e686D913/$1/g" config.yaml

	#prepare arbiter
  if [ "$(uname -m)" == "armv6l" ] || [ "$(uname -m)" == "armv7l" ] || [ "$(uname -m)" == "aarch64" ]; then
    echo "The current system architecture is ARM"
    echo_info "Downloading loan arbiter..."
    wget -O loan-arbiter-linux-arm64.tgz https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/loan-arbiter-linux-arm64.tgz
    tar xf loan-arbiter-linux-arm64.tgz
    #prepare keystore
    mv loan-arbiter-linux-arm64/keystore-generator .
    ./keystore-generator -c btc -s $2 -p $4 -o btcKey
    ./keystore-generator -c eth -s $3 -p $4 -o escKey
    mv btcKey escKey keys/
    #prepare arbiter
    echo_info "Replacing arbiter.."
    cp -v loan-arbiter-linux-arm64/arbiter ~/loan_arbiter/
    echo_info "Starting arbtier..."
    ./arbiter -p $4  > $SCRIPT_PATH/data/logs/arbiter.log 2>&1 &

    rm -f loan-arbiter-linux-arm64.tgz conf.tgz
  else
    echo "The current system architecture is x86"
    echo_info "Downloading loan arbiter..."
    wget -O loan-arbiter-linux-x86_64.tgz https://download.bel2.org/loan-arbiter/loan-arbiter-v0.0.1/loan-arbiter-linux-x86_64.tgz
    tar xf loan-arbiter-linux-x86_64.tgz
    #prepare keystore
    mv loan-arbiter-linux-x86_64/keystore-generator .
    ./keystore-generator -c btc -s $2 -p $4 -o btcKey
    ./keystore-generator -c eth -s $3 -p $4 -o escKey
    mv btcKey escKey keys/
    #prepare arbiter
    echo_info "Replacing arbiter.."
    cp -v loan-arbiter-linux-x86_64/arbiter ~/loan_arbiter/
    echo_info "Starting arbtier..."
    ./arbiter -p $4 > $SCRIPT_PATH/data/logs/arbiter.log 2>&1 &

    rm -f loan-arbiter-linux-x86_64.tgz conf.tgz
  fi

  check_status
  echo_info "Please check arbiter log via command: cat $SCRIPT_PATH/data/logs/arbiter.log"
}

SCRIPT_PATH=~/loan_arbiter
deploy_arbiter $1 $2 $3 $4
